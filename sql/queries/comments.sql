-- name: CreateComment :one
insert into comments(id, created_at, updated_at, comment, postid, userid)
values ($1, $2, $3, $4, $5, $6)
returning *;

-- name: DeleteComment :one
delete from comments where id = $1 and userid = $2
returning *;

-- name: UpdateComment :one
update comments
set updated_at = NOW(),
comment = $1
where id = $2 and userid = $3
returning *;

-- name: GetCommentsPost :many
select comments.id,comments.comment,comments.postid,users.id,users.name
from comments
inner join users
on comments.userid = users.id
where comments.postid = $1
order by comments.created_at;