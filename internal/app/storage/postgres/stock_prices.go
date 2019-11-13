package postgres

import "database/sql"

type (
	StockPriceWriter struct {
		db *sql.DB
	}

	StockPricesReader struct {
		db *sql.DB
	}
)

func NewStockPricesWriter(db *sql.DB) *StockPriceWriter {
	return &StockPriceWriter{db: db}
}

func (s *StockPriceWriter) Save() error {
	return nil
}

func NewStockPricesReader(db *sql.DB) *StockPricesReader {
	return &StockPricesReader{db: db}
}

func (s *StockPricesReader) FindByStockID(id string) error {
	return nil
}
