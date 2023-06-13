-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id          BIGINT          PRIMARY KEY,
    name        VARCHAR(255)    NOT NULL,
    password    VARCHAR(255)    NOT NULL,
    email       VARCHAR(255)    NOT NULL,
    created_at  TIMESTAMP       DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
