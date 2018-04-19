
gen:
	protoc -I/usr/local/include -I. \
	  -I$(GOPATH)/src \
	  -I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	  --go_out=plugins=grpc:. \
	  --grpc-gateway_out=logtostderr=true:. \
	  ./proto/bank.proto

	  mv ./proto/*.go ./api
