package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/t0239184/golearn/internal/database"
	"github.com/t0239184/golearn/internal/router"
)

func init() {
	log.Info("[main] init")
}

func main() {
	var err error
	db := database.InitDatabase()
	engine := router.New(db)

	addr := ":8080"
	server := http.Server{
		Addr:    addr,
		Handler: engine,
	}
	err = server.ListenAndServe()
	if err != nil {
		log.Fatalf("[main] run - http.ListenAndServe failed: %v", err)
	}

	fmt.Println("Started Listening for plain HTTP connection on " + addr)
}
