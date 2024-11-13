-- name: LikeAPost :one
insert into likes(id , created_at, updated_at, postid, userid)
values ($1, $2, $3, $4, $5)
returning *;

-- name: RemoveLike :one
delete from likes where postid = $1 and userid = $2
returning *;

-- name: GetNumLikesPost :one
select count(*) from likes where postid = $1;