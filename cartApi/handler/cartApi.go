package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	cart "git.imooc.com/cap1573/cartApi/proto/cart"
	cartApi "git.imooc.com/cap1573/cartApi/proto/cartApi"
	log "github.com/micro/go-micro/v2/logger"
)

type CartApi struct {
	CartService cart.CartService
}

// CartApi.Call 通过API向外暴露为 /cartApi/findAll 接受http请求
// 即：/cartApi/call 请求会调用go.micro.api.cartApi 服务的CartApi.Call方法
func (c *CartApi) FindAll(ctx context.Context, request *cartApi.Request, response *cartApi.Response) error {
	log.Info("接受到 /cartApi/findAll 访问请求")
	if _, ok := request.Get["user_id"]; !ok {
		//response.StatusCode = 500
		return errors.New("参数异常")
	}
	userIdString := request.Get["user_id"].Values[0]
	fmt.Println(userIdString)
	userId, err := strconv.ParseInt(userIdString, 10, 64)
	if err != nil {
		return err
	}
	cartAll, err := c.CartService.GetAll(context.TODO(), &cart.CartFindAll{UserId: userId})
	b, err := json.Marshal(cartAll)
	if err != nil {
		return err
	}
	response.StatusCode = 200
	response.Body = string(b)
	return nil
}
