package internal

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/chatex-com/process-manager"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/MarySmirnova/pereval/internal/config"
	"github.com/MarySmirnova/pereval/internal/rest"
	"github.com/MarySmirnova/pereval/pkg/storage/database"
)

type Application struct {
	sigChan <-chan os.Signal
	cfg     config.Application
	manager *process.Manager

	db *pgxpool.Pool
}

func NewApplication(cfg config.Application) (*Application, error) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	app := &Application{
		sigChan: sigChan,
		cfg:     cfg,
		manager: process.NewManager(),
	}

	if err := app.bootstrap(); err != nil {
		return nil, err
	}

	return app, nil
}

func (a *Application) bootstrap() error {
	//init database
	if err := a.initDatabase(); err != nil {
		return err
	}

	//init workers
	if err := a.bootstrapRestWorker(); err != nil {
		return err
	}

	return nil
}

func (a *Application) initDatabase() error {
	db, err := database.NewDBpg(a.cfg.Postgres)
	if err != nil {
		return err
	}

	a.db = db.GetPGXpool()
	return nil
}

func (a *Application) bootstrapRestWorker() error {
	worker := rest.NewWorker(a.cfg.REST)
	a.manager.AddWorker(process.NewServerWorker("httpServer", worker.GetHTTPServer()))

	return nil
}

func (a *Application) Run() {
	a.manager.StartAll()
	a.registerShutdown()
}

func (a *Application) registerShutdown() {
	defer a.db.Close()

	go func() {
		<-a.sigChan

		a.manager.StopAll()
	}()

	a.manager.AwaitAll()
}
