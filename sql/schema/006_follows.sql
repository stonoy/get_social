-- +goose Up
create table follows(
    id uuid primary key,
    created_at timestamp not null,
    updated_at timestamp not null,
    person uuid not null
    references users(id)
    on delete cascade,
    follower uuid not null
    references users(id)
    on delete cascade,
    unique (person, follower)
);

-- +goose Down
drop table follows;