package main

import (
		"fmt"
		"log"
		"net/http"
		"html/template"
		"io/ioutil"
		"strings"
		"net/url"
		"golang.org/x/net/html"
		)

var tpl *template.Template

type ReturnFuncData struct{ 
	HtmlVersion float64 
	pageTitle string
	h1Count int 
}

var headingTags[7] int

func init(){
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

type scrapeDataStore struct {
	Internal int `json:"internal"`
	External int `json:"external"`
}

func isInternal(parsedLink *url.URL, siteURL *url.URL, link string) bool {
	return parsedLink.Host == siteURL.Host || strings.Index(link, "#") == 0 || len(parsedLink.Host) == 0
}

func main(){

    fmt.Println("Now Listening on 8000")
    http.HandleFunc("/", serveFiles)
    http.HandleFunc("/process", processor)
    http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
    http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
    
    http.ListenAndServe(":8000", nil)
}


func serveFiles(w http.ResponseWriter, r *http.Request){

	tpl.ExecuteTemplate(w, "index.html", nil)
}

func processor(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST"{
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}	

	requested_url := r.FormValue("web_url")

	// Get the content of the URL and the return values
	funcData := getUrlContent(requested_url)

	//---- Get External & Internal links
	var internalLink, externalLink = linksCounter(requested_url)


	//---- Get HTML Heding tags
	fmt.Printf("%v", headingTags)
	fmt.Println("")

	headValues := make([]int, 7)

	// Copy headingTags array values to headvalues type
	for i := 0; i <= len(headingTags); i++{
	    copy(headValues, headingTags[:i])
	}

	// Printing out the result
	for index,element := range headValues{
        fmt.Println(index,"=>",element)
    }



	data := struct{
		Url string
		HtmlVersion float64
		Title string
		Heading1 int
		Internal int
		External int
		HeadValues []int
	}{
		Url: requested_url,
		HtmlVersion: funcData.HtmlVersion,
		Title: funcData.pageTitle,
		Heading1: funcData.h1Count,
		Internal: internalLink,
		External: externalLink,
		HeadValues: headValues,
	}
	
	
	tpl.ExecuteTemplate(w, "processor.html", data)
}


func getUrlContent(url string) ReturnFuncData{
	
	// Printing the HTML of URL
	fmt.Printf("HTML code of %s ...\n", url)
	resp, err := http.Get(url)

	// handle the error if there is one
	if err != nil {
		panic(err)
	}

	// do this now so it won't be forgotten
	defer resp.Body.Close()

	// reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)


	// Copy data from the response to standard output
	/*html, err := io.Copy(os.Stdout, resp.Body) //use package "io" and "os"*/

	if err != nil {
		panic(err)
	}

	//---- Get HTML Version --------

	html_str := string(html) // Convert []byte into a string
	var html_version float64 = getHtmlVersion(html_str)
	

	//---- Get HTML title --------
	var pageTitle string = getHtmlTagContent(html_str, "title")


	//---- Count the HTML tag --------
	var h1Count int = htmlTagCounter(html_str, "h1")
	
	
	for i := 0; i <= 6; i++{
		concatenatedTag := fmt.Sprintf("h%d", i)

    	if i == 0 { 
    		continue
    	}

		headingTags[i] = htmlTagCounter(html_str, concatenatedTag)
	}






	//---- Get Login Form --------
	//getLoginForm(html_str)

    //----------------------------

	
	return ReturnFuncData{html_version, pageTitle, h1Count}


}//getUrlContent



// To get the HTML version
func getHtmlVersion(html string) float64{
	
	var html_version float64 = 0

	if strings.Contains(html, "XHTML 1.0"){
		html_version = 1.0
	} else if strings.Contains(html, "XHTML 1.1"){
		html_version = 1.1
	} else if strings.Contains(html, "HTML 2.0"){
		html_version = 2.0
	} else if strings.Contains(html, "HTML 3.2"){
		html_version = 3.2
	} else if strings.Contains(html, "XHTML 4.01"){
		html_version = 4.01
	} else {
		html_version = 5
	}

	return html_version
}

// To get the HTML tag text content - Inside HTML tag
func getHtmlTagContent(HTMLString string, HtmlTag string) (title string){
  
    read := strings.NewReader(HTMLString)
    tokenizer := html.NewTokenizer(read)

    for {
        tokenType := tokenizer.Next() // get the next token type

        if tokenType == html.ErrorToken{ // handling erro
            return
        }

        if tokenType == html.StartTagToken{ // if it is a startTagToken
            token := tokenizer.Token() // Then get the token

            if token.Data == HtmlTag{ // if it matches with the name of HTML tag
                tokenType = tokenizer.Next() // get the HTML tag

                if tokenType == html.TextToken{ //just make sure it's actually a text token
                    token := tokenizer.Token() // get the token
                    title = token.Data // return the data
                    return
                }   
            }
        }
    }
}

func linksCounter(urlIn  string) (int, int){

	siteURL, parseErr := url.Parse(urlIn)

	if parseErr != nil {
		log.Fatalln(parseErr)
	}

	resp, err := http.Get(urlIn)
	if err != nil {
		log.Fatalln(err)
	}

	scrapeData := &scrapeDataStore{}

	tokenizer := html.NewTokenizer(resp.Body)
	end := false
	for {
		tt := tokenizer.Next()
		switch {
		case tt == html.StartTagToken:
			// fmt.Println(tt)
			token := tokenizer.Token()
			switch token.Data {
			case "a":

				for _, attr := range token.Attr {

					if attr.Key == "href" {
						link := attr.Val

						parsedLink, parseLinkErr := url.Parse(link)

						if parseLinkErr == nil {
							if isInternal(parsedLink, siteURL, link) {
								scrapeData.Internal++
							} else {
								scrapeData.External++
							}
						}

						if parseLinkErr != nil {
							fmt.Println("Can't parse: " + token.Data)
						}
					}
				}
				break
			}
		case tt == html.ErrorToken:
			end = true
			break
		}
		if end {
			break
		}
	}


	return scrapeData.Internal, scrapeData.External
}

func getLoginForm(html string){

	//var html_version float64 = 0
    
    if(strings.Contains(html, "login") || strings.Contains(html, "Login")){
 
    	fmt.Println("It has login text.")

    	if(strings.Contains(html, "form")){
			fmt.Println("It has form!")
		}

    }
	return
}

func htmlTagCounter(HTMLString string, HTMLTag string) int{

    read := strings.NewReader(HTMLString)
    tokenizer := html.NewTokenizer(read)
    var counter int = 0

    for {
        tokenType := tokenizer.Next() // get the next token type

        if tokenType == html.ErrorToken{ // handling erro
            break // break
        }

        if tokenType == html.StartTagToken{ // if it is a startTagToken
            token := tokenizer.Token() // Then get the token

            if token.Data == HTMLTag{ // if it matches with the name of HTML tag
                
                tokenType = tokenizer.Next() // get the HTML tag
                counter++ 	// increment the counter
            } 
        }
    }
    return counter
}