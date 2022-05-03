package superuser

// SuperUser はスーパーユーザーテーブルのモデル

type SuperUser struct {
	ID		int	`gorm:"type:int:not null"		json:"id"`
	UserID		string	`gorm:"type:varchar(200):not null"	json:"userid"`
	Password	string	`gorm:"type:varchar(200);not null"	json:"password"`
}
