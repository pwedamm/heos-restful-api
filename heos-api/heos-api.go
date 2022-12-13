package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"heos-restful-api/heos-api/handler"
	"net/http"
	"time"
)

var log = logrus.New()

func main() {

	log.SetLevel(logrus.DebugLevel)

	// creating mux handler https://github.com/gorilla/mux
	router := mux.NewRouter()
	router.HandleFunc("/hello", hello)

	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		// an example API handler
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	router.HandleFunc("/connect/{ip}", handler.ConnectHandler)

	server := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 30,
		Handler:      router,
	}

	log.Infof("starting webserver on: %s.", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Error(err)
	}
}

func hello(w http.ResponseWriter, req *http.Request) {

	log.Debug("entered hello handler")
	fmt.Fprintf(w, "hello\n")
}
