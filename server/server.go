package server

import (
	"rick-morty/api/data"

	"github.com/hashicorp/go-hclog"
)

type Server struct {
	l        hclog.Logger
	response *data.ResponseBD
}

func NewCharacter(l hclog.Logger, res *data.ResponseBD) *Server {
	return &Server{l, res}
}
