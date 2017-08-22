package main

import (
	"context"

	"github.com/micro/cli"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/registry"
	log "github.com/migege/anthill/proto/log"
)

func main() {
	app := cmd.App()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "s",
			Value: "ah.migege.com:8500",
			Usage: "anthill address info",
		},
	}

	var registry_address string
	before := app.Before
	app.Before = func(ctx *cli.Context) error {
		registry_address = ctx.String("s")
		return before(ctx)
	}

	cmd.Init(cmd.Name("ant client"), cmd.Version("11.0.4"))

	c := log.NewLoggerClient("migege.anthill.log", client.NewClient(client.Registry(registry.NewRegistry(registry.Addrs(registry_address)))))
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"X-User-Id": "mbp",
		"X-From-Id": "script",
	})

	n := 0
	for n < 10 {
		if _, err := c.Log(ctx, &log.LogRequest{
			Info: "你说什么",
		}); err != nil {
			panic(err)
		} else {
		}
		n += 1
	}
}
