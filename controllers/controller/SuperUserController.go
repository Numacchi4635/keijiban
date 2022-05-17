package controller

import ( 
	// Debug用
	"fmt"

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
fmt.Println("SuperUserPasswordCollation")
	InputPassword := c.Query("productPassword")
fmt.Println(InputPassword)

	resultSuperUser := db.FindSuperUser()
	if resultSuperUser == nil {
fmt.Println(resultSuperUser)
		c.JSON(http.StatusInternalServerError, nil)
	}

	if InputPassword == resultSuperUser[0].Password {
fmt.Println("パスワード一致")
		c.JSON(http.StatusOK, resultSuperUser)
	} else {
fmt.Println("パスワード不一致")
		c.JSON(http.StatusUnauthorized, nil)
	}
}
