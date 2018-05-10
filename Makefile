.PHONY : docker
docker: Dockerfile
	docker build -t go-grpc-awesome:latest .

.PHONY : grpc-awesome-osx
grpc-awesome-osx: main.go
	GOOS=darwin GOARCH=amd64 go build -ldflags '-s -w' -o $@

.PHONY : grpc-awesome-linux
grpc-awesome-linux: main.go
	GOOS=linux GOARCH=amd64 go build -ldflags '-s -w' -o $@
