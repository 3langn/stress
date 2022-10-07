package service

import (
	"context"
	"errors"
	"net/http"
	"time"
	"variant-service/internal/dto"
	"variant-service/internal/models"
	"variant-service/internal/repository"
	"variant-service/utils"

	"gorm.io/gorm"
)

type (
	VariantService interface {
		GetByID(ctx context.Context, id int64) (variant models.Variant, statusCode int, err error)
		CreateVariant(ctx context.Context, req *dto.CreateVariantRequest) (statusCode int, err error)
		SearchVariant(ctx context.Context, req *dto.SearchVariantRequest) (result []models.Variant, total *int64, err error)
		FindVariantsByIDs(ctx context.Context, req *dto.FindVariantsByIDsRequest) (result []models.Variant, total *int64, statusCode int, err error)
	}

	VariantServiceImpl struct {
		repo           repository.VariantRepository
		contextTimeout time.Duration
	}
)

func (s *VariantServiceImpl) FindVariantsByIDs(ctx context.Context, req *dto.FindVariantsByIDsRequest) (result []models.Variant, total *int64, statusCode int, err error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	result, total, err = s.repo.FindVariantsByIDs(ctx, req.IDs)
	if err != nil {
		return nil, total, http.StatusInternalServerError, err
	}
	return
}

func (s *VariantServiceImpl) SearchVariant(ctx context.Context, req *dto.SearchVariantRequest) (result []models.Variant, total *int64, err error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	result, total, err = s.repo.SearchVariant(ctx, repository.SearchVariantOptions{
		Value: req.Value,
		Name:  req.Name,
	})
	return
}

func NewVariantService(variantRepo repository.VariantRepository, timeout time.Duration) VariantService {
	if variantRepo == nil {
		panic("Variant Repository is nil")
	}
	if timeout == 0 {
		panic("Timeout is empty")
	}
	return &VariantServiceImpl{
		repo:           variantRepo,
		contextTimeout: timeout,
	}
}

func (s *VariantServiceImpl) GetByID(ctx context.Context, id int64) (variant models.Variant, statusCode int, err error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()
	variant, err = s.repo.GetByID(ctx, id)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			statusCode = http.StatusNotFound
			return
		}
		statusCode = http.StatusInternalServerError
	}
	return
}

func (s *VariantServiceImpl) CreateVariant(ctx context.Context, req *dto.CreateVariantRequest) (statusCode int, err error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	_, total, err := s.repo.SearchVariant(ctx, repository.SearchVariantOptions{
		Value: req.Value,
		Name:  req.Name,
	})
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if *total > 0 {
		return http.StatusConflict, errors.New(utils.ErrVariantExists)
	}

	v := models.Variant{
		Name:  req.Name,
		Value: req.Value,
	}

	err = s.repo.CreateVariant(ctx, v)
	return
}
