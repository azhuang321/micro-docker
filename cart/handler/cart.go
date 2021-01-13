package handler

import (
	"context"
	"git.imooc.com/cap1573/cart/common"
	"git.imooc.com/cap1573/cart/domain/model"
	"git.imooc.com/cap1573/cart/domain/service"
	"git.imooc.com/cap1573/cart/proto/cart"
)

type Cart struct {
	CartDataService service.ICartDataService
}

func (c *Cart) AddCart(ctx context.Context, request *cart.CartInfo, response *cart.ResponseAdd) error {
	carts := &model.Cart{}
	err := common.SwapTo(request, carts)
	if err != nil {
		return err
	}
	response.CartId, err = c.CartDataService.AddCart(carts)
	return nil
}

func (c *Cart) ClearCart(ctx context.Context, request *cart.Clean, response *cart.Response) error {
	if err := c.CartDataService.ClearCart(request.UserId); err != nil {
		return err
	}
	response.Msg = "购物车清空成功"
	return nil
}

func (c *Cart) Incr(ctx context.Context, request *cart.Item, response *cart.Response) error {
	if err := c.CartDataService.IncrNum(request.Id, request.ChangeNum); err != nil {
		return err
	}
	response.Msg = "购物车添加成功"
	return nil
}

func (c *Cart) Decr(ctx context.Context, request *cart.Item, response *cart.Response) error {
	if err := c.CartDataService.DecrNum(request.Id, request.ChangeNum); err != nil {
		return err
	}
	response.Msg = "购物车减少成功"
	return nil
}

func (c *Cart) DeleteItemByID(ctx context.Context, request *cart.CartId, response *cart.Response) error {
	if err := c.CartDataService.DeleteCart(request.Id); err != nil {
		return err
	}
	response.Msg = "购物车删除成功"
	return nil
}

func (c *Cart) GetAll(ctx context.Context, request *cart.CartFindAll, response *cart.CartAll) error {
	cartAll, err := c.CartDataService.FindAllCart(request.UserId)
	if err != nil {
		return err
	}
	for _, v := range cartAll {
		carts := &cart.CartInfo{}
		if err := common.SwapTo(v, carts); err != nil {
			return err
		}
		response.CartInfo = append(response.CartInfo, carts)
	}
	return nil
}
