package rest

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/MarySmirnova/pereval/internal/config"
	"github.com/MarySmirnova/pereval/internal/data"
	"github.com/MarySmirnova/pereval/pkg/storage/database"
)

type Worker struct {
	httpServer *http.Server
	storage    *database.Storage
}

func NewWorker(cfg config.REST, storage *database.Storage) *Worker {
	wr := &Worker{
		storage: storage,
	}

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
	var pereval *data.Pereval

	err := json.NewDecoder(r.Body).Decode(&pereval)
	if err != nil {
		log.WithError(err).Warn("unable to parse the request") // TODO: prepare public error description
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := wr.storage.SubmitData(pereval)
	if err != nil {
		log.WithError(err).Warn("unable to added data to DB") // TODO: prepare public error description
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	jsonID, err := json.Marshal(id)
	if err != nil {
		log.WithError(err).Warn("unable to parse response") // TODO: prepare public error description
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsonID)
}
