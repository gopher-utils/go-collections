test:
	go test -v ./... -count 1 -race

benchmark:
	go test -bench=. -benchmem