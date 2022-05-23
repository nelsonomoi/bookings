package main

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter,r *http.Request)  {
	
	fmt.Fprintf(w,"This is the home page.")

}

func About(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintf(w,"Welcome to the about page")
}

func main() {

	http.HandleFunc("/",Home)

	http.HandleFunc("/about",About)

	http.ListenAndServe(":8080",nil)

}