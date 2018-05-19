package main

import (
	"net/http"
	"log"
	"fmt"

	hello "macrodemo/srv/greeter"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-web"
	"context"
)

func main() {
	service := web.NewService(
		// 名字必须和服务端定义的一致
		web.Name("greeter"),
	)

	service.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()

			name := r.Form.Get("name")
			if len(name) == 0 {
				name = "World"
			}

			cl := hello.NewGreeterService("greeter", client.DefaultClient)
			rsp, err := cl.Hello(context.Background(), &hello.HelloRequest{
				Name: name,
			})

			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}

			w.Write([]byte(`<html><body><h1>` + rsp.Greeting + `</h1></body></html>`))
			return
		}

		fmt.Fprint(w, `<html><body><h1>Enter Name<h1><form method=post><input name=name type=text /></form></body></html>`)
	})

	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
