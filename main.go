package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func scanPort(protocol, hostname string, port int) bool {
	address := hostname + ":" + strconv.Itoa(port)
	conn, error := net.DialTimeout(protocol, address, 60*time.Second)

	if error != nil {
		return false
	}
	defer conn.Close()
	return true
}
func main() {
	fmt.Println("Port Scanning...")
	open := scanPort("tcp", "localhost", 8000)
	fmt.Println("Port Open: ", open)
}
