package main

import (
	"strconv"
	"time"

	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"
	ahuser "github.com/migege/anthill/proto/user"
	"golang.org/x/net/context"
)

const (
	SERVICE_NAME = "com.mayibot.ah.user"
)

var (
	ch_command map[uint64]chan ahuser.Command
)

type UserHandler struct{}

func (*UserHandler) Register(ctx context.Context, req *ahuser.User, rsp *ahuser.User) error {
	return nil
}

func (*UserHandler) Login(ctx context.Context, req *ahuser.User, rsp *ahuser.User) error {
	// TODO
	switch req.Uid {
	case 10001, 10003, 10004:
		rsp.Expired = 1830268799
		rsp.Status = ahuser.User_ACTIVATED
	case 10005:
		rsp.Expired = 0
		rsp.Status = ahuser.User_UNACTIVATED
	default:
		rsp.Expired = 0
		rsp.Status = ahuser.User_FROZEN
	}
	return nil
}

func (*UserHandler) NewQueen(ctx context.Context, req *ahuser.Queen, rsp *ahuser.Queen) error {
	// TODO
	rsp.Id = 30001
	rsp.Hostname = req.Hostname
	rsp.Pid = req.Pid
	rsp.Osname = req.Osname
	// TODO: return real data
	rsp.IpAddr = req.IpAddr
	rsp.Ctime = req.Ctime
	return nil
}

func (*UserHandler) FireCommand(ctx context.Context, req *ahuser.Command, rsp *ahuser.Response) error {
	if md, ok := metadata.FromContext(ctx); ok {
		if v, ok := md["Queen-Id"]; ok {
			queenId, _ := strconv.ParseUint(v, 10, 64)
			if ch, ok := ch_command[queenId]; ok {
				select {
				case ch <- *req:
					rsp.Code = 0
					rsp.Message = "ok"
				default:
					rsp.Code = 400
					rsp.Message = "worker is offline atm"
				}
			} else {
				rsp.Code = 401
				rsp.Message = "worker does not exist"
			}
		}
	}
	return nil
}

func (*UserHandler) OnCommand(ctx context.Context, req *ahuser.Queen, stream ahuser.UserService_OnCommandStream) error {
	ch := make(chan ahuser.Command)
	ch_command[req.Id] = ch

	for {
		select {
		case c := <-ch:
			if err := stream.Send(&c); err != nil {
				return err
			}
		}
	}

	return nil
}

func init() {
	ch_command = make(map[uint64]chan ahuser.Command)
}

func main() {
	service := micro.NewService(micro.Name(SERVICE_NAME), micro.RegisterTTL(30*time.Second), micro.RegisterInterval(10*time.Second))
	service.Init()

	ahuser.RegisterUserServiceHandler(service.Server(), new(UserHandler))
	if err := service.Run(); err != nil {
		panic(err)
	}
}
