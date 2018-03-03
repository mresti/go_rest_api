package app

import (
	"app/api"
	"flag"
	"fmt"
	"log"
	"net/http"
)

type App struct {
	Addr     string
	Port     int
	server   *http.Server
	handlers *http.ServeMux
}

func (a *App) readCmd() {
	// Read arg from command line for the port.
	portPtr := flag.Int("port", 9000, "a int")
	flag.Parse()
	a.Port = int(*portPtr)
}

func (a *App) Initialize() {
	a.readCmd()

	a.Addr = fmt.Sprintf(":%v", a.Port)
	a.handlers = api.Handlers()

	log.Printf("Now listening on %s\n", a.Addr)
	a.server = &http.Server{Handler: a.handlers, Addr: a.Addr}
}

func (a *App) Run() {
	fmt.Println("Server starting")
	log.Fatal(a.server.ListenAndServe())
}
