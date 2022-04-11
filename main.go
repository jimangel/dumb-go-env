package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func homePage(w http.ResponseWriter, r *http.Request) {

        fmt.Fprintf(w, "Blue/Green is currently: %s \n", os.Getenv("COLOR"))
        fmt.Fprintf(w, "Env is currently: %s", os.Getenv("ENV"))
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")) , nil))
}

func main() {
    handleRequests()
}