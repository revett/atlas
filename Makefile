build:
	go build -o sepias

install: build
	cp sepias ~/go/bin/
