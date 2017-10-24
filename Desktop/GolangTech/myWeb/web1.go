package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":80", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("p1.html")
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		r.ParseForm()
		flag := false
		if len(r.Form["f"]) > 0 {
			flag = true
		}
		dirTree(w, r.Form["dir1"][0], flag)
	}
}
