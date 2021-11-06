package main

import (
	"bufio"
	"context"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
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
	ipDatanodes = [3]string{"10.6.43.80", "10.6.43.77", "10.6.43.78"}
	mutex       sync.Mutex
)

func (s *namenodeServer) DevolverJugadasJug(ctx context.Context, in *pb.DevolverJugadasJugReq) (*pb.DevolverJugadasJugResp, error) {
	var msg string = ""
	// Buscar la ip del datanode con la informacion de la jugada en archivo de texto
	textFile, _ := os.Open("info.txt")
	scanner := bufio.NewScanner(textFile)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "Jugador_"+strconv.Itoa(int(in.NumJugador))+" Ronda") {
			// Hacer request al datanode para devolver la jugada
			ipDatanode := strings.Split(line, " ")[2]
			// Ver que ronda tiene el archivo de texto
			ronda, _ := strconv.Atoi(strings.Split(strings.Split(line, " ")[1], "_")[1])
			connData, err := grpc.Dial(ipDatanode+portDatanode, grpc.WithInsecure())
			if err != nil {
				log.Fatalf("No se pudo conectar con el datanode: %v", err)
			}
			clientDatanode := pb.NewDatanodeClient(connData)
			for {
				// Enviar request al datanode hasta que haya respuesta
				ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
				defer cancel()
				resp, err := clientDatanode.ObtenerJugadas(ctx, &pb.ObtenerJugadasReq{NumJugador: in.NumJugador, Ronda: int32(ronda)})
				if err != nil {
					time.Sleep(500 * time.Millisecond)
				} else {
					// log.Printf("El datanode %s ha retornado al jugador %d -> %s\n", ipDatanode, in.NumJugador, resp.Msg)
					msg += resp.Msg
					// log.Printf("Nuevo mensaje es %s \n", msg)
					break
				}
			}
		}
	}
	return &pb.DevolverJugadasJugResp{Msg: msg}, nil
}

func WriteText(ip string, ronda int32, numJugador int32) {
	mutex.Lock()
	// Escribir la ip del datanode con la informacion de la jugada en archivo de texto
	textFile, _ := os.OpenFile("info.txt", os.O_APPEND|os.O_WRONLY, 0644)
	_, err := textFile.WriteString("Jugador_" + strconv.Itoa(int(numJugador)) + " Ronda_" + strconv.Itoa(int(ronda)) + " " + ip + "\n")
	if err != nil {
		log.Fatalf("Error al escribir en el archivo: %v", err)
	}
	textFile.Close()
	mutex.Unlock()
}

func (s *namenodeServer) RegistrarJugada(ctx context.Context, req *pb.RegistrarJugadaReq) (*pb.RegistrarJugadaResp, error) {
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
		_, err := clientDatanode.GuardarJugada(ctx, &pb.GuardarJugadaReq{Jugada: req.Jugada, Ronda: req.Ronda, NumJugador: req.NumJugador})
		if err != nil {
			time.Sleep(500 * time.Millisecond)
		} else {
			break
		}
	}
	// Almacenar la ip del datanode con la informacion de la jugada en archivo de texto
	go WriteText(ipDatanode, req.Ronda, req.NumJugador)
	connData.Close()
	return &pb.RegistrarJugadaResp{}, nil
}

func main() {
	// Crear archivo de texto con las jugadas de cada jugador
	textFile, _ := os.Create("info.txt")
	textFile.Close()
	// Escuchar al lider
	liderListener, err := net.Listen("tcp", portServer)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterNamenodeServer(grpcServer, &namenodeServer{})
	log.Printf("server listening to lider at %v", liderListener.Addr())

	err = grpcServer.Serve(liderListener)
	if err != nil {
		log.Println(err)
	}

}
