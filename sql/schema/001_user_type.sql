-- +goose Up
create type user_type as enum ('admin', 'moderator', 'user');

-- +goose Down
drop type user_type;