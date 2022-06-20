package models

type Trade struct {
	ID       int     `json:"id"`
	MarketID int     `json:"market"`
	Price    float64 `json:"price"`
	Volume   float64 `json:"volume"`
	IsBuy    bool    `json:"is_buy"`
}
