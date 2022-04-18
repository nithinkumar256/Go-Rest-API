// main.go
package main

import (
	"fmt"
	"html"
	"net/http"
	"strconv"
	"strings"
	"time"
	"github.com/nithin/Data"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		urlPathElements := strings.Split(r.URL.Path, "/")

		if urlPathElements[1] == "getMonth" {
			number, _ := strconv.Atoi(strings.TrimSpace(urlPathElements[2]))
			if number == 0 || number > 12 {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("404 - Not Found"))
			} else {
				fmt.Fprintf(w, "%q", html.EscapeString(Data.Months[number]))
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Bad request"))
		}
	})
	// Create a server and run it on 8000 port
	s := &http.Server{
		Addr:           ":8000",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
