package main

import (
    "net/http"
    "log"
) 

func main() {
    fs_handler := http.FileServer(http.Dir("static_content/"))
    http.Handle("/", fs_handler)

    err := http.ListenAndServe(":1234", nil)
    if err != nil {
        log.Fatal(err)
    }
}
