package db

import (
	"item-stock/entity"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Use PostgreSQL in gorm
)

var (
    db  *gorm.DB
    err error
)

// Init is initialize db from main function
func Init() {
    db, err = gorm.Open("postgres", "host=0.0.0.0 port=5432 user=item_user dbname=stockdb password=pass sslmode=disable")
    if err != nil {
        panic(err)
    }

	autoMigration()
}

// GetDB is called in models
func GetDB() *gorm.DB {
    return db
}

// Close is closing db
func Close() {
    if err := db.Close(); err != nil {
        panic(err)
    }
}

func autoMigration()  {
	db.AutoMigrate(&entity.Item{})
}