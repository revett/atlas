build:
	go build -ldflags="-X 'main.BuildFlag=local'" -o atlas

install:
	go build -o atlas
	rm -f ~/go/bin/atlas
	cp atlas ~/go/bin/
