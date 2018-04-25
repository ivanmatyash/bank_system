NETWORK := net
POSTGRES_IMAGE_NAME := postgres:10.3-alpine

build:
	CGO_ENABLED=0 go build -o ./bin/bank ./src/bank
	CGO_ENABLED=0 go build -o ./bin/gw ./src/gw

image: build
	docker build -f ./docker/Dockerfile.bank -t bank:latest .
	docker build -f ./docker/Dockerfile.gw -t gw-bank:latest .

vendor:
	dep ensure

init:
	docker network create $(NETWORK)
	docker create --name pg-bank -p "5432:5432" -v $(CURDIR)/sqlstore:/tmp --network $(NETWORK) --network-alias pg-bank.$(NETWORK) $(POSTGRES_IMAGE_NAME)

pg-start:
	docker start pg-bank

up: image pg-start
	docker start pg-bank
	docker run --name bank_microservice -d -p "9091:91" --network $(NETWORK) --network-alias bank.$(NETWORK) bank:latest
	docker run --name gw_bank -d -p "8080:80" --network $(NETWORK) --network-alias gw-bank.$(NETWORK) gw-bank:latest

db-create:
	@echo "Creating database..."
	@docker exec pg-bank /tmp/db_create

down:
	docker stop gw_bank || true
	docker rm -f gw_bank || true
	docker image rm gw-bank || true

	docker stop bank_microservice || true
	docker rm -f bank_microservice || true
	docker image rm bank || true

clean:
	docker stop gw_bank || true
	docker rm -f gw_bank || true
	docker image rm gw-bank || true

	docker stop bank_microservice || true
	docker rm -f bank_microservice || true
	docker image rm bank || true

	docker stop pg-bank || true
	docker rm -f pg-bank || true
	
	docker network rm $(NETWORK) || true
	rm -rf ./bin

gen:
	protoc -I/usr/local/include -I. \
	  -I$(GOPATH)/src \
	  -I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	  --go_out=plugins=grpc:. \
	  --grpc-gateway_out=logtostderr=true:. \
	  ./proto/bank.proto
	  protoc-go-inject-tag -input=./proto/bank.pb.go

	  mv ./proto/*.go ./api
