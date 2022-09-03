package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/leonardo-amorim-rosa/go-rest-api/database"
	"github.com/leonardo-amorim-rosa/go-rest-api/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func BuscarTodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	var personalidades []models.Personalidade
	database.DB.Find(&personalidades) // busca todos os registros no banco, na tabela de personalidades e preenche a lista passada com parâmetro
	json.NewEncoder(w).Encode(personalidades)
}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// for _, personalidade := range models.Personalidades {
	// 	if strconv.Itoa(personalidade.Id) == id {
	// 		json.NewEncoder(w).Encode(personalidade)
	// 	}
	// }

	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)
	json.NewEncoder(w).Encode(personalidade)
}

func CriaUmaNovaPersonalidade(w http.ResponseWriter, r *http.Request) {
	var novaPersonalidade models.Personalidade
	json.NewDecoder(r.Body).Decode(&novaPersonalidade)
	database.DB.Create(&novaPersonalidade)
	json.NewEncoder(w).Encode(novaPersonalidade)
}

func DeletarUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personalidade models.Personalidade
	database.DB.Delete(&personalidade, id)
	json.NewEncoder(w).Encode(&personalidade)
}

func EditarPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)
	json.NewDecoder(r.Body).Decode(&personalidade)
	database.DB.Save(&personalidade)
	json.NewEncoder(w).Encode(personalidade)
}
