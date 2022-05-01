package controller

import (
	"fmt"
	// 文字列と基本データの変換
	strconv "strconv"

	// 乱数
	"crypto/rand"

	// エラー
	"errors"

	// Gin
	"github.com/gin-gonic/gin"

	// エンティティ(データベースのテーブルの行に対応)
	entity "github.com/Numacchi4635/keijiban/models/entity"

	// 認証モデル
//	authModel "github.com/Numacchi4635/keijiban/models/authModel"

	// DBアクセス用モジュール
	db "github.com/Numacchi4635/keijiban/models/db"
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
fmt.Println("AddProduct Start");
	productName := c.PostForm("productName")
	productMessage := c.PostForm("productMessage")
	productPassword, _ := MakeRandomStr(128)

	var product = entity.Product{
		Name:    productName,
		Message: productMessage,
		Password:productPassword,
	}

	db.InsertProduct(&product)
fmt.Println("AddProduct End");
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

// SuperUserLogin は スーパーユーザーのログイン認証を行う
//func SuperUserLogin(c *gin.Context) {
//	var request authModel.EmailLoginRequest
//	err := c.BindJSON(&request)
//	if err != nil {
//		c.Status(http.StatusBadRequest)
//	} else {
//		// メールアドレスでDBからユーザ取得
//		superuser, err := authRepository.GetUserByEmail(request.email)
//		// ハッシュ値でのパスワード比較
//		err = bcrypt.CompareHashAndPassword([]byte(superuser.Password), []byte(request.password))
//		if err != nil {
//			c.Status(http.StatusBadRequest)
//		} else {
//			session := sessions.Default(c)
//			// セッションに格納する為にユーザ情報をJson化
//			loginUser, err := json.Marshal(auth)
//			if err == nil {
//				session.Set("loginUser", string(loginUser))
//				session.Save()
//				c.Status(http.StatusOK)
//			} else {
//				c.Status(http.StatusInternalServerError)
//			}
//		}
//	}
//}
