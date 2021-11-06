package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"sync"
	"time"

	pb "github.com/axel-arroyo/sd-squid-game/gen/proto"

	"google.golang.org/grpc"
)

const (
	ipLider = "10.6.43.78"
	// Puerto del lider a donde enviar solicitudes
	portLider = ":50052"
)

// Retorna si un jugador fue eliminado en la ronda
func LuzVerdeLuzRoja(round int32, conn pb.LiderClient, bot bool, numJugador int32, suma *int32) bool {
	if bot {
		// Jugador bot
		// Escoger un número al azar entre 1 y 10 y enviar la jugada
		// opcion := rand.Int31n(10) + 1
		opcion := rand.Int31n(2) + 5
		respJugada, err := conn.EnviarJugada(context.Background(), &pb.EnviarJugadaReq{NumJugador: numJugador, Etapa: 1, Ronda: round, Jugada: opcion})
		if err != nil {
			fmt.Println(err)
		}
		// Si el jugador fue eliminado (escogió un numero mayor que el del Lider) retorna
		if respJugada.Eliminado {
			return true
		}
		*suma += opcion
		return false
	} else {
		// Jugador real
		var opcion int32
		fmt.Println("Ingrese un número del 1 al 10. Si es mayor o igual que el que escoja el Lider será eliminado, recuerda que el lider puede escoger un número entre 6 y 10")
		_, err := fmt.Scanf("%d", &opcion)
		if err != nil {
			fmt.Println("Error leyendo input")
		}
		// Enviar Req a Lider
		respJugada, err := conn.EnviarJugada(context.Background(), &pb.EnviarJugadaReq{NumJugador: numJugador, Etapa: 1, Ronda: round, Jugada: opcion})
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(respJugada.Msg)
		if respJugada.Eliminado {
			// Cagaste papito
			fmt.Println("Has sido eliminado")
			return true
		}
		*suma += opcion
		return false
	}
}

func TirarLaCuerda(conn pb.LiderClient, bot bool, numJugador int32) bool {
	var opcion int32
	// Escoger un número entre 1 y 4
	if bot {
		opcion = rand.Int31n(4) + 1
	} else {
		// Jugador real, leer opción del jugador
		fmt.Println("Ingrese un número entre 1 y 4")
		_, err := fmt.Scanf("%d", &opcion)
		if err != nil {
			fmt.Println("Error leyendo input")
		}
	}
	var respJugada *pb.EnviarJugadaResp
	var err error
	for {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		respJugada, err = conn.EnviarJugada(ctx, &pb.EnviarJugadaReq{NumJugador: numJugador, Etapa: 2, Ronda: 5, Jugada: opcion})
		if err != nil {
		} else {
			time.Sleep(1 * time.Second)
			break
		}
	}

	// fmt.Printf("Ha llegado respuesta al jugador %d\n", numJugador)
	if !bot {
		fmt.Println(respJugada.Msg)
	}
	return respJugada.Eliminado
}

func TodoONada(conn pb.LiderClient, bot bool, numJugador int32) bool {
	var opcion int32
	// Escoger un número entre 1 y 10
	if bot {
		opcion = rand.Int31n(10) + 1
	} else {
		// Jugador real, leer opción del jugador
		fmt.Println("Ingrese un número entre 1 y 10")
		_, err := fmt.Scanf("%d", &opcion)
		if err != nil {
			fmt.Println("Error leyendo input")
		}
	}
	var respJugada *pb.EnviarJugadaResp
	var err error
	for {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		respJugada, err = conn.EnviarJugada(ctx, &pb.EnviarJugadaReq{NumJugador: numJugador, Etapa: 3, Ronda: 6, Jugada: opcion})
		if err != nil {
		} else {
			time.Sleep(1 * time.Second)
			break
		}
	}

	// fmt.Printf("Ha llegado respuesta al jugador %d\n", numJugador)
	if !bot {
		fmt.Println(respJugada.Msg)
	}
	return respJugada.Eliminado
}

