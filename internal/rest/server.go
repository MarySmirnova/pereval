package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/MarySmirnova/pereval/internal/config"
	"github.com/MarySmirnova/pereval/pkg/storage/models"
)

type Worker struct {
	httpServer *http.Server
}

func NewWorker(cfg config.REST) *Worker {
	wr := &Worker{}

	handler := mux.NewRouter()
	handler.Name("submit_data").Methods(http.MethodPost).Path("/submitData").HandlerFunc(wr.submitDataHandler)

	wr.httpServer = &http.Server{
		Addr:         cfg.Listen,
		Handler:      handler,
		WriteTimeout: cfg.WriteTimeout,
		ReadTimeout:  cfg.ReadTimeout,
	}

	return wr
}

func (wr *Worker) GetHTTPServer() *http.Server {
	return wr.httpServer
}

func (wr *Worker) submitDataHandler(w http.ResponseWriter, r *http.Request) {
	var pereval models.Pereval

	err := json.NewDecoder(r.Body).Decode(&pereval)
	if err != nil {
		log.WithError(err).Warn("unable to parse the request") // TODO: prepare public error description
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(pereval)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
