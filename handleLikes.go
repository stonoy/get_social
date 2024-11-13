package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/stonoy/get_social/internal"
)

func (cfg *apiConfig) LikeAPost(w http.ResponseWriter, r *http.Request, user internal.User) {
	// get postId from url
	postIdStr := chi.URLParam(r, "postID")

	postId, err := uuid.Parse(postIdStr)
	if err != nil {
		respWithError(w, 400, fmt.Sprintf("error in parsing str -> uuid , %v", err))
		return
	}

	post, err := cfg.db.LikeAPost(r.Context(), internal.LikeAPostParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Postid:    postId,
		Userid:    user.ID,
	})
	if err != nil {
		respWithError(w, 500, fmt.Sprintf("error in LikeAPost , %v", err))
		return
	}

	type respStruct struct {
		Success bool `json:"success"`
	}

	respWithJson(w, 201, respStruct{
		Success: true,
	})

}
