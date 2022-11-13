package dto

type CategoryDto struct {
	IDcategory int    `json:"id"`
	Name       string `json:"name"`
}

type CategoriesDto []CategoryDto
