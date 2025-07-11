package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Connect() *gorm.DB {
	//dsn := "host=localhost dbname=db_blog port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	// 使用 postgres 用户而非 root
	dsn := "host=localhost user=root dbname=db_blog password=123456 port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	log.Println("Database connection established successfully")
	return DB
}
