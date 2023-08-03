CREATE TABLE note_categories
(
    id        SERIAL PRIMARY KEY,
    chat_id   INTEGER      NOT NULL,
    name      VARCHAR(255) NOT NULL,
    parent_id INTEGER,
    status    INTEGER      NOT NULL
);

CREATE INDEX notes_chat_id_idx ON note_categories (chat_id);
CREATE INDEX notes_parent_id_idx ON note_categories (parent_id);