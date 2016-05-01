.PHONY: build proto test build-linux build-darwin build-windows build-linux-arm clean

default: build

proto:
	protoc -I=./model/ --go_out=plugins=grpc:./model/ ./model/*.proto

testbuild: proto
	GOARCH=amd64 GOOS=linux go build -v -o bin/dp-x86_64-linux darkpassenger.go
	GOARCH=amd64 GOOS=darwin go build -v -o bin/dp-x86_64-darwin darkpassenger.go

build: build-linux build-darwin build-windows build-linux-arm

build-linux: proto
	GOARCH=amd64 GOOS=linux go build -v -o bin/dp-x86_64-linux darkpassenger.go
	GOARCH=386 GOOS=linux go build -v -o bin/dp-x86_32-linux darkpassenger.go
	
build-linux-arm: proto
	GOARCH=arm GOARM=7 GOOS=linux go build -v -o bin/dp-armv7-linux darkpassenger.go
	GOARCH=arm GOARM=6 GOOS=linux go build -v -o bin/dp-armv6-linux darkpassenger.go
	GOARCH=arm GOARM=5 GOOS=linux go build -v -o bin/dp-armv5-linux darkpassenger.go

build-darwin: proto
	GOARCH=amd64 GOOS=darwin go build -v -o bin/dp-x86_64-darwin darkpassenger.go

build-windows: proto
	GOARCH=amd64 GOOS=windows go build -v -o bin/dp-x86_64-windows.exe darkpassenger.go
	GOARCH=386 GOOS=windows go build -v -o bin/dp-x86_32-windows.exe darkpassenger.go

test:
	go test ./...

clean:
	rm -r bin/*
