-- +goose Up
-- +goose StatementBegin

CREATE TABLE users
(
    id         bigserial           not null,
    first_name varchar(255),
    last_name  varchar(255),
    email      varchar(255) unique NOT NULL,
    password   varchar(255)        NOT NULL,
    age        int8,
    gender     varchar(10),
    city       varchar(255),
    interests  text,
    created_at date                NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
