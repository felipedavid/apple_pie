package main

import (
	"flag"
	"log"
	"net/http"
)

type app struct {
}

func main() {
	addr := flag.String("addr", ":8000", "Interface and port number of the server")
	flag.Parse()

	a := app{}

	log.Printf("Broker serving running on %s\n", *addr)
	err := http.ListenAndServe(*addr, a.routes())
	log.Panic(err.Error())
}
