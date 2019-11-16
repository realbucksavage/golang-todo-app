package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

// Run migrations on the database.
func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&Todo{})

	fmt.Println("AutoMigrations completed.")
}
