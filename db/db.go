package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	// 環境変数を読み込む
	if os.Getenv("GO_ENV") == "dev" {
		err := godotenv.Load()
		// エラーが発生した場合はログを出力して終了
		if err != nil {
			log.Fatalln(err)
		}
	}

	// データベースに接続
	url := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Connected to database")
	return db
}

// データベースを閉じる
func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()

	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err)
	}
}
