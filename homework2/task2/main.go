package main

import (
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	s := &http.Server{
		Addr: ":9210",
		Handler: handler(),
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}

func handler() http.HandlerFunc {
	return myFooBarHandler
}

func myFooBarHandler(w http.ResponseWriter, req *http.Request) {
	n, err := strconv.ParseInt(req.URL.Query().Get("n"), 10, 64)

	if req.Method != "GET" || err != nil || n <= 0 {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		res := ""
		for i := 1; i <= int(n); i++ {
			if i % 15 == 0 {
				res += "buz"
			} else if i % 3 == 0 {
				res += "foo"
			} else if i % 5 == 0 {
				res += "bar"
			} else {
				res += strconv.FormatInt(int64(i), 10)
			}

			if i != int(n) {
				res += " "
			}
		}

		w.Write([]byte(res))
	}
}
