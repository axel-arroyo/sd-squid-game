package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	pb "github.com/axel-arroyo/sd-squid-game/gen/proto"

	"github.com/streadway/amqp"
	"google.golang.org/grpc"
)

const (
	ipNamenode   = "10.6.43.79"
	portNamenode = ":50057"
	ipPozo       = "10.6.43.77"
	portPozo     = ":50051"
	portServer   = ":50052"
)

// Variables Globales
var startedStage [3]bool = [3]bool{false, false, false}
var jugadoresVivosArray []int32 = []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
var playerListLock sync.Mutex
var ganadores []int32

var estadoRecibido [][]bool = [][]bool{
	{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false}}

var ended = false
var etapa int32 = 1
var jugadoresVivos int32 = 16
var jugadoresUnidos int32 = 0
var jugadoresListos [3]int32 = [3]int32{0, 0, 0}
var jugadoresFinalizados [3]int32 = [3]int32{0, 0, 0}
var menuAlreadyPrinted bool = false
var menuLock sync.Mutex

// estadoFinalEtapa2 -> 0: no se sabe, 1: eliminado, 2: no eliminado
var estadoFinalEtapa2 [16]int32 = [16]int32{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
var numLuzVerde [4]int32
var numTirarCuerda int32

// Tirar la cuerda
var jugadaRecibida [16]int32 = [16]int32{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
var jugadoresListosCuerda int32 = 0
var equipo1 []int32
var equipo2 []int32
var jugadorEliminadoTirarCuerda int32 = 0
var sumaEquipo1 int32 = 0
var sumaEquipo2 int32 = 0
var eliminarEquipo1 bool = rand.Float32() < 0.5

// Todo o nada
var jugadorEliminadoTodoNada int32 = 0
var numeroTodoONada int32 = rand.Int31n(10) + 1
var parejas [][]int32
var jugadasTodoONada [16]int32 = [16]int32{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
var jugadaRecibida3 [16]int32 = [16]int32{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type liderServer struct {
	pb.UnimplementedLiderServer
}

// RabbitMQ
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func avisarPozo(numJugador int32, ronda int32) {
	conn, err := amqp.Dial("amqp://admin:admin@" + ipPozo + ":5672/")
	failOnError(err, "Falla en conectarse a RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// El mensaje lleva el numero de jugador y la ronda separados por coma
	body := strconv.Itoa(int(numJugador)) + "," + strconv.Itoa(int(ronda))
	err = ch.Publish(
		"",                    // exchange
		"jugadoresEliminados", // routing key
		false,                 // mandatory
		false,                 // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
}

func EscogerNumerosLuzVerde() {
	// Escoger un numero por el Lider
	for i := 0; i < 4; i++ {
		numLuzVerde[i] = rand.Int31n(5) + 6
	}
}

func jugadorEnEquipo(a int32, list []int32) bool {
	for _, b := range list {
		if b == a {
			// log.Printf("Revisando %d en %v retorna true\n", a, list)
			return true
		}
	}
	// log.Printf("Revisando %d en %v retorna false\n", a, list)
	return false
}

func SumarArray(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func EliminarJugador(numJugador int32, ronda int32) {
	playerListLock.Lock()
	avisarPozo(numJugador, ronda)
	atomic.AddInt32(&jugadoresVivos, -1)
	// Eliminar jugador de la lista
	for i, v := range jugadoresVivosArray {
		if v == numJugador {
			jugadoresVivosArray[i] = jugadoresVivosArray[len(jugadoresVivosArray)-1]
			jugadoresVivosArray = jugadoresVivosArray[:len(jugadoresVivosArray)-1]
			break
		}
	}
	log.Printf("El jugador %d ha sido eliminado \n", numJugador)
	playerListLock.Unlock()
	// Request al pozo para actualizar pozo
}

func shufflePlayers() {
	// Reordenar jugadores al azar
	for i := 0; i < len(jugadoresVivosArray); i++ {
		j := rand.Intn(len(jugadoresVivosArray))
		jugadoresVivosArray[i], jugadoresVivosArray[j] = jugadoresVivosArray[j], jugadoresVivosArray[i]
	}
}

func TirarLaCuerda() {
	// Escoger un numero al azar entre 1 y 4
	numTirarCuerda = rand.Int31n(4) + 1
	if jugadoresVivos%2 == 1 {
		// Eliminar un jugador para tener paridad
		randomIdx := rand.Intn(len(jugadoresVivosArray))
		jugadorEliminadoTirarCuerda = jugadoresVivosArray[randomIdx]
		EliminarJugador(jugadorEliminadoTirarCuerda, 5)
	}
	// Dividir equipos
	shufflePlayers()
	equipo1 = make([]int32, len(jugadoresVivosArray)/2)
	equipo2 = make([]int32, len(jugadoresVivosArray)/2)
	copy(equipo1, jugadoresVivosArray[:len(jugadoresVivosArray)/2])
	copy(equipo2, jugadoresVivosArray[len(jugadoresVivosArray)/2:])
	log.Printf("Equipo 1: %v	-	Equipo 2: %v\n", equipo1, equipo2)
}

func TodoONada() {
	if len(jugadoresVivosArray)%2 == 1 {
		// Eliminar un jugador para tener paridad
		randomIdx := rand.Intn(len(jugadoresVivosArray))
		numJugadorEliminado := jugadoresVivosArray[randomIdx]
		EliminarJugador(numJugadorEliminado, 6)
		jugadorEliminadoTodoNada = numJugadorEliminado
	}
	// Armar parejas
	shufflePlayers()
	for i := 0; i < len(jugadoresVivosArray); i += 2 {
		end := i + 2
		if end > len(jugadoresVivosArray) {
			end = len(jugadoresVivosArray)
		}
		newpareja := make([]int32, 2)
		copy(newpareja, jugadoresVivosArray[i:end])
		parejas = append(parejas, newpareja)
	}
}

const (
	monstrarVivos  = 1
	jugadasJugador = 2
	siguienteEtapa = 3
)

func IniciarEtapa() {
	switch etapa {
	case 1:
		log.Println("Ha comenzado la etapa 1: Luz verde luz roja")
		EscogerNumerosLuzVerde()
	case 2:
		log.Println("Ha comenzado la etapa 2: Tirar la cuerda")
		TirarLaCuerda()
	case 3:
		log.Println("Ha comenzado la etapa 3: Todo o nada")
		TodoONada()
		log.Printf("Las parejas son: %v\n", parejas)
	}
}

func DevolverJugadas(numJugador int32) {
	// Conexion con namenode
	conn, err := grpc.Dial(ipNamenode+portNamenode, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to namenode: %v", err)
	}
	clientNamenode := pb.NewNamenodeClient(conn)
	// Enviar request
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()
	resp, err := clientNamenode.DevolverJugadasJug(ctx, &pb.DevolverJugadasJugReq{
		NumJugador: numJugador,
	})
	if err != nil {
		log.Fatalf("Error en el request a namenode: %v", err)
	} else {
		log.Printf("Jugadas de jugador %d: \n%s\n", numJugador, resp.Msg)
	}
	conn.Close()
}

func MenuFinal() {
	log.Println("El juego ha terminado")
	log.Println("Ingrese opción: \n 1. Mostrar los ganadores \n 2. Mostrar las jugadas de un Jugador \n 3. Terminar")
	var opcion int
	fmt.Scanf("%d", &opcion)
	switch opcion {
	case 1:
		log.Printf("Los ganadores son: %v\n", jugadoresVivosArray)
		MenuFinal()
	case 2:
		log.Println("Ingrese el numero de jugador: (1-16)")
		var numJugador int32
		fmt.Scanf("%d", &numJugador)
		DevolverJugadas(numJugador)
		MenuFinal()
	case 3:
		os.Exit(0)
	}
}

func MenuJuego() {
	if jugadoresVivos == 0 {
		log.Println("No hay jugadores vivos")
		log.Println("Ingrese opción: \n 2. Mostrar las jugadas de un Jugador \n 4. Terminar")
	}
	// Revisar si existe solo 1 jugador (terminar antes de etapa 3)
	if jugadoresVivos == 1 {
		log.Printf("El jugador %d ha ganado\n", jugadoresVivosArray[0])
		ended = true
		MenuFinal()
	} else {
		log.Printf("Ha finalizado la etapa %d\n", etapa)
		if etapa == 2 {
			log.Printf("Equipo1: %d 		Equipo2:%d\n", sumaEquipo1, sumaEquipo2)
		}
		if etapa == 3 {
			// Terminó el juego
			MenuFinal()
		} else {
			log.Println("Ingrese opción: \n 1. Mostrar los jugadores vivos \n 2. Mostrar las jugadas de un Jugador \n 3. Pasar a la siguiente Etapa")
		}
	}
	// Leer opcion
	var opcion int32
	_, err := fmt.Scanf("%d", &opcion)
	if err != nil {
		log.Println("Error leyendo opción")
	}
	switch opcion {
	case siguienteEtapa:
		atomic.AddInt32(&etapa, 1)
		IniciarEtapa()
		startedStage[etapa-1] = true
	case monstrarVivos:
		log.Println(jugadoresVivosArray)
		MenuJuego()
	case jugadasJugador:
		log.Println("Ingrese el numero de jugador (1-16)")
		_, err := fmt.Scanf("%d", &opcion)
		if err != nil {
			log.Println("Error leyendo opción")
		}
		if opcion <= 16 && opcion >= 1 {
			DevolverJugadas(opcion)
		} else {
			log.Println("El numero de jugador no es valido")
		}
		MenuJuego()
	case 4:
		os.Exit(0)
	}
}

func (s *liderServer) UnirseAJuego(req *pb.UnirseReq, stream pb.Lider_UnirseAJuegoServer) error {
	log.Println("El jugador " + strconv.Itoa(int(req.NumJugador)) + " se ha unido")

	res := &pb.UnirseResp{Msg: "Te has unido al squirt game.\nEsperando a que el Líder inicie el juego."}
	err := stream.Send(res)
	if err != nil {
		return err
	}

	atomic.AddInt32(&jugadoresUnidos, 1)
	if atomic.LoadInt32(&jugadoresUnidos) == 16 {
		log.Println("Presione cualquier tecla para iniciar el juego")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		startedStage[0] = true
		IniciarEtapa()
	}
	for {
		if startedStage[0] {
			stream.Send(&pb.UnirseResp{Msg: "El juego ha comenzado"})
			break
		}
	}

	return nil
}

func sum(array [16]int32) int32 {
	var result int32 = 0
	for _, v := range array {
		result += v
	}
	return result
}

func resultCuerda(numJugador int32) *pb.EnviarJugadaResp {
	var paridadLider int32 = numTirarCuerda % 2
	var paridadEquipo1 int32 = int32(sumaEquipo1) % 2
	var paridadEquipo2 int32 = int32(sumaEquipo2) % 2

	if paridadEquipo1 == paridadEquipo2 {
		// Empate de los dos equipos, revisar si uno es igual al del lider, si es así, ambos pasan, si no, se elimina un equipo al azar
		if paridadEquipo1 == paridadLider {
			// Ganaron, ambos pasan
			return &pb.EnviarJugadaResp{Eliminado: false, Msg: "Ambos equipos pasan"}
		}
		// Ambos perdieron, eliminar uno al azar
		if eliminarEquipo1 {
			// Eliminar equipo 1
			if jugadorEnEquipo(numJugador, equipo1) {
				// El jugador pertenece al equipo 1, eliminarlo
				EliminarJugador(numJugador, 5)
				return &pb.EnviarJugadaResp{Eliminado: true, Msg: "El lider ha escogido " + strconv.Itoa(int(numTirarCuerda)) + " y tu equipo ha sumado " + strconv.Itoa(int(sumaEquipo1)) + ". Han sido eliminados al azar (empate)"}
			} else {
				// El jugador es del equipo 2, no eliminarlo
				estadoFinalEtapa2[numJugador-1] = 2
				return &pb.EnviarJugadaResp{Eliminado: false, Msg: "El lider ha escogido " + strconv.Itoa(int(numTirarCuerda)) + " y tu equipo ha sumado " + strconv.Itoa(int(sumaEquipo2)) + ". Han pasado la etapa al azar (empate)"}
			}
		} else {
			// Eliminar equipo 2
			if jugadorEnEquipo(numJugador, equipo2) {
				EliminarJugador(numJugador, 5)
				return &pb.EnviarJugadaResp{Eliminado: true, Msg: "El lider ha escogido " + strconv.Itoa(int(numTirarCuerda)) + " y tu equipo ha sumado " + strconv.Itoa(int(sumaEquipo2)) + ". Han sido eliminados al azar (empate)"}
			} else {
				// El jugador es del equipo 1, no eliminarlo
				return &pb.EnviarJugadaResp{Eliminado: false, Msg: "El lider ha escogido " + strconv.Itoa(int(numTirarCuerda)) + " y tu equipo ha sumado " + strconv.Itoa(int(sumaEquipo1)) + ". Han pasado la etapa al azar (empate)"}
			}
		}
	} else {
		// Un equipo gana, el otro pierde, revisar cual
		// Gana equipo 1
		if paridadLider == paridadEquipo1 {
			// Eliminar equipo 2
			if jugadorEnEquipo(numJugador, equipo2) {
				// El jugador pertenece al equipo 2, eliminarlo
				EliminarJugador(numJugador, 5)
				return &pb.EnviarJugadaResp{Eliminado: true, Msg: "El lider ha escogido " + strconv.Itoa(int(numTirarCuerda)) + " y tu equipo ha sumado " + strconv.Itoa(int(sumaEquipo2))}
			} else {
				// El jugador pertenece al equipo 1, no eliminarlo
				return &pb.EnviarJugadaResp{Eliminado: false, Msg: "El lider ha escogido " + strconv.Itoa(int(numTirarCuerda)) + " y tu equipo ha sumado " + strconv.Itoa(int(sumaEquipo1))}
			}
		} else {
			// Eliminar equipo 1
			if jugadorEnEquipo(numJugador, equipo1) {
				// El jugador pertenece al equipo 1, eliminarlo
				EliminarJugador(numJugador, 5)
				return &pb.EnviarJugadaResp{Eliminado: true, Msg: "El lider ha escogido " + strconv.Itoa(int(numTirarCuerda)) + " y tu equipo ha sumado " + strconv.Itoa(int(sumaEquipo1))}
			} else {
				return &pb.EnviarJugadaResp{Eliminado: false, Msg: "El lider ha escogido " + strconv.Itoa(int(numTirarCuerda)) + " y tu equipo ha sumado " + strconv.Itoa(int(sumaEquipo2))}
			}
		}
	}
}

func resultTodoNada(numJugador int32) (*pb.EnviarJugadaResp, error) {
	// Revisar error donde ambos miembros de una pareja pierden (xd)
	var jug1 int32
	var jug2 int32
	var dif1 float64
	var dif2 float64
	for _, pareja := range parejas {
		jug1 = pareja[0]
		jug2 = pareja[1]
		// Resta entre las jugadas del jugador y del lider
		dif1 = math.Abs(float64(numeroTodoONada - jugadasTodoONada[jug1-1]))
		dif2 = math.Abs(float64(numeroTodoONada - jugadasTodoONada[jug2-1]))
		// El que está revisando es el jug1 (primero en la pareja)
		if jug1 == numJugador {
			// Caso base, los dos números son iguales, ambos ganan
			if jugadasTodoONada[jug2-1] == jugadasTodoONada[jug1-1] {
				ganadores = append(ganadores, numJugador)
				return &pb.EnviarJugadaResp{Eliminado: false, Msg: "Ambos jugadores ganan"}, nil
			}
			// Revisar quien está más cerca del número del lider
			if dif1 < dif2 {
				// Gana el jugador 1, pierde el jugador 2
				return &pb.EnviarJugadaResp{Eliminado: false, Msg: "¡Has ganado!"}, nil
			} else {
				EliminarJugador(numJugador, 6)
				return &pb.EnviarJugadaResp{Eliminado: true, Msg: "Has sido eliminado"}, nil
			}
		}
		// El que está revisando es el jug2 (segundo en la pareja)
		if jug2 == numJugador {
			// Caso base, los dos números son iguales, ambos ganan
			if jugadasTodoONada[jug2-1] == jugadasTodoONada[jug1-1] {
				return &pb.EnviarJugadaResp{Eliminado: false, Msg: "Ambos jugadores ganan"}, nil
			}
			if dif1 < dif2 {
				// Gana el jugador 1, pierde el jugador 2
				EliminarJugador(numJugador, 6)
				return &pb.EnviarJugadaResp{Eliminado: true, Msg: "Has sido eliminado"}, nil
			} else {
				return &pb.EnviarJugadaResp{Eliminado: false, Msg: "¡Has ganado!"}, nil
			}
		}
	}
	return nil, errors.New("no se encontró el jugador en una pareja")
}

func RegistrarJugada(jugada int32, numJugador int32, ronda int32) {
	// Conectar al Namenode
	connNamenode, err := grpc.Dial(ipNamenode+portNamenode, grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}
	clientNamenode := pb.NewNamenodeClient(connNamenode)
	// Hacer request al namenode hasta que la almacene
	for {
		_, err = clientNamenode.RegistrarJugada(context.Background(), &pb.RegistrarJugadaReq{Jugada: jugada, NumJugador: numJugador, Ronda: ronda})
		if err != nil {
			time.Sleep(time.Second * 2)
		} else {
			break
		}
	}
	connNamenode.Close()
}

func (s *liderServer) EnviarJugada(ctx context.Context, req *pb.EnviarJugadaReq) (*pb.EnviarJugadaResp, error) {
	switch req.Etapa {
	case 1:
		menuAlreadyPrinted = false
		// LuzVerdeLuzRoja
		numLuzVerdeLuzRoja := numLuzVerde[req.Ronda-1]
		// Registrar jugada en namenode
		RegistrarJugada(req.Jugada, req.NumJugador, req.Ronda)
		eliminado := req.Jugada >= int32(numLuzVerdeLuzRoja)
		if eliminado {
			EliminarJugador(req.NumJugador, req.Ronda)
		}
		return &pb.EnviarJugadaResp{Eliminado: eliminado, Msg: "El lider ha ingresado " + strconv.Itoa(int(numLuzVerdeLuzRoja))}, nil
	case 2:
		menuAlreadyPrinted = false
		// Si el jugador fue eliminado al azar, avisarle que fue eliminado
		if req.NumJugador == jugadorEliminadoTirarCuerda {
			return &pb.EnviarJugadaResp{Eliminado: true, Msg: "Has sido eliminado al azar"}, nil
		}
		if etapa != 2 {
			return nil, errors.New("aún no se definen los equipos")
		}
		if jugadaRecibida[req.NumJugador-1] == 1 {
			// Ya se recibió la jugada del jugador anteriormente
			if sum(jugadaRecibida) >= int32(len(jugadoresVivosArray)) {
				// Se recibieron todas las jugadas, calcular el resultado
				return resultCuerda(req.NumJugador), nil
			} else {
				// Aún no se recibieron todas las jugadas
				return nil, errors.New("aún no se recibieron todas las jugadas")
			}
		}
		// Buscar a que equipo pertenece el jugador y sumarle la jugada a su equipo
		if jugadorEnEquipo(req.NumJugador, equipo1) {
			atomic.AddInt32(&sumaEquipo1, req.Jugada)
		} else {
			atomic.AddInt32(&sumaEquipo2, req.Jugada)
		}
		if jugadaRecibida[req.NumJugador-1] == 0 {
			atomic.AddInt32(&jugadoresListosCuerda, 1)
		}
		atomic.StoreInt32(&jugadaRecibida[req.NumJugador-1], 1)
		RegistrarJugada(req.Jugada, req.NumJugador, req.Ronda)
		return nil, errors.New("jugada recibida")

	case 3:
		// Todo o Nada
		menuAlreadyPrinted = false
		if req.NumJugador == jugadorEliminadoTodoNada {
			return &pb.EnviarJugadaResp{Eliminado: true, Msg: "Has sido eliminado al azar"}, nil
		}
		if etapa != 3 {
			return nil, errors.New("aún no se definen las parejas")
		}
		if jugadaRecibida3[req.NumJugador-1] == 1 {
			// Ya se recibió la jugada del jugador anteriormente
			if sum(jugadaRecibida3) >= int32(len(jugadoresVivosArray)) {
				// Se recibieron ambas jugadas de la pareja
				return resultTodoNada(req.NumJugador)
			} else {
				// Aún no se recibieron las jugadas de la pareja
				return nil, errors.New("aún no se reciben ambas jugadas")
			}
		}
		for _, pareja := range parejas {
			if pareja[0] == req.NumJugador || pareja[1] == req.NumJugador {
				// El jugador pertenece a la pareja
				jugadaRecibida3[req.NumJugador-1] = 1
				jugadasTodoONada[req.NumJugador-1] = req.Jugada
				RegistrarJugada(req.Jugada, req.NumJugador, req.Ronda)
				return nil, errors.New("jugada recibida")
			}
		}
		return nil, errors.New("wtf")
	}
	return &pb.EnviarJugadaResp{
		Msg: "Se ha recibido la jugada correctamente"}, nil
}

func printMenuOnce() bool {
	menuLock.Lock()
	var canPrint bool
	if menuAlreadyPrinted {
		canPrint = false
	} else {
		canPrint = true
		menuAlreadyPrinted = true
	}
	menuLock.Unlock()
	return canPrint
}

func (s *liderServer) EnviarEstado(ctx context.Context, req *pb.EnviarEstadoReq) (*pb.EnviarEstadoResp, error) {
	// Todos las etapas han finalizado, revisar ganadores
	if req.Etapa == 4 {
		if !jugadorEnEquipo(req.NumJugador, ganadores) {
			ganadores = append(ganadores, req.NumJugador)
		}
		// Esperar todas las request
		for {
			if jugadoresVivos == int32(len(ganadores)) {
				if printMenuOnce() {
					MenuJuego()
				}
				break
			}
		}
		pozo := strconv.Itoa(int(1600000000-jugadoresVivos*100000000) / (int(jugadoresVivos)))
		return &pb.EnviarEstadoResp{
			Msg: "Han ganado " + strconv.Itoa(int(jugadoresVivos)) + " jugadores, tu premio es de " + pozo}, nil
	}
	switch req.Estado {
	case true:
		if ended {
			// El juego terminó, revisar si es de los ganadores
			if jugadorEnEquipo(req.NumJugador, jugadoresVivosArray) {
				// Eliminarlo de la lista de jugadores vivos
				return &pb.EnviarEstadoResp{Msg: "¡Has ganado!"}, nil
			}
		}
		// Revisar si el estado ya había sido recibido
		if estadoRecibido[req.Etapa-1][req.NumJugador-1] {
			// Si empezó la etapa, avisarle al jugador que puede seguir
			if startedStage[req.Etapa] {
				return &pb.EnviarEstadoResp{Msg: "si"}, nil
			}
			return nil, errors.New("aún no empieza la siguiente etapa")
		}

		atomic.AddInt32(&jugadoresFinalizados[req.Etapa-1], 1)
		estadoRecibido[req.Etapa-1][req.NumJugador-1] = true
		// Esperar a que todos los jugadores estén listos
		for {
			if jugadoresFinalizados[req.Etapa-1] == jugadoresVivos {
				if printMenuOnce() {
					MenuJuego()
				}
				if startedStage[req.Etapa] {
					break
				}
			}
		}
		return &pb.EnviarEstadoResp{Msg: "si"}, nil
	case false:
		EliminarJugador(req.NumJugador, req.Ronda)
		return &pb.EnviarEstadoResp{Msg: "no"}, nil
	}
	return nil, nil
}

func (s *liderServer) PedirPozo(ctx context.Context, req *pb.PedirPozoReq) (*pb.PedirPozoResp, error) {
	// Conectar al Pozo
	connPozo, err := grpc.Dial(ipPozo+portPozo, grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}
	clientPozo := pb.NewPozoClient(connPozo)
	// Hacer request al pozo
	respPozo, err := clientPozo.Amount(context.Background(), &pb.AmountReq{})
	if err != nil {
		log.Println("Error pidiendo el pozo")
	}
	// Mandar respuesta al jugador
	return &pb.PedirPozoResp{Valor: respPozo.Valor}, nil
}

func main() {
	// Escuchar al Jugador
	jugadorListener, err := net.Listen("tcp", portServer)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Crear servidor
	grpcServer := grpc.NewServer()
	pb.RegisterLiderServer(grpcServer, &liderServer{})
	log.Printf("server listening to player at %v", jugadorListener.Addr())

	err = grpcServer.Serve(jugadorListener)
	if err != nil {
		log.Println(err)
	}

}
