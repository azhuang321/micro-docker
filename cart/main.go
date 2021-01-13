package main

import (
	"git.imooc.com/cap1573/cart/common"
	"git.imooc.com/cap1573/cart/domain/repositiry"
	service2 "git.imooc.com/cap1573/cart/domain/service"
	"git.imooc.com/cap1573/cart/handler"
	"git.imooc.com/cap1573/cart/proto/cart"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	ratelimit "github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var QPS = 100

func main() {
	//配置中心
	consulConfig, err := common.GetConsulConfig("127.0.0.1", 8500, "/micro/config")
	if err != nil {
		log.Error(err)
	}
	//注册中心
	consulRegister := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})
	//todo 链路追踪

	mysqlInfo := common.GetMysqlFromConsul(consulConfig, "mysql")
	db, err := gorm.Open("mysql", mysqlInfo.User+":"+mysqlInfo.Pwd+"@/"+mysqlInfo.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Error(err)
	}
	defer db.Close()
	db.SingularTable(true)

	//初始化一次
	err = repositiry.NewCartRepository(db).InitTable()
	if err != nil {
		log.Error(err)
	}

	service := micro.NewService(
		micro.Name("go.micro.service.cart"),
		micro.Version("latest"),
		//暴露服务地址
		micro.Address("0.0.0.0:8087"),
		//注册中心
		micro.Registry(consulRegister),
		//todo 链路追踪

		//添加限流
		micro.WrapHandler(ratelimit.NewHandlerWrapper(QPS)),
	)

	service.Init()

	cartDataService := service2.NewCartDataService(repositiry.NewCartRepository(db))

	cart.RegisterCartHandler(service.Server(), &handler.Cart{CartDataService: cartDataService})

	if err := service.Run(); err != nil {
		log.Error(err)
	}
}
