package main

import ("fmt"
		"net/http"
		"io/ioutil")

func main(){
	resp, _ := http.Get("https://google.com")
	bytes, _ := ioutil.ReadAll(resp.Body)
	string_body := string(bytes)
	fmt.Println(string_body)
	resp.Body.Close()
}