// cmd/main.go
package main

import (
    "log"
    "net/http"
    "omprakash/library_api/handlers"
)

func main() {
    router := handlers.NewRouter()
    log.Fatal(http.ListenAndServe(":8080", router))
}
