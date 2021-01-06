package main

import (
	"git.imooc.com/cap1573/user/handler"
	pb "git.imooc.com/cap1573/user/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("user"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterUserHandler(srv.Server(), new(handler.User))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
