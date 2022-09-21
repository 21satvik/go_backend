package models

import "time"

type Order struct {
	Id           uint `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time
	ProductRefer int     `json:"product_id"`
	Product      Product `json:"product" gorm:"foreignKey:ProductRefer"`
	UserRefer    int     `json:"user_id"`
	User         User    `json:"user" gorm:"foreignKey:UserRefer"`
}
