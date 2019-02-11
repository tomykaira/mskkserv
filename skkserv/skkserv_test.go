package skkserv

import (
	"bytes"
	"net"
	"testing"
	"time"
)

func createSkkserv(t *testing.T, readable string) (*Skkserv, MockConn) {
	var config Config
	config.Engines = append([]Engine{}, TestDb)

	var mock MockConn
	mock.t = t
	e, _ := StringToEuc(readable)
	mock.readable = bytes.NewBuffer(e)
	mock.written = bytes.NewBuffer(nil)
	mock.closed = new(bool)
	*mock.closed = false

	return New(config), mock
}

type MockAddr struct {
}

func (a MockAddr) Network() string {
	return "127.0.0.1:0"
}

func (a MockAddr) String() string {
	return "127.0.0.1:0"
}

type MockConn struct {
	t        *testing.T
	readable *bytes.Buffer
	written  *bytes.Buffer
	closed   *bool
}

func (c MockConn) Read(b []byte) (n int, err error) {
	return c.readable.Read(b)
}

func (c MockConn) Write(b []byte) (n int, err error) {
	return c.written.Write(b)
}

func (c MockConn) Close() error {
	*c.closed = true
	return nil
}

func (c MockConn) LocalAddr() net.Addr {
	var m MockAddr
	return m
}

func (c MockConn) RemoteAddr() net.Addr {
	var m MockAddr
	return m
}

func (c MockConn) SetDeadline(t time.Time) error {
	c.t.Fatalf("SetDeadLine is not supported")
	return nil
}

func (c MockConn) SetReadDeadline(t time.Time) error {
	c.t.Fatalf("SetReadDeadLine is not supported")
	return nil
}

func (c MockConn) SetWriteDeadline(t time.Time) error {
	c.t.Fatalf("SetWriteDeadLine is not supported")
	return nil
}

func TestHandleRequestEnd(t *testing.T) {
	serv, conn := createSkkserv(t, "0\n")
	serv.HandleRequest(conn)
	if !*conn.closed {
		t.Error("0 does not close connection")
	}
	if conn.written.Len() != 0 {
		t.Errorf("connection get unexpected data %v", conn.written.Bytes())
	}
}

func TestHandleRequestVersionThenEnd(t *testing.T) {
	serv, conn := createSkkserv(t, "2\n0\n")
	serv.HandleRequest(conn)
	if !*conn.closed {
		t.Error("0 does not close connection")
	}
	written := EucToString(conn.written.Bytes())
	expected := "mskkserv " + VERSION + " \n"
	if written != expected {
		t.Errorf("Version returned unexpected string %v (expected: %v)", written, expected)
	}
}

func TestHandleRequestVersionNoLF(t *testing.T) {
	serv, conn := createSkkserv(t, "2")
	serv.HandleRequest(conn)
	written := EucToString(conn.written.Bytes())
	expected := "mskkserv " + VERSION + " \n"
	if written != expected {
		t.Errorf("Version returned unexpected string %v (expected: %v)", written, expected)
	}
}

func TestHandleRequestVersionTwiceThenEnd(t *testing.T) {
	serv, conn := createSkkserv(t, "2\n2\n0\n")
	serv.HandleRequest(conn)
	if !*conn.closed {
		t.Error("0 does not close connection")
	}
	written := EucToString(conn.written.Bytes())
	expected := "mskkserv " + VERSION + " \n"
	if written != expected+expected {
		t.Errorf("Version returned unexpected string %v", written)
	}
}

func TestHandleRequestRequestFailEnd(t *testing.T) {
	serv, conn := createSkkserv(t, "1みちご   \n0\n")
	serv.HandleRequest(conn)
	if !*conn.closed {
		t.Error("0 does not close connection")
	}
	written := EucToString(conn.written.Bytes())
	expected := "4みちご \n"
	if written != expected {
		t.Errorf("Version returned unexpected string %v (expected: %v)", written, expected)
	}
}

func TestHandleRequestRequestSuccessEnd(t *testing.T) {
	serv, conn := createSkkserv(t, "1わりもどs \n0\n")
	serv.HandleRequest(conn)
	if !*conn.closed {
		t.Error("0 does not close connection")
	}
	written := EucToString(conn.written.Bytes())
	expected := "1/割り戻/割戻/\n"
	if written != expected {
		t.Errorf("Version returned unexpected string %v (expected: %v)", written, expected)
	}
}

func TestHandleRequestRequestNoLFSuccessEnd(t *testing.T) {
	serv, conn := createSkkserv(t, "1わりもどs ")
	serv.HandleRequest(conn)
	if !*conn.closed {
		t.Error("0 does not close connection")
	}
	written := EucToString(conn.written.Bytes())
	expected := "1/割り戻/割戻/\n"
	if written != expected {
		t.Errorf("Version returned unexpected string %v (expected: %v)", written, expected)
	}
}
