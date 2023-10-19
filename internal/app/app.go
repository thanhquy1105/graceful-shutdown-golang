package app

import (
	"context"
	"log"
	"net/http"

	"github.com/thanhquy1105/graceful-shutdown-golang/internal/router"
)

type (
	app struct {
		Server *http.Server
	}

	App interface {
		Start() error
		Stop(ctx context.Context) error
	}
)

const (
	ADDR = ":8080"
)

func New() App {
	router := router.New()

	httpServer := &http.Server{
		Addr:    ADDR,
		Handler: router,
	}
	return app{
		Server: httpServer,
	}
}

func (a app) Start() error {
	log.Printf("Server is listening at %s", ADDR)
	return a.Server.ListenAndServe()
}

func (a app) Stop(ctx context.Context) error {
	return a.Server.Shutdown(ctx)
}
