package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// criar variavel de conexao com o banco de dados
var (
	db * gorm.DB
)
// funcao de conexao database
func Connect() {
	database, err := gorm.Open("mysql", "root:dev@2025/bookmanager?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Falha ao conectar ao banco de dados!")
	}

	db = database
}
// funcao para retornar a conexao com o banco de dados
func GetDB() * gorm.DB {
	return db
}