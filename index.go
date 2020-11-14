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
	var html_version, pageTitle = getUrlContent(requested_url)
	

	//---- Get External & Internal links
	var internalLink, externalLink = linksCounter(requested_url)
	

	data := struct{
		Url string
		HtmlVersion float64
		Title string
		Internal int
		External int
	}{
		Url: requested_url,
		HtmlVersion: html_version,
		Title: pageTitle,
		Internal: internalLink,
		External: externalLink,
	}
	
	
	tpl.ExecuteTemplate(w, "processor.html", data)
}


func getUrlContent(url string) (float64, string){
	
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

	// show the HTML code as a string %s
	//fmt.Printf("%s\n", html)
	

	//---- Get HTML Version --------

	html_str := string(html) // Convert []byte into a string
	var html_version float64 = getHtmlVersion(html_str)
	

	//---- Get HTML title --------
	var title string = getHtmlTagContent(html_str, "title")


	//---- Get Login Form --------
	//getLoginForm(html_str)


    //----------------------------
	
	return html_version, title


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

    r := strings.NewReader(HTMLString)
    z := html.NewTokenizer(r)

    var i int
    for {
        tt := z.Next()

        i++
        if i > 100 {
            return
        }

        switch {
        case tt == html.ErrorToken:
            return
        case tt == html.StartTagToken:
            t := z.Token()

            // Check the the HTML tag
            if t.Data != HtmlTag {
                continue
            }

            tt := z.Next()

            if tt == html.TextToken {
                t := z.Token()
                title = t.Data
                return
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