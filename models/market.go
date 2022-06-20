package models

type Market struct {
	ID                     int     `json:"id"`
	AvgPrice               float64 `json:"avg_price"`
	AvgVolume              float64 `json:"avg_volume"`
	TotalBuys              int     `json:"-"`
	TotalPrice             float64 `json:"-"`
	TotalTrades            int     `json:"-"`
	TotalVolume            float64 `json:"total_volume"`
	PercentBuys            float64 `json:"percent_buys"`
	VolumeWeightedAvgPrice float64 `json:"volume_weighted_avg_price"`
}
