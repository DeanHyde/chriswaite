package handlers

import (
	"fmt"
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

	tmpl, err := template.ParseFiles("templates/dashboard.html.tmpl", "templates/contact.html.tmpl")
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}

	urlData := Url{"contact"}
	tmpl.ExecuteTemplate(w, "dashboard.html.tmpl", urlData)
}

func PostSendMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	name := r.FormValue("name")
	email := r.FormValue("email")
	subject := r.FormValue("subject")
	message := r.FormValue("message")

	fmt.Fprintf(w, "Name: %q<br />", name)
	fmt.Fprintf(w, "Email: %q<br />", email)
	fmt.Fprintf(w, "Subject: %q<br />", subject)
	fmt.Fprintf(w, "Message: %q<br />", message)

	// tmpl, err := template.ParseFiles("templates/dashboard.html.tmpl", "templates/contact.html.tmpl")
	// if err != nil {
	// 	libhttp.HandleErrorJson(w, err)
	// 	return
	// }

	// urlData := Url{"contact"}
	// tmpl.ExecuteTemplate(w, "dashboard.html.tmpl", urlData)
}
