// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: posts.sql

package internal

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createPosts = `-- name: CreatePosts :one
insert into posts(id, created_at, updated_at, content, author)
values ($1, $2, $3, $4, $5)
returning id, created_at, updated_at, content, author, likes, comments
`

type CreatePostsParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Content   string
	Author    uuid.UUID
}

func (q *Queries) CreatePosts(ctx context.Context, arg CreatePostsParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createPosts,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Content,
		arg.Author,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Content,
		&i.Author,
		&i.Likes,
		&i.Comments,
	)
	return i, err
}

const deletePost = `-- name: DeletePost :one
delete from posts where id = $1 and author = $2
returning id, created_at, updated_at, content, author, likes, comments
`

type DeletePostParams struct {
	ID     uuid.UUID
	Author uuid.UUID
}

func (q *Queries) DeletePost(ctx context.Context, arg DeletePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, deletePost, arg.ID, arg.Author)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Content,
		&i.Author,
		&i.Likes,
		&i.Comments,
	)
	return i, err
}

const getNumPostsByIUser = `-- name: GetNumPostsByIUser :one
select count(*) from posts where author = $1
`

func (q *Queries) GetNumPostsByIUser(ctx context.Context, author uuid.UUID) (int64, error) {
	row := q.db.QueryRowContext(ctx, getNumPostsByIUser, author)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getPostById = `-- name: GetPostById :one
select posts.id, posts.created_at, posts.updated_at, posts.content, posts.author, posts.likes, posts.comments, users.name from posts
inner join users
on posts.author = users.id
where posts.id = $1
`

type GetPostByIdRow struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Content   string
	Author    uuid.UUID
	Likes     int32
	Comments  int32
	Name      string
}

func (q *Queries) GetPostById(ctx context.Context, id uuid.UUID) (GetPostByIdRow, error) {
	row := q.db.QueryRowContext(ctx, getPostById, id)
	var i GetPostByIdRow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Content,
		&i.Author,
		&i.Likes,
		&i.Comments,
		&i.Name,
	)
	return i, err
}

const getPostsByIUser = `-- name: GetPostsByIUser :many
select posts.id, posts.created_at, posts.updated_at, posts.content, posts.author, posts.likes, posts.comments, users.name from posts
inner join users
on posts.author = users.id
where posts.author = $1
order by posts.created_at desc
limit $2 offset $3
`

type GetPostsByIUserParams struct {
	Author uuid.UUID
	Limit  int32
	Offset int32
}

type GetPostsByIUserRow struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Content   string
	Author    uuid.UUID
	Likes     int32
	Comments  int32
	Name      string
}

func (q *Queries) GetPostsByIUser(ctx context.Context, arg GetPostsByIUserParams) ([]GetPostsByIUserRow, error) {
	rows, err := q.db.QueryContext(ctx, getPostsByIUser, arg.Author, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPostsByIUserRow
	for rows.Next() {
		var i GetPostsByIUserRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Content,
			&i.Author,
			&i.Likes,
			&i.Comments,
			&i.Name,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const handlePostComments = `-- name: HandlePostComments :one
update posts
set updated_at = NOW(),
comments = greatest(0, comments + $1)
where id = $2
returning id, created_at, updated_at, content, author, likes, comments
`

type HandlePostCommentsParams struct {
	Comments int32
	ID       uuid.UUID
}

func (q *Queries) HandlePostComments(ctx context.Context, arg HandlePostCommentsParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, handlePostComments, arg.Comments, arg.ID)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Content,
		&i.Author,
		&i.Likes,
		&i.Comments,
	)
	return i, err
}

const handlePostLike = `-- name: HandlePostLike :one
update posts
set updated_at = NOW(),
likes = greatest(0, likes + $1)
where id = $2
returning id, created_at, updated_at, content, author, likes, comments
`

type HandlePostLikeParams struct {
	Likes int32
	ID    uuid.UUID
}

func (q *Queries) HandlePostLike(ctx context.Context, arg HandlePostLikeParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, handlePostLike, arg.Likes, arg.ID)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Content,
		&i.Author,
		&i.Likes,
		&i.Comments,
	)
	return i, err
}

const numPostSuggestions = `-- name: NumPostSuggestions :one
select count(*) from posts
where author in (
    select person from follows
    where follower = $1
)
`

func (q *Queries) NumPostSuggestions(ctx context.Context, follower uuid.UUID) (int64, error) {
	row := q.db.QueryRowContext(ctx, numPostSuggestions, follower)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const postSuggestions = `-- name: PostSuggestions :many
select posts.id, posts.created_at, posts.updated_at, posts.content, posts.author, posts.likes, posts.comments, users.name from posts
inner join users
on posts.author = users.id
where posts.author in (
    select person from follows
    where follower = $1
) and posts.created_at between $2 and $3
order by posts.created_at desc
limit $4 offset $5
`

type PostSuggestionsParams struct {
	Follower    uuid.UUID
	CreatedAt   time.Time
	CreatedAt_2 time.Time
	Limit       int32
	Offset      int32
}

type PostSuggestionsRow struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Content   string
	Author    uuid.UUID
	Likes     int32
	Comments  int32
	Name      string
}

func (q *Queries) PostSuggestions(ctx context.Context, arg PostSuggestionsParams) ([]PostSuggestionsRow, error) {
	rows, err := q.db.QueryContext(ctx, postSuggestions,
		arg.Follower,
		arg.CreatedAt,
		arg.CreatedAt_2,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []PostSuggestionsRow
	for rows.Next() {
		var i PostSuggestionsRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Content,
			&i.Author,
			&i.Likes,
			&i.Comments,
			&i.Name,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updatePost = `-- name: UpdatePost :one
update posts
set updated_at = NOW(),
content = $1
where id = $2 and author = $3
returning id, created_at, updated_at, content, author, likes, comments
`

type UpdatePostParams struct {
	Content string
	ID      uuid.UUID
	Author  uuid.UUID
}

func (q *Queries) UpdatePost(ctx context.Context, arg UpdatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, updatePost, arg.Content, arg.ID, arg.Author)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Content,
		&i.Author,
		&i.Likes,
		&i.Comments,
	)
	return i, err
}
