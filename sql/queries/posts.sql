-- name: CreatePosts :one
insert into posts(id, created_at, updated_at, content, author)
values ($1, $2, $3, $4, $5)
returning *;

-- name: GetPostById :one
select posts.*, users.name from posts
inner join users
on posts.author = users.id
where posts.id = $1;

-- name: GetPostsByIUser :many
select posts.*, users.name from posts
inner join users
on posts.author = users.id
where posts.author = $1
order by posts.created_at desc
limit $2 offset $3;

-- name: GetNumPostsByIUser :one
select count(*) from posts where author = $1;

-- name: DeletePost :one
delete from posts where id = $1 and author = $2
returning *;

-- name: UpdatePost :one
update posts
set updated_at = NOW(),
content = $1
where id = $2 and author = $3
returning *;

-- name: PostSuggestions :many
select posts.*, users.name from posts
inner join users
on posts.author = users.id
where posts.author in (
    select person from follows
    where follower = $1
) and posts.created_at between $2 and $3
order by posts.created_at desc
limit $4 offset $5;

-- name: NumPostSuggestions :one
select count(*) from posts
where author in (
    select person from follows
    where follower = $1
);

-- name: HandlePostLike :one
update posts
set updated_at = NOW(),
likes = greatest(0, likes + $1)
where id = $2
returning *;

-- name: HandlePostComments :one
update posts
set updated_at = NOW(),
comments = greatest(0, comments + $1)
where id = $2
returning *;
