package main

import (
	"fmt"
	"net/http"
	"text/template"

	"web_mustapha/funnc_work"
)

func rootpage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404: Page not found.", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "405: Method not allowed.", http.StatusNotFound)
		return
	}
	t, err := template.ParseFiles("page/index.html")
	if err != nil {
		http.Error(w, "500: Internal Server Error.", http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, "500: Internal Server Error.", http.StatusInternalServerError)
		return
	}
}

func result_page(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/result" {
		http.Error(w, "404: page not found.", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "405: Method not allowed.", http.StatusMethodNotAllowed)
		return
	}
	first_number := r.PostFormValue("nbr1")
	secand_number := r.PostFormValue("nbr2")
	result, err := funnc_work.Add(first_number, secand_number)
	if err != nil {
		http.Error(w, "500: internal server error.", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Result: %d", result)
}

func main() {
	http.HandleFunc("/", rootpage)
	http.HandleFunc("/result", result_page)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
