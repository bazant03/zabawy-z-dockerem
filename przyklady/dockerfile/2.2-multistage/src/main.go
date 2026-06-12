package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, req *http.Request) {
	res := "<h1 style=\"color:blue;text-align:center\">Teraz serwer w Go</h1>" + 
		   "<h2 style=\"color:blue;text-align:center\">Budowany jako multi-stage</h2>"
	fmt.Fprintf(w, res)
}

func main() {
	http.HandleFunc("/", handler);
	http.ListenAndServe(":8080", nil);
}