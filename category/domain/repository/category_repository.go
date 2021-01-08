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
	FindAll() ([]model.Category, error)
	FindCategoryByName(string) (*model.Category, error)
	FindCategoryByLevel(uint32) ([]model.Category, error)
	FindCategoryByParent(int64) ([]model.Category, error)
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

func (u *CategoryRepository) FindAll() (categoryAll []model.Category, err error) {
	return categoryAll, u.mysqlDb.Find(&categoryAll).Error
}

func (u *CategoryRepository) FindCategoryByName(categoryName string) (category *model.Category, err error) {
	//todo 是否需要创建新的对象
	return category, u.mysqlDb.Where("category_name = ?", categoryName).Find(category).Error
}

func (u *CategoryRepository) FindCategoryByLevel(level uint32) (category []model.Category, err error) {
	return category, u.mysqlDb.Where("category_level = ?", level).Find(category).Error
}

func (u *CategoryRepository) FindCategoryByParent(parent int64) (category []model.Category, err error) {
	return category, u.mysqlDb.Where("category_parent = ?", parent).Find(category).Error
}
