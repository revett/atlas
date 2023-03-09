build:
	go build -o atlas

install: build
	rm -f ~/go/bin/atlas
	cp atlas ~/go/bin/
