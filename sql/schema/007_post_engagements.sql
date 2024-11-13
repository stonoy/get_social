-- +goose Up
alter table posts
    add column likes integer default 0 not null,
    add column comments integer default 0 not null;

-- +goose Down
alter table posts
    drop column likes,
    drop column comments;

