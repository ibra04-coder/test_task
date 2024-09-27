package db

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
    "os"
    "music-library/models"
)

var DB *gorm.DB

func ConnectDatabase() {
    var err error
    dsn := "host=" + os.Getenv("DB_HOST") + 
            " user=" + os.Getenv("DB_USER") + 
            " password=" + os.Getenv("DB_PASSWORD") + 
            " dbname=" + os.Getenv("DB_NAME") + 
            " port=" + os.Getenv("DB_PORT") + 
            " sslmode=disable"
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }
}

func Migrate() {
    DB.AutoMigrate(&models.Song{})
}
