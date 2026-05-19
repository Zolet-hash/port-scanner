package tcp

import (
	"fmt"
	"net"
)

func ScanUDPPort(host string, port int) bool {
	address := fmt.Sprintf("%s%d", host, port)
	conn, err := net.Dial("udp4", address)

	if err != nil {
		return false
	}
	conn.Close()
	return true

}
