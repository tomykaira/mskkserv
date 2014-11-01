package skkserv

import (
	"bytes"
	"errors"
	"net/http"
	"strings"
	"testing"
)

type MockReadCloser struct {
	r *bytes.Buffer
}

func (m MockReadCloser) Read(p []byte) (n int, err error) {
	return m.r.Read(p)
}

func (m MockReadCloser) Close() (err error) {
	return
}

func newMockReadCloser(data string) MockReadCloser {
	return MockReadCloser{r: bytes.NewBufferString(data)}
}

func TestGoogleSearchGetsError(t *testing.T) {
	g := &GoogleTrans{getter: func(url string) (*http.Response, error) {
		if !strings.Contains(url, "%E3%81%97%E3%81%91%E3%82%93") {
			t.Fatalf("Request url does not contain encoded query %v", url)
		}
		return nil, errors.New("unknown error")
	}}
	res := g.Search("しけん")
	if res != nil {
		t.Errorf("Returns value when HTTP request failed %v", res)
	}
}

func TestGoogleSearchFindEntries(t *testing.T) {
	var g = &GoogleTrans{getter: func(url string) (*http.Response, error) {
		if !strings.Contains(url, "%E3%81%82%E3%81%95") {
			t.Fatalf("Request url does not contain encoded query %v", url)
		}
		return &http.Response{Body: newMockReadCloser(`[["あさ",["朝","麻","アサ","あさ","厚狭"]]]`)}, nil
	}}
	assertEqualList(t, []string{"朝", "麻", "厚狭"}, g.Search("あさ"))
}

func TestGoogleSearchReturnNilWhenAllFiltered(t *testing.T) {
	var g = &GoogleTrans{getter: func(url string) (*http.Response, error) {
		return &http.Response{Body: newMockReadCloser(`[["あさ",["あさ","アサ"]]]`)}, nil
	}}
	res := g.Search("あさ")
	if res != nil {
		t.Errorf("Returns value when all filtered %v", res)
	}
}

func TestGoogleSearchReturnNilForBrokenData(t *testing.T) {
	var g = &GoogleTrans{getter: func(url string) (*http.Response, error) {
		if !strings.Contains(url, "%E3%81%82%E3%81%95") {
			t.Fatalf("Request url does not contain encoded query %v", url)
		}
		return &http.Response{Body: newMockReadCloser(`[["あ",["朝","麻","アサ","あさ","厚狭"]]]`)}, nil
	}}
	res := g.Search("あさ")
	if res != nil {
		t.Errorf("Returns value when all filtered %v", res)
	}
}
