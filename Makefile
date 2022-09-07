build:
	go build -ldflags="-X 'main.BuildFlag=local'" -o sepia

install:
	go build -o sepia
	rm -f ~/go/bin/sepia
	cp sepia ~/go/bin/
