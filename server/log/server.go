package main

import (
	"os"
	"time"

	"golang.org/x/net/context"

	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"
	ahlog "github.com/migege/anthill/proto/log"
)

var (
	writer *Writer
)

type Logger struct{}

func (this *Logger) Log(ctx context.Context, req *ahlog.LogRequest, rsp *ahlog.LogResponse) error {
	if md, ok := metadata.FromContext(ctx); ok {
		user := md["X-User-Id"]
		writer.Write(user, req.Info)
	}
	return nil
}

func init() {
	writer = &Writer{
		UserFileMap: make(map[string]*os.File),
	}
}

func main() {
	service := micro.NewService(micro.Name("migege.anthill.log"), micro.RegisterTTL(30*time.Second), micro.RegisterInterval(10*time.Second))
	service.Init()

	ahlog.RegisterLoggerHandler(service.Server(), new(Logger))

	if err := service.Run(); err != nil {
		panic(err)
	}
}
