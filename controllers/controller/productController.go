package controller

import (
	// 文字列と基本データの変換
	strconv "strconv"

	// 乱数
	"crypto/rand"

	// エラー
	"errors"

	// Debug
	"fmt"

	// Gin
	"github.com/gin-gonic/gin"

	// エンティティ(データベースのテーブルの行に対応)
	entity "../../models/entity"

	// DBアクセス用モジュール
	db "../../models/db"
)

// FetchAllProducts は 全ての掲示板情報を取得する
func FetchAllProducts(c *gin.Context) {
	resultProducts := db.FindAllProducts()

	// URLへのアクセスに対してJSONを返す
	c.JSON(200, resultProducts)
}

// FindProduct は 指定したIDの掲示板情報を取得する
func FindProduct(c *gin.Context) {
	productIDStr := c.Query("productID")

	productID, _ := strconv.Atoi(productIDStr)

	resultProduct, _ := db.FindProduct(productID)

	// URLへのアクセスに対してJSONを返す
	c.JSON(200, resultProduct)
}

// AddProduct は 掲示板情報をDBへ登録する
func AddProduct(c *gin.Context) {
	productName := c.PostForm("productName")
	productMessage := c.PostForm("productMessage")
	productPassword, _ := MakeRandomStr(128)

	var product = entity.Product{
		Name:    productName,
		Message: productMessage,
		Password:productPassword,
	}

	db.InsertProduct(&product)
}

// 推測不可能なパスワード生成
func MakeRandomStr(digit uint32) (string, error){
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// 乱数を生成
	b := make([]byte, digit)
	if _, err := rand.Read(b); err != nil {
		return "", errors.New("unexpected error...")
	}

	// letters からランダムに取り出して文字列を生成
	var result string
	for _, v := range b {
		// index が letters の長さに収まるように調整
		result += string(letters[int(v)%len(letters)])
	}
	return result, nil
}

// DeleteProduct は 掲示板情報をDBから削除する
func DeleteProduct(c *gin.Context) {
fmt.Println("DeleteProduct Start");
	productIDStr := c.PostForm("productID")
fmt.Println(productIDStr);

	productID, _ := strconv.Atoi(productIDStr)

fmt.Println(productID);

	db.DeleteProduct(productID)
}
