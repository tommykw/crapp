package main

import (
     "fmt"
     "strings"
     "time"
     "github.com/PuerkitoBio/goquery"
)

// google play url
const APP_STORE_URL = ""

func getVersion() string {
    ver := ""
    doc, _ := goquery.NewDocument(APP_STORE_URL)
    doc.Find("div").Each(func(_ int, s *goquery.Selection) {
        itemprop, _ := s.Attr("itemprop")
        classname, _ := s.Attr("class")
        if itemprop == "softwareVersion" && classname == "content" {
            ver = strings.TrimSpace(s.Text())
        }
    })

    return ver
}

func main() {
    ver := getVersion()
    if ver == "" {
        fmt.Println("cannot get current version")
        return
    }
    fmt.Printf("current ver %s\n", ver)
   
    for { 
        current_ver := getVersion()
        if ver != current_ver {
            fmt.Printf("updated to ver %s\n", current_ver)
            break
        }
        fmt.Printf("current ver %s\n", current_ver)
        time.Sleep(60 * 1000 * time.Millisecond)
    }
}
