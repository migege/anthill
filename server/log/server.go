package main

import (
	"fmt"
	"os"
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
	writer         *Writer
	status_channel map[string]chan ahlog.Info
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
		if _, ok := status_channel[key]; !ok {
			status_channel[key] = make(chan ahlog.Info, 10)
		}

		select {
		case status_channel[key] <- *req:
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
		for {
			select {
			case c := <-status_channel[key]:
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
	status_channel = make(map[string]chan ahlog.Info)
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
