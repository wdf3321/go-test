package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// 設定 MongoDB 服務器的地址和端口
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// 建立一個新的客戶端
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 檢查連接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// 之後你可以使用 'client' 進行更多操作，例如創建、讀取、更新和刪除文檔

	// 最後，別忘了斷開與 MongoDB 的連接
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}
