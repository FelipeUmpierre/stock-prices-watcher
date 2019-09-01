package stock

import "time"

type (
	Stocks struct {
		Series map[time.Time]Stock `json:"Time Series (1min)"`
	}

	Stock struct {
		Open   string `json:"1. open"`
		High   string `json:"2. high"`
		Low    string `json:"3. low"`
		Close  string `json:"4. close"`
		Volume string `json:"5. volume"`
	}
)
