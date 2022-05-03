package db

import (
	"fmt"

	// Go言語のORM
	"github.com/jinzhu/gorm"

	// AuthModel(スーパーユーザーデータベースのテーブルの行に対応
	superuser "github.com/Numacchi4635/keijiban/models/superuser"
)

// DB接続する
func openSuperUser() *gorm.DB {
	DBMS :="mysql"
//	CONNECT := os.Getenv("CONNECT")
// fmt.Println(DBMS);
// fmt.Println(CONNECT);
	USER := "root"
	PASS := "Dms1234a"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "superuser"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME

	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		fmt.Println(err);
		panic(err.Error())
	}

	// DBエンジンを「InnoDB」に設定
	db.Set("gorm:table_options", "ENGINE=InnoDB")

	// 詳細なログを表示
	db.LogMode(true)

	// 登録するテーブル名を単数形にする（デフォルトは複数形）
	db.SingularTable(true)

	// マイグレーション（テーブルが無い時は自動生成）
	db.AutoMigrate(&superuser.SuperUser{})

	fmt.Println("db connected: ", &db)
	return db
}

// FindSuperUser は スーパーユーザーテーブルのレコードを全件(登録上限は1件)取得する
func FindSuperUser() []superuser.SuperUser {
	superusers := []superuser.SuperUser{}

	db := openSuperUser()
	// select
	db.Order("ID asc").Find(&superusers)

	// defer 関数がreturnする時に実行される
	defer db.Close()

	return superusers
}

// InsertSuperUser はスーパーユーザーテーブルにレコードを追加する
// (1件しか登録できないテーブルなので、実質新規登録になる
// 既にテーブルにデータが存在している場合はUpDateになる
func InsertSuperUser(registerSuperUser *superuser.SuperUser) {
fmt.Println("InsertSuperUser Start");
	db := openSuperUser()
	// insert
	db.Create(&registerSuperUser)
	defer db.Close()
fmt.Println("InsertSuperUser End");
}
