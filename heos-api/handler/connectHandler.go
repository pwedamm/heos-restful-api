package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"heos-restful-api/heos-api/internal"
	"net/http"
)

var log = logrus.New()

func ConnectHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	ip := vars["ip"]
	log.Infof("Got ip for connection with heos speaker: ", ip)

	if !internal.ValidIP4(ip) {

		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ip is not valid: %v\\n", ip)
	}

	internal.NewHeos(ip)

	heos := internal.NewHeos(ip)
	err := heos.Connect()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "connection failed due to error: %v\\n", err.Error())
	} else {

		w.WriteHeader(http.StatusOK)
		log.Infof("conncected successful to heos speaker with ip. %s ", ip)
	}
}
