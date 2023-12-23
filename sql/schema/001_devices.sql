-- +goose Up
CREATE TABLE devices (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    codetag VARCHAR(12) NOT NULL
);


-- +goose Down

DROP TABLE devices;