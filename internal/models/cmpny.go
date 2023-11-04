package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name     string `json:"name" gorm:"unique" validate:"required"`
	Location string `json:"location" validate:"required"`
	Field    string `json:"field" validate:"required"`
}

type Jobs struct {
	gorm.Model
	Company      Company `json:"-" gorm:"ForeignKey:cid"`
	Cid          uint    `json:"cid"`
	Name         string  `json:"name"`
	Salary       string  `json:"salary"`
	NoticePeriod string  `json:"notice_period"`
}
