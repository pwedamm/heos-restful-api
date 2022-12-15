package heos_api

import (
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"os"
)

type Configuration struct {
	HeosIp     string
	ServerIp   string
	ServerPort string
	Username   string
	Password   string
	LogLevel   string
}

func GetConfiguration() (Configuration, error) {

	file, err := os.Open("config.json")
	if err != nil {
		return Configuration{}, errors.New("Configuration file open error " + err.Error())
	}
	defer file.Close()
	dec := json.NewDecoder(file)
	configuration := Configuration{}
	errno := dec.Decode(&configuration)
	if err != nil {
		return Configuration{}, errors.New("Configuration parsing error " + errno.Error())
	}

	return configuration, nil
}

func (c Configuration) GetIp() string {
	return c.HeosIp
}

func (c Configuration) GetUsername() string {
	return c.Username
}

func (c Configuration) GetPassword() string {
	return c.Password
}

func (c Configuration) GetServerPort() string {
	return c.ServerPort
}

func (c Configuration) GetServerIp() string {
	return c.ServerIp
}

func (c Configuration) GetLogLever() logrus.Level {

	if c.LogLevel == "INFO" {
		return logrus.InfoLevel
	}

	if c.LogLevel == "DEBUG" {
		return logrus.DebugLevel
	}

	if c.LogLevel == "TRACE" {
		return logrus.TraceLevel
	}

	return logrus.ErrorLevel
}
