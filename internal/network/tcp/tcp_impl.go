package tcp

import (
	"net"

	"github.com/eXoterr/ggINX/internal/config"
)

type TCPImpl struct {
	listener net.Listener
}

func New() TCP {
	return &TCPImpl{}
}

func (tcp *TCPImpl) Setup(addr string, conf config.HTTP) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	tcp.listener = listener
	// tcp.listener.(*net.TCPListener).SetDeadline(time.Now().Add(time.Duration(conf.ReadTimeout) * time.Second))

	return nil
}

func (tcp *TCPImpl) Listen(stop <-chan struct{}) {
	defer tcp.listener.Close()

	for {
		conn, err := tcp.listener.Accept()
		if err != nil {
			continue
		}

		session := &Session{Conn: conn}

		go session.Info()
	}

}
