package main

import (
    "fmt"
    "strings"

    "golang.org/x/net/html"
)

type headingTags struct {
    Name string
    Value int
}

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
    <body><h1>H11</h1><h1>H12</h1><h1>H12</h1><h1>H12</h1><h1>H12</h1><img src="#" /></body>`

    title := getTitle(HTMLString)
    fmt.Println(title)


    /*tags := headingTags{}*/

    var tags headingTags
    fmt.Printf("%+v\n", tags) // Print with Variable Name

    
    for i := 0; i <= 6; i++{
        if i == 0{
            continue
        }
        
        concatenatedTag := fmt.Sprintf("h%dName", i)
        concatenatedValue := fmt.Sprintf("h%vValue", i)
        fmt.Println(concatenatedTag)
        fmt.Println(concatenatedValue)
        
        tags.Name = concatenatedTag
        tags.Value = i


    }

    fmt.Printf("%+v\n", tags) // Print with Variable Name
/*
        emp := Employee{Id:1200, Name: "Mark Taylor"}
    
    fmt.Printf("%v\n",emp)  // Without Variable Name
    fmt.Printf("%d\n",emp.Id)
    fmt.Printf("%s\n",emp.Name)*/

    //fmt.Println("AFTER: ", headingTags)

}

func getTitle(HTMLString string) int  {

    read := strings.NewReader(HTMLString)
    tokenizer := html.NewTokenizer(read)
    var counter int = 0

    for {
        tokenType := tokenizer.Next() // get the next token type

        if tokenType == html.ErrorToken{ // handling erro
            break
        }

        if tokenType == html.StartTagToken{ // if it is a startTagToken
            token := tokenizer.Token() // Then get the token

            if token.Data == "h1"{ // if it matches with the name of HTML tag
                
                tokenType = tokenizer.Next() // get the HTML tag
                counter++ 
            } 
        }

    }
    return counter
}
