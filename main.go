package returnhandling

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(rw http.ResponseWriter, statusCode int, data interface{}) {

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(statusCode)

	if data != nil {
		if err := json.NewEncoder(rw).Encode(data); err != nil {
			log.Fatal(err)
		}
	}
}

func Erro(rw http.ResponseWriter, statusCode int, erro error) {
	JSON(rw, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}
