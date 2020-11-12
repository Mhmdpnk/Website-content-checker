package main

import ("fmt"
		"net/http"
		/*"io/ioutil"*/)

func main(){

	http.HandleFunc("/", index_handler) // Handling index page
	http.ListenAndServe(":8000", nil)

}

// w stands for writer, r stands for request
func index_handler(w http.ResponseWriter, r *http.Request){

	fmt.Fprintf(w, `
		<!DOCTYPE html>
<html>
<head>
	<title>Golang Website Checker</title>
	<link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css" integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">

</head>
<body>
	<div class="container">
		<div class="jumbotron text-center">
			<h1>Website Content Checker</h1>
			<p>Insert the website address and check the content of the page</p> 
	
			<form action="" class="form-row justify-content-center mt-4">
				<div class="col-5">
					<input type="text" class="form-control" id="web_url" placeholder="Website URL">
				</div>
			  	<button type="submit" name="check_url" class="btn btn-primary mb-2">Check</button>
			</form>
		</div>
	</div>
</body>
</html>
		`)

	/*resp, _ := http.Get("https://google.com")
	bytes, _ := ioutil.ReadAll(resp.Body)
	string_body := string(bytes)
	fmt.Println(string_body)
	resp.Body.Close()*/
}
