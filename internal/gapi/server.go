package gapi

import (
	"github.com/zvash/bgmood-file-service/internal/filepb"
	"github.com/zvash/bgmood-file-service/internal/util"
)

// Server serves gRPC requests for our api gateway.
type Server struct {
	filepb.UnimplementedFileServer
	config util.Config
}

func NewServer(config util.Config) (*Server, error) {
	server := &Server{
		config: config,
	}
	return server, nil
}
