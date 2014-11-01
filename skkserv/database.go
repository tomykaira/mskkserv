package skkserv

import (
	"bytes"
	"strings"

	"github.com/jbarham/go-cdb"
)

type Database struct {
	cdb *cdb.Cdb
}

// Load SKK CDB database. Return error if open failed.
//
// The database must be encoded in EUC-JP.
func LoadDatabase(cdbPath string) (db *Database, err error) {
	db = new(Database)
	db.cdb, err = cdb.Open(cdbPath)
	return db, err
}

func (d *Database) Search(query string) (cands []string) {
	eucQuery := StringToEuc(query)
	rdata, err := d.cdb.Find(eucQuery)
	if err != nil {
		return
	}
	var b bytes.Buffer
	b.ReadFrom(rdata)
	string := EucToString(b.Bytes())
	return strings.Split(strings.Trim(string, "/"), "/")
}
