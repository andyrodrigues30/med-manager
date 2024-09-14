-- +goose Up
-- +goose StatementBegin

CREATE table users (
    id int NOT NULL, 
    name text NOT NULL,
    username text NOT NULL,
    password text NOT NULL,
    joined timestamp NOT NULL,
    is_admin boolean,
    PRIMARY KEY (id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP table users;

-- +goose StatementEnd
