killport:
	fuser -k 50054/tcp && fuser -k 50052/tcp

run: killport
	go run datanode/datanode.go &
	go run lider/lider.go

clean:
	rm ./datanode/plays/*.txt