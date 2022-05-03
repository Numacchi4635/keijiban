package controller

import ( 
	// debug用
	"fmt"

	// エラー
//	"errors"

	// Gin
	"github.com/gin-gonic/gin"

	// スーパーユーザー（データベースのテーブルの行に対応）
	superuser "github.com/Numacchi4635/keijiban/models/superuser"
	// DBアクセス用モジュール
	db "github.com/Numacchi4635/keijiban/models/db"
)

func FindSuperUser(c *gin.Context) {
fmt.Println("FindSuperUser Start")
	resultSuperUser := db.FindSuperUser();

	// URLへのアクセスに対してJSONを返す
	c.JSON(200, resultSuperUser);
}

func AddSuperUser(c *gin.Context) {
fmt.Println("AddSuperUser Start");
	superUserID := c.PostForm("UserID")
	superUserPassword := c.PostForm("Password")
	var SuperUser = superuser.SuperUser{
		ID	:1,
		UserID	:superUserID,
		Password:superUserPassword,
	}
	db.InsertSuperUser(&SuperUser)
fmt.Println("AddSuperSuser End")
}