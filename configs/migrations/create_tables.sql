CREATE TABLE stocks (
    id UUID NOT NULL,
    name VARCHAR NOT NULL UNIQUE,
    CONSTRAINT "pk_stocks_id" PRIMARY KEY ("id")
);

CREATE TABLE stock_prices (
    id UUID,
    stocks_id UUID NOT NULL,
    open REAL,
    high REAL,
    low REAL,
    close REAL,
    volume REAL,
    CONSTRAINT "pk_stock_prices_id" PRIMARY KEY ("id"),
    CONSTRAINT "fk_stocks_id" FOREIGN KEY ("stocks_id") REFERENCES stocks ("id")
);
