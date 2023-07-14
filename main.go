package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/jt_dev?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Println("err:", err)
	}
	var a int64
	err = db.Count(&a).Error
	if err != nil {
		log.Println("err:", err)
	}
	log.Println("count:",a)

}
