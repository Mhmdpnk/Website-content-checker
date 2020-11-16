package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func main() {

	HTMLString := `<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01//EN"
   "http://www.w3.org/TR/html4/strict.dtd">
<html itemscope itemtype="http://schema.org/QAPage">

<head>

<title>go - Golang parse HTML, extract all content with &lt;body&gt; &lt;/body&gt; tags - Stack Overflow</title>
    <link rel="shortcut icon" href="//cdn.sstatic.net/Sites/stackoverflow/img/favicon.ico?v=4f32ecc8f43d">
    <link rel="apple-touch-icon image_src" href="//cdn.sstatic.net/Sites/stackoverflow/img/apple-touch-icon.png?v=c78bd457575a">
    <link rel="search" type="application/opensearchdescription+xml" title="Stack Overflow" href="/opensearch.xml">
    <meta name="twitter:card" content="summary">
    <meta name="twitter:domain" content="stackoverflow.com"/>
    <meta property="og:type" content="website" /><input type="password" > <img src="" /><body></body>`


	var pageTag int = getHtmlTagContent(HTMLString, "DOCTYPE")
	fmt.Println(pageTag)
}

func getHtmlTagContent(HTMLString string, HtmlTag string) int{

    read := strings.NewReader(HTMLString)
    tokenizer := html.NewTokenizer(read)
    var counter int = 0

    for {
        tokenType := tokenizer.Next() // get the next token type

        if tokenType == html.ErrorToken{ // handling erro
            break // break
        }

        if tokenType == html.DoctypeToken{ // if it is a startTagToken
            token := tokenizer.Token() // Then get the token
			counter++ 
			fmt.Println("token:", token)
        }
    }
    return counter
}
