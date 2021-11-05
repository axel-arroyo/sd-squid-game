package main

import (
	"context"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"
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

// https://stackoverflow.com/questions/55300117/how-do-i-find-all-files-that-have-a-certain-extension-in-go-regardless-of-depth
func playerFiles(player int32) []string {
	var matches []string
	var pattern = "jugador_" + strconv.Itoa(int(player)) + "__" + "*.txt"
	err := filepath.Walk("/home/alumno/sd-squid-game/datanode/plays/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		return nil
	}
	return matches
}

func (s *datanodeServer) ObtenerJugadas(ctx context.Context, in *pb.ObtenerJugadasReq) (*pb.ObtenerJugadasResp, error) {
	// var jugadasPath = playerFiles(in.NumJugador)
	var msg string = ""
	var filename = "/home/alumno/sd-squid-game/datanode/plays/jugador_" + strconv.Itoa(int(in.NumJugador)) + "__ronda_" + strconv.Itoa(int(in.Ronda)) + ".txt"
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("failed to read %s: %v\n", filename, err)
	} else {
		msg = string(dat)
	}
	return &pb.ObtenerJugadasResp{Msg: msg}, nil
}

func writePlay(jugador int32, ronda int32, jugada int32) {
	// textFile, _ := os.Create(filepath.Join("/plays/jugador_"+strconv.Itoa(int(jugador))+"__ronda_"+strconv.Itoa(int(ronda))+".txt", ))
	textFile, _ := os.Create("/home/alumno/sd-squid-game/datanode/plays/jugador_" + strconv.Itoa(int(jugador)) + "__ronda_" + strconv.Itoa(int(ronda)) + ".txt")
	textFile.WriteString(strconv.Itoa(int(jugada)) + "\n")
	textFile.Close()
	log.Printf("Escrita jugada de jugador %d", jugador)
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
