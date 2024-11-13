package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/stonoy/get_social/internal"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Location  string    `json:"location"`
	Age       int32     `json:"age"`
	Username  string    `json:"username"`
	Bio       string    `json:"bio"`
	Role      string    `json:"role"`
}

type Post struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Content   string    `json:"content"`
	Author    uuid.UUID `json:"author"`
}

func postDbToResp(posts []internal.Post) []Post {
	final := []Post{}

	for _, post := range posts {
		final = append(final, Post{
			ID:        post.ID,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
			Content:   post.Content,
			Author:    post.Author,
		})
	}
	return final
}
