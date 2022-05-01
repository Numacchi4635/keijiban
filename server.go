package main

import (
	// ロギングを行うパッケージ
	"log"

	// HTTPを扱うパッケージ
	"net/http"

	// Gin
	"github.com/gin-gonic/gin"

	// MySQL用ドライバ
	_ "github.com/jinzhu/gorm/dialects/mysql"

	// コントローラ
	controller "github.com/Numacchi4635/keijiban/controllers/controller"
)

func main() {
	// サーバーを起動する
	server()
}

func server() {
	// デフォルトのミドルウェアでginのルーターを作成
	// Logger と アプリケーションクラッシュをキャッチするRecoveryミドルウェアを保有しています
	router := gin.Default()

/*
	// 静的ファイルのパスを指定
	router.Static("/views", "./views")

	// ルーターの設定
	// URLへのアクセスに対して静的ページを返す
	router.StaticFS("/keijiban", http.Dir("./views/static"))

	// 全ての掲示板情報のJSONを返す
	router.GET("/fetchAllProducts", controller.FetchAllProducts)

	// 1つの掲示板情報のJSONを返す
	router.GET("/fetchProduct", controller.FindProduct)

	// 掲示板情報をDBへ登録する
	router.POST("/addProduct", controller.AddProduct)

	// 掲示板情報を変更する
//	router.POST("/changeStateProduct", controller.ChangeStateProduct)

	// 掲示板情報を削除する
	router.POST("/deleteProduct", controller.DeleteProduct)
*/
	if err := router.Run(":8000"); err != nil {
		log.Fatal("Server Run Failed.: ", err)
	}
}
