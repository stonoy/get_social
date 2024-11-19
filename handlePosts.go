package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/stonoy/get_social/internal"
)

func (cfg *apiConfig) CreatePosts(w http.ResponseWriter, r *http.Request, user internal.User) {
	type reqStruct struct {
		Content string `json:"content"`
	}

	decoder := json.NewDecoder(r.Body)
	reqObj := reqStruct{}
	err := decoder.Decode(&reqObj)
	if err != nil {
		respWithError(w, 400, fmt.Sprintf("error in decoding json -> %v", err))
		return
	}

	post, err := cfg.db.CreatePosts(r.Context(), internal.CreatePostsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Content:   reqObj.Content,
		Author:    user.ID,
	})
	if err != nil {
		respWithError(w, 500, fmt.Sprintf("error in CreatePosts -> %v", err))
		return
	}

	type respStruct struct {
		Success bool `json:"success"`
		Post    Post `json:"post"`
	}

	respWithJson(w, 201, respStruct{
		Success: true,
		Post: Post{
			ID:        post.ID,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
			Content:   post.Content,
			Author:    post.Author,
		},
	})
}

func (cfg *apiConfig) GetSinglePost(w http.ResponseWriter, r *http.Request) {
	postIdStr := chi.URLParam(r, "postID")

	postId, err := uuid.Parse(postIdStr)
	if err != nil {
		respWithError(w, 400, fmt.Sprintf("error in parsing uuid -> %v", err))
		return
	}

	post, err := cfg.db.GetPostById(r.Context(), postId)
	if err != nil {
		if err == sql.ErrNoRows {
			respWithError(w, 400, "No such post exist")
			return
		} else {
			respWithError(w, 500, fmt.Sprintf("error in GetPostById -> %v", err))
			return
		}
	}

	type respStruct struct {
		Post PostWithUser `json:"post"`
	}

	respWithJson(w, 201, respStruct{
		Post: PostWithUser{
			ID:        post.ID,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
			Content:   post.Content,
			Author:    post.Author,
			Name:      post.Name,
		},
	})
}

func (cfg *apiConfig) DeletePost(w http.ResponseWriter, r *http.Request, user internal.User) {
	postIdStr := chi.URLParam(r, "postID")

	postId, err := uuid.Parse(postIdStr)
	if err != nil {
		respWithError(w, 400, fmt.Sprintf("error in parsing uuid -> %v", err))
		return
	}

	post, err := cfg.db.DeletePost(r.Context(), internal.DeletePostParams{
		ID:     postId,
		Author: user.ID,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			respWithError(w, 400, "No such post exist")
			return
		} else {
			respWithError(w, 500, fmt.Sprintf("error in DeletePost -> %v", err))
			return
		}
	}

	type respStruct struct {
		Success bool `json:"success"`
		Post    Post `json:"post"`
	}

	respWithJson(w, 201, respStruct{
		Success: true,
		Post: Post{
			ID:        post.ID,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
			Content:   post.Content,
			Author:    post.Author,
		},
	})
}

func (cfg *apiConfig) UpdatePost(w http.ResponseWriter, r *http.Request, user internal.User) {
	postIdStr := chi.URLParam(r, "postID")

	postId, err := uuid.Parse(postIdStr)
	if err != nil {
		respWithError(w, 400, fmt.Sprintf("error in parsing uuid -> %v", err))
		return
	}

	type reqStruct struct {
		Content string `json:"content"`
	}

	decoder := json.NewDecoder(r.Body)
	reqObj := reqStruct{}
	err = decoder.Decode(&reqObj)
	if err != nil {
		respWithError(w, 400, fmt.Sprintf("error in decoding json -> %v", err))
		return
	}

	post, err := cfg.db.UpdatePost(r.Context(), internal.UpdatePostParams{
		Content: reqObj.Content,
		ID:      postId,
		Author:  user.ID,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			respWithError(w, 400, "No such post exist")
			return
		} else {
			respWithError(w, 500, fmt.Sprintf("error in UpdatePost -> %v", err))
			return
		}
	}

	type respStruct struct {
		Success bool `json:"success"`
		Post    Post `json:"post"`
	}

	respWithJson(w, 201, respStruct{
		Success: true,
		Post: Post{
			ID:        post.ID,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
			Content:   post.Content,
			Author:    post.Author,
		},
	})
}

