package internal

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"net"
)

// Heos port
var port = "1255"

// NewHeos returns Heos instancec
func NewHeos(host string) Heos {
	var heos Heos
	heos.host = host + ":" + port

	return heos
}

// Connect initialises connection to the speaker
func (heos *Heos) Connect() error {
	conn, err := net.Dial("tcp", heos.host)
	if err != nil {
		return err
	}
	heos.conn = conn

	return nil
}

// Disconnect disconnects from the speaker
func (heos *Heos) Disconnect() error {
	return heos.conn.Close()
}

// SendCmdToHeosSpeaker sends given command with parameters to the speaker
func (heos *Heos) SendCmdToHeosSpeaker(cmd Command, params map[string]string) ([]byte, error) {

	logrus.Info("sending: "+"heos://%s/%s?%s\r\n", cmd.Group, cmd.Command, paramsToStr(params))

	_, err := fmt.Fprintf(heos.conn, "heos://%s/%s?%s\r\n", cmd.Group, cmd.Command, paramsToStr(params))
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(heos.conn)

	if !scanner.Scan() {
		return nil, fmt.Errorf("no response")
	}

	if err := scanner.Err(); err != nil {
		return nil, errors.New("Error in response: " + err.Error())
	}

	event := bytes.TrimRight([]byte(scanner.Text()), "\x00")
	return event, nil
}

func paramsToStr(params map[string]string) string {
	var str string

	first := true
	for k, v := range params {
		if first {
			first = false
			str = fmt.Sprintf("%s=%s", k, v)
			continue
		}
		str += fmt.Sprintf("&%s=%s", k, v)
	}

	return str
}
