package main

import (
	"adp-backend/app"
	"flag"
	"net/http"
	"time"
)

type Data struct {
	Number  int
	Name    string
	Address string
}

func main() {

	config := app.Config{
		DB:     app.DBConfig(),
		Router: app.ConfigRouter(),
	}

	config.ConfigAPIs()
	s := &http.Server{
		Addr:         getPortNumber(),
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}

	config.Router.Logger.Fatal(config.Router.StartServer(s))
}

func getPortNumber() string {
	port := flag.String("port", "8000", "Port number")
	flag.Parse()
	if port == nil || len(*port) == 0 {
		panic("Please enter port")
	}
	return ":" + *port
}
