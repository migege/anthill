package main

import (
	"context"
	"log"
	"strings"

	restful "github.com/emicklei/go-restful"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	web "github.com/micro/go-web"
	ahlog "github.com/migege/anthill/proto/log"
)

const (
	API_NAME     = "com.mayibot.ah.api.log"
	SERVICE_NAME = "com.mayibot.ah.log"
)

type Logger struct {
	client ahlog.LoggerClient
}

func (this *Logger) Status(req *restful.Request, rsp *restful.Response) {
	name := req.PathParameter("name")
	elements := strings.Split(name, ",")
	if len(elements) < 3 {
		rsp.WriteEntity(&ahlog.Info{})
		return
	}

	ctx := metadata.NewContext(context.Background(), map[string]string{
		"X-User-Id":    elements[0],
		"X-Host":       elements[1],
		"X-Process-Id": elements[2],
	})

	response, err := this.client.GetStatus(ctx, &ahlog.Info{})
	if err != nil {
		rsp.WriteError(500, err)
		return
	}

	rsp.WriteEntity(response)
}

func main() {
	service := web.NewService(web.Name(API_NAME))
	service.Init()

	logger := &Logger{client: ahlog.NewLoggerClient(SERVICE_NAME, client.DefaultClient)}
	wc := restful.NewContainer()

	ws := new(restful.WebService)
	ws.Consumes(restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON)
	ws.Path("/status")
	ws.Route(ws.GET("/{name}").To(logger.Status))

	wc.Add(ws)
	service.Handle("/", wc)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
