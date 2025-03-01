package main

import (
	"fmt"
	"go-rest-api/db"
	"go-rest-api/model"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	// 環境変数を読み込む
	if os.Getenv("GO_ENV") == "dev" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalln(err)
		}
	}
	// データベースに接続
	dbConn := db.NewDB()
	defer fmt.Println("Successfully migrated")

	// データベースを閉じる
	defer db.CloseDB(dbConn)

	// マイグレーションを実行
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}
