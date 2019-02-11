package skkserv

import (
	"bytes"
	"io"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func EucToString(euc []byte) string {
	var buf bytes.Buffer
	reader := transform.NewReader(bytes.NewReader(euc), japanese.EUCJP.NewDecoder())
	buf.ReadFrom(reader)
	return buf.String()
}

func StringToEuc(str string) ([]byte, error) {
	var buf bytes.Buffer
	writer := transform.NewWriter(&buf, japanese.EUCJP.NewEncoder())
	_, err := io.WriteString(writer, str)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
