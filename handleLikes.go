package main

import (
	"database/sql"
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

	_, err = cfg.db.LikeAPost(r.Context(), internal.LikeAPostParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Postid:    postId,
		Userid:    user.ID,
	})
	if err != nil {
		if isUniqueViolation(err) {
			// delete the like
			_, err = cfg.db.RemoveLike(r.Context(), internal.RemoveLikeParams{
				Postid: postId,
				Userid: user.ID,
			})
			if err != nil {
				if err == sql.ErrNoRows {
					respWithError(w, 400, "No such post like exist")
					return
				} else {
					respWithError(w, 500, fmt.Sprintf("error in RemoveLike -> %v", err))
					return
				}
			}

			// remove like to the post
			_, err = cfg.db.HandlePostLike(r.Context(), internal.HandlePostLikeParams{
				ID:    postId,
				Likes: -1,
			})
			if err != nil {
				if err == sql.ErrNoRows {
					respWithError(w, 400, "No such post exist")
					return
				} else {
					respWithError(w, 500, fmt.Sprintf("error in HandlePostLike -> %v", err))
					return
				}
			}

			type respStruct struct {
				Deleted bool `json:"deleted"`
			}

			respWithJson(w, 200, respStruct{
				Deleted: true,
			})
			return
		} else {
			respWithError(w, 500, fmt.Sprintf("error in LikeAPost -> %v", err))
			return
		}
	}

	// add like to the post
	_, err = cfg.db.HandlePostLike(r.Context(), internal.HandlePostLikeParams{
		ID:    postId,
		Likes: 1,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			respWithError(w, 400, "No such post exist")
			return
		} else {
			respWithError(w, 500, fmt.Sprintf("error in HandlePostLike -> %v", err))
			return
		}
	}

	type respStruct struct {
		Success bool `json:"success"`
	}

	respWithJson(w, 201, respStruct{
		Success: true,
	})

}

// remove like obsolate 19.11.24
func (cfg *apiConfig) RemoveLike(w http.ResponseWriter, r *http.Request, user internal.User) {
	// get postId from url
	postIdStr := chi.URLParam(r, "postID")

	postId, err := uuid.Parse(postIdStr)
	if err != nil {
		respWithError(w, 400, fmt.Sprintf("error in parsing str -> uuid , %v", err))
		return
	}

	_, err = cfg.db.RemoveLike(r.Context(), internal.RemoveLikeParams{
		Postid: postId,
		Userid: user.ID,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			respWithError(w, 400, "No such post like exist")
			return
		} else {
			respWithError(w, 500, fmt.Sprintf("error in RemoveLike -> %v", err))
			return
		}
	}

	// remove like to the post
	_, err = cfg.db.HandlePostLike(r.Context(), internal.HandlePostLikeParams{
		ID:    postId,
		Likes: -1,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			respWithError(w, 400, "No such post exist")
			return
		} else {
			respWithError(w, 500, fmt.Sprintf("error in HandlePostLike -> %v", err))
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
