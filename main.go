package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Board [9]int // 9 for board, 10 for winner
const (
	E = 0
	O = 1
	X = -1
)

func blank() Board {
	return Board{E, E, E, E, E, E, E, E}
}

var boards []Board = []Board{}

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//b := blank()
	// js, err := json.Marshal(b)
	js, err := json.Marshal(boards)
	w.Write(js)
	if err != nil {
		log.Fatal(err)
	}

}

func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "post called"}`))
}

func put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"message": "put called"}`))
}

func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "delete called"}`))
}

func InitBoard() {
	boards = append(boards, blank())
	boards = append(boards, Board{E, E, E, E, E, E, X, E, E})
	boards = append(boards, Board{E, E, E, E, O, E, X, E, E})

}

func main() {

	InitBoard()
	println("len(boards) ", len(boards))

	fmt.Println("start")
	//handleRequests()
	r := mux.NewRouter()
	r.HandleFunc("/board", get).Methods(http.MethodGet)
	r.HandleFunc("/board", post).Methods(http.MethodPost)
	r.HandleFunc("/board", put).Methods(http.MethodPut)
	r.HandleFunc("/board", delete).Methods(http.MethodDelete)
	//r.HandleFunc("/board", )
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", r))
	fmt.Println("done")
}

// TEST
