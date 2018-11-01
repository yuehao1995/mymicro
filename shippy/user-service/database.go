package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func CreateConnection() (*gorm.DB, error) {
	host := "localhost"
	user := "postgres"
	password := "123456"
	dbName := "postgres"

	return gorm.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s  user=%s dbname=%s sslmode=disable password=%s",
			host, user, dbName, password,
		),
	)
}
