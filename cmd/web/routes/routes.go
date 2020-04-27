package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wesleimp/random-names-generator/pkg/generators/animals"
	"github.com/wesleimp/random-names-generator/pkg/generators/fruits"
	"github.com/wesleimp/random-names-generator/pkg/generators/heroes"
	"github.com/wesleimp/random-names-generator/pkg/generators/vegetables"
)

// Setup all routes
func Setup(r *mux.Router) {
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.HandleFunc("/api/heroes", func(w http.ResponseWriter, r *http.Request) {
		name := heroes.Generate()
		w.Write([]byte(name))
	})

	r.HandleFunc("/api/animals", func(w http.ResponseWriter, r *http.Request) {
		name := animals.Generate()
		w.Write([]byte(name))
	})

	r.HandleFunc("/api/fruits", func(w http.ResponseWriter, r *http.Request) {
		name := fruits.Generate()
		w.Write([]byte(name))
	})

	r.HandleFunc("/api/vegetables", func(w http.ResponseWriter, r *http.Request) {
		name := vegetables.Generate()
		w.Write([]byte(name))
	})
}
