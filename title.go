package main

import (
    "fmt"
    "strings"

    "golang.org/x/net/html"
)

func main() {

    HTMLString := `<!DOCTYPE html>
<html itemscope itemtype="http://schema.org/QAPage">

<head>

<title>go - Golang parse HTML, extract all content with &lt;body&gt; &lt;/body&gt; tags - Stack Overflow</title>
    <link rel="shortcut icon" href="//cdn.sstatic.net/Sites/stackoverflow/img/favicon.ico?v=4f32ecc8f43d">
    <link rel="apple-touch-icon image_src" href="//cdn.sstatic.net/Sites/stackoverflow/img/apple-touch-icon.png?v=c78bd457575a">
    <link rel="search" type="application/opensearchdescription+xml" title="Stack Overflow" href="/opensearch.xml">
    <meta name="twitter:card" content="summary">
    <meta name="twitter:domain" content="stackoverflow.com"/>
    <meta property="og:type" content="website" />
    <body><input type="password" /><img src="#" /></body>`

    title := getTitle(HTMLString)

    fmt.Println(title)
}

func getTitle(HTMLString string) (title string) {

    r := strings.NewReader(HTMLString)
    z := html.NewTokenizer(r)

    var i int
    for {
        tt := z.Next()

        i++
        if i > 100 { // Title should be one of the first tags
            return
        }

        switch {
        case tt == html.ErrorToken:
            fmt.Println("Error")
            // End of the document, we're done
            return
        case tt == html.SelfClosingTagToken:
            t := z.Token()

            // Check if the token is an <title> tag
            if t.Data != "input" {
                continue
            }

            // fmt.Printf("%+v\n%v\n%v\n%v\n", t, t, t.Type.String(), t.Attr)
            tt := z.Next()

            if tt == html.TextToken {
                t := z.Token()
                title = t.Data
                return
                // fmt.Printf("%+v\n%v\n", t, t.Data)
            }
        }
    }
}
