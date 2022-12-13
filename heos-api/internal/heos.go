package internal

import "net"

// Heos is a representation of HEOS Speaker
type Heos struct {
	conn net.Conn
	host string
}

func (heos *Heos) GetConnection() net.Conn {
	return heos.conn
}
