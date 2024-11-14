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

type Comment struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Comment   string    `json:"comment"`
	Userid    uuid.UUID `json:"user_id"`
	Postid    uuid.UUID `json:"post_id"`
}

type PostComments struct {
	ID      uuid.UUID `json:"id"`
	Comment string    `json:"comment"`
	Postid  uuid.UUID `json:"post_id"`
	UserId  uuid.UUID `json:"user_id"`
	Name    string    `json:"name"`
}

type FollowSuggestions struct {
	PersonId  uuid.UUID `json:"person_id"`
	Name      string    `json:"name"`
	Followers int64     `json:"followers"`
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

func commentDbToResp(comments []internal.GetCommentsPostRow) []PostComments {
	final := []PostComments{}

	for _, comment := range comments {
		final = append(final, PostComments{
			ID:      comment.ID,
			Comment: comment.Comment,
			Postid:  comment.Postid,
			UserId:  comment.Postid,
			Name:    comment.Name,
		})
	}

	return final
}

func followSuggestionsDbToResp(suggestions []internal.FollowSuggestionsRow) []FollowSuggestions {
	final := []FollowSuggestions{}

	for _, suggestion := range suggestions {
		final = append(final, FollowSuggestions{
			PersonId:  suggestion.ID,
			Name:      suggestion.Name,
			Followers: suggestion.Followers,
		})
	}

	return final
}
