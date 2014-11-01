package skkserv

import (
	"bytes"
	"errors"
	"io"
	"log"
	"net"

	"code.google.com/p/go.text/encoding/japanese"
	"code.google.com/p/go.text/transform"
)

const (
	COMMAND_END     = '0'
	COMMAND_REQUEST = '1'
	COMMAND_VERSION = '2'
	COMMAND_HOST    = '3'
)

func HandleRequest(conn net.Conn) {
	var buffer bytes.Buffer
	for {
		err := readFromConnection(conn, &buffer)
		if err != nil {
			if err != io.EOF {
				log.Println("Error reading from conn", err.Error())
			}
			conn.Close()
			return
		}

		toClose, err := processCommand(conn, &buffer)
		if err != nil {
			log.Println("Unexpected error processing command", err.Error())
		}
		if toClose {
			conn.Close()
			log.Println("Closed conenction from " + conn.RemoteAddr().String())
			return
		}
	}
}

func readFromConnection(conn net.Conn, buffer *bytes.Buffer) (err error) {
	const TEMP_BUFFER_LEN = 1024
	readbuf := make([]byte, TEMP_BUFFER_LEN)
	for {
		len, err := conn.Read(readbuf)
		if err != nil {
			return err
		}
		buffer.Write(readbuf[0:len])
		if len < TEMP_BUFFER_LEN || readbuf[len-1] == '\n' {
			break
		}
	}
	return nil
}

func processCommand(conn net.Conn, buffer *bytes.Buffer) (toClose bool, err error) {
	for buffer.Len() > 0 {
		command, err := buffer.ReadBytes('\n')
		if err == io.EOF {
			// command does not completed. wait for next chunk
			buffer.Write(command)
			return false, nil
		} else if err != nil {
			return true, err
		}

		switch command[0] {
		case COMMAND_END:
			return true, nil
		case COMMAND_REQUEST:
			writer := transform.NewWriter(conn, japanese.EUCJP.NewEncoder())
			io.WriteString(writer, "1/あ/い/う/\n")
		case COMMAND_VERSION:
			conn.Write([]byte("mskkserv " + VERSION + " \n"))
		case COMMAND_HOST:
			conn.Write([]byte(conn.LocalAddr().String() + " \n"))
		default:
			return true, errors.New("Unknown command")
		}
	}

	return false, nil
}
