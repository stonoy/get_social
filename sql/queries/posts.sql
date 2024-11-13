-- name: CreatePosts :one
insert into posts(id, created_at, updated_at, content, author)
values ($1, $2, $3, $4, $5)
returning *;

-- name: GetPostById :one
select * from posts where id = $1;

-- name: GetPostsByIUser :many
select * from posts 
where author = $1
limit $2 offset $3;

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
select * from posts
where author in (
    select person from follows
    where follower = $1
)
order by created_at
limit $2 offset $3;

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