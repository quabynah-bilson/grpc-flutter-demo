package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	// created grpc server
	server := grpc.NewServer()

	// @todo -> register service

	listen, err := net.Listen("tcp", ":6000")
	if err != nil {
		log.Fatalf("failed to listen: %s", err.Error()) // fatal error, exit
	}

	log.Printf("starting grpc server on port %s", ":6000")
	if err = server.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %s", err.Error())
	}
}
