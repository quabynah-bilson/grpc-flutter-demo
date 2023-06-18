PROTO_PATH=proto  # proto directory
SERVER_DIR=server # server directory

# output directory for generated proto code
mkdir -p $SERVER_DIR/gen

# generate proto code for golang
protoc -I=$PROTO_PATH --go_out=$SERVER_DIR/gen --go_opt=paths=source_relative \
  --go-grpc_out=$SERVER_DIR/gen --go-grpc_opt=paths=source_relative \
  $(find $PROTO_PATH -name '*.proto')
