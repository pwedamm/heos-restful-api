package internal

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
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
func (heos *Heos) SendCmdToHeosSpeaker(cmd Command, params map[string]string) (string, error) {

	_, err := fmt.Fprintf(heos.conn, "heos://%s/%s?%s\r\n", cmd.Group, cmd.Command, params)
	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(heos.conn)

	if !scanner.Scan() {
		return "", fmt.Errorf("no response")
	}

	if err := scanner.Err(); err != nil {
		return "", errors.New("Error in response: " + err.Error())
	}

	event := bytes.TrimRight([]byte(scanner.Text()), "\x00")

	return string(event), nil
}
