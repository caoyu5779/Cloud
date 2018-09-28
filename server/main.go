package main

import (
"net/http"
"io"
)

func firstPage(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w, "<h1> hello , this is my first page! </h1>")
}

func main()  {
	http.HandleFunc("/", firstPage)
	http.ListenAndServe(":7878", nil)
}