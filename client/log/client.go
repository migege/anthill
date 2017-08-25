package main

import (
	"context"
	"fmt"
	"time"

	"github.com/micro/cli"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/registry"
	log "github.com/migege/anthill/proto/log"
)

var registry_address string
var ant_host string
var ant_user string
var ant_pid string

func init() {
	app := cmd.App()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "s",
			Value: "ah.mayibot.com:8500",
			Usage: "anthill address info",
		},
		cli.StringFlag{
			Name:  "H",
			Usage: "ant hostname",
		},
		cli.StringFlag{
			Name:  "u",
			Usage: "ant user",
		},
		cli.StringFlag{
			Name:  "p",
			Usage: "ant pid",
		},
	}

	before := app.Before
	app.Before = func(ctx *cli.Context) error {
		registry_address = ctx.String("s")
		ant_host = ctx.String("H")
		ant_user = ctx.String("u")
		ant_pid = ctx.String("p")
		return before(ctx)
	}
}

func main() {
	cmd.Init(cmd.Name("ant client"), cmd.Version("15.0.0"))

	c := log.NewLoggerClient("com.mayibot.ah.log", client.NewClient(client.Registry(registry.NewRegistry(registry.Addrs(registry_address)))))
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"X-User-Id":    ant_user,
		"X-From-Id":    "ant",
		"X-Host":       ant_host,
		"X-Process-Id": ant_pid,
	})

	var lastTs int64 = 0
	for {
		if info, err := c.GetStatus(ctx, &log.Info{}); err == nil && info != nil && len(info.Info) > 0 && lastTs < info.Ts {
			lastTs = info.Ts
			fmt.Println(info.Info)
		}
		time.Sleep(1 * time.Second)
	}
}
