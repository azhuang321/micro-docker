package service

import (
	"git.imooc.com/cap1573/cart/domain/model"
	"git.imooc.com/cap1573/cart/domain/repositiry"
)

type ICartDataService interface {
	AddCart(*model.Cart) (int64, error)
	DeleteCart(int64) error
	UpdateCart(*model.Cart) error
	FindCartByID(int64) (*model.Cart, error)
	FindAllCart(int64) ([]model.Cart, error)

	ClearCart(int64) error
	DecrNum(int64, int64) error
	IncrNum(int64, int64) error
}

func NewCartDataService(CartRepository repositiry.ICartRepository) ICartDataService {
	return &CartDataService{CartRepository: CartRepository}
}

type CartDataService struct {
	CartRepository repositiry.ICartRepository
}

func (c *CartDataService) AddCart(cart *model.Cart) (int64, error) {
	return c.CartRepository.CreateCart(cart)
}

func (c *CartDataService) DeleteCart(cartId int64) error {
	return c.CartRepository.DeleteCartByID(cartId)
}

func (c *CartDataService) UpdateCart(cart *model.Cart) error {
	return c.CartRepository.UpdateCart(cart)
}

func (c *CartDataService) FindCartByID(cartId int64) (*model.Cart, error) {
	return c.CartRepository.FindCartByID(cartId)
}

func (c *CartDataService) FindAllCart(userId int64) ([]model.Cart, error) {
	return c.CartRepository.FindAll(userId)
}

func (c *CartDataService) ClearCart(userId int64) error {
	return c.CartRepository.ClearCart(userId)
}

func (c *CartDataService) DecrNum(cartId int64, num int64) error {
	return c.CartRepository.DecrNum(cartId, num)
}

func (c *CartDataService) IncrNum(cartId int64, num int64) error {
	return c.CartRepository.IncrNum(cartId, num)
}
