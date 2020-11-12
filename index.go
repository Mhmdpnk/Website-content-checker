package main

import (
		"fmt"
		"net/http"
		"html/template"
		)

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main(){

    fmt.Println("Now Listening on 8000")
    http.HandleFunc("/", serveFiles)
    http.HandleFunc("/process", processor)
    http.ListenAndServe(":8000", nil)
}


func serveFiles(w http.ResponseWriter, r *http.Request){

	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func processor(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST"{
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}	

	requested_url := r.FormValue("web_url")
	name := r.FormValue("full_name")

	fmt.Println(requested_url)
	fmt.Println(name)


	/*d := struct{
		url string
		fName string
	}{
		url: requested_url,
		fName: name,
	}*/
	
	
	tpl.ExecuteTemplate(w, "/processor.gohtml", requested_url)
}
