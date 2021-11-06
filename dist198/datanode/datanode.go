package main

import (
	"context"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"

	pb "github.com/axel-arroyo/sd-squid-game/gen/proto"
	"google.golang.org/grpc"
)

type datanodeServer struct {
	pb.UnimplementedDatanodeServer
}

const (
	portServer = ":50054"
)

func (s *datanodeServer) ObtenerJugadas(ctx context.Context, in *pb.ObtenerJugadasReq) (*pb.ObtenerJugadasResp, error) {
	path, _ := os.Getwd()
	// var jugadasPath = playerFiles(in.NumJugador)
	var msg string = ""
	var filename = path + "/plays/jugador_" + strconv.Itoa(int(in.NumJugador)) + "__ronda_" + strconv.Itoa(int(in.Ronda)) + ".txt"
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("failed to read %s: %v\n", filename, err)
	} else {
		msg = "Ronda " + strconv.Itoa(int(in.Ronda)) + ": " + string(dat)
	}
	return &pb.ObtenerJugadasResp{Msg: msg}, nil
}

func writePlay(jugador int32, ronda int32, jugada int32) {
	path, _ := os.Getwd()
	// textFile, _ := os.Create(filepath.Join("/plays/jugador_"+strconv.Itoa(int(jugador))+"__ronda_"+strconv.Itoa(int(ronda))+".txt", ))
	textFile, _ := os.Create(path + "/plays/jugador_" + strconv.Itoa(int(jugador)) + "__ronda_" + strconv.Itoa(int(ronda)) + ".txt")
	textFile.WriteString(strconv.Itoa(int(jugada)) + "\n")
	textFile.Close()
	// log.Printf("Escrita jugada de jugador %d", jugador)
}

func (s *datanodeServer) GuardarJugada(ctx context.Context, in *pb.GuardarJugadaReq) (*pb.GuardarJugadaResp, error) {
	// Guardar jugada en otro proceso y retornar instantaneo
	go writePlay(in.NumJugador, in.Ronda, in.Jugada)
	return &pb.GuardarJugadaResp{}, nil
}

func main() {
	// Escuchar al namenode
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
