package database

import (
	"fmt"
	"numtostr/gotodo/config"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB is the underlying database connection
var DB *gorm.DB

// Connect initiate the database connection and migrate all the tables
func Connect() {
	db, err := gorm.Open(sqlite.Open(config.DB), &gorm.Config{
		NowFunc: func() time.Time { return time.Now().Local() },
	})

	if err != nil {
		fmt.Println("[DATABASE]::CONNECTION_ERROR")
		panic(err)
	}

	// Setting the database connection to use in routes
	DB = db

	fmt.Println("[DATABASE]::CONNECTED")
}
