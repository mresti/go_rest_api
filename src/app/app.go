package app

import (
	"app/api"
	"flag"
	"fmt"
	"log"
	"net/http"
)

const addr = ""

type App struct {
	Addr   string
	Port   string
	server *http.Server
}

func (a *App) Initialize() {
	// Read arg from command line for the port.
	portPtr := flag.String("port", "9000", "a string")
	flag.Parse()

	log.Printf("API Demo is running in port: %q", *portPtr)
	apiPort := string(*portPtr)
	a.Addr = fmt.Sprintf("%v:%v", addr, apiPort)
	mux := api.Handlers()

	log.Printf("Now listening on %s...\n", a.Addr)
	a.server = &http.Server{Handler: mux, Addr: a.Addr}
}

func (a *App) Run() {
	fmt.Println("Server starting")
	log.Fatal(a.server.ListenAndServe())
}
