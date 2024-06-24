#!/bin/bash


export GOPATH=$HOME/go;
export GO_LIBS=$HOME/go_libs;
export PATH=$PATH:$GOPATH/bin;
export PATH=$PATH:$GO_LIBS/bin;
protoc -I ./protobuf \
  --go_out=./pkg/proto --go_opt=paths=source_relative \
  --go-grpc_out=./pkg/proto --go-grpc_opt=paths=source_relative \
  --grpc-gateway_out=./pkg/proto --grpc-gateway_opt=paths=source_relative \
  --openapiv2_out=./swagger-ui --openapiv2_opt=use_go_templates=true \
  protobuf/keeper.proto
