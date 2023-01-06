run:  
	go run ./cmd/twitter | jq 

fmt:
	goimports -w .

lint:
	golangci-lint run

test:
	go test ./...

mocks:
	minimock -s .go -i github.com/demimurg/twitter/internal/usecase.* -o ./internal/usecase/mock

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pkg/proto/twitter.proto

# prepare for commit
prepare: fmt lint test

# generate mocks for public interfaces and protobuf for api
generate: mocks proto