run:  
	go run ./cmd/twitter | jq 

fmt:
	goimports -w .

lint:
	golangci-lint run

test:
	go test ./...

test-e2e:
	docker compose up -d --wait
	go test --tags=e2e ./...
	docker compose down

mocks:
	minimock -s .go -i github.com/demimurg/twitter/internal/usecase.* -o ./internal/usecase/mock

proto:
	buf generate

# prepare for commit
prepare: fmt lint test

# generate mocks for public interfaces and protobuf for api
generate: mocks proto