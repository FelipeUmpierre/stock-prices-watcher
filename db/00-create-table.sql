CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE events (
    id SERIAL PRIMARY KEY,
    aggregate_id UUID NOT NULL,
    aggregate_version INTEGER NOT NULL DEFAULT 1,
    handle VARCHAR NOT NULL,
    metadata JSONB,
    payload JSONB,

    UNIQUE(aggregate_id, aggregate_version)
);

CREATE TABLE symbols (
    id SERIAL PRIMARY KEY,
    aggregate_id UUID DEFAULT uuid_generate_v4(),
    symbol VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    UNIQUE(symbol)
);