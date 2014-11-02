package main

import (
	"flag"
	"log"
	"net"
	"strconv"
	"strings"

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
		if config.Debug {
			log.Println("Accepting connection from " + conn.RemoteAddr().String())
		}
		go serv.HandleRequest(conn)
	}
}

func parseFlags() (config skkserv.Config) {
	flag.StringVar(&config.Host, "host", "127.0.0.1", "Host to bind")
	flag.IntVar(&config.Port, "port", 1178, "Listening port")
	flag.BoolVar(&config.Debug, "debug", false, "Enable debug mode")
	flag.Parse()
	config.Engines = initialzeEngines(flag.Args())
	return config
}

func initialzeEngines(args []string) (engines []skkserv.Engine) {
	for _, arg := range args {
		pair := strings.SplitN(arg, "=", 2)
		switch pair[0] {
		case "database":
			if len(pair) == 1 {
				log.Fatalln("database engine requires CDB file: database=/path/to/SKK-JISYO.L.cdb")
			}
			database, err := skkserv.LoadDatabase(pair[1])
			if err != nil {
				log.Fatalln("Failed to load database")
			}
			engines = append(engines, database)
			log.Println("Using database engine " + pair[1])
		case "googletrans":
			engines = append(engines, skkserv.NewGoogleTrans())
			log.Println("Using Google transliterate engine")
		}
	}
	if len(engines) == 0 {
		log.Fatalln("No engines selected")
	}
	return engines
}
