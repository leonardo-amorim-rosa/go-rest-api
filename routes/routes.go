package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/leonardo-amorim-rosa/go-rest-api/controllers"
	"github.com/leonardo-amorim-rosa/go-rest-api/middlewares"
)

func HandleRequests() {
	r := mux.NewRouter()
	r.Use(middlewares.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.BuscarTodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.CriaUmaNovaPersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletarUmaPersonalidade).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}", controllers.EditarPersonalidade).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
