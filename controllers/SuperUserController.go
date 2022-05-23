package controllers

import (
	// http
	"net/http"

	// Gin
	"github.com/gin-gonic/gin"

	// スーパーユーザー（データベースのテーブルの行に対応）
	superuser "github.com/Numacchi4635/keijiban/models/superuser"

	// DBアクセス用モジュール
	db "github.com/Numacchi4635/keijiban/models/db"
)

// スーパーユーザーのID番号は1固定
const SuperUserIDNo = 1

func AddSuperUser(c *gin.Context) {
	superUserID := c.PostForm("superUserID")
	superUserPassword := c.PostForm("superUserPassword")
	var SuperUser = superuser.SuperUser{
		ID:       SuperUserIDNo,
		UserID:   superUserID,
		Password: superUserPassword,
	}
	db.InsertSuperUser(&SuperUser)
}

func SuperUserPasswordCollation(c *gin.Context) {
	inputPassword := c.Query("productPassword")
	resultSuperUser := db.FindSuperUser()
	if resultSuperUser == nil {
		c.JSON(http.StatusInternalServerError, nil)
	}

	if inputPassword == resultSuperUser[0].Password {
		c.JSON(http.StatusOK, resultSuperUser)
	} else {
		c.JSON(http.StatusUnauthorized, nil)
	}
}

func SuperUserPasswordCollationDB(inputPassword string) int {
	resultSuperUser := db.FindSuperUser()
	if resultSuperUser == nil {
		return http.StatusInternalServerError
	}

	if inputPassword == resultSuperUser[0].Password {
		return http.StatusOK
	} else {
		return http.StatusUnauthorized
	}
}
