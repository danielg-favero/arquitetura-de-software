package server

import (
	"github.com/codeedu/go-hexagonal/application"
	"net/http"
	"os"
	"time"
	"log"
	"github.com/codeedu/go-hexagonal/adapters/web/handler"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

type WebServer struct {
	Service application.ProductServiceInterface
}

func NewWebServer() *WebServer {
	return &WebServer{}
}

func (webServer WebServer) Serve() {
	r := mux.NewRouter() // Fazer roteamento
	n := negroni.New(negroni.NewLogger()) // Middleware

	handler.MakeProductHandlers(r, n, webServer.Service)
	http.Handle("/", r)

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr: ":9000",
		Handler: http.DefaultServeMux,
		ErrorLog: log.New(os.Stderr, "log: ", log.Lshortfile),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}