package main

import (
	"flag"
	"log"
	"net"
	"strconv"

	"github.com/tomykaira/mskkserv/skkserv"
)

const (
	SERVER_TYPE = "tcp"
)

func main() {
	config := parseFlags()

	addr := config.Host + ":" + strconv.Itoa(config.Port)
	listener, err := net.Listen(SERVER_TYPE, addr)
	if err != nil {
		log.Fatalln("Error listening: ", err.Error())
	}
	defer listener.Close()
	log.Println("Listening on " + addr)

	serv := skkserv.New(config)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection: ", err.Error())
			continue
		}
		log.Println("Accepting connection from " + conn.RemoteAddr().String())
		go serv.HandleRequest(conn)
	}
}

func parseFlags() (config skkserv.Config) {
	flag.StringVar(&config.Host, "host", "127.0.0.1", "Host to bind")
	flag.IntVar(&config.Port, "port", 1178, "Listening port")
	flag.StringVar(&config.DatabaseFile, "database", "/usr/local/share/skk/SKK-JISYO.L.cdb", "Listening port")
	flag.Parse()
	return config
}
