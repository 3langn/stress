package repository

import (
	"context"
	"fmt"
	"variant-service/internal/lib/db"
	"variant-service/internal/models"
	"variant-service/utils"

	"gorm.io/gorm"
)

type FindOneVariantOptions struct {
	Name  string
	Value string
}
type GetByIDsOptions struct {
	IDs  []int
	Name string
}

type SearchVariantOptions struct {
	Value string
	Name  string
}

type VariantRepository interface {
	GetByID(ctx context.Context, id int64) (res models.Variant, err error)
	CreateVariant(ctx context.Context, category models.Variant) (err error)
	FindOne(ctx context.Context, opts FindOneVariantOptions) (res []models.Variant, err error)
	SearchVariant(ctx context.Context, opts SearchVariantOptions) (res []models.Variant, total *int64, err error)
	GetByIDs(ctx context.Context, opts GetByIDsOptions) (res []models.Variant, err error)
	FindVariantsByIDs(ctx context.Context, ids []int64) (res []models.Variant, total *int64, err error)
}

type VariantRepositoryImpl struct {
	db *db.Database
}

func (m *VariantRepositoryImpl) FindVariantsByIDs(ctx context.Context, ids []int64) (res []models.Variant, total *int64, err error) {
	total = new(int64)
	err = m.db.Debug().WithContext(ctx).Where("id IN ?", ids).Find(&res).Count(total).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return res, total, fmt.Errorf(utils.VariantNotFound)
		}
		return []models.Variant{}, total, err
	}
	return
}

func NewVariantRepository(engine *db.Database) VariantRepository {
	if engine == nil {
		panic("Database engine is null")
	}
	return &VariantRepositoryImpl{db: engine}
}

func (m *VariantRepositoryImpl) SearchVariant(ctx context.Context, opts SearchVariantOptions) (res []models.Variant, total *int64, err error) {
	total = new(int64)

	q := m.db.WithContext(ctx)
	if opts.Name != "" {
		q = q.Where("name LIKE ?", "%"+opts.Name+"%")
	}
	if opts.Value != "" {
		q = q.Where("value LIKE ?", "%"+opts.Value+"%")
	}

	q = q.Find(&res).Count(total)
	if q.Error != nil {
		if *total == 0 {
			return res, total, fmt.Errorf(utils.VariantNotFound)
		}
		return []models.Variant{}, total, err
	}
	return
}

func (m *VariantRepositoryImpl) FindOne(ctx context.Context, opts FindOneVariantOptions) (res []models.Variant, err error) {
	q := m.db.WithContext(ctx).Where("name = ? and value = ?", opts.Name, opts.Value).Find(&res)

	if q.Error != nil {
		return []models.Variant{}, err
	}
	if q.RowsAffected == 0 {
		return res, fmt.Errorf(utils.VariantNotFound)
	}
	return
}

func (m *VariantRepositoryImpl) GetByIDs(ctx context.Context, opts GetByIDsOptions) (res []models.Variant, err error) {
	err = m.db.WithContext(ctx).Where("id IN ?", opts.IDs).Find(&res).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return res, fmt.Errorf(utils.VariantNotFound)
		}
		return []models.Variant{}, err
	}
	return
}

func (m *VariantRepositoryImpl) GetByID(ctx context.Context, id int64) (res models.Variant, err error) {
	err = m.db.WithContext(ctx).Where("id = ?", id).First(&res).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return res, fmt.Errorf(utils.VariantNotFound)
		}
		return models.Variant{}, err
	}
	return
}

func (m *VariantRepositoryImpl) CreateVariant(ctx context.Context, category models.Variant) (err error) {

	_, err = m.FindOne(ctx, FindOneVariantOptions{Name: category.Name, Value: category.Value})
	if err == nil {
		return fmt.Errorf(utils.ErrVariantExists)
	}

	err = m.db.WithContext(ctx).Create(&category).Error
	return
}
