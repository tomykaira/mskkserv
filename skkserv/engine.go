package skkserv

type Engine interface {
	Search(query string) (cands []string)
}
