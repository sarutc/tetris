package main

import (
    "html/template"
    "log"
    "net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("index.html")
    t.Execute(w, nil) 
}

func main() {
    http.HandleFunc("/", HomePage)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
