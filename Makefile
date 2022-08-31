build:
	go build -ldflags="-X 'main.BuildFlag=local'" -o sepias

install:
	go build -o sepias
	rm -f ~/go/bin/sepias
	cp sepias ~/go/bin/
