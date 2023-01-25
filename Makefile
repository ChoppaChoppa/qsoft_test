build:
	go build cmd/qsoft/main.go

run: build
	./main

clean:
	go clean
	rm linux
