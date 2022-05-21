package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"fmt"

	// ロギングを行うパッケージ
	"log"

	// HTTPを扱うパッケージ
	"net/http"

	// Go言語のORM
	"github.com/jinzhu/gorm"

	// Gin
	"github.com/gin-gonic/gin"

	// MySQL用ドライバ
	_ "github.com/jinzhu/gorm/dialects/mysql"

	// コントローラ
	controller "github.com/Numacchi4635/keijiban/controllers/controller"

	// スーパーユーザー（データベースのテーブルの行に対応）
	superuser "github.com/Numacchi4635/keijiban/models/superuser"

	// グローバル変数
	globaldata "github.com/Numacchi4635/keijiban/model/globaldata"
)

func main() {
	// プログラム終了時にDBをCloseする
	defer gorm.Close()

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

fmt.Println(path);
fileInfos, err := ioutil.ReadDir(path+"/views")
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
	router.GET("/fetchAllProducts", controller.FetchAllProducts)

	// 1つの掲示板情報のJSONを返す
	router.GET("/fetchProduct", controller.FindProduct)

	// サーバー側の環境変数を返す
	router.GET("/responseServerEnv", controller.ResponseServerEnv)

	// 管理者パスワードの照合を行う
	router.GET("/superUserPasswordCollation", controller.SuperUserPasswordCollation);

	// 掲示板情報をDBへ登録する
	router.POST("/addProduct", controller.AddProduct)

	// 掲示板情報を変更する
//	router.POST("/changeStateProduct", controller.ChangeStateProduct)

	// 掲示板情報を削除する
	router.POST("/deleteProduct", controller.DeleteProduct)

	// ユーザーパスワードの照合
	router.POST("/UserPasswordCollation", controller.UserPasswordCollation)

	// SuperUserパスワードを設定・変更
	router.POST("/InsertSuperUserPassword", controller.AddSuperUser)

	if DEBUG_MODE == "true" {
		if err := router.Run(":8080"); err != nil {
			log.Fatal("Server Run Failed.: ", err)
		}
	} else {
		fmt.Println(os.Getenv("PORT"))
		if err := router.Run(":"+os.Getenv("PORT")); err != nil {
			log.Fatal("Server Run Failed.: ", err)
		}
	}
}

// SuperUserTableとDB接続する
func openSuperUser() (*gorm.DB, error) {
	DBMS :="mysql"
	CONNECT := os.Getenv("CONNECT")

	globaldata.SuperUserDb, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		fmt.Println(err);
		return nil, err
	}

	// DBエンジンを「InnoDB」に設定
	globaldata.SuperUserDb.Set("gorm:table_options", "ENGINE=InnoDB")

	// 詳細なログを表示
	globaldata.SuperUserDb.LogMode(true)

	// 登録するテーブル名を単数形にする（デフォルトは複数形）
	globaldata.SuperUserDb.SingularTable(true)

	// マイグレーション（テーブルが無い時は自動生成）
	globaldata.SuperUserDb.AutoMigrate(&superuser.SuperUser{})

	fmt.Println("db connected: ", &globaldata.SuperUserDb)
	return globaldata.SuperUserDb, nil
}
