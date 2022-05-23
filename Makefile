build:
	go build -o sepias

install: build
	rm ~/go/bin/sepias
	cp sepias ~/go/bin/
