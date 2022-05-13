package controller

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

func FindSuperUser(c *gin.Context) {
	resultSuperUser := db.FindSuperUser();
	if resultSuperUser == nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	// URLへのアクセスに対してJSONを返す
	c.JSON(http.StatusOK, resultSuperUser);
}

func AddSuperUser(c *gin.Context) {
	superUserID := c.PostForm("UserID")
	superUserPassword := c.PostForm("Password")
	var SuperUser = superuser.SuperUser{
		ID	:SuperUserIDNo,
		UserID	:superUserID,
		Password:superUserPassword,
	}
	db.InsertSuperUser(&SuperUser)
}

func SuperUserPasswordCollation(c *gin.Context) {
	InputPassword := c.PostForm("superUserPassword")

	resultSuperUser := db.FindSuperUser()
	if resultSuperUser == nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	if InputPassword == resultSuperUser[0].Password {
		c.JSON(http.StatusOK, resultSuperUser)
	} else {
		c.JSON(http.StatusUnauthorized, resultSuperUser)
	}
}
