run:
	go run main.go

build:
	go build

cli-c:
	echo "Compiling for every OS and Platform"; \
	GOOS=freebsd GOARCH=386 go build -o bin/pith-freebsd-386 main.go;\
	GOOS=linux GOARCH=386 go build -o bin/pith-linux-386 main.go; \
	GOOS=windows GOARCH=386 go build -o bin/pith-windows-386 main.go