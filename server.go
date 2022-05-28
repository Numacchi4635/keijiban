package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	// ロギングを行うパッケージ
	"log"

	// HTTPを扱うパッケージ
	"net/http"

	// Gin
	"github.com/gin-gonic/gin"

	// MySQL用ドライバ
	_ "github.com/jinzhu/gorm/dialects/mysql"

	// コントローラ
	"github.com/Numacchi4635/keijiban/controllers"
)

func main() {
	// サーバーを起動する
	server()
}

func server() {
	// デフォルトのミドルウェアでginのルーターを作成
	// Logger と アプリケーションクラッシュをキャッチするRecoveryミドルウェアを保有しています
	router := gin.Default()

	path, err := filepath.Abs("./")
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	fmt.Println(path)
	fileInfos, err := ioutil.ReadDir(path + "/views")
	if err != nil {
		log.Fatal(err)
	}

	for _, fileInfo := range fileInfos {
		fmt.Println(fileInfo.Name())
	}

	// 環境変数からDEBUG_MODE取得
	DEBUG_MODE := os.Getenv("DEBUG_MODE")
	fmt.Println(DEBUG_MODE)

	// 静的ファイルのパスを指定
	router.Static("/views", path+"/views")

	// ルーターの設定
	// URLへのアクセスに対して静的ページを返す
	router.StaticFS("/keijiban", http.Dir("./views/static"))

	// 全ての掲示板情報のJSONを返す
	router.GET("/fetchAllProducts", controllers.FetchAllProducts)

	// 1つの掲示板情報のJSONを返す
	router.GET("/fetchProduct", controllers.FindProduct)

	// サーバー側の環境変数を返す
	router.GET("/responseServerEnv", controllers.ResponseServerEnv)

	// 管理者パスワードの照合を行う
	router.GET("/superUserPasswordCollation", controllers.SuperUserPasswordCollation)

	// ユーザーパスワードの照合を行う
	router.GET("/userPasswordCollation", controllers.UserPasswordCollation)

	// 掲示板情報をDBへ登録する
	router.POST("/addProduct", controllers.AddProduct)

	// 掲示板情報を削除する
	router.POST("/deleteProduct", controllers.DeleteProduct)

	// SuperUserパスワードを設定・変更
	router.POST("/InsertSuperUserPassword", controllers.AddSuperUser)

	if DEBUG_MODE == "true" {
		if err := router.Run(":8080"); err != nil {
			log.Fatal("Server Run Failed.: ", err)
		}
	} else {
		fmt.Println(os.Getenv("PORT"))
		if err := router.Run(":" + os.Getenv("PORT")); err != nil {
			log.Fatal("Server Run Failed.: ", err)
		}
	}
}
