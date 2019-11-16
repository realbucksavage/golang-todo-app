package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/realbucksavage/todos/database/models"
	"github.com/realbucksavage/todos/lib"
	"strconv"
	"time"
)

// InitDb connects to our postgres database. This function also runs migrations after a successful connection.
func InitDb() *gorm.DB {
	db := openConnection()

	db.LogMode(true)

	models.RunMigrations(db)
	return db
}

func openConnection() *gorm.DB {
	// A connection attempt will be made `maxRetries` time.
	maxRetries, _ := strconv.Atoi(lib.GetEnv("DB_MAX_RETRIES", "3"))

	// The next connection attempt will be made after `waitTime` seconds.
	waitTime, _ := strconv.Atoi(lib.GetEnv("DB_WAIT_TIME", "5"))

	for i := 1; i <= maxRetries; i++ {
		fmt.Printf("Opening Connection; Attempt %d of %d...\n", i, maxRetries)
		dbConfig := "sslmode=disable host=db port=5432 dbname=todos user=tduser password=tdpass"

		db, err := gorm.Open("postgres", dbConfig)
		if err != nil {
			// No need to sleep on the last attempt.
			if i != maxRetries {
				fmt.Printf("Cannot open connection (retrying in %ds): %v\n", waitTime, err)
				time.Sleep(time.Duration(waitTime) * time.Second)
			}
			continue
		}

		return db
	}

	panic(fmt.Errorf("Cannot open database connection after %d retries!\n", maxRetries))
}
