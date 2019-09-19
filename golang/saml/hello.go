package main

import (
    "fmt"
    "net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World!")
}

func main() {
    app := http.HandlerFunc( hello )
    http.Handle("/hello", app)
    http.ListenAndServe(":8000", nil)
}
