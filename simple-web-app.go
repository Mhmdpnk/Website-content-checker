package main

import ("fmt"
	"net/http"
)

func main(){

	http.HandleFunc("/", index_handler) // Handling index page
	http.ListenAndServe(":8000", nil)
}

// w stands for writer, r stands for request
func index_handler(w http.ResponseWriter, r *http.Request){

	fmt.Fprintf(w, "Whoa, Go is neat!")
}
