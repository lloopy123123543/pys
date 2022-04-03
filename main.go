package main

import (
	"html/template"
	"net/http"
)

func home_page(page http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("tamplates/index.html")
	tmpl.Execute(page, nil)

}
func rules_page(page http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("tamplates/rules.html")
	tmpl.Execute(page, nil)
}
func handleRequest() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/", home_page)
	http.HandleFunc("/rules/", rules_page)
	http.ListenAndServe(":8080", nil)

}
func main() {

	handleRequest()

}