func (cfg *apiConfig) PostSuggestion(w http.ResponseWriter, r *http.Request, user internal.User) {
	// get query from request url
	queries := r.URL.Query()

	// set default queries
	page := 1
	startTime := time.Now().AddDate(-1, 0, 0)
	endTime := time.Now().AddDate(1, 0, 0)

	// update quries from url
	pageStr := queries.Get("page")
	if pageStr != "" {
		pageInt, err := strconv.Atoi(pageStr)
		if err != nil {
			respWithError(w, 400, fmt.Sprintf("error in converting str to int page -> %v", err))
			return
		}

		page = pageInt
	}

	startTimeQ := queries.Get("startTime")
	if startTimeQ != "" {
		// parse time
		theTime, err := GetTimeFromStr(startTimeQ)
		if err != nil {
			respWithError(w, 400, fmt.Sprintf("error in GetTimeFromStr -> %v", err))
			return
		}
		startTime = theTime
	}

	endTimeQ := queries.Get("endTime")
	if endTimeQ != "" {
		// parse time
		theTime, err := GetTimeFromStr(endTimeQ)
		if err != nil {
			respWithError(w, 400, fmt.Sprintf("error in GetTimeFromStr -> %v", err))
			return
		}
		endTime = theTime
	}

	limit := 2
	offset := (page - 1) * limit

	posts, err := cfg.db.PostSuggestions(r.Context(), internal.PostSuggestionsParams{
		Follower:    user.ID,
		Limit:       int32(limit),
		Offset:      int32(offset),
		CreatedAt:   startTime,
		CreatedAt_2: endTime,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			respWithError(w, 400, "No such post exist")
			return
		} else {
			respWithError(w, 500, fmt.Sprintf("error in PostSuggestions -> %v", err))
			return
		}
	}

	numOfPosts, err := cfg.db.NumPostSuggestions(r.Context(), user.ID)
	if err != nil {
		respWithError(w, 500, fmt.Sprintf("error in NumPostSuggestions -> %v", err))
		return
	}

	numOfPages := math.Ceil(float64(numOfPosts) / float64(limit))

	type respStruct struct {
		Posts      []PostWithUser `json:"posts"`
		NumOfPages float64        `json:"numOfPages"`
		Page       int            `json:"page"`
	}

	respWithJson(w, 200, respStruct{
		Posts:      postDbToResp2(posts),
		NumOfPages: numOfPages,
		Page:       page,
	})
}

func (cfg *apiConfig) GetPostsByUser(w http.ResponseWriter, r *http.Request, user internal.User) {
	// get query from request url
	queries := r.URL.Query()

	// set default queries
	page := 1

	// update quries from url
	pageStr := queries.Get("page")
	if pageStr != "" {
		pageInt, err := strconv.Atoi(pageStr)
		if err != nil {
			respWithError(w, 400, fmt.Sprintf("error in converting str to int page -> %v", err))
			return
		}

		page = pageInt
	}

	limit := 2
	offset := (page - 1) * limit

	posts, err := cfg.db.GetPostsByIUser(r.Context(), internal.GetPostsByIUserParams{
		Author: user.ID,
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		if err == sql.ErrNoRows {
			respWithError(w, 400, "No such post exist")
			return
		} else {
			respWithError(w, 500, fmt.Sprintf("error in GetPostsByIUser -> %v", err))
			return
		}
	}

	numOfPosts, err := cfg.db.GetNumPostsByIUser(r.Context(), user.ID)
	if err != nil {
		respWithError(w, 500, fmt.Sprintf("error in GetNumPostsByIUser -> %v", err))
		return
	}

	numOfPages := math.Ceil(float64(numOfPosts) / float64(limit))

	type respStruct struct {
		Posts      []PostWithUser `json:"posts"`
		NumOfPages float64        `json:"numOfPages"`
		Page       int            `json:"page"`
	}

	respWithJson(w, 200, respStruct{
		Posts:      postDbToResp1(posts),
		NumOfPages: numOfPages,
		Page:       page,
	})
}
