package database

import (
	"fmt"
	"log"
	"sukvij/galenfers/configs"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(cfg configs.Config) (*gorm.DB, error) {
	var err error
	fmt.Println("haha bro cfg ", cfg)
	for i := 0; i < 10; i++ {
		database, err := gorm.Open(postgres.Open(cfg.DBUrl), &gorm.Config{})
		if err == nil {
			return database, nil
		}
		log.Println("Waiting for database...")
		time.Sleep(2 * time.Second)
	}
	return nil, err
}
