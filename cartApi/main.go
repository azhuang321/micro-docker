package main

import (
	"context"
	"fmt"
	"git.imooc.com/cap1573/cartApi/handler"
	"git.imooc.com/cap1573/cartApi/proto/cart"
	"git.imooc.com/cap1573/cartApi/proto/cartApi"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"github.com/micro/go-plugins/wrapper/select/roundrobin/v2"
	"github.com/micro/micro/v3/service/logger"
	"net"
	"net/http"
)

func main() {
	// 注册中心
	consulRegister := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})
	//todo 添加链路追踪

	//熔断器
	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()
	//启动端口上报
	go func() {
		err := http.ListenAndServe(net.JoinHostPort("0.0.0.0", "9096"), hystrixStreamHandler)
		if err != nil {
			log.Error(err)
		}
	}()

	// Create service
	service := micro.NewService(
		micro.Name("go.micro.api.cartApi"),
		micro.Version("latest"),
		micro.Address("0.0.0.0:8086"),
		//添加注册中心
		micro.Registry(consulRegister),
		//todo 添加链路追踪

		//添加熔断
		micro.WrapClient(NewClientHystrixWrapper()),
		//添加负载均衡
		micro.WrapClient(roundrobin.NewClientWrapper()),
	)
	service.Init()

	cartService := cart.NewCartService("go.micro.service.cart", service.Client())
	err := cartApi.RegisterCartApiHandler(service.Server(), &handler.CartApi{CartService: cartService})
	if err != nil {
		log.Error(err)
	}
	// Run service
	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}

type clientWrapper struct {
	client.Client
}

func (c *clientWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	return hystrix.Do(req.Service()+"."+req.Endpoint(), func() error {
		//run 正常执行裸机
		fmt.Println(req.Service() + "." + req.Endpoint())
		return c.Client.Call(ctx, req, rsp, opts...)
	}, func(err error) error {
		fmt.Println(err)
		return err
	})
}

func NewClientHystrixWrapper() client.Wrapper {
	return func(c client.Client) client.Client {
		return &clientWrapper{c}
	}
}
