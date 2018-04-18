package main 

import "github.com/gocql/gocql"

type Server struct{
	tag 	TagVal
	session *gocql.Session
}

func (s *Server) getFromServer(done chan TagVal) {
	done <- queryGet(s.session, done)
}

func (s *Server) setToServer(tv TagVal, done chan bool) {
	if s.tag.smaller(tv) {
		s.tag = tv
		querySet(tv, s.session)
	}
	done <- true
}
