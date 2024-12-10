-- +goose Up
CREATE TABLE todos
(
    id       serial,
    author   varchar NOT NULL,
    text     varchar NOT NULL,
    complete boolean NOT NULL,

    PRIMARY KEY (id)
);

-- +goose Down
DROP TABLE todos;