-- +goose Up
create table comments(
    id uuid primary key,
    created_at timestamp not null,
    updated_at timestamp not null,
    comment text not null,
    userid uuid not null
    references users(id)
    on delete cascade,
    postid uuid not null
    references posts(id)
    on delete cascade
);

-- +goose Down
drop table comments;