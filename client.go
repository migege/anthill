package main

import (
	"context"
	"fmt"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/metadata"
	log "github.com/migege/anthill/proto/log"
)

func main() {
	cmd.Init()

	cli := log.NewLoggerClient("migege.anthill.log", client.DefaultClient)
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"X-User-Id": "mbp",
		"X-From-Id": "script",
	})

	if rsp, err := cli.Log(ctx, &log.LogRequest{
		Info: "你说什么",
	}); err != nil {
		panic(err)
	} else {
		fmt.Println(rsp.Msg)
	}
}
