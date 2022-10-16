package repository

import (
	"search-service/internal/dto"
	"search-service/internal/lib/db"
)

type SearchSearchOption struct {
	dto.Paging
	Keyword string `json:"keyword"`
}

type SearchRepository interface {
}

type SearchRepositoryImpl struct {
	db *db.Database
}

// func (m *SearchRepositoryImpl) Create(ctx context.Context, req models.Search) (res models.Search, err error) {
// 	err = m.db.WithContext(ctx).Create(&req).Error
// 	if err != nil {
// 		return models.Search{}, err
// 	}
// 	return req, nil
// }

// func (m *SearchRepositoryImpl) Search(ctx context.Context, req SearchSearchOption) (res []models.Search, total *int64, err error) {
// 	total = new(int64)
// 	err = m.db.Debug().
// 		Model(&models.Search{}).
// 		WithContext(ctx).
// 		Select("id,name,price,discount").
// 		Where("name ILIKE ?", "%"+req.Keyword+"%").
// 		Limit(req.Limit).Offset(req.Limit * (req.Page - 1)).
// 		Order(req.Sort).
// 		Count(total).
// 		Find(&res).Error
// 	if err != nil {
// 		return []models.Search{}, total, err
// 	}
// 	return
// }

func NewSearchRepository(engine *db.Database) SearchRepository {
	if engine == nil {
		panic("Database engine is null")
	}
	return &SearchRepositoryImpl{db: engine}
}

// func (m *SearchRepositoryImpl) GetByID(ctx context.Context, id int64) (res models.Search, err error) {
// 	err = m.db.WithContext(ctx).Where("id = ?", id).First(&res).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return res, fmt.Errorf(utils.SearchNotFound)
// 		}
// 		return models.Search{}, err
// 	}
// 	return
// }
