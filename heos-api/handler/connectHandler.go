package handler

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"heos-restful-api/heos-api/internal"
	"net/http"
)

var log = logrus.New()

var globalHeos internal.Heos
var defaultIp = "192.168.178.76"

func ConnectHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	ip := vars["ip"]
	log.SetLevel(logrus.DebugLevel)
	log.Infof("Got ip for connection with heos speaker: ", ip)

	if !internal.ValidIP4(ip) {

		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ip is not valid: %v\\n", ip)
		return
	}

	heos := internal.NewHeos(ip)
	err := heos.Connect()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "connection failed due to error: %v\\n", err.Error())
		return
	} else {

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "connection successful to: %v", ip)
		log.Infof("conncected successful to heos speaker with ip. %s ", ip)
		globalHeos = heos
	}
}

func GetConnectedHeos() (internal.Heos, error) {

	if globalHeos.GetConnection() != nil {
		log.Debug("found connected heos client.")
		return globalHeos, nil
	}

	// Try to connect with predefined ip addresses
	log.Debugf("trying to connect with default client address %s", defaultIp)
	heos := internal.NewHeos(defaultIp)
	err := heos.Connect()
	if err == nil {
		globalHeos = heos
		log.Debugf("connected successfully to default client with ip %s", defaultIp)
		return heos, nil
	}
	return globalHeos, errors.New("client not connected")
}
