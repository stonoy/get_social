-- name: CreateUser :one
insert into users(id, created_at, updated_at, name, email, password, location, age, role, username, bio)
values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
returning *;

-- name: IsAdmin :one
select
case
    when count(*) = 0 then true
    else false
end as user_count
from users;

-- name: GetUserByEmail :one
select * from users where email = $1;

-- name: GetUserById :one
select * from users where id = $1;

-- name: UpdateUserDetails :one
update users
set updated_at = NOW(),
name = $1,
location = $2,
age = $3,
username = $4,
bio = $5
where id = $6
returning *;

-- name: GetUsers :many
select * from users
where (name like $1) and (location like $2);