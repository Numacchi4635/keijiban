package db

import (
	// 環境変数取得用
	"os"

	// 標準出力用（エラーログ）
	"fmt"

	// Go言語のORM
	"github.com/jinzhu/gorm"

	// AuthModel(スーパーユーザーデータベースのテーブルの行に対応
	superuser "github.com/Numacchi4635/keijiban/models/superuser"
)

// DB接続する
func openSuperUser() (*gorm.DB, error) {
	DBMS :="mysql"
	CONNECT := os.Getenv("CONNECT")

	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		fmt.Println(err);
		return nil, err
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
	return db, nil
}

// FindSuperUser は スーパーユーザーテーブルのレコードを全件(登録上限は1件)取得する
func FindSuperUser() []superuser.SuperUser {
	superusers := []superuser.SuperUser{}

	db, err := openSuperUser()
	if err != nil {
		fmt.Println(err)
		return nil;
	}

	// select
	db.Order("ID asc").Find(&superusers)

	// defer 関数がreturnする時に実行される
	defer db.Close()

	return superusers
}

// InsertSuperUser はスーパーユーザーテーブルにレコードを追加する
// (1件しか登録できないテーブルなので、実質新規登録になる
// 既にテーブルにデータが存在している場合はUpDateになる
func InsertSuperUser(registerSuperUser *superuser.SuperUser) error {

	// スーパーユーザーテーブルの検索
	existSuperUser := FindSuperUser()

	if len(existSuperUser) >= 1 {
		// テーブルのレコードが存在する場合は消去する
		for i := 0; i < len(existSuperUser); i++{
			DeleteSuperUser(existSuperUser[i].ID)
		}
	}
	db, err := openSuperUser()
	if db == nil {
		fmt.Println(err)
		return err
	}

	// insert
	db.Create(&registerSuperUser)
	defer db.Close()
	return nil
}

// DeleteSuperUser は スーパーユーザーテーブルの指定したレコードを削除する
func DeleteSuperUser(ID int) error {
	SuperUser := []superuser.SuperUser{}

	db, err := openSuperUser()
	if err != nil {
		fmt.Println(err)
		return err
	}

	// delete
	db.Delete(&SuperUser, ID)
	defer db.Close()
	return nil
}
