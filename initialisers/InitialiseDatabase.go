package initializers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitialiseDatabase() {
	dsn := "host=ep-purple-wave-ad9d2mzp-pooler.c-2.us-east-1.aws.neon.tech user=neondb_owner password=npg_BqlmWbdVYh53 dbname=neondb port=5432 sslmode=require"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	DB = db
}
