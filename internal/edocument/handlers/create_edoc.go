package handlers

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func (h Handlers) CreateEdoc(w http.ResponseWriter, r *http.Request) {

	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var data map[string]any

	jsonData := r.FormValue("json")

	err = json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if data != nil {
		err = h.repo.CreateEdoc(context.Background(), bson.M{"json": data})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	w.WriteHeader(http.StatusCreated)

}
