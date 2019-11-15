package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"
)

type (
	Stocks struct {
		Series map[string]Stock `json:"Time Series (1min)"`
	}

	Stock struct {
		Open   string `json:"1. open"`
		High   string `json:"2. high"`
		Low    string `json:"3. low"`
		Close  string `json:"4. close"`
		Volume string `json:"5. volume"`
	}
)

func main() {
	t := time.Now()

	resp, err := http.Get("https://www.alphavantage.co/query?function=TIME_SERIES_INTRADAY&symbol=MSFT&interval=1min&apikey=demo")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var stocks Stocks
	if err := json.NewDecoder(resp.Body).Decode(&stocks); err != nil {
		panic(err)
	}

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

	log.Printf(`
	%+v
	`, stocks)

	log.Fatalln(time.Since(t))
}