func Player(conn pb.LiderClient, wg *sync.WaitGroup, numJugador int32) {
	defer wg.Done()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stream, err := conn.UnirseAJuego(ctx, &pb.UnirseReq{Msg: "Quiero unirme al squirt game", NumJugador: numJugador})
	if err != nil {
		log.Fatalf("%v.UnirseAJuego(_) = _, %v", conn, err)
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error")
		}
		log.Printf(resp.Msg)
	}

	//RPC normal
	// LuzVerdeLuzRoja
	var suma int32 = 0
	fmt.Println("Etapa 1: Luz verde luz roja")
	fmt.Println("Debes sumar a lo menos 21 en cuatro rondas o menos, si al terminar la última ronda no logras sumar 21 serás eliminado.")
	for ronda := 1; ronda <= 4; ronda++ {
		eliminado := LuzVerdeLuzRoja(int32(ronda), conn, false, numJugador, &suma)
		if eliminado {
			// El jugador quedó eliminado (Escoger un número mayor)
			return
		}
		if suma >= 21 {
			// El jugador gana la etapa aunque no hayan pasado todas las rondas
			fmt.Println("Felicidades, lograste sumar más de 21, has pasado a la siguiente etapa.")
			break
		}
		fmt.Printf("Ronda %d: Tu suma total es %d \n", ronda, suma)
	}
	if suma < 21 {
		fmt.Println("No lograste sumar 21 en las cuatro rondas. Has sido eliminado.")
		// Avisar al lider para que actualize el pozo
		_, _ = conn.EnviarEstado(context.Background(), &pb.EnviarEstadoReq{NumJugador: numJugador, Estado: false, Etapa: 1, Ronda: 4})
		return
	}
	// El jugador pasó la etapa 1
	// Al término de la etapa, se le da la opción al jugador de ver o no el pozo acumulado
	fmt.Print("Deseas ver el pozo acumulado? (y/n) ")
	var opcion string
	_, err = fmt.Scanf("%s", &opcion)
	if err != nil {
		fmt.Print("Error leyendo la opción")
	}
	if opcion == "y" {
		respPozo, err := conn.PedirPozo(context.Background(), &pb.PedirPozoReq{})
		if err != nil {
			fmt.Println("Error pidiendo el pozo")
		}
		fmt.Printf("El pozo acumulado es %d \n", respPozo.Valor)
	}

	fmt.Println("Esperando a que el lider inice la siguiente etapa")
	// Avisar al lider que el jugador está vivo
	for {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		respEstado, err := conn.EnviarEstado(ctx, &pb.EnviarEstadoReq{NumJugador: numJugador, Estado: true, Etapa: 1, Ronda: 4})
		if err != nil {
		} else {
			if respEstado.Msg == "¡Has ganado!" {
				fmt.Println(respEstado.Msg)
				return
			}
			break
		}
	}
	fmt.Println("El lider ha comenzado la siguiente etapa.")
	fmt.Println("Etapa 2: Tirar la cuerda")
	fmt.Println("Debes elegir un número entre 1 y 4, si tu equipo logra sumar un número de la misma paridad que el escogido por el Líder, pasarás a la siguiente etapa.")

	eliminado := TirarLaCuerda(conn, false, numJugador)
	if eliminado {
		fmt.Println("Has sido eliminado")
		return
	}

	fmt.Println("Has pasado a la siguiente etapa.")
	fmt.Print("Deseas ver el pozo acumulado? (y/n) ")
	_, err = fmt.Scanf("%s", &opcion)
	if err != nil {
		fmt.Print("Error leyendo la opción")
	}
	if opcion == "y" {
		respPozo, err := conn.PedirPozo(context.Background(), &pb.PedirPozoReq{})
		if err != nil {
			fmt.Println("Error pidiendo el pozo")
		}
		fmt.Printf("El pozo acumulado es %d \n", respPozo.Valor)
	}
	fmt.Println("Esperando a que el lider inicie la siguiente etapa")

	// Avisar al Lider que el jugador está vivo para iniciar la etapa 3
	for {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_, err := conn.EnviarEstado(ctx, &pb.EnviarEstadoReq{NumJugador: numJugador, Estado: true, Etapa: 2, Ronda: 5})
		if err != nil {
			log.Println(err)
			time.Sleep(2 * time.Second)
		} else {
			break
		}
	}

	fmt.Println("El lider ha comenzado la siguiente etapa.")
	fmt.Println("Etapa 3: Todo o nada")
	fmt.Println("Competencia en parejas: El jugador que escoja el número más cercano al lider será ganador del Juego del Calamar.")

	eliminado = TodoONada(conn, false, numJugador)
	if eliminado {
		return
	}

	// Avisar que ganó
	for {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		respJugada, err := conn.EnviarEstado(ctx, &pb.EnviarEstadoReq{NumJugador: numJugador, Estado: true, Etapa: 4})
		if err != nil {
			time.Sleep(2 * time.Second)
		} else {
			fmt.Println(respJugada.Msg)
			break
		}
	}
}

