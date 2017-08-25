package main

import (
	"time"

	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"
	ahuser "github.com/migege/anthill/proto/user"
	"golang.org/x/net/context"
)

type UserHandler struct{}

func (*UserHandler) Register(ctx context.Context, req *ahuser.User, rsp *ahuser.Account) error {
	return nil
}

func (*UserHandler) Login(ctx context.Context, req *ahuser.User, rsp *ahuser.Account) error {
	user := req.GetUsername()
	if md, ok := metadata.FromContext(ctx); ok && md["X-User-Id"] == user {
		switch user {
		case "aib", "test", "zh":
			rsp.TsExpire = 1830268799
			rsp.Frozen = false
		case "mbp":
			rsp.TsExpire = 0
			rsp.Frozen = false
		default:
			rsp.Frozen = true
		}
	}
	return nil
}

func main() {
	service := micro.NewService(micro.Name("migege.anthill.user"), micro.RegisterTTL(30*time.Second), micro.RegisterInterval(10*time.Second))
	service.Init()

	ahuser.RegisterUserServiceHandler(service.Server(), new(UserHandler))
	if err := service.Run(); err != nil {
		panic(err)
	}
}
