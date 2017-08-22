package main

import (
	"time"

	micro "github.com/micro/go-micro"
	ahtime "github.com/migege/anthill/proto/time"
	"golang.org/x/net/context"
)

type TimeHandler struct{}

func (t *TimeHandler) Now(ctx context.Context, req *ahtime.Time, rsp *ahtime.Time) error {
	rsp.Ts = time.Now().Unix()
	rsp.TsUtc = time.Now().UTC().Unix()
	return nil
}

func main() {
	service := micro.NewService(micro.Name("migege.anthill.time"), micro.RegisterTTL(30*time.Second), micro.RegisterInterval(10*time.Second))
	service.Init()

	ahtime.RegisterTimeServiceHandler(service.Server(), new(TimeHandler))
	if err := service.Run(); err != nil {
		panic(err)
	}
}
