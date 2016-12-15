package core

import (
	"github.com/jinzhu/gorm"
	"fmt"
)

func GetDb() *gorm.DB{
	db, err := gorm.Open("mysql", "root:root@/go_do_task?charset=utf8&parseTime=True&loc=Local")
	if (err != nil) {
		fmt.Sprintf("Cant conect to db %s", err);
	}
	return db;
}