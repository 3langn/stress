package repository

import (
	"category-service/internal/lib/db"
	"category-service/internal/models"
	"category-service/utils"
	"context"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type FindOneCategoryOptions struct {
	Name string
}
type GetByIDsOptions struct {
	IDs  []int
	Name string
}

type SearchCategoryOptions struct {
	Keyword string
}

type CategoryRepository interface {
	GetByID(ctx context.Context, id int64) (res models.Category, statusCode int, err error)
	CreateCategory(ctx context.Context, category models.Category) (err error)
	FindOne(ctx context.Context, opts FindOneCategoryOptions) (res []models.Category, err error)
	SearchCategory(ctx context.Context, opts SearchCategoryOptions) (res []models.Category, total *int64, err error)
	GetByIDs(ctx context.Context, opts GetByIDsOptions) (res []models.Category, err error)
}

type CategoryRepositoryImpl struct {
	db *db.Database
}

func NewCategoryRepository(engine *db.Database) CategoryRepository {
	if engine == nil {
		panic("Database engine is null")
	}
	return &CategoryRepositoryImpl{db: engine}
}

func (m *CategoryRepositoryImpl) SearchCategory(ctx context.Context, opts SearchCategoryOptions) (res []models.Category, total *int64, err error) {
	total = new(int64)

	err = m.db.WithContext(ctx).Where("name ILIKE ?", "%"+opts.Keyword+"%").Find(&res).Count(total).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return res, total, fmt.Errorf(utils.ErrCatNotFound)
		}
		return []models.Category{}, total, err
	}
	return
}

func (m *CategoryRepositoryImpl) FindOne(ctx context.Context, opts FindOneCategoryOptions) (res []models.Category, err error) {
	q := m.db.WithContext(ctx).Where("name = ?", opts.Name).Find(&res)

	if q.Error != nil {
		return []models.Category{}, err
	}
	if q.RowsAffected == 0 {
		return res, fmt.Errorf(utils.ErrCatNotFound)
	}
	return
}

func (m *CategoryRepositoryImpl) GetByIDs(ctx context.Context, opts GetByIDsOptions) (res []models.Category, err error) {
	err = m.db.WithContext(ctx).Where("id IN ?", opts.IDs).Find(&res).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return res, fmt.Errorf(utils.ErrCatNotFound)
		}
		return []models.Category{}, err
	}
	return
}

func (m *CategoryRepositoryImpl) GetByID(ctx context.Context, id int64) (res models.Category, statusCode int, err error) {
	err = m.db.WithContext(ctx).Where("id = ?", id).First(&res).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return res, http.StatusNotFound, fmt.Errorf(utils.ErrCatNotFound)
		}
		return models.Category{}, http.StatusInternalServerError, err
	}
	return
}

func (m *CategoryRepositoryImpl) CreateCategory(ctx context.Context, category models.Category) (err error) {

	_, err = m.FindOne(ctx, FindOneCategoryOptions{Name: category.Name})
	if err == nil {
		return fmt.Errorf(utils.ErrCatAlreadyExists)
	}

	err = m.db.WithContext(ctx).Create(&category).Error
	return
}
