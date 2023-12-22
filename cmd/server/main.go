package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/go-chi/chi/v5"
	root "github.com/marvinarlt/go-x-astro"
	"github.com/marvinarlt/go-x-astro/astro"
)

type Data struct {
	Person Person
}

type Person struct {
	Id        int
	FirstName string
	LastName  string
	Age       uint
	Email     string
}

var people = []Person{
	{
		Id:        1,
		FirstName: "Gannon",
		LastName:  "Olson",
		Age:       44,
		Email:     "et.libero.proin@aol.ca",
	},
	{
		Id:        2,
		FirstName: "Camilla",
		LastName:  "Ortiz",
		Age:       24,
		Email:     "ante@hotmail.edu",
	},
	{
		Id:        3,
		FirstName: "Maggy",
		LastName:  "Mcdaniel",
		Age:       36,
		Email:     "suspendisse.aliquet.molestie@aol.couk",
	},
}

func main() {
	router := chi.NewRouter()
	client := astro.New(&root.DistFileSystem)

	if err := client.LoadTemplates("."); err != nil {
		log.Fatalf("could not load templates: %v\n", err)
	}

	if err := client.ParseTemplates(createTemplateParsingHandler(router)); err != nil {
		log.Fatalf("could not parse template: %v\n", err)
	}

	log.Println("starting web server on port 1337")

	if err := http.ListenAndServe(":1337", router); err != nil {
		log.Fatalf("could not start web server: %v\n", err)
	}
}

func createTemplateParsingHandler(router *chi.Mux) astro.ParseHandlerFunc {
	return func(t *astro.Template, tmpl *template.Template) error {
		log.Printf("adding route: GET %s\n", t.Pattern)
		router.Get(t.Pattern, createTemplateRequestHandler(tmpl))

		return nil
	}
}

func createTemplateRequestHandler(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data Data

		idParam := chi.URLParam(r, "id")

		if idParam != "" {
			id, err := strconv.Atoi(idParam)

			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte(err.Error()))
				return
			}

			data.Person, err = findPerson(id)

			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("person not found"))
				return
			}
		}

		if err := tmpl.Execute(w, data); err != nil {
			w.Write([]byte(err.Error()))
		}
	}
}

func findPerson(id int) (Person, error) {
	for _, person := range people {
		if person.Id == id {
			return person, nil
		}
	}

	return Person{}, fmt.Errorf("could not find person")
}
