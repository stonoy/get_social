package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/stonoy/get_social/internal"
	"golang.org/x/crypto/bcrypt"
)

func (cfg *apiConfig) register(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var reqStruct struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := decoder.Decode(&reqStruct)
	if err != nil {
		respWithError(w, 400, fmt.Sprintf("error in decoding json -> %v", err))
		return
	}

	if reqStruct.Email == "" || reqStruct.Name == "" || reqStruct.Password == "" {
		respWithError(w, 400, "enter correct credentials")
		return
	}

	// encrypt password
	passwordByte, err := bcrypt.GenerateFromPassword([]byte(reqStruct.Password), 14)
	if err != nil {
		respWithError(w, 400, fmt.Sprintf("error in encrypting password -> %v", err))
		return
	}

	// define default role
	var role internal.UserType = internal.UserTypeUser

	// check admin
	isAdmin, err := cfg.db.IsAdmin(r.Context())
	if err != nil {
		respWithError(w, 500, fmt.Sprintf("error in IsAdmin -> %v", err))
		return
	}

	if isAdmin {
		role = internal.UserTypeAdmin
	}

	// create user
	theUser, err := cfg.db.CreateUser(r.Context(), internal.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      reqStruct.Name,
		Email:     reqStruct.Email,
		Password:  string(passwordByte),
		Username:  reqStruct.Email,
		Bio:       "",
		Age:       18,
		Location:  "",
		Role:      role,
	})

	if err != nil {
		respWithError(w, 500, fmt.Sprintf("error in CreateUser -> %v", err))
		return
	}

	// generate token
	token, err := GenerateToken(cfg.jwt_secret, theUser)
	if err != nil {
		respWithError(w, 500, fmt.Sprintf("error in GenerateToken -> %v", err))
		return
	}

	type respStruct struct {
		Token string `json:"token"`
		User  User   `json:"user"`
	}

	respWithJson(w, 201, respStruct{
		Token: token,
		User: User{
			ID:        theUser.ID,
			CreatedAt: theUser.CreatedAt,
			UpdatedAt: theUser.UpdatedAt,
			Name:      theUser.Name,
			Email:     theUser.Email,
			Username:  theUser.Username,
			Location:  theUser.Location,
			Age:       theUser.Age,
			Bio:       theUser.Bio,
			Role:      string(theUser.Role),
		},
	})

}

func (cfg *apiConfig) login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var reqStruct struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := decoder.Decode(&reqStruct)
	if err != nil {
		respWithError(w, 400, fmt.Sprintf("error in decoding json -> %v", err))
		return
	}

	if reqStruct.Email == "" || reqStruct.Password == "" {
		respWithError(w, 400, "enter correct credentials")
		return
	}

	// get the user
	theUser, err := cfg.db.GetUserByEmail(r.Context(), reqStruct.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			respWithError(w, 400, "No such user exist")
			return
		} else {
			respWithError(w, 500, fmt.Sprintf("error in GetUserByEmail -> %v", err))
			return
		}
	}

	// check password
	err = bcrypt.CompareHashAndPassword([]byte(theUser.Password), []byte(reqStruct.Password))
	if err != nil {
		respWithError(w, 400, fmt.Sprintf("password not matched"))
		return
	}

	// generate token
	token, err := GenerateToken(cfg.jwt_secret, theUser)
	if err != nil {
		respWithError(w, 500, fmt.Sprintf("error in GenerateToken -> %v", err))
		return
	}

	type respStruct struct {
		Token string `json:"token"`
		User  User   `json:"user"`
	}

	respWithJson(w, 201, respStruct{
		Token: token,
		User: User{
			ID:        theUser.ID,
			CreatedAt: theUser.CreatedAt,
			UpdatedAt: theUser.UpdatedAt,
			Name:      theUser.Name,
			Email:     theUser.Email,
			Username:  theUser.Username,
			Location:  theUser.Location,
			Age:       theUser.Age,
			Bio:       theUser.Bio,
			Role:      string(theUser.Role),
		},
	})

}
