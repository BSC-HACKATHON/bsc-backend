PROTO_DIR=./
OUT_DIR=./

proto:
	mkdir -p $(OUT_DIR)
	protoc --go_out=$(OUT_DIR) --go-grpc_out=$(OUT_DIR) $(PROTO_DIR)/*.proto

clean:
	rm -rf $(OUT_DIR)

