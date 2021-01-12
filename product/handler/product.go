package handler

import (
	"context"
	"git.imooc.com/cap1573/product/common"
	"git.imooc.com/cap1573/product/domain/model"
	"git.imooc.com/cap1573/product/domain/service"
	"git.imooc.com/cap1573/product/proto/product"
)

type Product struct {
	ProductDataService service.IProductDataService
}

func (p *Product) AddProduct(ctx context.Context, request *product.ProductInfo, response *product.ResponseProduct) error {
	productAdd := &model.Product{}
	if err := common.SwapTo(request, productAdd); err != nil {
		return err
	}
	productId, err := p.ProductDataService.AddProduct(productAdd)
	if err != nil {
		return err
	}
	response.ProductId = productId
	return nil
}

func (p *Product) FindProductById(ctx context.Context, request *product.RequestID, response *product.ProductInfo) error {
	productData, err := p.ProductDataService.FindProductByID(request.ProductId)
	if err != nil {
		return err
	}
	if err := common.SwapTo(productData, response); err != nil {
		return err
	}
	return nil
}

func (p *Product) UpdateProduct(ctx context.Context, request *product.ProductInfo, response *product.Response) error {
	productAdd := &model.Product{}
	if err := common.SwapTo(request, productAdd); err != nil {
		return err
	}
	err := p.ProductDataService.UpdateProduct(productAdd)
	if err != nil {
		return err
	}
	response.Msg = "更新成功"
	return nil
}

func (p *Product) DeleteProductByID(ctx context.Context, request *product.RequestID, response *product.Response) error {
	if err := p.ProductDataService.DeleteProduct(request.ProductId); err != nil {
		return err
	}
	response.Msg = "删除成功"
	return nil
}

func (p *Product) FindAllProduct(ctx context.Context, request *product.RequestAll, response *product.AllProduct) error {
	productAll, err := p.ProductDataService.FindAllProduct()
	if err != nil {
		return err
	}
	for _, v := range productAll {
		productInfo := &product.ProductInfo{}
		err := common.SwapTo(v, productInfo)
		if err != nil {
			return err
		}
		response.ProductInfo = append(response.ProductInfo, productInfo)
	}
	return nil
}
