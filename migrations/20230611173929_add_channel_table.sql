-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS channels (
    id      BIGINT          PRIMARY KEY,
    name    VARCHAR(64)     NOT NULL,
    created_at TIMESTAMP    DEFAULT NOW()
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
