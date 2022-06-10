build:
	go build -o sepias

install: build
	rm -f ~/go/bin/sepias
	cp sepias ~/go/bin/
