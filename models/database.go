// database.go
package models

import (
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
)

var DB *gorm.DB

func Connect() *gorm.DB {
    database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect to the database")
    }

    // Migrate the database without specifying a default value for the id field

    //need to migrate book


    //// will be done next
    err = database.AutoMigrate(&User{})
    if err != nil {
        panic("failed to auto-migrate the database")
    }

    DB = database
    return DB
}