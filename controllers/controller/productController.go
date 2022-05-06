package controller

import (
	"C"
	"os"
	// Debug用
	"fmt"

	// 文字列と基本データの変換
	strconv "strconv"

	// 乱数
	"crypto/rand"

	// エラー
	"errors"

	// JSON
//	"encoding/json"

	// Gin
	"github.com/gin-gonic/gin"

	// エンティティ(データベースのテーブルの行に対応)
	entity "github.com/Numacchi4635/keijiban/models/entity"

	// DBアクセス用モジュール
	db "github.com/Numacchi4635/keijiban/models/db"
)

// 環境変数PUBLIC_MODEを含む返信用struct
type resultResponse struct {
	ID		int	`gorm:"primary_key;not null"		json:"id"`
	Name		string	`gorm:"type:varchar(200);not null	json:"name"`
	Message		string	`gorm:"type:varchar(400);not null	json:"message"`
	Password	string	`gorm:"type:varchar(400):not null	json:"password"`
	PublicMode	string
}

// FetchAllProducts は 全ての掲示板情報を取得する
func FetchAllProducts(c *gin.Context) {

	resultProducts := db.FindAllProducts()
	var result_response[len(resultProducts)] resultResponse

fmt.Println(len(resultProducts))

	// 環境変数PUBLIC_MODE取得
	PUBLIC_MODE :=  os.Getenv("PUBLIC_MODE")
fmt.Println(PUBLIC_MODE)
	for i := 0 ; i < len(resultProducts) ; i++ {
		C.memcpy(result_response[i], resultProducts[i], len(resultProducts[i]))
		result[i].puclic_mode = PUBLIC_MODE
	}

//	resultResponse	*result_response;


//	// 環境変数PUBLIC_MODEをJSONに変換
//	json_public_mode, err := json.Marshal(PUBLIC_MODE)
//	if err != nil {
//		fmt.Println(err)
//		panic(err.Error())
//	}
//fmt.Println("JSON Result Products = ", resultProducts)
//fmt.Println("JSON PUBLIC MODE = ", json_public_mode)

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
	productIDStr := c.PostForm("productID")

	productID, _ := strconv.Atoi(productIDStr)

	db.DeleteProduct(productID)
}

// パスワード照合
func UserPasswordCollation(c *gin.Context) {
	InputIDStr := c.PostForm("productID")
	InputID, _ := strconv.Atoi(InputIDStr)
	InputPassword := c.PostForm("productPassword")
	resultFindProduct, _ := db.FindProduct(InputID)

	if InputPassword == resultFindProduct[0].Password {
		c.JSON(200, resultFindProduct)
	} else {
		c.JSON(201, resultFindProduct)
	}
}
