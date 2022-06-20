package calc

func AvgPrice(TotalPrice, TotalVolume float64) float64 {
	return TotalPrice / TotalVolume
}

func AvgVolume(TotalVolume float64, TotalTrades int) float64 {
	return TotalVolume / float64(TotalTrades)
}

func PercentBuys(TotalBuys, TotalTrades int) float64 {
	return float64(TotalBuys) / float64(TotalTrades) * 100
}

func VolumeWeightedAvgPrice(AvgPrice, TotalVolume float64) float64 {
	return (AvgPrice * TotalVolume) / TotalVolume
}
