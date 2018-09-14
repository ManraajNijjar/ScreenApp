package main

import (
	"html/template"
	"log"
	"net/http"
)

type pageVariables struct {
}

func main() {
	//http.HandleFunc("/", serveIndex)
	//http.HandleFunc("/app.js", serveApp)
	//http.ListenAndServe(":8080", nil)

}

func serveIndex(writer http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("index.html")

	if err != nil {
		log.Fatalf("Failed to parse")
	}

	err = t.Execute(writer, pageVariables{})

	if err != nil {
		log.Fatalf("Failed to execute")
	}
}

func serveApp(writer http.ResponseWriter, req *http.Request) {
	http.ServeFile(writer, req, "app.js")
}
