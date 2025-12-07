package config

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
    dsn := "root:password@tcp(127.0.0.1:3306)/bookdb?charset=utf8mb4&parseTime=True&loc=Local"
    database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Erro ao conectar ao banco de dados")
    }

    DB = database
}

func GetDB() *gorm.DB {
	return DB
}