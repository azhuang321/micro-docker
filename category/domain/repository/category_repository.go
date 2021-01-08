package repository

import (
	"git.imooc.com/cap1573/category/domain/model"
	"github.com/jinzhu/gorm"
)

type ICategoryRepository interface {
	//初始化数据表
	InitTable() error
	FindCategoryByID(int64) (*model.Category, error)
	CreateCategory(category *model.Category) (int64, error)
	DeleteCategoryByID(int64) error
	UpdateCategory(*model.Category) error
	FindAll() ([]*model.Category, error)
}

func NewCategoryRepository(db *gorm.DB) ICategoryRepository {
	return &CategoryRepository{
		mysqlDb: db,
	}
}

type CategoryRepository struct {
	mysqlDb *gorm.DB
}

//初始化表
func (u *CategoryRepository) InitTable() error {
	return u.mysqlDb.CreateTable(&model.Category{}).Error
}

func (u *CategoryRepository) FindCategoryByID(categoryId int64) (category *model.Category, err error) {
	category = &model.Category{}
	return category, u.mysqlDb.First(category, categoryId).Error
}

func (u *CategoryRepository) CreateCategory(category *model.Category) (categoryId int64, err error) {
	return categoryId, u.mysqlDb.Create(category).Error
}

func (u *CategoryRepository) DeleteCategoryByID(categoryId int64) error {
	return u.mysqlDb.Where("id = ?", categoryId).Error
}

func (u *CategoryRepository) UpdateCategory(category *model.Category) error {
	return u.mysqlDb.Model(category).Update(&category).Error
}

func (u *CategoryRepository) FindAll() (categoryAll []*model.Category, err error) {
	return categoryAll, u.mysqlDb.Find(&categoryAll).Error
}
