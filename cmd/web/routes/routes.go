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

	api := r.PathPrefix("/api").Subrouter()
	{

		api.HandleFunc("/heroes", func(w http.ResponseWriter, r *http.Request) {
			name := heroes.Generate()
			w.Write([]byte(name))
		})

		api.HandleFunc("/animals", func(w http.ResponseWriter, r *http.Request) {
			name := animals.Generate()
			w.Write([]byte(name))
		})

		api.HandleFunc("/fruits", func(w http.ResponseWriter, r *http.Request) {
			name := fruits.Generate()
			w.Write([]byte(name))
		})

		api.HandleFunc("/vegetables", func(w http.ResponseWriter, r *http.Request) {
			name := vegetables.Generate()
			w.Write([]byte(name))
		})
	}
}
