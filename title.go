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
    <body><h1>H11</h1><h1>H12</h1><h1>H12</h1><h1>H12</h1><h1>H12</h1><form><input type="password"></form></body>`




    if getTitle(HTMLString){
        fmt.Printf("true")
    }else{
        fmt.Printf("False")
    }

    //var inputType = "type=\"password\""

    if strings.Contains(HTMLString, "type=\"password\""){
        fmt.Println("Find it!")
    }

}

func getTitle(HTMLString string) bool  {

    read := strings.NewReader(HTMLString)
    tokenizer := html.NewTokenizer(read)
    var itHas bool

    for {
        tokenType := tokenizer.Next() // get the next token type

        if tokenType == html.ErrorToken{ // handling erro
            break // break
        }

        if tokenType == html.StartTagToken{ // if it is a startTagToken
            token := tokenizer.Token() // Then get the token

            if token.Data == "input"{ // if it matches with the name of HTML tag
                
                tokenType = tokenizer.Next() // get the HTML tag
                itHas = true
            }
        }
    }
    return itHas
}
