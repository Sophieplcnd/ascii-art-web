package main

import (
	a "ascii-web-dockerize/ascii-art"
	"fmt"
	"html/template"
	"net/http"
)

var tmpl *template.Template

type ShowError struct {
	Code         int
	ErrorMessage string
}

var (
	BadReqError   = &ShowError{Code: http.StatusBadRequest, ErrorMessage: http.StatusText(http.StatusBadRequest)}
	IntServError  = &ShowError{Code: http.StatusInternalServerError, ErrorMessage: http.StatusText(http.StatusInternalServerError)}
	NotFoundError = &ShowError{Code: http.StatusNotFound, ErrorMessage: http.StatusText(http.StatusNotFound)}
)

func init() {
	tmpl = template.Must(template.ParseGlob("static/*.html"))
}

func main() {
	path := "static"
	fs := http.FileServer(http.Dir(path))
	/*instantiate a file server object by passing
	the directory where all our static files are placed */
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", Home)
	http.HandleFunc("/result", result)
	http.HandleFunc("/404", Err404)
	http.HandleFunc("/400", Err400)
	http.HandleFunc("/500", Err500)
	fmt.Printf("Fetching server...")
	http.ListenAndServe(":8234", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		tmpl.ExecuteTemplate(w, "404.html", nil)
		return
	}
	tmpl.ExecuteTemplate(w, "index.html", nil)

	
}

func result(w http.ResponseWriter, r *http.Request) {
	bannerop := r.FormValue("font")  //for the banner choice
	textinput := r.FormValue("text") // for the text input

	// check if characters are applicable
	runes := []rune(textinput)
	for i := range runes {
		if runes[i] > 127 {
			http.ServeFile(w, r, "static/400.html")
			return
		}
	}

	text ,error500 := a.Ascii(bannerop, textinput) //runs our ascii-code over the banner and text
	if error500{

		tmpl.ExecuteTemplate(w, "500.html", nil)

		return

	}
		

	 
	tmpl.ExecuteTemplate(w, "result.html", text)
}

func Err404(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "404.html", nil)
}

func Err500(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "500.html", nil)
}

func Err400(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "400.html", nil)
}
