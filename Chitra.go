package main

import (
	"chitra/net"
)

func main() {
	go func() {
		net.StartUIServer()
	}()

	net.StartAPIServer()
}

