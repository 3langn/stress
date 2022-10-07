package dto

type CreateVariantRequest struct {
	Name  string `json:"name" validate:"required"`
	Value string `json:"value" validate:"required"`
}

type UpdateVariantRequest struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
type SearchVariantRequest struct {
	Value string `json:"value"`
	Name  string `json:"name"`
}

type FindVariantsByIDsRequest struct {
	IDs []int64 `json:"ids"`
}
