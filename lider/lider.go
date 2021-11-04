package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	pb "github.com/axel-arroyo/sd-squid-game/gen/proto"

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

var started = false
var etapa int32 = 1
var jugadoresVivos int32 = 16
var jugadoresUnidos int32 = 0
var jugadoresListos [3]int32 = [3]int32{0, 0, 0}
var jugadaRecibida [16]int32 = [16]int32{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

var jugadoresFinalizados [3]int32 = [3]int32{0, 0, 0}
var menuAlreadyPrinted bool = false
var menuLock sync.Mutex

// Tirar la cuerda
var jugadoresListosCuerda int32 = 0

// estadoFinalEtapa2 -> 0: no se sabe, 1: eliminado, 2: no eliminado
var estadoFinalEtapa2 [16]int32 = [16]int32{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
var numLuzVerde [4]int32
var numTirarCuerda int32
var jugadoresArray []int32
var equipo1 []int32
var equipo2 []int32
var jugadorEliminadoTirarCuerda int32 = 0
var sumaEquipo1 int32 = 0
var sumaEquipo2 int32 = 0
var eliminarEquipo1 bool = rand.Float32() < 0.5

type liderServer struct {
	pb.UnimplementedLiderServer
}

func EscogerNumerosLuzVerde() {
	// Escoger un numero por el Lider
	for i := 0; i < 4; i++ {
		numLuzVerde[i] = rand.Int31n(5) + 6
	}
}

func EliminarJugadorSobrante(i int) {
	jugadoresArray[i] = jugadoresArray[len(jugadoresArray)-1]
	jugadoresArray = jugadoresArray[:len(jugadoresArray)-1]
}

func jugadorEnEquipo(a int32, list []int32) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func SumarArray(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func EliminarJugador(numJugador int32) {
	atomic.AddInt32(&jugadoresVivos, -1)
	log.Printf("El jugador %d ha sido eliminado \n", numJugador)
	// Request al pozo para actualizar pozo
}

func TirarLaCuerda() {
	// Escoger un numero al azar entre 1 y 4
	numTirarCuerda = rand.Int31n(4) + 1
	if jugadoresVivos%2 == 1 {
		// Eliminar un jugador para tener paridad
		randomIdx := rand.Intn(len(jugadoresArray))
		jugadorEliminadoTirarCuerda = jugadoresArray[randomIdx]
		// Avisar al jugador (?)
		EliminarJugadorSobrante(randomIdx)
		EliminarJugador(jugadorEliminadoTirarCuerda)
	}
	// Dividir equipos
	equipo1 = jugadoresArray[:len(jugadoresArray)/2]
	equipo2 = jugadoresArray[len(jugadoresArray)/2:]
	log.Printf("Equipo 1: %v	-	Equipo 2: %v\n", equipo1, equipo2)
}

func TodoONada() {

}

const (
	siguienteEtapa = 1
	monstrarVivos  = 2
	jugadasJugador = 3
)

func IniciarEtapa() {
	switch etapa {
	case 1:
		log.Println("Ha comenzado la etapa 1: Luz verde luz roja")
		EscogerNumerosLuzVerde()
		break
	case 2:
		log.Println("Ha comenzado la etapa 2: Tirar la cuerda")
		TirarLaCuerda()
		break
	case 3:
		log.Println("Ha comenzado la etapa 3: Todo o nada")
		TodoONada()
		break
	}
}

func MenuJuego() {
	log.Printf("Ha finalizado la etapa %d\n", etapa)
	if etapa == 2 {
		log.Printf("Equipo1: %d 		Equipo2:%d\n", sumaEquipo1, sumaEquipo2)
	}
	log.Println("Ingrese opción: \n 1. Pasar a la siguiente Etapa \n 2. Mostrar los jugadores vivos \n 3. Mostrar las jugadas de un Jugador")
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
		break
	case monstrarVivos:
		log.Println(jugadoresArray)
		MenuJuego()
		break
	case jugadasJugador:
		MenuJuego()
		break
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
		started = true
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
				EliminarJugador(numJugador)
				return &pb.EnviarJugadaResp{Eliminado: true, Msg: "El lider ha escogido " + strconv.Itoa(int(numTirarCuerda)) + " y tu equipo ha sumado " + strconv.Itoa(int(sumaEquipo1))}
			} else {
				// El jugador es del equipo 2, no eliminarlo
				estadoFinalEtapa2[numJugador-1] = 2
				return &pb.EnviarJugadaResp{Eliminado: false, Msg: "El lider ha escogido " + strconv.Itoa(int(numTirarCuerda)) + " y tu equipo ha sumado " + strconv.Itoa(int(sumaEquipo2))}
			}
		} else {
			// Eliminar equipo 2
			if jugadorEnEquipo(numJugador, equipo2) {
				EliminarJugador(numJugador)
				return &pb.EnviarJugadaResp{Eliminado: true, Msg: "El lider ha escogido " + strconv.Itoa(int(numTirarCuerda)) + " y tu equipo ha sumado " + strconv.Itoa(int(sumaEquipo2))}
			} else {
				// El jugador es del equipo 1, no eliminarlo
				return &pb.EnviarJugadaResp{Eliminado: false, Msg: "El lider ha escogido " + strconv.Itoa(int(numTirarCuerda)) + " y tu equipo ha sumado " + strconv.Itoa(int(sumaEquipo1))}
			}
		}
	} else {
		// Un equipo gana, el otro pierde, revisar cual
		// Gana equipo 1
		if paridadLider == paridadEquipo1 {
			// Eliminar equipo 2
			if jugadorEnEquipo(numJugador, equipo2) {
				// El jugador pertenece al equipo 2, eliminarlo
				EliminarJugador(numJugador)
				return &pb.EnviarJugadaResp{Eliminado: true, Msg: "El lider ha escogido " + strconv.Itoa(int(numTirarCuerda)) + " y tu equipo ha sumado " + strconv.Itoa(int(sumaEquipo2))}
			} else {
				// El jugador pertenece al equipo 1, no eliminarlo
				return &pb.EnviarJugadaResp{Eliminado: false, Msg: "El lider ha escogido " + strconv.Itoa(int(numTirarCuerda)) + " y tu equipo ha sumado " + strconv.Itoa(int(sumaEquipo1))}
			}
		} else {
			// Eliminar equipo 1
			if jugadorEnEquipo(numJugador, equipo1) {
				// El jugador pertenece al equipo 1, eliminarlo
				EliminarJugador(numJugador)
				return &pb.EnviarJugadaResp{Eliminado: true, Msg: "El lider ha escogido " + strconv.Itoa(int(numTirarCuerda)) + " y tu equipo ha sumado " + strconv.Itoa(int(sumaEquipo1))}
			} else {
				return &pb.EnviarJugadaResp{Eliminado: false, Msg: "El lider ha escogido " + strconv.Itoa(int(numTirarCuerda)) + " y tu equipo ha sumado " + strconv.Itoa(int(sumaEquipo2))}
			}
		}
	}
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
			break
		} else {
			time.Sleep(time.Second * 2)
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
		// RegistrarJugada()
		eliminado := req.Jugada >= int32(numLuzVerdeLuzRoja)
		if eliminado {
			EliminarJugador(req.NumJugador)
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
			if sum(jugadaRecibida) == int32(len(jugadoresArray)) {
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
		return nil, errors.New("jugada recibida")

	case 3:
		// Todo o Nada
		break
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
	switch req.Estado {
	case true:
		// Caso especial req.Etapa == 2 -> Etapa 3
		if req.Etapa == 2 {
			// Revisar si esta en la lista de ganadoresEtapa2
			atomic.AddInt32(&jugadoresFinalizados[req.Etapa-1], 1)
		}
		if jugadorEnEquipo(req.NumJugador, jugadoresArray) {
			if startedStage[req.Etapa] {
				return &pb.EnviarEstadoResp{Msg: "si"}, nil
			}
			return nil, errors.New("aún no empieza la siguiente etapa")
		}
		atomic.AddInt32(&jugadoresFinalizados[req.Etapa-1], 1)
		// Agregar jugador a la lista de jugadores vivos para la etapa 2
		if req.Etapa == 1 {
			jugadoresArray = append(jugadoresArray, req.NumJugador)
		}
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
		EliminarJugador(req.NumJugador)
		return &pb.EnviarEstadoResp{Msg: "no"}, nil
	}
	return nil, nil
}

func (s *liderServer) PedirPozo(ctx context.Context, req *pb.PedirPozoReq) (*pb.PedirPozoResp, error) {
	// Conectar al Pozo
	connPozo, err := grpc.Dial("localhost"+portPozo, grpc.WithInsecure())
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
