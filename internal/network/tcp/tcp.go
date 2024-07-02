package tcp

import "github.com/eXoterr/ggINX/internal/config"

type TCP interface {
	Listen()
	Setup(addr string, conf config.HTTP) error
}
