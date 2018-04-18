package main 

import "github.com/gocql/gocql"

type Server struct{
	tag 	TagVal
	session *gocql.Session
}

func (s Server) getFromServer(key int, done chan TagVal) {
	s.tag.Val = queryGet(key, s.session).Val		
	done <- s.tag
}

func (s Server) setToServer(key int, tv TagVal, done chan bool) {
	if s.tag.smaller(tv) {
		s.tag = tv
		querySet(key, tv, s.session)
	}
	done <- true
}
