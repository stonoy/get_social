package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/stonoy/get_social/internal"
)

func (cfg *apiConfig) CreateComment(w http.ResponseWriter, r *http.Request, user internal.User) {
	decoder := json.NewDecoder(r.Body)
	var reqStruct struct {
		Comment string `json:"comment"`
		PostId  string `json:"post_id"`
	}
	err := decoder.Decode(&reqStruct)
	if err != nil {
		respWithError(w, 400, fmt.Sprintf("error in decoding req body -> %v", err))
		return
	}

	postId, err := uuid.Parse(reqStruct.PostId)
	if err != nil {
		respWithError(w, 400, fmt.Sprintf("error in parsing uuid -> %v", err))
		return
	}

	comment, err := cfg.db.CreateComment(r.Context(), internal.CreateCommentParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Comment:   reqStruct.Comment,
		Postid:    postId,
		Userid:    user.ID,
	})
	if err != nil {
		respWithError(w, 500, fmt.Sprintf("error in CreateComment -> %v", err))
		return
	}

	type respStruct struct {
		Success bool    `json:"success"`
		Comment Comment `json:"comment"`
	}

	respWithJson(w, 201, respStruct{
		Success: true,
		Comment: Comment{
			ID:        comment.ID,
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
			Comment:   comment.Comment,
			Userid:    comment.Userid,
			Postid:    comment.Postid,
		},
	})
}

func (cfg *apiConfig) DeleteComment(w http.ResponseWriter, r *http.Request, user internal.User) {
	commentIdStr := chi.URLParam(r, "commentID")

	commentId, err := uuid.Parse(commentIdStr)
	if err != nil {
		respWithError(w, 400, fmt.Sprintf("error in parsing uuid -> %v", err))
		return
	}

	_, err = cfg.db.DeleteComment(r.Context(), internal.DeleteCommentParams{
		ID:     commentId,
		Userid: user.ID,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			respWithError(w, 400, "No such comment exist")
			return
		} else {
			respWithError(w, 500, fmt.Sprintf("error in DeleteComment -> %v", err))
			return
		}
	}

	type respStruct struct {
		Success bool `json:"success"`
	}

	respWithJson(w, 200, respStruct{
		Success: true,
	})
}

func (cfg *apiConfig) UpdateComment(w http.ResponseWriter, r *http.Request, user internal.User) {
	commentIdStr := chi.URLParam(r, "commentID")

	commentId, err := uuid.Parse(commentIdStr)
	if err != nil {
		respWithError(w, 400, fmt.Sprintf("error in parsing uuid -> %v", err))
		return
	}

	decoder := json.NewDecoder(r.Body)
	var reqStruct struct {
		Comment string `json:"comment"`
	}
	err = decoder.Decode(&reqStruct)
	if err != nil {
		respWithError(w, 400, fmt.Sprintf("error in decoding req body -> %v", err))
		return
	}

	updatedComment, err := cfg.db.UpdateComment(r.Context(), internal.UpdateCommentParams{
		Comment: reqStruct.Comment,
		ID:      commentId,
		Userid:  user.ID,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			respWithError(w, 400, "No such comment exist")
			return
		} else {
			respWithError(w, 500, fmt.Sprintf("error in UpdateComment -> %v", err))
			return
		}
	}

	type respStruct struct {
		Success bool    `json:"success"`
		Comment Comment `json:"comment"`
	}

	respWithJson(w, 200, respStruct{
		Success: true,
		Comment: Comment{
			ID:        updatedComment.ID,
			CreatedAt: updatedComment.CreatedAt,
			UpdatedAt: updatedComment.UpdatedAt,
			Comment:   updatedComment.Comment,
			Userid:    updatedComment.Userid,
			Postid:    updatedComment.Postid,
		},
	})
}
