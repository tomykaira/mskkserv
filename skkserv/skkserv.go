package skkserv

import (
	"bytes"
	"io"
	"log"
	"net"

	"code.google.com/p/go.text/encoding/japanese"
	"code.google.com/p/go.text/transform"

	"github.com/tomykaira/mskkserv"
)

const (
	COMMAND_CLIENT_END     = '0'
	COMMAND_CLIENT_REQUEST = '1'
	COMMAND_VERSION        = '2'
	COMMAND_HOST           = '3'
)

func HandleRequest(conn net.Conn) {
	buffer := bytes.Buffer
	for {
		err := readFromConnection(buffer)
		if err != nil {
			if err != io.EOF {
				mskkserv.Logger.Print("Error reading from conn", err.Error())
			}
			conn.Close()
			return
		}

		toClose, err := processCommand(conn, buffer)
		if err != nil {
			mskkserv.Logger.Print("Unexpected error processing command", err.Error())
		}
		if toClose {
			conn.Close()
			return
		}
	}
}

func readFromConnection(buffer bytes.Buffer) (err error) {
	const TEMP_BUFFER_LEN = 1024
	readbuf := make([]byte, TEMP_BUFFER_LEN)
	for {
		len, err := conn.Read(readbuf)
		if err != nil {
			return err
		}
		buffer.Write(readbuf)
		if len < TEMP_BUFFER_LEN || readbuf[len-1] == '\n' {
			break
		}
	}
	return nil
}

func processCommand(conn net.Conn, buffer bytes.Buffer) (toClose bool, err error) {
	command, err := buffer.ReadBytes('\n')
	if err == io.EOF {
		// command does not completed. wait for next chunk
		buffer.Write(command)
		return false, nil
	} else if err != nil {
		return true, err
	}

	switch command[0] {
	case COMMAND_CLIENT_END:
		return true, nil
	case COMMAND_CLIENT_REQUEST:
		writer = transform.NewWriter(conn, japanese.EUCJP.NewDecoder())
		io.WriteString(writer, "1/あ/い/う/\n")
	case COMMAND_CLIENT_VERSION:
		conn.Write([]byte("mskkserv " + mskkserv.VERSION + " "))
	case COMMAND_CLIENT_HOST:
		conn.Write([]byte(conn.LocalAddr().String() + " "))
	default:
		return true, errors.New("Unknown command")
	}

	return false, nil
}
