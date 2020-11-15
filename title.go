package main

import (
    "fmt"        
    "log"
    "github.com/PuerkitoBio/goquery"
)

func main() {

    doc, err := goquery.NewDocument("https://en.wikipedia.org/wiki/Example.com")
    
    if err != nil {
        log.Fatal(err)
    }
    
    doc.Find("a[href]").Each(func(index int, item *goquery.Selection) {
        href, _ := item.Attr("href")
        fmt.Printf("link: %s - anchor text: %s\n", href, item.Text())
        
    })
    
}