package models

import "time"

type Product struct {
	Id        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	Name      string `json:"name"`
	SerialNo  string `json:"serial_number"`
}
