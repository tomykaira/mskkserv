package skkserv

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

type Getter func(url string) (resp *http.Response, err error)

type GoogleTrans struct {
	getter Getter
}

func NewGoogleTrans() *GoogleTrans {
	return &GoogleTrans{getter: http.Get}
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

// Decode response like [["あさ",["朝","麻","アサ","あさ","厚狭"]]] to
// {"朝","麻","厚狭"}
func decodeResponse(query string, data string) (cands []string) {
	re := regexp.MustCompile("\\[\\s*\"" + query + "\"\\s*,\\s*\\[(.*?)\\]\\s*\\]")
	matches := re.FindStringSubmatch(data)

	if len(matches) <= 1 {
		return nil
	}
	filterRe := regexp.MustCompile("(([ァ-ヾ]+)|([ぁ-ゞ]+)|([ｦ-ﾟ]+)|([a-zA-Z]+))")
	for _, cand := range strings.Split(matches[1], ",") {
		cand = strings.Trim(cand, " \"")
		if !filterRe.MatchString(cand) {
			cands = append(cands, cand)
		}
	}
	return cands
}
