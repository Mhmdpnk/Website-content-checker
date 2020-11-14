package main

import (
		"fmt"
		"net/http"
		"html/template"
		"io/ioutil"
		"strings"
		"golang.org/x/net/html"
		)

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseGlob("templates/*.html"))
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
	email := r.FormValue("email")


	// Get the content of the URL and the return values
	var html_version, pageTitle = getUrlContent(requested_url)
	

	data := struct{
		Email string
		Url string
		HtmlVersion float64
		Title string
	}{
		Email: email,
		Url: requested_url,
		HtmlVersion: html_version,
		Title: pageTitle,
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
	var title string = getHtmlTitle(html_str)

    fmt.Println(title)


    //----------------------------
	
	return html_version, title


}//getUrlContent



// To get the HTML version
func getHtmlVersion(html string) float64{
	
	var url_length int = len(html) // Find the length of HTML
	var html_version float64 = 0
    
	//fmt.Println(strings.Contains(html, "doctype")) 


	for i := 0; i <= url_length; i++{

		if strings.Contains(html, "XHTML 1.0"){
			html_version = 1.0
		} else if strings.Contains(html, "XHTML 1.1"){
			html_version = 1.1
		} else if strings.Contains(html, "XHTML 4.01"){
			html_version = 4.01
		} else {
			html_version = 5
		}

	}//for

	return html_version
}

// To get the HTML title
func getHtmlTitle(HTMLString string) (title string) {

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
            // End of the document, we're done
            return
        case tt == html.StartTagToken:
            t := z.Token()

            // Check if the token is an <title> tag
            if t.Data != "title" {
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