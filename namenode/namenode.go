package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"sync"
	"time"

	pb "github.com/axel-arroyo/sd-squid-game/gen/proto"
	"google.golang.org/grpc"
)

type namenodeServer struct {
	pb.UnimplementedNamenodeServer
}

const (
	portDatanode = ":50054"
	portServer   = ":50057"
)

var (
	ipDatanodes = [3]string{"localhost", "localhost", "localhost"}
	mutex       sync.Mutex
)

func (s *namenodeServer) DevolverJugadasJug(ctx context.Context, in *pb.DevolverJugadasJugReq) (*pb.DevolverJugadasJugResp, error) {
	var msg string
	// Conectar con todos los datanodes
	for i := 0; i < 3; i++ {
		ipDatanode := ipDatanodes[i]
		connData, err := grpc.Dial(ipDatanode+portDatanode, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("No se pudo conectar con el datanode: %v", err)
		}
		clientDatanode := pb.NewDatanodeClient(connData)
		// Enviar la peticion al datanode
		for {
			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()
			resp, err := clientDatanode.RegistrarJugada(ctx, &pb.RegistrarJugadaReq{NumJugador: in.NumJugador})
			if err != nil {
				time.Sleep(500 * time.Millisecond)
			} else {
				msg = msg + resp.Jugada
				break
			}
		}
		connData.Close()
	}
}

func WriteText(ip string, ronda int32, numJugador int32) {
	mutex.Lock()
	// Escribir la ip del datanode con la informacion de la jugada en archivo de texto
	textFile, _ := os.Open("info.txt")
	_, err := textFile.WriteString("Jugador_" + strconv.Itoa(int(numJugador)) + " Ronda_" + strconv.Itoa(int(ronda)) + " " + ip + "\n")
	if err != nil {
		log.Fatalf("Error al escribir en el archivo: %v", err)
	}
	mutex.Unlock()
}

func (s *namenodeServer) RegistrarJugada(ctx context.Context, in *pb.RegistrarJugadaReq) (*pb.RegistrarJugadaResp, error) {
	// Escoger un datanode al azar para almacenar la jugada
	random := rand.Intn(3)
	ipDatanode := ipDatanodes[random]
	connData, err := grpc.Dial(ipDatanode+portDatanode, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar con el datanode: %v", err)
	}
	clientDatanode := pb.NewDatanodeClient(connData)
	// Enviar la jugada al datanode
	for {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		_, err := clientDatanode.GuardarJugada(ctx, &pb.GuardarJugadaReq{Jugada: in.Jugada, Ronda: in.Ronda, NumJugador: in.NumJugador})
		if err != nil {
			time.Sleep(500 * time.Millisecond)
		} else {
			break
		}
	}
	// Almacenar la ip del datanode con la informacion de la jugada en archivo de texto
	go WriteText(ipDatanode, in.Ronda, in.NumJugador)
	connData.Close()
	return &pb.RegistrarJugadaResp{}, nil
}

func main() {
	// Escuchar al lider
	liderListener, err := net.Listen("tcp", portServer)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterNamenodeServer(grpcServer, &namenodeServer{})
	log.Printf("server listening to lider at %v", liderListener.Addr())

}
