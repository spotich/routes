run: build start

build:
		go build -o bin/main cmd/main.go

start:
		cd bin && ./main