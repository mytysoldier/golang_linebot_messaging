package main

import (
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("templates/index.html")
	if err != nil {
		panic(err.Error())
	}
	if err = temp.Execute(w, nil); err != nil {
		panic(err.Error())
	}
}
