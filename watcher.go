package main

import (
     "fmt"
     "net/http"
)

type Result struct {
     Url string
}

func FetchPage(url string) []Result {
    res, err := http.Get(url)
    if err != nil {
        fmt.Println(err)
    }
    defer res.Body.Close()
    return nil
}

func main() {
    url := ""
}
