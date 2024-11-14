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

func (cfg *apiConfig) Follow(w http.ResponseWriter, r *http.Request, user internal.User) {
	type reqStruct struct {
		Person string `json:"person"`
	}

	decoder := json.NewDecoder(r.Body)
	reqObj := &reqStruct{}
	err := decoder.Decode(reqObj)
	if err != nil {
		respWithError(w, 400, fmt.Sprintf("eror in decoding req body -> %v", err))
		return
	}

	personId, err := uuid.Parse(reqObj.Person)
	if err != nil {
		respWithError(w, 400, fmt.Sprintf("error in parsing uuid -> %v", err))
		return
	}

	follow, err := cfg.db.Follow(r.Context(), internal.FollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Person:    personId,
		Follower:  user.ID,
	})
	if err != nil {
		respWithError(w, 500, fmt.Sprintf("error in Follow -> %v", err))
		return
	}

	type respStruct struct {
		SuccessMsg string `json:"success_msg"`
	}

	respWithJson(w, 201, respStruct{
		SuccessMsg: fmt.Sprintf("%v user followed %v with follow Id -> %v", user.ID, personId, follow.ID),
	})
}

func (cfg *apiConfig) Unfollow(w http.ResponseWriter, r *http.Request, user internal.User) {
	personIdStr := chi.URLParam(r, "personID")

	personId, err := uuid.Parse(personIdStr)
	if err != nil {
		respWithError(w, 400, fmt.Sprintf("error in parsing uuid -> %v", err))
		return
	}

	follow, err := cfg.db.Unfollow(r.Context(), internal.UnfollowParams{
		Person:   personId,
		Follower: user.ID,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			respWithError(w, 400, "No such follow exist")
			return
		} else {
			respWithError(w, 500, fmt.Sprintf("error in Unfollow -> %v", err))
			return
		}
	}

	type respStruct struct {
		SuccessMsg string `json:"success_msg"`
	}

	respWithJson(w, 201, respStruct{
		SuccessMsg: fmt.Sprintf("%v user unfollowed %v with follow Id -> %v", user.ID, personId, follow.ID),
	})
}

func (cfg *apiConfig) FollowSuggestions(w http.ResponseWriter, r *http.Request, user internal.User) {
	suggestions, err := cfg.db.FollowSuggestions(r.Context(), user.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			respWithError(w, 400, "No such follows exist")
			return
		} else {
			respWithError(w, 500, fmt.Sprintf("error in Unfollow -> %v", err))
			return
		}
	}

	type respStruct struct {
		FollowSuggestions []FollowSuggestions `json:"follow_suggestions"`
	}

	respWithJson(w, 200, respStruct{
		FollowSuggestions: followSuggestionsDbToResp(suggestions),
	})
}
