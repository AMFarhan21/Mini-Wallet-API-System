package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDBConnection(user, password, host, db string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:5440/%s?sslmode=disable", user, password, host, db)

	gormDB, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return &gorm.DB{}, err
	}

	log.Println("DB:", dsn)
	log.Println("Connected to the database")
	return gormDB, nil
}
