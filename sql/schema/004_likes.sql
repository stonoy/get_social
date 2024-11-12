-- +goose Up
create table likes(
    id uuid primary key,
    created_at timestamp not null,
    updated_at timestamp not null,
    userid uuid not null
    references users(id)
    on delete cascade,
    postid uuid not null
    references posts(id)
    on delete cascade,
    unique(userid, postid)
);

-- +goose Down
drop table likes;