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

var (
	writer      *Writer
	statusCache map[string]ahlog.Info
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
		statusCache[key] = *req
	}
	return nil
}

func (this *Logger) GetStatus(ctx context.Context, req *ahlog.Info, rsp *ahlog.Info) error {
	if md, ok := metadata.FromContext(ctx); ok {
		user := md["X-User-Id"]
		pid := md["X-Process-Id"]
		host := md["X-Host"]

		key := fmt.Sprintf("%s@%s/%s", user, host, pid)
		if c, ok := statusCache[key]; ok {
			rsp.Info = c.Info
			rsp.Ts = c.Ts
		}
	}
	return nil
}

func init() {
	writer = &Writer{UserFileMap: make(map[string]struct {
		F  *os.File
		Ts int64
	})}
	statusCache = make(map[string]ahlog.Info)
}

func main() {
	service := micro.NewService(micro.Name("migege.anthill.log"), micro.RegisterTTL(30*time.Second), micro.RegisterInterval(10*time.Second))
	service.Init()

	ahlog.RegisterLoggerHandler(service.Server(), new(Logger))

	ticker := time.NewTicker(1 * time.Minute)
	go func() {
		select {
		case <-ticker.C:
			writer.GC()
		}
	}()

	if err := service.Run(); err != nil {
		panic(err)
	}
}