func BotPlayer(conn pb.LiderClient, wg *sync.WaitGroup, numJugador int32) {
	defer wg.Done()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	stream, err := conn.UnirseAJuego(ctx, &pb.UnirseReq{Msg: "Quiero unirme al squirt game", NumJugador: numJugador})
	if err != nil {
		log.Fatalf("%v.UnirseAJuego(_) = _, %v", conn, err)
	}

	for {
		_, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error in bot player")
		}
	}

	//RPC normal
	// LuzVerdeLuzRoja
	var suma int32 = 0
	for round := 1; round <= 4; round++ {
		eliminado := LuzVerdeLuzRoja(int32(round), conn, true, numJugador, &suma)
		if eliminado {
			// El jugador quedó eliminado (Escoger un número mayor)
			// El server ya actualizó el pozo, no es necesario avisar
			return
		}
		if suma >= 21 {
			break
		}
	}
	if suma < 21 {
		_, err = conn.EnviarEstado(context.Background(), &pb.EnviarEstadoReq{NumJugador: numJugador, Estado: false, Etapa: 1, Ronda: 4})
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	// Avisar al Lider que el jugador está vivo al finalizar la etapa 1
	for {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_, err := conn.EnviarEstado(ctx, &pb.EnviarEstadoReq{NumJugador: numJugador, Estado: true, Etapa: 1})
		if err != nil {
		} else {
			break
		}
		time.Sleep(1 * time.Second)
	}

	// TirarLaCuerda
	eliminado := TirarLaCuerda(conn, true, numJugador)
	if eliminado {
		return
	}

	// Avisar al Lider que el jugador está vivo para empezar la etapa 3
	for {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_, err := conn.EnviarEstado(ctx, &pb.EnviarEstadoReq{NumJugador: numJugador, Estado: true, Etapa: 2})
		if err != nil {
			time.Sleep(1 * time.Second)
		} else {
			break
		}
	}

	// Todo o nada
	eliminado = TodoONada(conn, true, numJugador)
	if eliminado {
		return
	}
	// Avisar que ganó, se asume que llega
	for {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_, err = conn.EnviarEstado(ctx, &pb.EnviarEstadoReq{NumJugador: numJugador, Estado: true, Etapa: 4})
		if err != nil {
			time.Sleep(2 * time.Second)
		} else {
			break
		}
	}

}

func main() {
	// Conectar al Lider
	connLider, err := grpc.Dial(ipLider+portLider, grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}

	clientLider := pb.NewLiderClient(connLider)

	var wg sync.WaitGroup
	wg.Add(16)

	go Player(clientLider, &wg, 1)
	go BotPlayer(clientLider, &wg, 2)
	go BotPlayer(clientLider, &wg, 3)
	go BotPlayer(clientLider, &wg, 4)
	go BotPlayer(clientLider, &wg, 5)
	go BotPlayer(clientLider, &wg, 6)
	go BotPlayer(clientLider, &wg, 7)
	go BotPlayer(clientLider, &wg, 8)
	go BotPlayer(clientLider, &wg, 9)
	go BotPlayer(clientLider, &wg, 10)
	go BotPlayer(clientLider, &wg, 11)
	go BotPlayer(clientLider, &wg, 12)
	go BotPlayer(clientLider, &wg, 13)
	go BotPlayer(clientLider, &wg, 14)
	go BotPlayer(clientLider, &wg, 15)
	go BotPlayer(clientLider, &wg, 16)

	wg.Wait()
}
