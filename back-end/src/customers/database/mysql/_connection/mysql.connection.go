package mysqlConnection

import (
	env "customers/utils/env"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init() *gorm.DB {

	host := env.ReadFile("DB_HOST")
	port := env.ReadFile("DB_PORT")
	user := env.ReadFile("DB_USERNAME")
	pass := env.ReadFile("DB_PASSWORD")
	db_name := env.ReadFile("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, db_name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return db
}
