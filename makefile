SRC_DIR=./core/
DST_DIR=./core/


proto:
	protoc -I=$(SRC_DIR) --go_out=plugins=grpc:$(DST_DIR) $(SRC_DIR)/drone.proto

build:
	go build -race -o ./bin/drone ./cmd/main.go
