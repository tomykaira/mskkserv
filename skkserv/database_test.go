package skkserv

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io"
	"io/ioutil"
	"log"
	"testing"
)

const (
	/*
		  Base64 encoded, gzipped sample database
			+9,13:わりもどs->/割り戻/割戻/
			+5,4:わりt->/割/
	*/
	TEST_DATA = "H4sICIupVFQAA291dC5jZGIA0+dgAAP9UXqUHqVH6VF6lB6lBxXNBKTtoXxsNEjeH8onleYEYl4gXvJ+yaslj5acLNbf/GTJqzMPgRSQYAXKsEBlS4BC+l9ajvyRg2oFgR+vDogyIPEBi/ziNk8IAAA="
)

var TestDb *Database

func init() {
	file, err := ioutil.TempFile("", "skk_testdic")
	if err != nil {
		log.Fatalln(err)
	}
	decoded, err := base64.StdEncoding.DecodeString(TEST_DATA)
	if err != nil {
		log.Fatalln(err)
	}
	reader, err := gzip.NewReader(bytes.NewReader(decoded))
	if err != nil {
		log.Fatalln(err)
	}
	io.Copy(file, reader)
	err = file.Close()
	if err != nil {
		log.Fatalln(err)
	}
	TestDb, err = LoadDatabase(file.Name())
	if err != nil {
		log.Fatalln(err)
	}
}

func TestLoadDatabaseFails(t *testing.T) {
	_, err := LoadDatabase("notfound")
	if err == nil {
		t.Errorf("LoadDatabase must return error for invalid database path")
	}
}

func TestSearchNotFound(t *testing.T) {
	if TestDb.Search("わ") != nil {
		t.Errorf("Search() must return nil for a not found entry")
	}
}

func TestSearchFindOneEntry(t *testing.T) {
	assertEqualList(t, []string{"割"}, TestDb.Search("わりt"))
}

func TestSearchFindTwoEntries(t *testing.T) {
	assertEqualList(t, []string{"割り戻", "割戻"}, TestDb.Search("わりもどs"))
}

func assertEqualList(t *testing.T, x, y []string) {
	if len(x) != len(y) {
		t.Errorf("lists sizes are different %v != %v", x, y)
		return
	}

	for i := 0; i < len(x); i++ {
		if x[i] != y[i] {
			t.Errorf("Lists are different at %i: %v != %v", i, x[i], y[i])
		}
	}
}
