package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Daniel-Vinicius/go-ports-adapters/adapters/web/handler"
	"github.com/Daniel-Vinicius/go-ports-adapters/application"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

type Webserver struct {
	Service application.ProductServiceInterface
}

func MakeNewWebserver() *Webserver {
	return &Webserver{}
}

func (webserver Webserver) Serve() {
	router := mux.NewRouter()
	negroni := negroni.New(negroni.NewLogger())

	handler.MakeProductHandlers(router, negroni, webserver.Service)
	http.Handle("/", router)

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":9000",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log: ", log.Lshortfile),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
