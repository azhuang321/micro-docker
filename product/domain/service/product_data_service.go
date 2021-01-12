package service

import (
	"git.imooc.com/cap1573/product/domain/model"
	"git.imooc.com/cap1573/product/domain/repository"
)

type IProductDataService interface {
	AddProduct(*model.Product) (int64, error)
	DeleteProduct(int64) error
	UpdateProduct(*model.Product) error
	FindProductByID(int64) (*model.Product, error)
	FindAllProduct() ([]model.Product, error)
}

func NewProductDataService(repository repository.IProductRepository) IProductDataService {
	return &ProductDataService{repository}
}

type ProductDataService struct {
	ProductRepository repository.IProductRepository
}

func (u *ProductDataService) AddProduct(product *model.Product) (int64, error) {
	return u.ProductRepository.CreateProduct(product)
}

func (u *ProductDataService) DeleteProduct(productID int64) error {
	return u.ProductRepository.DeleteProduct(productID)
}

func (u *ProductDataService) UpdateProduct(product *model.Product) error {
	return u.ProductRepository.UpdateProduct(product)
}

func (u *ProductDataService) FindProductByID(productID int64) (*model.Product, error) {
	return u.ProductRepository.FindProductByID(productID)
}

func (u *ProductDataService) FindAllProduct() ([]model.Product, error) {
	return u.ProductRepository.FindAll()
}
