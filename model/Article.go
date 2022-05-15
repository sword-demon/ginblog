package model

import "gorm.io/gorm"

type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null" json:"title"`
	Cid     int    `gorm:"type:bigint;not null" json:"cid"`
	Desc    string `gorm:"type:varchar(255);not null" json:"desc"`
	Content string `gorm:"type:text;not null" json:"content"`
	Img     string `gorm:"type:varchar(255);not null" json:"img"`
}
