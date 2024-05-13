package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"regexp"
	"text/template"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600) // Escribe en un archivo con permisos 0600 que indica que se crea con permisos de lectura y escritura.
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// Expresión regular para validar la URL.
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)                         // Si no encuentra la página, responde con un mensaje de error.
		return "", errors.New("Invalid Page Title") // Retorna un error.
	}
	return m[2], nil
}

// Capturar la petición del cliente y responder con un mensaje.
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r) // Captura el título de la página.
	if err != nil {
		return
	}
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound) // Redirecciona a la página de edición.
		return
	}
	renderTemplate(w, "view", p) // Renderiza el template en el navegador.
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r) // Captura el título de la página.
	if err != nil {
		return
	}
	p, err := loadPage(title) // Carga la página con el título capturado.
	if err != nil {
		p = &Page{Title: title} // Si no existe la página, crea una nueva.
	}
	renderTemplate(w, "edit", p) // Renderiza el template edit.html.
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r) // Captura el título de la página.
	if err != nil {
		return
	}
	body := r.FormValue("body") // Captura el contenido de la página.
	p := &Page{Title: title, Body: []byte(body)}
	err = p.save() // Guarda la página.
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // Si hay un error, responde con un mensaje de error.
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound) // Redirecciona a la página de visualización.
}

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

// Renderiza el template en el navegador.
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p) // Ejecuta el template en el navegador.
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // Si hay un error, responde con un mensaje de error.
		return
	}
}

func main() {
	// p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	// p1.save()
	// p2, _ := loadPage("TestPage")
	// fmt.Println(string(p2.Body))

	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	// Levanto el servidor en el puerto 8080.
	log.Fatal(http.ListenAndServe(":8080", nil))
}
