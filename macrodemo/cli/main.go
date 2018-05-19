package main

import (
	"context"
	"fmt"

	hello "macrodemo/srv/greeter"
	"github.com/micro/go-micro"
)

func main() {
	// create a new service
	service := micro.NewService()

	// parse command line flags
	service.Init()

	// Use the generated client stub
	cl := hello.NewGreeterService("greeter", service.Client())

	// Make request
	rsp, err := cl.Hello(context.TODO(), &hello.HelloRequest{
		Name: "John",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp.Greeting)
}
