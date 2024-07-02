package tcp

import (
	"fmt"
	"net"
)

type Session struct {
	Conn net.Conn
}

func (s *Session) Info() {
	fmt.Println(s.Conn.RemoteAddr())
	s.Conn.Close()
}
