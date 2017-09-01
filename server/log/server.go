package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"golang.org/x/net/context"

	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"
	ahlog "github.com/migege/anthill/proto/log"
)

const (
	SERVICE_NAME = "com.mayibot.ah.log"
)

var (
	writer          *Writer
	ch_sub_status   map[string][]chan ahlog.Info
	lock_sub_status sync.Mutex
)

type Logger struct{}

func (this *Logger) Log(ctx context.Context, req *ahlog.Info, rsp *ahlog.Response) error {
	if md, ok := metadata.FromContext(ctx); ok {
		fmt.Println(md)
		antId := md["Ant-Id"]
		writer.Write(antId, req.Info)
	}
	return nil
}

func (this *Logger) LogStatus(ctx context.Context, req *ahlog.Info, rsp *ahlog.Response) error {
	if md, ok := metadata.FromContext(ctx); ok {
		fmt.Println(md)
		antId := md["Ant-Id"]
		lock_sub_status.Lock()
		ch_subs := ch_sub_status[antId]
		lock_sub_status.Unlock()
		for _, ch_sub := range ch_subs {
			select {
			case ch_sub <- *req:
			default:
			}
		}
	}
	return nil
}

func (this *Logger) LogProfit(ctx context.Context, req *ahlog.Profit, rsp *ahlog.Response) error {
	err := this.Log(ctx, &ahlog.Info{Info: req.Info, Ts: req.Ts}, rsp)
	return err
}

func (this *Logger) Status(ctx context.Context, req *ahlog.Info, stream ahlog.Logger_StatusStream) error {
	if md, ok := metadata.FromContext(ctx); ok {
		antId := md["Ant-Id"]
		ch_sub := make(chan ahlog.Info, 1)
		lock_sub_status.Lock()
		ch_sub_status[antId] = append(ch_sub_status[antId], ch_sub)
		lock_sub_status.Unlock()

		for {
			select {
			case c := <-ch_sub:
				if err := stream.Send(&c); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func init() {
	writer = &Writer{UserFileMap: make(map[string]struct {
		F  *os.File
		Ts int64
	})}
	ch_sub_status = make(map[string][]chan ahlog.Info)
}

func main() {
	service := micro.NewService(micro.Name(SERVICE_NAME), micro.RegisterTTL(30*time.Second), micro.RegisterInterval(10*time.Second))
	service.Init()

	ahlog.RegisterLoggerHandler(service.Server(), new(Logger))

	ticker := time.NewTicker(1 * time.Minute)
	go func() {
		for {
			select {
			case <-ticker.C:
				writer.GC()
			}
		}
	}()

	if err := service.Run(); err != nil {
		panic(err)
	}
}
