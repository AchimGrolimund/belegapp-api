VERSION=0.0.2

.Phony: build-docker
build-docker:
	docker buildx build --platform=linux/amd64,linux/arm64,linux/arm/v6,linux/arm/v7,linux/386 -t grolimundachim/test:${VERSION} . --push

.Phony: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags "-w -s" -a -o ./build/main .

.Phony: clean
clean:
	rm -rf ./build

.Phony: test
test:
	go test -v ./... -coverprofile=coverage.out -covermode=atomic && go tool cover -html=coverage.out -o coverage.html


