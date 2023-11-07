package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/google/uuid"
)

type CreateCampaignRequest struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

type CreateCampaignResponse struct {
	ID uuid.UUID `json:"id"`
}

var m sync.Map

func main() {

	http.HandleFunc("/create-campaign", func(writer http.ResponseWriter, request *http.Request) {
		var (
			req  CreateCampaignRequest
			resp CreateCampaignResponse
		)

		if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		id, _ := m.LoadOrStore(req.Name, uuid.New())

		resp.ID = id.(uuid.UUID)
		req.ID = resp.ID

		log.Printf("received request to create campaign, request: %+v", req)

		writer.Header().Add("Content-Type", "application/json")
		err := json.NewEncoder(writer).Encode(resp)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	})

	log.Println("adserver is running")
	log.Fatal(http.ListenAndServe(":9091", nil))
}
