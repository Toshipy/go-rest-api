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

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// データベースに接続
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, dbname)

	fmt.Println(url)

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
