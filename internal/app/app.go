package app

import (
	"hw1/configs"
	"hw1/internal/pkg/endpoints"
	"hw1/internal/pkg/repository"
	"hw1/internal/pkg/services"
	"log"
	"net/http"
)

type App struct {
	end  *endpoints.Endpoint
	serv *services.Service
	rep  *repository.Rep
}

func NewApp() (*App, error) {
	a := &App{}
	a.rep = repository.NewRep()
	a.serv = services.NewService(a.rep)
	a.end = endpoints.NewEndpoint(a.serv)
	http.HandleFunc("/info", a.end.InfoHandler)
	http.HandleFunc("/info/currency", a.end.CurrencyHandler)
	return a, nil
}

func (a *App) Run() error {
	log.Println("server running")
	port := configs.GetEnv("PORT")
	if port[0] != ':' {
		port = ":" + port
	}
	err := http.ListenAndServe(port, nil)
	if err != nil {
		return err
	}
	return nil
}
