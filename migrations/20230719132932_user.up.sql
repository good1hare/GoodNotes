CREATE TABLE IF NOT EXISTS users
(
    id        SERIAL PRIMARY KEY,
    user_name VARCHAR(255) NOT NULL,
    chat_id   BIGINT       NOT NULL
);

CREATE UNIQUE INDEX idx_users_chat_id ON users(chat_id);
