package skkserv

import (
	"bytes"
	"io"

	"code.google.com/p/go.text/encoding/japanese"
	"code.google.com/p/go.text/transform"
)

func EucToString(euc []byte) string {
	var buf bytes.Buffer
	reader := transform.NewReader(bytes.NewReader(euc), japanese.EUCJP.NewDecoder())
	buf.ReadFrom(reader)
	return buf.String()
}

func StringToEuc(str string) []byte {
	var buf bytes.Buffer
	writer := transform.NewWriter(&buf, japanese.EUCJP.NewEncoder())
	io.WriteString(writer, str)
	return buf.Bytes()
}
