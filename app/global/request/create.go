package request

type CreateStockOption struct {
	Symbol string `json:"symbol" validate:"required"`
	Name   string `json:"name" validate:"required"`
}

type CreateDailyRecordOption struct {
	Date       string  `json:"date" validate:"required"`
	ClosePrice float32 `json:"close_price" validate:"required"`
	HighPrice  float32 `json:"high_price" validate:"required"`
	LowPrice   float32 `json:"low_price" validate:"required"`
	StockID    int     `json:"stock_id" validate:"required"`
}

type CreateUserOption struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Age      int    `json:"age" validate:"required"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
