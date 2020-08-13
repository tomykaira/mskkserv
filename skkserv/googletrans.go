package skkserv

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

type Getter func(url string) (resp *http.Response, err error)

type GoogleTrans struct {
	getter Getter
}

func NewGoogleTrans() *GoogleTrans {
	client := http.Client{
		Timeout: 800 * time.Millisecond,
	}
	return &GoogleTrans{getter: client.Get}
}

func (g *GoogleTrans) Search(query string) (cands []string) {
	escaped := url.QueryEscape(query)
	resp, err := g.getter("https://inputtools.google.com/request?text=" + escaped + "&itc=ja-t-ja-hira-i0-und&num=100")
	if err != nil {
		log.Println("Google transliterate request failed ", err)
		return nil
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	return decodeResponse(query, bytes.NewBuffer(data).String())
}

// Decode response like ["SUCCESS",[["あさ",["朝","麻","アサ","あさ","厚狭"],[],{"candidate_type":[0,0,0,0,0]}]]] to
// {"朝","麻","厚狭"}
func decodeResponse(query string, response string) (cands []string) {
	var data interface{}
	dec := json.NewDecoder(bytes.NewBufferString(response))
	dec.Decode(&data)

	if arr, ok := data.([]interface{}); ok && len(arr) >= 2 {
		data = arr[1]
	} else {
		return nil
	}
	if arr, ok := data.([]interface{}); ok && len(arr) >= 1 {
		data = arr[0]
	} else {
		return nil
	}
	if arr, ok := data.([]interface{}); ok && len(arr) >= 2 && arr[0] == query {
		data = arr[1]
	} else {
		return nil
	}
	rawCands, ok := data.([]interface{})
	if !ok {
		return nil
	}

	filterRe := regexp.MustCompile("\\A(([ァ-ヾ]+)|([ぁ-ゞ]+)|([ｦ-ﾟ]+)|([a-zA-Z]+))\\z")
	for _, cand := range rawCands {
		if str, ok := cand.(string); ok {
			str = strings.Trim(str, " ")
			if !filterRe.MatchString(str) {
				cands = append(cands, str)
			}
		}
	}
	return cands
}
