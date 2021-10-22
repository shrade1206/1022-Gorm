package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Main_DB *gorm.DB

func main() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               "root:root@tcp(127.0.0.1:3306)/camdb?charset=utf8mb4&parseTime=True&loc=Local",
		DefaultStringSize: 171,
	}), &gorm.Config{
		SkipDefaultTransaction:                   false,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatal("Failed To Connect : ", err.Error())
	}
	sqlDB, _ := db.DB()
	sqlDB.SetConnMaxIdleTime(20) //
	sqlDB.SetConnMaxLifetime(20)
	defer sqlDB.Close()
	// m := db.Migrator()
	// 建立新table
	// m.CreateTable(&Cam{})
	// 查詢table是否存在
	// m.HasTable(&Cam{})
	// 指定Table改名
	// m.RenameTable(&Cam{}, "aaaaa")

	// db.LogMode(true)
	Main_DB = db
	TestCreate()
	// Find()
}
