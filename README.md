# gRPC echo
echo gRPC service written in Golang

## Build gRPC echo service go package
```bash
 protoc --go_out=./pkg/grpc/echo/ \
        --go_opt=paths=source_relative \
        --go-grpc_out=./pkg/grpc/echo \
        --go-grpc_opt=paths=source_relative \
        --proto_path=$PWD/definition \
            $PWD/definition/echo.proto
```

## Build echo service image
```bash
docker build  --platform linux/amd64 -t fgarciacode/grpc-echo:v0.1.2 .
```

## Run echo service
```bash
docker run --name echo-service -d -p 5001:5001 fgarciacode/grpc-echo:v0.1.2
```