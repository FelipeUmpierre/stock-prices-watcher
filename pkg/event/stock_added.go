package event

type StockAdded struct {
	Open   int64 `json:"open"`
	High   int64 `json:"high"`
	Low    int64 `json:"low"`
	Close  int64 `json:"close"`
	Volume int64 `json:"volume"`
}

func NewStockAdded(open, high, low, close, volume int64) *StockAdded {
	return &StockAdded{
		Open:   open,
		High:   high,
		Low:    low,
		Close:  close,
		Volume: volume,
	}
}
