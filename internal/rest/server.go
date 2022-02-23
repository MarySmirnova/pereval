package rest

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/MarySmirnova/pereval/internal/config"
)

type Worker struct {
	httpServer *http.Server
}

func NewWorker(cfg config.REST) *Worker {
	wr := &Worker{}

	handler := mux.NewRouter()
	handler.Name("submit_data").Methods(http.MethodPost).Path("/submitData").HandlerFunc(wr.submitData)

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

func (wr *Worker) submitData(w http.ResponseWriter, r *http.Request) {

}
