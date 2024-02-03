-- +goose Up
-- +goose StatementBegin
CREATE TABLE members (
    id SERIAL PRIMARY KEY,
    tg_id VARCHAR(255) NOT NULL,
    chat_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS members;
-- +goose StatementEnd
