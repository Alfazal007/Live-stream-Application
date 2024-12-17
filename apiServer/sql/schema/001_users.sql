-- +goose Up
create table users (
    id uuid primary key,
    username text not null unique,
    password text not null
);

-- +goose Down
drop table users;
