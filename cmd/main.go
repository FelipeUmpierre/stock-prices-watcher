package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"
)

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

func (s *Stocks) UnmarshalJSON(b []byte) error {
	dec := json.NewDecoder(bytes.NewReader(b))
	for i := 0; i < 17; i++ {
		if _, err := dec.Token(); err != nil {
			log.Println("failed to get token", err)
		}
	}

	var ss Stocks
	for dec.More() {
		var m map[string]Stock
		if err := dec.Decode(&m); err != nil {
			log.Println("failed to decode", err)
		}

		sm := make(map[time.Time]Stock, len(m))
		for t, st := range m {
			f, _ := time.Parse("2006-01-02 15:04:05", t)
			sm[f] = st
		}
	}

	ss = sm

	return nil
}

func main() {
	resp, err := http.Get("https://www.alphavantage.co/query?function=TIME_SERIES_INTRADAY&symbol=MSFT&interval=1min&apikey=demo")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var stocks Stocks
	if err := json.NewDecoder(resp.Body).Decode(&stocks); err != nil {
		panic(err)
	}

	// stocksSlice := make([]string, 0, len(stocks.Series))
	// for t := range stocks.Series {
	// 	stocksSlice = append(stocksSlice, t)
	// }
	// sort.Strings(stocksSlice)

	ctx := context.Background()
	eg, _ := errgroup.WithContext(ctx)

	// id := uuid.New()

	for t, stock := range stocks.Series {
		t := t
		stock := stock

		eg.Go(func() error {
			log.Printf("%s: %+v\n", t, stock.Open)
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		panic(err)
	}

	log.Fatalf(`
	%+v
	`, stocks)
}
