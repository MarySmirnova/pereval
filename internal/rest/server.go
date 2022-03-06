package rest

import (
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/MarySmirnova/pereval/internal/config"
	"github.com/MarySmirnova/pereval/pkg/storage/database"

	_ "github.com/MarySmirnova/pereval/docs"
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
	handler.Name("post_data").Methods(http.MethodPost).Path("/submitData").HandlerFunc(wr.postDataHandler)
	handler.Name("get_status").Methods(http.MethodGet).Path("/submitData/{id:[0-9]+}/status").HandlerFunc(wr.getStatusHandler)
	handler.Name("change_data").Methods(http.MethodPut).Path("/submitData/{id:[0-9]+}").HandlerFunc(wr.changeDataHandler)
	handler.Name("get_all_data").Methods(http.MethodGet).Path("/submitData/").HandlerFunc(wr.getAllDataHandler)
	handler.Name("get_data").Methods(http.MethodGet).Path("/submitData/{id:[0-9]+}").HandlerFunc(wr.getDataHandler)

	handler.PathPrefix("/swagger").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://178.154.220.204:8080/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	))

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
