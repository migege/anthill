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
	ch_pub_status   map[string]chan ahlog.Info
	lock_pub_status sync.Mutex
	ch_sub_status   map[string][]chan ahlog.Info
	lock_sub_status sync.Mutex
)

type Logger struct{}

func (this *Logger) Log(ctx context.Context, req *ahlog.Info, rsp *ahlog.Info) error {
	if md, ok := metadata.FromContext(ctx); ok {
		user := md["X-User-Id"]
		pid := md["X-Process-Id"]
		host := md["X-Host"]
		writer.Write(user, host, pid, req.Info)
	}
	return nil
}

func (this *Logger) LogStatus(ctx context.Context, req *ahlog.Info, rsp *ahlog.Info) error {
	if md, ok := metadata.FromContext(ctx); ok {
		user := md["X-User-Id"]
		pid := md["X-Process-Id"]
		host := md["X-Host"]

		key := fmt.Sprintf("%s@%s/%s", user, host, pid)
		lock_pub_status.Lock()
		if _, ok := ch_pub_status[key]; !ok {
			ch_pub_status[key] = make(chan ahlog.Info)
		}
		lock_pub_status.Unlock()

		select {
		case ch_pub_status[key] <- *req:
			lock_sub_status.Lock()
			ch_subs := ch_sub_status[key]
			lock_sub_status.Unlock()
			for _, ch_sub := range ch_subs {
				select {
				case ch_sub <- *req:
				default:
				}
			}
			return nil
		default:
			return nil
		}
	}
	return nil
}

func (this *Logger) Status(ctx context.Context, req *ahlog.Info, stream ahlog.Logger_StatusStream) error {
	if md, ok := metadata.FromContext(ctx); ok {
		user := md["X-User-Id"]
		pid := md["X-Process-Id"]
		host := md["X-Host"]

		key := fmt.Sprintf("%s@%s/%s", user, host, pid)
		ch_sub := make(chan ahlog.Info)
		lock_sub_status.Lock()
		ch_sub_status[key] = append(ch_sub_status[key], ch_sub)
		lock_sub_status.Unlock()

		for {
			select {
			case c := <-ch_sub:
				if err := stream.Send(&ahlog.Info{Info: c.Info, Ts: c.Ts}); err != nil {
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
	ch_pub_status = make(map[string]chan ahlog.Info)
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
