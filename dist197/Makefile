killport:
	fuser -k 50057/tcp && fuser -k 50051/tcp

run: killport
	go run pozo/pozo.go &
	go run datanode/datanode.go

clean:
	rm ./datanode/plays/*.txt
	rm ./pozo/pozo.txt