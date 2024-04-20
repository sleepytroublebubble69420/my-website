package main

import (
    "net/http"
    "log"
) 

func main() {
    http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
        log.Printf("Got request, Method: %s", r.Method)
        err := r.ParseForm()
        if err != nil {
            log.Fatal(err)
        }
        log.Println(r.Form)
        log.Println(r.PostForm)
    })

    fs_handler := http.FileServer(http.Dir("static_content/"))
    http.Handle("/", fs_handler)

    err := http.ListenAndServe(":1234", nil)
    if err != nil {
        log.Fatal(err)
    }
}
