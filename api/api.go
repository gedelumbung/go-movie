package api

import (
	"github.com/gedelumbung/go-movie/config"
	"github.com/gedelumbung/go-movie/repository"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

type API struct {
	config *conf.Configuration
	web    *echo.Echo
	log    *log.Logger
	db     repository.Repository
}

func (a *API) ListenAndServe() {
	a.web.Logger.Fatal(a.web.Start(a.config.API.Host))
}

func NewAPI(config *conf.Configuration, db repository.Repository, log *log.Logger) *API {
	a := &API{config: config, web: echo.New(), db: db, log: log}
	a.registerRoutes()
	return a
}
