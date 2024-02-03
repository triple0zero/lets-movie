-- +goose Up
-- +goose StatementBegin
CREATE TABLE movies (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    url VARCHAR(255) NOT NULL,
    description TEXT,
    kprating  FLOAT,
    imdbrating FLOAT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Перенести в отдельную таблицу, которая относится именно с логике формирования списка просомра, рейтинга и истории
-- CREATE TABLE movies (
--     id SERIAL PRIMARY KEY,
--     chat_id INT NOT NULL,
--     name VARCHAR(255) NOT NULL,
--     url VARCHAR(255) NOT NULL,
--     rating INT NOT NULL,
--     created_at TIMESTAMP NOT NULL DEFAULT NOW(),
--     updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
--     wivedAt TIMESTAMP DEFAULT NULL
-- );

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS movies;
-- +goose StatementEnd
