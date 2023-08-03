CREATE TABLE notes
(
    id          SERIAL PRIMARY KEY,
    chat_id     INTEGER      NOT NULL,
    name        VARCHAR(255) NOT NULL,
    description TEXT,
    category_id INTEGER      NOT NULL,
    status      INTEGER      NOT NULL
);

CREATE INDEX notes_chat_id_idx ON notes (chat_id);
CREATE INDEX notes_category_id_idx ON notes (category_id);