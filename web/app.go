package web

import (
	"App-CloudBase-mcdull-mall/env"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type App struct {
	Conf   *env.Conf
	Router *mux.Router
}

func (app *App) Initialize() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	app.Conf = env.LoadConf()
	app.Router = NewRouter(context.Background())
}

func (app *App) Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	log.Printf("App-CloudBase-mcdull-mall runs on :%s\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), app.Router)
	if err != nil {
		panic(err)
	}
}
