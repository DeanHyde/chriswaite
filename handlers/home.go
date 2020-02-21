package handlers

import (
	"html/template"
	"net/http"

	"github.com/DeanHyde/chriswaite/libhttp"
)

type Url struct {
	Path string
}

func GetHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	tmpl, err := template.ParseFiles("templates/dashboard.html.tmpl", "templates/home.html.tmpl")
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}

	urlData := Url{"home"}
	tmpl.ExecuteTemplate(w, "dashboard.html.tmpl", urlData)
}

func GetServices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	tmpl, err := template.ParseFiles("templates/dashboard.html.tmpl", "templates/services.html.tmpl")
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}

	urlData := Url{"services"}
	tmpl.ExecuteTemplate(w, "dashboard.html.tmpl", urlData)
}

func GetGallery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	tmpl, err := template.ParseFiles("templates/dashboard.html.tmpl", "templates/gallery.html.tmpl")
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}

	urlData := Url{"gallery"}
	tmpl.ExecuteTemplate(w, "dashboard.html.tmpl", urlData)
}

func GetFaq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	tmpl, err := template.ParseFiles("templates/dashboard.html.tmpl", "templates/faq.html.tmpl")
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}

	urlData := Url{"faq"}
	tmpl.ExecuteTemplate(w, "dashboard.html.tmpl", urlData)
}

func GetContactUs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	tmpl, err := template.ParseFiles("templates/dashboard.html.tmpl", "templates/home.html.tmpl")
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}

	urlData := Url{"contact"}
	tmpl.ExecuteTemplate(w, "dashboard.html.tmpl", urlData)
}
