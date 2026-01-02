-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS products (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    price_in_₹ INTEGER NOT NULL CHECK (price_in_₹ >= 0),
    quantity INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
