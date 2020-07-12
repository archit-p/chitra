package main

import (
	"gitlab.com/archit-p/chitra/net"
)

func main() {
	go func() {
		net.StartUIServer()
	}()

	net.StartAPIServer()
}

