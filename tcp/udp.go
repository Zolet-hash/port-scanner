package tcp

import (
	"fmt"
	"net"
	"time"
)

func ScanUDPPort(host string, port int, timeout time.Duration) {
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialUDP("udp", nil, addre)

}
