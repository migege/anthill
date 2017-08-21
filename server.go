package main

import (
	"time"

	"golang.org/x/net/context"

	"github.com/micro/go-micro"
	log "github.com/migege/anthill/proto/log"
)

type Logger struct{}

func (this *Logger) Log(ctx context.Context, req *log.LogRequest, rsp *log.LogResponse) error {
	rsp.Msg = req.Info + " logged."
	return nil
}

func main() {
	service := micro.NewService(micro.Name("migege.anthill.log"), micro.RegisterTTL(time.Second*30), micro.RegisterInterval(time.Second*10))
	service.Init()

	log.RegisterLoggerHandler(service.Server(), new(Logger))

	if err := service.Run(); err != nil {
		panic(err)
	}
}
