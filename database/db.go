package database

import (
	"log"
	"my-gin-app/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	dsn := "host=localhost user=KokiAoyagi password=koouki0802 dbname=test port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("データベース接続に失敗", err)
	}
	log.Println("データベース接続に成功！", DB)

	DB = db

	err = DB.AutoMigrate(&models.Post{})
	if err != nil {
		log.Fatal("マイグレーションに失敗:", err)
	}
	log.Println("マイグレーションに成功！")

	// samplePost := models.Post{
	// 	Title:     "初めての投稿",
	// 	Content:   "これはシードデータ",
	// 	CreatedAt: time.Now(),
	// 	UpdatedAt: time.Now(),
	// }

	// result := DB.Create(&samplePost)
	// if result.Error != nil {
	// 	log.Fatal("シードデータの投入に失敗:", result.Error)
	// }
	// log.Println("シードデータの投入完了！")

}
