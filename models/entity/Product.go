package entity

// Product はテーブルのモデル

type Product struct {
	ID       int    `gorm:"primary_key;not null"		json:"id"`
	Name     string `gorm:"type:varchar(200);not null	json:"name"`
	Message  string `gorm:"type:varchar(400);not null	json:"message"`
	Password string `gorm:"type:varchar(400):not null	json:"password"`
}
