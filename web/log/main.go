package main

import (
	"context"
	"errors"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	web "github.com/micro/go-web"
	ahlog "github.com/migege/anthill/proto/log"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func StatusStream(cli ahlog.LoggerClient, ws *websocket.Conn) error {
	var req ahlog.Info
	err := ws.ReadJSON(&req)
	if err != nil {
		return err
	}

	infos := strings.Split(req.Info, ",")
	if len(infos) < 3 {
		return errors.New("invalid request")
	}

	go func() {
		for {
			if _, _, err := ws.NextReader(); err != nil {
				break
			}
		}
	}()

	ctx := metadata.NewContext(context.Background(), map[string]string{
		"X-User-Id":    infos[0],
		"X-Host":       infos[1],
		"X-Process-Id": infos[2],
	})
	stream, err := cli.Status(ctx, &req)
	if err != nil {
		return err
	}
	defer stream.Close()

	for {
		rsp, err := stream.Recv()
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}

		err = ws.WriteJSON(rsp)
		if err != nil {
			if isExpectedClose(err) {
				break
			} else {
				return err
			}
		}
	}
	return nil
}

func isExpectedClose(err error) bool {
	if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
		return false
	}
	return true
}

func main() {
	service := web.NewService(web.Name("com.mayibot.ah.web.log"))
	service.Init()

	cli := ahlog.NewLoggerClient("com.mayibot.ah.log", client.NewClient(client.RequestTimeout(time.Second*120)))

	service.Handle("/", http.FileServer(http.Dir("html")))
	service.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Fatal("Upgrade:", err)
			return
		}
		defer conn.Close()

		if err := StatusStream(cli, conn); err != nil {
			log.Println("Echo:", err)
			return
		}
	})

	if err := service.Run(); err != nil {
		log.Fatal("Run:", err)
	}
}
