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

func writePlay(jugador int32, ronda int32) {
	prevMont := getAmount()
	textFile, _ := os.OpenFile("pozo.txt", os.O_APPEND|os.O_WRONLY, 0644)
	newAmount := prevMont + 100000000
	_, err := textFile.WriteString("Jugador_" + strconv.Itoa(int(jugador)) + " Ronda_" + strconv.Itoa(int(ronda)) + " " + strconv.Itoa(int(newAmount)) + "\n")
	if err != nil {
		log.Fatalf("Error al escribir en el archivo: %v", err)
	}
	textFile.Close()
}

func main() {
	// Crear archivo de texto con las jugadas
	textFile, _ := os.Create("pozo.txt")
	textFile.Close()
	// textFile.WriteString("Jugador_0 Ronda_0 0\n")

	writePlay(1, 1)
	writePlay(2, 2)

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
