package db

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	db, err := gorm.Open(mysql.Open(os.Getenv("DSN")), &gorm.Config{})

	if err != nil {
		fmt.Println("failed to gorm open Database:", err)
		return
	}

	DB = db

	// db.AutoMigrate(&model.User{})
	// db.AutoMigrate(&model.UserToken{})
	// db.AutoMigrate(&model.UserTemporaryPassword{})
	// db.AutoMigrate(&model.UserRating{})
	// db.AutoMigrate(&model.Post{})
	// db.AutoMigrate(&model.PostLike{})
	// db.AutoMigrate(&model.Report{})
	// db.AutoMigrate(&model.Comment{})
}
