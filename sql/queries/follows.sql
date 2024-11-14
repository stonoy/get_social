-- name: Follow :one
insert into follows(id, created_at, updated_at, person, follower)
values ($1, $2, $3, $4, $5)
returning *;

-- name: Unfollow :one
delete from follows where person = $1 and follower = $2
returning *;

-- name: FollowSuggestions :many
select users.id, users.name, count(follows.*) as followers
from follows
inner join users
on follows.person = users.id
where follows.follower != $1
group by users.id
limit 3;

