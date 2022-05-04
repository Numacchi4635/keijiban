package controller

import ( 
	// debug用
//	"fmt"

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
	resultSuperUser := db.FindSuperUser();

	// URLへのアクセスに対してJSONを返す
	c.JSON(200, resultSuperUser);
}

func AddSuperUser(c *gin.Context) {
	superUserID := c.PostForm("UserID")
	superUserPassword := c.PostForm("Password")
	var SuperUser = superuser.SuperUser{
		ID	:1,
		UserID	:superUserID,
		Password:superUserPassword,
	}
	db.InsertSuperUser(&SuperUser)
}

func SuperUserPasswordCollation(c *gin.Context) {
	InputPassword := c.PostForm("superUserPassword")

	resultSuperUser := db.FindSuperUser()

	if InputPassword == resultSuperUser[0].Password {
		c.JSON(200, resultSuperUser)
	} else {
		c.JSON(201, resultSuperUser)
	}
}
