package stock

import (
	"github.com/google/uuid"
)

type (
	StockPrices struct {
		ID       uuid.UUID `json:"id"`
		StocksID uuid.UUID `json:"stocks_id"`
		Open     float64   `json:"open"`
		High     float64   `json:"high"`
		Low      float64   `json:"low"`
		Close    float64   `json:"close"`
		Volume   float64   `json:"volume"`
	}
)

func NewStockPrices(id uuid.UUID) *StockPrices {
	return &StockPrices{ID: id}
}

func (s *StockPrices) WithStocksID(id uuid.UUID) {
	s.StocksID = id
}

func (s *StockPrices) WithPrices(open, high, low, close, vlm float64) {
	s.Open = open
	s.High = high
	s.Low = low
	s.Close = close
	s.Volume = vlm
}
