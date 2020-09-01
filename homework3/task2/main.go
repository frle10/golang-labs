package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type webResponseEven struct {
	Number int `json:"n"`
	Even bool `json:"even"`
}

type webResponseSquare struct {
	Number int `json:"n"`
	Square int `json:"s"`
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", hello)
	r.Get("/even/{n}", even)
	r.Get("/square/{n}", square)
	log.Panic(http.ListenAndServe(":3333", r))
}

func hello(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("fergo!"))
}

func even(w http.ResponseWriter, r *http.Request) {
	n, err := strconv.ParseInt(chi.URLParam(r, "n"), 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	even := false
	if n % 2 != 0 {
		even = false
	} else {
		even = true
	}

	evenResponse := webResponseEven{int(n), even}
	js, _ := json.Marshal(evenResponse)
	
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func square(w http.ResponseWriter, r *http.Request) {
	n, err := strconv.ParseInt(chi.URLParam(r, "n"), 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	square := n * n
	squareResponse := webResponseSquare{int(n), int(square)}
	js, _ := json.Marshal(squareResponse)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
