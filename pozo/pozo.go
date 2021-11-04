package main

import (
	"bufio"
	"context"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	pb "github.com/axel-arroyo/sd-squid-game/gen/proto"

	"google.golang.org/grpc"
)

const (
	// Puerto en que escucha al Lider
	portServer = ":50051"
	// Puero en que el Lider escucha al Pozo
	portClient = "x"
)

type pozoServer struct {
	pb.UnimplementedPozoServer
}

func (s *pozoServer) Amount(ctx context.Context, req *pb.AmountReq) (*pb.AmountResp, error) {
	return &pb.AmountResp{Valor: int32(getAmount())}, nil
}

func getAmount() int {
	// Obtener ultima palabra del texto
	file, ferr := os.Open("pozo.txt")
	if ferr != nil {
		panic(ferr)
	}
	scanner := bufio.NewScanner(file)
	var line string
	for scanner.Scan() {
		line = scanner.Text()
	}
	last := line[strings.LastIndex(line, " ")+1:]
	i, err := strconv.Atoi(last)
	if err != nil {
		return 0
	}
	return i
}

func writePlay(textFile os.File, jugador int32, ronda int32) {
	prevMont := getAmount()
	newAmount := prevMont + 100000000
	textFile.WriteString("Jugador_" + strconv.Itoa(int(jugador)) + " Ronda_" + strconv.Itoa(int(ronda)) + " " + strconv.Itoa(int(newAmount)) + "\n")
}

func main() {
	// Crear archivo de texto con las jugadas
	textFile, err := os.Create("pozo.txt")
	writePlay(*textFile, 1, 1)

	listener, err := net.Listen("tcp", portServer)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	pb.RegisterPozoServer(grpcServer, &pozoServer{})
	log.Printf("server listening to lider  at %v", listener.Addr())

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Println(err)
	}
}
