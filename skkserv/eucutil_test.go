package skkserv

import (
	"bytes"
	"testing"
)

func TestStringToEucSuccess(t *testing.T) {
	shikenEuc, err := StringToEuc("しけん")
	if err != nil {
		t.Errorf("StringToEuc(\"しけん\") must succeed but got error %v", err)
	}
	answer := []byte{'x', 'y', 'z'}
	if bytes.Compare(shikenEuc, answer) == 0 {
		t.Errorf("StringToEuc(\"しけん\") must be %v but got %v", answer, shikenEuc)
	}
}

func TestStringToEucFailure(t *testing.T) {
	_, err := StringToEuc("Ԕ")
	if err == nil {
		t.Errorf("StringToEuc(\"(hebrew)\") must ret error")
	}
}
