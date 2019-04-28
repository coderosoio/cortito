LDFLAGS := -ldflags="-s -w"
TEST := `go list ./... | awk '!/cmd/ && !/helper/ && !/middleware/'`

build:
	go build ${LDFLAGS} -o bin/api cmd/api/main.go

run:
	dogo

test:
	export APP_ENV=test
	cd ../account && dogo &
	cd ../shortener && dogo &
	cd ../api
	go test -v -p=1 -coverprofile=coverage.out ${TEST}
	go tool cover -html=coverage.out -o coverage.html

.PHONY: build

