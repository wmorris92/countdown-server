package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wmorris92/countdown-server/solver"
)

func solveCountdown(w http.ResponseWriter, r *http.Request) {
	type countdownResult struct {
		Words []string `json:"words"`
	}
	type errorResult struct {
		Error string `json:"error"`
	}
	vars := mux.Vars(r)
	solutions, err := solver.FindWordsForLetters(vars["letters"])
	var response []byte

	if err != nil {
		response, _ = json.Marshal(errorResult{err.Error()})
	} else {
		response, _ = json.Marshal(countdownResult{solutions})
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte(string(response)))
}

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/countdown/{letters}", solveCountdown)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8080", r))
}
