package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/JSON")
	w.WriteHeader(statusCode)

	if data != nil {
		if erro := json.NewEncoder(w).Encode(data); erro != nil {
			log.Fatal(erro)
		}
	}
}

func Erro(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		Erro string `json: "error"`
	}{
		Erro: erro.Error(),
	})
}
