-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS group_messages (
    channel_id  BIGINT      NOT NULL,
    message_id  BIGINT      NOT NULL,
    user_id     BIGINT      NOT NULL,
    content     TEXT        NOT NULL,
    created_at  TIMESTAMP   DEFAULT NOW(),
    PRIMARY KEY (channel_id, message_id) 
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
