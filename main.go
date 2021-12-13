package main

import (
        "fmt"
        "html"
        "log"
        "net/http"
)

func main() {

        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
                fmt.Fprintf(w, "SELAMAT DATANG!!!, %q", html.EscapeString(r.URL.Path))
        })

        http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
                fmt.Fprintf(w, "Semoga Lancar")
        })

        log.Fatal(http.ListenAndServe(":8081", nil))
}

