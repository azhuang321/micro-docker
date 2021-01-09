package handler

import (
	"context"
	"git.imooc.com/cap1573/category/common"
	"git.imooc.com/cap1573/category/domain/model"
	"git.imooc.com/cap1573/category/domain/service"
	"git.imooc.com/cap1573/category/proto/category"
	"github.com/micro/go-micro/util/log"
)

type Category struct {
	CateDataService service.ICategoryDataService
}

//提供创建分类
func (c *Category) CreateCategory(ctx context.Context, request *category.CategoryRequest, response *category.CreateCategoryResponse) error {
	category := &model.Category{}
	err := common.SwapTo(request, category)
	if err != nil {
		return err
	}
	categoryId, err := c.CateDataService.AddCategory(category)
	if err != nil {
		return err
	}
	response.Message = "分类添加成功"
	response.CategoryId = categoryId
	return nil
}

//更新分类
func (c *Category) UpdateCategory(ctx context.Context, request *category.CategoryRequest, response *category.UpdateCategoryResponse) error {
	category := &model.Category{}
	err := common.SwapTo(request, category)
	if err != nil {
		return err
	}
	err = c.CateDataService.UpdateCategory(category)
	if err != nil {
		return err
	}
	response.Message = "分类更新成功"
	return nil
}

//分类删除
func (c *Category) DeleteCategory(ctx context.Context, request *category.DeleteCategoryRequest, response *category.DeleteCategoryResponse) error {
	err := c.CateDataService.DeleteCategory(request.CategoryId)
	if err != nil {
		return err
	}
	response.Message = "删除成功"
	return nil
}

//根据分类名称查找
func (c *Category) FindCategoryByName(ctx context.Context, request *category.FindByNameRequest, response *category.CategoryResponse) error {
	categoryByName, err := c.CateDataService.FindCategoryByName(request.CategoryName)
	if err != nil {
		return err
	}
	return common.SwapTo(categoryByName, response)
}

//根据ID查找分类
func (c *Category) FindCategoryById(ctx context.Context, request *category.FindByIdRequest, response *category.CategoryResponse) error {
	categoryById, err := c.CateDataService.FindCategoryById(request.CategoryId)
	if err != nil {
		return err
	}
	return common.SwapTo(categoryById, response)
}

//通过层级查找分类
func (c *Category) FindCategoryByLevel(ctx context.Context, request *category.FindByLevelRequest, response *category.FindAllResponse) error {
	categoryByLevel, err := c.CateDataService.FindCategoryByLevel(request.Level)
	if err != nil {
		return err
	}
	categoryToResponse(categoryByLevel, response)
	return nil
}

//转化函数
func categoryToResponse(categorySlice []model.Category, response *category.FindAllResponse) {
	for _, cg := range categorySlice {
		cr := &category.CategoryResponse{}
		err := common.SwapTo(cg, cr)
		if err != nil {
			log.Error(err)
			break
		}
		response.Category = append(response.Category, cr)
	}
}

func (c *Category) FindCategoryByParent(ctx context.Context, request *category.FindByParentRequest, response *category.FindAllResponse) error {
	categoryByParent, err := c.CateDataService.FindCategoryByParent(request.ParentId)
	if err != nil {
		return err
	}
	categoryToResponse(categoryByParent, response)
	return nil
}

func (c *Category) FindAllCategory(ctx context.Context, request *category.FindAllRequest, response *category.FindAllResponse) error {
	categoryAll, err := c.CateDataService.FindAllCategory()
	if err != nil {
		return err
	}
	categoryToResponse(categoryAll, response)
	return nil
}
