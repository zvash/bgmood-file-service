package main

import (
	"github.com/zvash/bgmood-file-service/internal/filepb"
	"github.com/zvash/bgmood-file-service/internal/gapi"
	"github.com/zvash/bgmood-file-service/internal/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	runGrpcServer(config)
}

func runGrpcServer(config util.Config) {
	server, err := gapi.NewServer(config)
	if err != nil {
		log.Fatal("cannot create gRPC server")
	}

	grpcServer := grpc.NewServer()
	filepb.RegisterFileServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot create listener")
	}

	log.Printf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server")
	}
}
