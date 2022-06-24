.PHONY : mac windows linux pi all mkdir
mac: perpare
	 CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -trimpath -o ./build/nAgent_Mac/nAgent nAgent.go

windows: perpare
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -trimpath -o ./build/nAgent_Windows/nAgent.exe nAgent.go

linux: perpare
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -trimpath -o ./build/nAgent_Linux/nAgent nAgent.go

pi: perpare
	CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -ldflags "-s -w" -trimpath -o ./build/nAgent_RaspberryPi/nAgent nAgent.go
all: mac windows linux pi
perpare:
	mkdir -p ./build/
	mkdir -p ./build/nAgent_Mac
	mkdir -p ./build/nAgent_Windows
	mkdir -p ./build/nAgent_Linux
	mkdir -p ./build/nAgent_RaspberryPi