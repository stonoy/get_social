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
where follows.person not in (
	select person from follows
	where follows.follower = $1
)
group by users.id;

-- name: PersonsIFollow :many
select users.id, users.name 
from follows
inner join users
on follows.person = users.id
where follower = $1;

-- name: MyFollowers :many
select users.id, users.name 
from follows
inner join users
on follows.follower = users.id
where person = $1;

