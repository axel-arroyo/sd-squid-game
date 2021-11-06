killport:
	fuser -k 50054/tcp

run: killport
	go run datanode/datanode.go

clean:
	rm ./datanode/plays/*.txt