package structs

type StockObj struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type UserObj struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
