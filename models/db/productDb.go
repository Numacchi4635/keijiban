package db

import (
	// フォーマットI/O
	"fmt"

	"os"

	// Go言語のORM
	"github.com/jinzhu/gorm"

	// エンティティ(データベースのテーブルの行に対応)
	entity "github.com/Numacchi4635/keijiban/models/entity"
)

// DB接続する
func open() (*gorm.DB, error) {

	// 環境変数からDEBUG_MODE取得

	DBMS := "mysql"
	CONNECT := os.Getenv("CONNECT")

	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}

	// DBエンジンを「InnoDB」に設定
	db.Set("gorm:table_options", "ENGINE=InnoDB")

	// 詳細なログを表示
	db.LogMode(true)

	// 登録するテーブル名を単数形にする（デフォルトは複数形）
	db.SingularTable(true)

	// マイグレーション（テーブルが無い時は自動生成）
	db.AutoMigrate(&entity.Product{})

	fmt.Println("db connected: ", &db)
	return db, nil
}

// FindAllProducts は 掲示板テーブルのレコードを全件取得する
func FindAllProducts() []entity.Product {
	products := []entity.Product{}

	db, err := open()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	// select
	db.Order("ID asc").Find(&products)

	// defer 関数がreturnする時に実行される
	defer db.Close()

	return products
}

// FindProduct は 掲示板テーブルのレコードを1件取得する
func FindProduct(password string) (entity.Product, error) {
	product := entity.Product{}

	db, err := open()
	if err != nil {
		fmt.Println(err)
		return product, err
	}
	// select
	db.Find(&product, "password=?", password)
	defer db.Close()

	return product, nil
}

// InsertProduct は 掲示板テーブルにレコードを追加する
func InsertProduct(registerProduct *entity.Product) error {
	fmt.Println("InsertProduct Start")
	db, err := open()
	if err != nil {
		fmt.Println(err)
		return err
	}

	// insert
	db.Create(&registerProduct)
	defer db.Close()
	return nil
}

// DeleteProduct は 掲示板テーブルの指定したレコードを削除する
func DeleteProduct(productID int) error {
	product := []entity.Product{}

	db, err := open()
	if err != nil {
		fmt.Println(err)
		return err
	}

	// delete
	err = db.Delete(&product, productID).Error
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer db.Close()
	return nil
}
