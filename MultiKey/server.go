package main 

import (
	"fmt"
	"github.com/gocql/gocql"
)

type Server struct{
	tag 	TagVal
	session *gocql.Session
}

func (s *Server) getFromServer(key int, done chan TagVal) {
	s.tag = queryGet(0, s.session)
	if key != 0 {
		s.tag.Val = queryGet(key, s.session).Val		
	}
//	fmt.Println(s.tag)
	done <- s.tag
}

func (s *Server) setToServer(key int, tv TagVal, done chan bool) {
	if s.tag.smaller(tv) {
		fmt.Println("do update")
		s.tag = tv
		querySet(0, tv, s.session)
		querySet(key, tv, s.session)
	}
	done <- true
}
