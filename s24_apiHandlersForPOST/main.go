package main

import (
	"fmt"
	"mod24/api"
	"mod24/data"
	"net/http"
	"text/template"
)

func handleTemplate(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("./template/index.template")
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	html.Execute(w, data.GetAll())
}

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/api/get-exhibitions", api.HandleGet)
	// server.HandleFunc("/api/make-exhibition", api.HandlePost)
	server.HandleFunc("POST /api/make-exhibition", api.HandlePost)
	// If you look here, when we are registering the function,we don't say that we want this only for one method (methods like POST,GET, DELETE).It's for all the methods.There are some libraries out there that can let you orthere are ways that you can hook also the method here, but for now it's like this.
	//NOTE :  In Go version 1.22, you can register for one METHOD. More Info : https://go.dev/blog/routing-enhancements
	//example : server.HandleFunc("POST /api/make-exhibition", api.HandlePost)

	server.HandleFunc("/template", handleTemplate)

	fs := http.FileServer(http.Dir("./public"))
	server.Handle("/", fs)

	err := http.ListenAndServe(":3335", server)
	if err != nil {
		fmt.Println(err)
	}
}
