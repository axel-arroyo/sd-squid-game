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

	"github.com/streadway/amqp"
	"google.golang.org/grpc"
)

const (
	// Puerto en que escucha al Lider
	portServer = ":50051"
)

type pozoServer struct {
	pb.UnimplementedPozoServer
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func (s *pozoServer) Amount(ctx context.Context, req *pb.AmountReq) (*pb.AmountResp, error) {
	return &pb.AmountResp{Valor: int32(getAmount())}, nil
}

func getAmount() int {
	// Obtener ultima palabra del texto
	file, ferr := os.Open("pozo/pozo.txt")
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
	textFile, _ := os.OpenFile("pozo/pozo.txt", os.O_APPEND|os.O_WRONLY, 0644)
	newAmount := prevMont + 100000000
	_, err := textFile.WriteString("Jugador_" + strconv.Itoa(int(jugador)) + " Ronda_" + strconv.Itoa(int(ronda)) + " " + strconv.Itoa(int(newAmount)) + "\n")
	if err != nil {
		log.Fatalf("Error al escribir en el archivo: %v", err)
	}
	textFile.Close()
}

func rabbit() {
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"jugadoresEliminados", // name
		false,                 // durable
		false,                 // delete when unused
		false,                 // exclusive
		false,                 // no-wait
		nil,                   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			msg := strings.Split(string(d.Body), ",")
			numJugador, ronda := msg[0], msg[1]
			num, _ := strconv.Atoi(numJugador)
			round, _ := strconv.Atoi(ronda)
			writePlay(int32(num), int32(round))
		}
	}()

	log.Printf(" [*] Esperando notificación de jugadores eliminados.")
	<-forever
}

func main() {
	go rabbit()
	// Crear archivo de texto con las jugadas
	textFile, _ := os.Create("pozo/pozo.txt")
	textFile.Close()

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
