package model

import (
	"time"
)

type DailyRecord struct {
	ID         int       `json:"id" gorm:"column:id;type:int(11) unsigned auto_increment;NOT NULL;primary_key"`
	Date       time.Time `json:"date" gorm:"column:date;type:datetime; not null"`
	ClosePrice float32   `json:"close_price" gorm:"column:close_price;type:decimal(10,2); not null"`
	HighPrice  float32   `json:"high_price" gorm:"column:high_price;type:decimal(10,2); not null"`
	LowPrice   float32   `json:"low_price" gorm:"column:low_price;type:decimal(10,2); not null"`
	StockID    uint      `json:"stock_id" gorm:"column:stock_id;not null"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:created_at;type:TIMESTAMP; default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"column:updated_at;type:TIMESTAMP; not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
