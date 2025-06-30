CLIENT_CMD=./cmd/client
SERVER_CMD=./cmd/server
BIN_DIR=./bin
CONFIG=./internal/config/config.yml

CLIENT_BIN=$(BIN_DIR)/client
SERVER_BIN=$(BIN_DIR)/server

.PHONY: all build run-client run-server clean test lint

build: $(CLIENT_BIN) $(SERVER_BIN)

$(CLIENT_BIN): $(CLIENT_CMD)/*.go
	go build -o $(CLIENT_BIN) $(CLIENT_CMD)

$(SERVER_BIN): $(SERVER_CMD)/*.go
	go build -o $(SERVER_BIN) $(SERVER_CMD)

run-client: $(CLIENT_BIN)
	$(CLIENT_BIN) --config-file=$(CONFIG)

run-server: $(SERVER_BIN)
	$(SERVER_BIN) --config-file=$(CONFIG)

clean:
	rm -rf $(BIN_DIR)

test:
	go test ./...

lint:
	golangci-lint run ./...
