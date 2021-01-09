package main

import (
	"git.imooc.com/cap1573/category/common"
	"git.imooc.com/cap1573/category/domain/repository"
	"git.imooc.com/cap1573/category/domain/service"
	"git.imooc.com/cap1573/category/handler"
	category "git.imooc.com/cap1573/category/proto/category"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-plugins/registry/consul/v2"
)

func main() {
	/**
	{
	    "host":"127.0.0.1",
	    "user":"root",
	    "pwd":"root",
	    "database":"user",
	    "port":3306
	}
	*/

	//配置中心
	consulConfig, err := common.GetConsulConfig("127.0.0.1", 8500, "/micro/config")

	if err != nil {
		log.Error(err)
	}
	//注册中心
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	// Create service
	srv := micro.NewService(
		micro.Name("go.micro.service.category"),
		micro.Version("latest"),
		//这里设置地址和需要暴露的端口
		micro.Address("127.0.0.1:8082"),
		//添加consul作为注册中心
		micro.Registry(consulRegistry),
	)

	//获取mysql配置,路径中不带前缀
	mysqlInfo := common.GetMysqlFromConsul(consulConfig, "mysql")

	//初始化数据库
	db, err := gorm.Open("mysql", mysqlInfo.User+":"+mysqlInfo.Pwd+"@/"+mysqlInfo.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Error(err)
	}
	defer db.Close()
	db.SingularTable(true)

	//只执行一次
	//rp := repository.NewCategoryRepository(db)
	//rp.InitTable()

	srv.Init()

	categoryDataService := service.NewCategoryDataService(repository.NewCategoryRepository(db))
	err = category.RegisterCategoryHandler(srv.Server(), &handler.Category{CateDataService: categoryDataService})
	if err != nil {
		log.Error(err)
	}

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
