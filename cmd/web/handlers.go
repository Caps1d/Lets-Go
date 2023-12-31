package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"text/template"

	// "html/template"
	"net/http"
	"strconv"

	"github.com/Caps1d/Lets-Go/internal/models"
)

func (app *applicaiton) home(w http.ResponseWriter, r *http.Request) {
	// since this handles a subtree path, it will match any requests
	// that start with "/", hence we need to add a check to handle unwanted behaviour
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	app.infoLog.Println("Home endpoint reached")

	snippets, err := app.snippets.Latest()

	if err != nil {
		app.serverError(w, err)
		return
	}

	for _, snippet := range snippets {
		fmt.Fprintf(w, "%+v\n", snippet)
	}

	//
	// files := []string{
	// 	"./ui/html/pages/base.tmpl.html",
	// 	"./ui/html/pages/home.tmpl.html",
	// 	"./ui/html/partials/nav.tmpl.html",
	// }
	//
	// // go can use ParseFiles to read the template file into a template set
	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// 	app.serverError(w, err)
	// }
	//
	// // we use ExecuteTemplate to write the content of the "base" template
	// // from the template set into the response body. We have 4 templates in the template set:
	// // "base", "title", "main", "nav" where "base" invokes the other 3
	// // The last parameter represents any dynamic content
	// // that we would like to pass to the template - will use it later
	// err = ts.ExecuteTemplate(w, "base", nil)
	// if err != nil {
	// 	app.serverError(w, err)
	// }

	// js, err := json.Marshal(snippets)
	//
	// if err != nil {
	// 	app.serverError(w, err)
	// }
	//
	// js = append(js, '\n')
	//
	// w.Header().Set("Content-Type", "application/json")
	// w.Write(js)
}

func (app *applicaiton) snippetView(w http.ResponseWriter, r *http.Request) {

	// getting url query string parameters
	// we also want to make sure that the id is an int
	// we parse the str and convert it to an int
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		w.Header().Set("allow-id", ">=1")
		app.notFound(w)
		return
	}

	s, err := app.snippets.Get(id)

	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	files := []string{
		"./ui/html/pages/base.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/view.tmpl.html",
	}

	// Parse the template files...
	ts, err := template.ParseFiles(files...)

	if err != nil {
		app.serverError(w, err)
		return
	}

	// And then execute them. Notice how we are passing in the snippet // data (a models.Snippet struct) as the final parameter?
	err = ts.ExecuteTemplate(w, "base", s)
	if err != nil {
		app.serverError(w, err)
	}

	js, err := json.Marshal(s)

	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("Displaying snippet with ID %d...", id)

	js = append(js, ' ')

	fmt.Fprintf(w, "Displaying snippet with ID %d...", id)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (app *applicaiton) snippetCreate(w http.ResponseWriter, r *http.Request) {
	// we can write "POST" or use constants from net/http
	if r.Method != http.MethodPost {
		// adds a custom header Allow: POST
		w.Header().Set("Allow", http.MethodPost)
		// w.WriteHeader(http.StatusMethodNotAllowed)
		// w.Write([]byte("Request Method Not Allowed"))
		// A shortcut to this is the http.Error helper func -> using a custom helper now, defined in helpers.go
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	// snippet struct
	var s models.Snippet

	// decode the requests body into our post struct declared as p
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		app.serverError(w, err)
		return
	}

	app.infoLog.Printf("Received post: %v", s)

	lastId, err := app.snippets.Insert(s.Title, s.Content, s.ExpiresInt)

	if err != nil {
		app.serverError(w, err)
	}
	app.infoLog.Printf("New snippet created with id %d", lastId)

	w.Write([]byte("Create a new snippet..."))
}
