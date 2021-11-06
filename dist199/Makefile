killport:
	fuser -k 50057/tcp

run: killport
	go run namenode/namenode.go &
	go run jugadores/jugador.go

clean:
	rm ./namenode/info.txt