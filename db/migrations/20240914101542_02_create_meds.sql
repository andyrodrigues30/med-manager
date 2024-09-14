-- +goose Up
-- +goose StatementBegin

CREATE table meds (
    id int NOT NULL, 
    name text NOT NULL,
    PRIMARY KEY (id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP table meds;
-- +goose StatementEnd
