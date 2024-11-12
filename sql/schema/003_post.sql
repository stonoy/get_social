-- +goose Up
create table posts(
    id uuid primary key,
    created_at timestamp not null,
    updated_at timestamp not null,
    content text not null,
    author uuid not null
    references users(id)
    on delete cascade
);

-- +goose Down
drop table posts;