package repositiry

import (
	"errors"
	"git.imooc.com/cap1573/cart/domain/model"
	"github.com/jinzhu/gorm"
)

type ICartRepository interface {
	InitTable() error
	FindCartByID(int64) (*model.Cart, error)
	CreateCart(*model.Cart) (int64, error)
	DeleteCartByID(int64) error
	UpdateCart(*model.Cart) error
	FindAll(int64) ([]model.Cart, error)

	ClearCart(int64) error
	IncrNum(int64, int64) error
	DecrNum(int64, int64) error
}

func NewCartRepository(db *gorm.DB) ICartRepository {
	return &CartRepository{mysqlDb: db}
}

type CartRepository struct {
	mysqlDb *gorm.DB
}

func (c *CartRepository) InitTable() error {
	return c.mysqlDb.CreateTable(&model.Cart{}).Error
}

func (c *CartRepository) FindCartByID(cartId int64) (cart *model.Cart, err error) {
	cart = &model.Cart{}
	return cart, c.mysqlDb.First(cart, cartId).Error
}

func (c *CartRepository) CreateCart(cart *model.Cart) (int64, error) {
	db := c.mysqlDb.FirstOrCreate(cart, model.Cart{ProductID: cart.ProductID, SizeID: cart.SizeID, UserID: cart.UserID})
	if db.Error != nil {
		return 0, db.Error
	}
	if db.RowsAffected == 0 {
		return 0, errors.New("购物车插入失败")
	}
	return cart.ID, nil
}

func (c *CartRepository) DeleteCartByID(cartId int64) error {
	return c.mysqlDb.Where("id = ?", cartId).Delete(&model.Cart{}).Error
}

func (c *CartRepository) UpdateCart(cart *model.Cart) error {
	return c.mysqlDb.Model(cart).Update(cart).Error
}

func (c *CartRepository) FindAll(userId int64) (cartAll []model.Cart, err error) {
	return cartAll, c.mysqlDb.Where("user_id = ?", userId).Find(&cartAll).Error
}

func (c *CartRepository) ClearCart(userId int64) error {
	return c.mysqlDb.Where("user_id = ?", userId).Delete(&model.Cart{}).Error
}

func (c *CartRepository) IncrNum(cartId int64, num int64) error {
	cart := model.Cart{ID: cartId}
	return c.mysqlDb.Model(cart).UpdateColumn("num", gorm.Expr("num + ? ", num)).Error
}

func (c *CartRepository) DecrNum(cartId int64, num int64) error {
	cart := model.Cart{ID: cartId}
	db := c.mysqlDb.Model(cart).UpdateColumn("num", gorm.Expr("num + ? ", num))
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return errors.New("减少失败")
	}
	return nil
}
