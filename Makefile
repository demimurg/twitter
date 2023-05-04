# run application with pretty json logs output
run:
	go run ./cmd/twitter | jq 

# format code and imports
fmt:
	goimports -w .

# lint go code with default configuration for golangci-lint
lint:
	golangci-lint run

# run unit-tests for this project with race condition check
test:
	go test -race ./...

# start external docker dependencies and run end-to-end tests for app
test-e2e:
	docker compose up -d --wait
	go test --tags=e2e ./...
	docker compose down

# generate code to mock interfaces from internal/usecase package
mocks:
	minimock -s .go -i github.com/demimurg/twitter/internal/usecase.* -o ./internal/usecase/mock

# generate code from proto files using buf
proto:
	buf generate

# add all necessary helm repositories
helm-repo:
	helm repo add grafana-labs https://grafana.github.io/helm-charts

# fetch new versions and charts for grafana-labs helm repository
helm-update:
	helm repo update grafana-labs

# download dependencies from Chart.yaml to deploy/charts (use if you change version)
helm-download:
	helm dependency update deploy

# prepare for commit
prepare: fmt lint test

# generate mocks for public interfaces and protobuf for api
generate: mocks proto