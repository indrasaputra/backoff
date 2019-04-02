test:
	go test -v -race ./...

dep:
	dep ensure

vendor:
	env GO111MODULE=on go mod vendor

cover:
	go test -coverprofile=coverage.out && go tool cover -html=coverage.out