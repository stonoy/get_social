-- name: CreateUser :one
insert into users(id, created_at, updated_at, name, email, password, location, age, role, username, bio)
values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
returning *;