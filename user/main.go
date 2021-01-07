package main

import (
	"fmt"
	"git.imooc.com/cap1573/user/domain/repository"
	service2 "git.imooc.com/cap1573/user/domain/service"
	"git.imooc.com/cap1573/user/handler"
	user "git.imooc.com/cap1573/user/proto/user"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
)

func main() {
	//服务参数设置
	srv := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Version("latest"),
	)
	//初始化服务
	srv.Init()

	//创建数据库连接
	db, err := gorm.Open("mysql", "root:root@/micro?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	db.SingularTable(true)

	//只执行一次
	//rp := repository.NewUserRepository(db)
	//rp.InitTable()

	//创建服务实例
	userDataService := service2.NewUserDataService(repository.NewUserRepository(db))

	err = user.RegisterUserHandler(srv.Server(), &handler.User{UserDataService: userDataService})

	if err != nil {
		fmt.Println(err)
	}

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}

}
