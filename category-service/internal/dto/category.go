package dto

type CreateCategoryRequest struct {
	Name     string `json:"name" validate:"required"`
	SubID    []int  `json:"sub"`
	ParentID *int   `json:"parent_id"`
}

type UpdateCategoryRequest struct {
	ID       int    `json:"id" validate:"required"`
	Name     string `json:"name"`
	SubID    []int  `json:"sub"`
	ParentID int    `json:"parent_id"`
}
type SearchCategoryRequest struct {
	Keyword string `json:"keyword"`
}
