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