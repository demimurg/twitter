mocks:
	minimock -s .go -i github.com/demimurg/twitter/internal/usecase.* -o ./internal/usecase/mock

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pkg/proto/twitter.proto

generate: mocks proto