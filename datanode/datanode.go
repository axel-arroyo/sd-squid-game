package main

import (
	"log"
	"net"
	"os"

	pb "github.com/axel-arroyo/sd-squid-game/gen/proto"
	"google.golang.org/grpc"
)

type datanodeServer struct {
	pb.UnimplementedDatanodeServer
}

const (
	portServer = ":50054"
)

func writePlay(jugador int32, etapa int32) {
	var textFile os.File
	textFile.WriteString("")
}

func main() {
	namenodeListener, err := net.Listen("tcp", portServer)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterDatanodeServer(grpcServer, &datanodeServer{})
	log.Printf("server listening to namenode at %v", namenodeListener.Addr())
	err = grpcServer.Serve(namenodeListener)
	if err != nil {
		log.Println(err)
	}
}
