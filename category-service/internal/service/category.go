package service

import (
	"category-service/internal/dto"
	"category-service/internal/models"
	"category-service/internal/repository"
	"context"
	"time"
)

type (
	CategoryService interface {
		GetByID(ctx context.Context, id int64) (businessGroup models.Category, statusCode int, err error)
		CreateCategory(ctx context.Context, req *dto.CreateCategoryRequest) (err error)
		SearchCategory(ctx context.Context, req *dto.SearchCategoryRequest) (result []models.Category, total *int64, err error)
	}

	CategoryServiceImpl struct {
		repo           repository.CategoryRepository
		contextTimeout time.Duration
	}
)

func (s *CategoryServiceImpl) SearchCategory(ctx context.Context, req *dto.SearchCategoryRequest) (result []models.Category, total *int64, err error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	result, total, err = s.repo.SearchCategory(ctx, repository.SearchCategoryOptions{
		Keyword: req.Keyword,
	})
	return
}

func NewCategoryService(businessGroupRepo repository.CategoryRepository, timeout time.Duration) CategoryService {
	if businessGroupRepo == nil {
		panic("Category Repository is nil")
	}
	if timeout == 0 {
		panic("Timeout is empty")
	}
	return &CategoryServiceImpl{
		repo:           businessGroupRepo,
		contextTimeout: timeout,
	}
}

func (s *CategoryServiceImpl) GetByID(ctx context.Context, id int64) (businessGroup models.Category, statusCode int, err error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()
	businessGroup, statusCode, err = s.repo.GetByID(ctx, id)
	return
}

func (s *CategoryServiceImpl) CreateCategory(ctx context.Context, req *dto.CreateCategoryRequest) (err error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()
	var result []models.Category

	if len(req.SubID) != 0 {
		result, err = s.repo.GetByIDs(ctx, repository.GetByIDsOptions{
			IDs: req.SubID,
		})

		if err != nil {
			return
		}
	}

	category := models.Category{
		Name:     req.Name,
		Sub:      result,
		ParentID: req.ParentID,
	}

	err = s.repo.CreateCategory(ctx, category)
	return
}
