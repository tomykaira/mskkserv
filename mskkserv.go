package main

import (
	"log"
	"net"
	"os"

	"github.com/tomykaira/mskkserv/skkserv"
)

// TODO: configure addr via options
const (
	SERVER_TYPE = "tcp"
	SERVER_ADDR = "localhost:1178"
	SERVER_PORT = "1178" // skkserv
	VERSION     = "0.1.0"
)

var (
	Logger = log.Logger.New(os.Stdout, "mskkserv: ", log.Ldate|log.Ltime|log.Llongfile)
)

func main() {
	listener, err := net.Listen(SERVER_TYPE, SERVER_ADDR)
	if err != nil {
		Logger.Fatal("Error listening: ", err.Error())
	}
	defer l.Close()
	Logger.Print("Listening on " + SERVER_ADDR)

	for {
		conn, err := listener.Accept()
		if err != nil {
			Logger.Print("Error accepting connection: ", err.Error())
			continue
		}
		go skkserv.HandleRequest(conn)
	}
}
