package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func main() {
    http.HandleFunc("/", hello)
    port := os.Getenv("PORT")
    if port == "" {
        port = "80"
    }
    fmt.Printf("Listening on port ", port)
    err := http.ListenAndServe(":"+port, nil)
    if err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}

func hello(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(w, "hello")
}
