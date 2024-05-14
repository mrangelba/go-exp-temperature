package rest

import (
	"log"
	"net/http"

	"github.com/mrangelba/go-exp-temperature/internal/di"

	"github.com/go-chi/chi"
)

func Initialize() {
	webController := di.ConfigWebController()

	r := chi.NewRouter()

	r.Get("/cep/{cep}", webController.Get)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", r)
}
