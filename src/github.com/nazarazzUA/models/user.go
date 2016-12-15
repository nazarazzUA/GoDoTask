package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Email  string  `gorm:"type:varchar(100);unique_index"`
	FirstName string  `gorm:"index:idx_name_firs"`
	LastName  string  `gorm:"index:idx_name_last"`
	Password string `gorm:"type:varchar(255)"`
}