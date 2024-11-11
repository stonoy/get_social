package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func respWithError(w http.ResponseWriter, code int, msg string) {
	msgStruct := struct {
		Msg string `json:"msg"`
	}{
		Msg: msg,
	}

	if code > 499 {
		log.Printf("error in server -> %v", code)
	}

	respWithJson(w, code, msgStruct)
}

func respWithJson(w http.ResponseWriter, code int, msg interface{}) {
	// convert/marshal into byte
	dat, err := json.Marshal(msg)
	if err != nil {
		log.Panicf("can not marshal the msg staruct -> %v", err)
		respWithError(w, 500, fmt.Sprintf("can not marshal the msg staruct -> %v", err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}
