CREATE TABLE IF NOT EXISTS users
(
    id        SERIAL PRIMARY KEY,
    chat_id   BIGINT       NOT NULL,
    type      VARCHAR(255) NOT NULL,
    value     VARCHAR(255) NOT NULL
);
