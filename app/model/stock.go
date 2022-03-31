package model

import (
	"time"
)

type Stock struct {
	ID          int       `json:"id" gorm:"column:id;type:int(11) unsigned auto_increment;NOT NULL;primary_key"`
	Symbol      string    `json:"symbol" gorm:"column:symbol;type:varchar(255);not null"`
	Name        string    `json:"name" gorm:"column:name;type:varchar(255);not null"`
	LatestFresh time.Time `json:"latest_fresh" gorm:"column:latest_fresh;type:TIMESTAMP; default:CURRENT_TIMESTAMP"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at;type:TIMESTAMP; default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at;type:TIMESTAMP; not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`

	DailyRecord []DailyRecord `gorm:"foreignKey:StockID"`
}

//type User struct {
//	gorm.Model
//	CreditCards []CreditCard `gorm:"foreignKey:UserRefer"`
//}
//
//type CreditCard struct {
//	gorm.Model
//	Number    string
//	UserRefer uint
//}
