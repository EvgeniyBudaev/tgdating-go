SERVER_BINARY=aggregationApp

## build_server: builds the server binary as a linux executable
build_server:
	@echo "Building server binary..."
	chdir ./app && set GOOS=linux && set GOARCH=amd64 && set CGO_ENABLED=0 && go build -o ./bin/${SERVER_BINARY} ./cmd
	@echo "Done!"