// Handler que usaremos para los dispositivos
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/matialvarez7/television/internal/database"
)

func (apiCfg *apiConfig) handlerCreateDevice(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Codetag string `json:"codetag"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("error parsing JSON: %v", err))
	}

	if len(params.Codetag) > 11 {
		respondWithError(w, 400, fmt.Sprintln("El código del dispositivo no puede tener más de 12 caracteres"))
		return
	}

	device, err := apiCfg.DB.CreateDevice(r.Context(), database.CreateDeviceParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Codetag:   params.Codetag,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("couldn't create device: %v", err))
	}
	respondWithJSON(w, 201, device)
}
