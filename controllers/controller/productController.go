package controller

import (
	// 標準出力（エラーログ）
	"fmt"

	// 環境変数取得用
	"os"

	// 文字列と基本データの変換
	strconv "strconv"

	// 乱数
	"crypto/rand"

	// http
	"net/http"

	// Gin
	"github.com/gin-gonic/gin"

	// エンティティ(データベースのテーブルの行に対応)
	entity "github.com/Numacchi4635/keijiban/models/entity"

	// DBアクセス用モジュール
	db "github.com/Numacchi4635/keijiban/models/db"
)

// 環境変数publicModeを含む返信用struct
type resultResponse struct {
	products	[]entity.Product
	publicMode	string
}

// FetchAllProducts は 全ての掲示板情報を取得する
func FetchAllProducts(c *gin.Context) {

	resultProducts := db.FindAllProducts()
	if resultProducts == nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	// 環境変数publicMode取得
	publicMode :=  os.Getenv("publicMode")

	resultresponse := resultResponse {
		products:	resultProducts,
		publicMode:	publicMode,
	}


	// URLへのアクセスに対してJSONを返す
	c.JSON(http.StatusOK, resultresponse)
}

// FindProduct は 指定したIDの掲示板情報を取得する
func FindProduct(c *gin.Context) {
	productIDStr := c.Query("productID")

	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	productPassword := c.Query("productPassword")
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	resultProduct, err := db.FindProduct(productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	// 入力パスワードが違う場合
	if resultProduct.Password != productPassword {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	// URLへのアクセスに対してJSONを返す
	c.JSON(http.StatusOK, resultProduct)
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
		fmt.Println(err)
		return "", err
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

	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	err = db.DeleteProduct(productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	// 削除成功時はStatusNoContent(204)を返す
	c.JSON(http.StatusNoContent, nil)
}

// パスワード照合
func UserPasswordCollation(c *gin.Context) {
	InputIDStr := c.PostForm("productID")
	InputID, err := strconv.Atoi(InputIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	InputPassword := c.PostForm("productPassword")
	resultFindProduct, _ := db.FindProduct(InputID)

	if InputPassword == resultFindProduct.Password {
		c.JSON(http.StatusOK, resultFindProduct)
	} else {
		c.JSON(http.StatusUnauthorized, nil)
	}
}

// 環境変数を返す
func ResponseServerEnv(c *gin.Context) {
	// 環境変数PUBLIC_MODE取得
	publicMode :=  os.Getenv("publicMode")

	resultresponse := resultResponse {
		publicMode:	publicMode,
	}

	// URLへのアクセスに対してJSONを返す
	c.JSON(http.StatusOK, resultresponse)
}
