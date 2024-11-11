-- +goose Up
create table users(
    id uuid primary key,
    created_at timestamp not null,
    updated_at timestamp not null,
    name text not null,
    email text not null unique,
    password text not null,
    location text not null,
    age integer not null,
    username text not null unique,
    bio text not null,
    role user_type not null
);

-- +goose Down
drop table users;