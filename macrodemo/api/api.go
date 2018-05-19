package main

import (
	hello "macrodemo/srv/greeter"
	"context"
	api "github.com/micro/micro/api/proto"
	"log"
	"github.com/micro/go-micro/errors"
	"strings"
	"encoding/json"
	"github.com/micro/go-micro"
)

type Say struct {
	Client hello.GreeterService
}

func (s *Say) Hello(ctx context.Context, req *api.Request, rsp *api.Response ) error {
	log.Print("Received Say.Hello API request")

	name, ok := req.Get["name"]
	if !ok || len(name.Values) == 0 {
		return errors.BadRequest("greeter", "Name cannot be blank")
	}

	response, err := s.Client.Hello(ctx, &hello.HelloRequest{
		Name: strings.Join(name.Values, " "),
	})

	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]string{
		"message": response.Greeting,
	})
	rsp.Body = string(b)
	return nil
}

func main() {

	service := micro.NewService(
		micro.Name("go.micro.api.greeter"),
	)

	// parse command line flags
	service.Init()

	service.Server().Handle(
		service.Server().NewHandler(
			&Say{Client: hello.NewGreeterService("greeter", service.Client())},
		),
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
