package main

import (
	"github.com/micro/cli"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro"
	"log"

	"./db"
	"./handler"
	"../user-srv/config"
	"../pb"

)

func main() {


	// 创建Service，并定义一些参数
	service := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)

	// 定义Service动作操作
	service.Init(
		micro.Action(func(c *cli.Context) {
			log.Println("micro.Action test ...")
			// 先注册db
			db.Init(config.MysqlDSN)
			log.Println("注冊db成功")
			pb.RegisterUserServiceHandler(service.Server(), handler.NewUserHandler(), server.InternalHandler(true))
		}),
		micro.AfterStop(func() error {
			log.Println("micro.AfterStop test ...")
			return nil
		}),
		micro.AfterStart(func() error {
			log.Println("micro.AfterStart test ...")
			return nil
		}),
	)

	log.Println("启动user-srv服务 ...")

	//启动service
	if err := service.Run(); err != nil {
		log.Panic("user-srv服务启动失败 ...")
	}
}

