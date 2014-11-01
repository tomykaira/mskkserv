package main

import (
	"flag"
	"log"
	"net"

	"github.com/tomykaira/mskkserv/skkserv"
)

// TODO: configure addr via options
const (
	SERVER_TYPE = "tcp"
	SERVER_ADDR = "localhost:1178"
	SERVER_PORT = "1178" // skkserv
)

func main() {
	listener, err := net.Listen(SERVER_TYPE, SERVER_ADDR)
	if err != nil {
		log.Fatalln("Error listening: ", err.Error())
	}
	defer listener.Close()
	log.Println("Listening on " + SERVER_ADDR)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection: ", err.Error())
			continue
		}
		go skkserv.HandleRequest(conn)
	}
}

func parseFlags() {

}
