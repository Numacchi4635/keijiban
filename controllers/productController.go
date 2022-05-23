package controllers

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
	Products	[]entity.Product
	PublicMode	string
}

// FetchAllProducts は 全ての掲示板情報を取得する
func FetchAllProducts(c *gin.Context) {

	inputPassword := c.Query("productPassword")

	// 管理者パスワード照合
	rtn := SuperUserPasswordCollationDB(inputPassword)
	if ( rtn == http.StatusInternalServerError ){
		fmt.Println("Internal Server Error")
		c.JSON(http.StatusInternalServerError, nil)
	} else if (rtn == http.StatusUnauthorized) {
		fmt.Println("管理者パスワード不一致")
		c.JSON(http.StatusUnauthorized, nil)
	}

	ResultProducts := db.FindAllProducts()
	if ResultProducts == nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	// 環境変数publicMode取得
	PublicMode :=  os.Getenv("PUBLIC_MODE")

	ResultResponse := resultResponse {
		Products:	ResultProducts,
		PublicMode:	PublicMode,
	}

	// URLへのアクセスに対してJSONを返す
	c.JSON(http.StatusOK, ResultResponse)
}

// FindProduct は 指定したIDの掲示板情報を取得する
func FindProduct(c *gin.Context) {
	ProductIDStr := c.Query("productID")

	ProductID, err := strconv.Atoi(ProductIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	ProductPassword := c.Query("productPassword")
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	ResultProduct, err := db.FindProduct(ProductID)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	// 入力パスワードが違う場合
	if ResultProduct.Password != ProductPassword {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	// URLへのアクセスに対してJSONを返す
	c.JSON(http.StatusOK, ResultProduct)
}

// AddProduct は 掲示板情報をDBへ登録する
func AddProduct(c *gin.Context) {
	ProductName := c.PostForm("productName")
	ProductMessage := c.PostForm("productMessage")
	InputPassword := c.PostForm("superUserPassword")

	// 管理者パスワード照合
	rtn := SuperUserPasswordCollationDB(InputPassword)
	if ( rtn == http.StatusInternalServerError ){
		fmt.Println("Internal Server Error")
		c.JSON(http.StatusInternalServerError, nil)
	} else if (rtn == http.StatusUnauthorized) {
		fmt.Println("管理者パスワード不一致")
		c.JSON(http.StatusUnauthorized, nil)
	}

	fmt.Println("AddProduct 管理者パスワード一致")

	// 一般ユーザー側のパスワード生成
	ProductPassword, _ := MakeRandomStr(128)

	var Product = entity.Product{
		Name:    ProductName,
		Message: ProductMessage,
		Password:ProductPassword,
	}

	db.InsertProduct(&Product)
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
	ProductIDStr := c.PostForm("productID")
	inputPassword := c.PostForm("superUserPassword")

	ProductID, err := strconv.Atoi(ProductIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	// 管理者パスワード照合
	rtn := SuperUserPasswordCollationDB(inputPassword)
	if ( rtn == http.StatusInternalServerError ){
		fmt.Println("Internal Server Error")
		c.JSON(http.StatusInternalServerError, nil)
	} else if (rtn == http.StatusUnauthorized) {
		fmt.Println("管理者パスワード不一致")
		c.JSON(http.StatusUnauthorized, nil)
	}


	err = db.DeleteProduct(ProductID)
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
	InputPassword := c.PostForm("ProductPassword")
	ResultFindProduct, _ := db.FindProduct(InputID)

	if InputPassword == ResultFindProduct.Password {
		c.JSON(http.StatusOK, ResultFindProduct)
	} else {
		c.JSON(http.StatusUnauthorized, nil)
	}
}

// 環境変数を返す
func ResponseServerEnv(c *gin.Context) {
	// 環境変数PUBLIC_MODE取得
	PublicMode :=  os.Getenv("PUBLIC_MODE")

	ResultResponse := resultResponse {
		PublicMode:	PublicMode,
	}

	// URLへのアクセスに対してJSONを返す
	c.JSON(http.StatusOK, ResultResponse)
}
