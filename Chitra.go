package main

import (
	"flag"
	"os"

	"gitlab.com/archit-p/chitra/net"
	"gitlab.com/archit-p/chitra/repo"
)

func main() {
	// parse the command line options
	var config repo.ConfigFile

	pwd, _ := os.Getwd()

	flag.StringVar(&config.ServerDir, "sdir", "~/Videos/", "directory to serve videos from")
	flag.StringVar(&config.ClientDir, "cdir", pwd + "/build/client", "directory for client app")
	flag.StringVar(&config.ServerPort, "sport", "8080", "port for API server")
	flag.StringVar(&config.ClientPort, "cport", "5000", "port for UI client")

	flag.Parse()

	go func() {
		net.StartUIServer(config.ClientDir, config.ClientPort)
	}()

	net.StartAPIServer(config.ServerDir, config.ServerPort)
}
