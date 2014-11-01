package main

import (
	"flag"
	"log"
	"net"
	"strconv"

	"github.com/tomykaira/mskkserv/skkserv"
)

// TODO: configure addr via options
const (
	SERVER_TYPE = "tcp"
)

type Config struct {
	host string
	port int
}

func main() {
	config := parseFlags()
	addr := config.host + ":" + strconv.Itoa(config.port)
	listener, err := net.Listen(SERVER_TYPE, addr)
	if err != nil {
		log.Fatalln("Error listening: ", err.Error())
	}
	defer listener.Close()
	log.Println("Listening on " + addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection: ", err.Error())
			continue
		}
		log.Println("Accepting connection from " + conn.RemoteAddr().String())
		go skkserv.HandleRequest(conn)
	}
}

func parseFlags() (config Config) {
	flag.StringVar(&config.host, "host", "127.0.0.1", "Host to bind")
	flag.IntVar(&config.port, "port", 1178, "Listening port")
	flag.Parse()
	return config
}
