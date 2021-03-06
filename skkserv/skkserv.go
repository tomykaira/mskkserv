package skkserv

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"net"
)

const (
	COMMAND_END     = '0'
	COMMAND_REQUEST = '1'
	COMMAND_VERSION = '2'
	COMMAND_HOST    = '3'
)

type Skkserv struct {
	config Config
}

func New(config Config) *Skkserv {
	return &Skkserv{config: config}
}

func (s *Skkserv) HandleRequest(conn net.Conn) {
	buf := bufio.NewReader(conn)
	for {
		command, err := buf.ReadByte()
		if err != nil {
			if err != io.EOF {
				log.Println("Error reading command", err.Error())
			}
			conn.Close()
			return
		}

	Process:
		switch command {
		case COMMAND_END:
			conn.Close()
			if s.config.Debug {
				log.Println("Disconnect " + conn.RemoteAddr().String() + " due to COMMAND_END")
			}
			return
		case COMMAND_REQUEST:
			queryBytes, err := buf.ReadBytes(' ')
			if err != nil {
				log.Println("Error reading query string", err.Error())
				conn.Close()
			}
			query := EucToString(queryBytes[0 : len(queryBytes)-1])

			for _, engine := range s.config.Engines {
				cands := engine.Search(query)
				if s.config.Debug {
					log.Printf("Trying %v, %v, %v\n", engine, query, len(cands))
				}
				emit := false
				var retbuf bytes.Buffer
				head, _ := StringToEuc("1")
				retbuf.Write(head)
				for _, cand := range cands {
					euc, err := StringToEuc("/" + cand)
					if err == nil {
						emit = true
						retbuf.Write(euc)
					}
				}
				trail, _ := StringToEuc("/\n")
				retbuf.Write(trail)
				if emit {
					log.Println(EucToString(retbuf.Bytes()))
					conn.Write(retbuf.Bytes())
					break Process
				}
			}
			ret, _ := StringToEuc("4" + query + " \n")
			conn.Write(ret)
		case COMMAND_VERSION:
			conn.Write([]byte("mskkserv " + VERSION + " \n"))
		case COMMAND_HOST:
			conn.Write([]byte(conn.LocalAddr().String() + " \n"))
		case ' ', '\n':
			// skip indifferent separators
		default:
			log.Printf("Unknown command %v. Skipping\n", command)
		}
	}
}
