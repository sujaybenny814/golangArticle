package model

type Article struct {
	Id          int    `json:"id" gorm:"primaryKey"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Email       string `json:"email" validate:"required"`
	Status      string `json:"status" gorm:"default:pending"`
}
