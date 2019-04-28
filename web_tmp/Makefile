LDFLAGS := -ldflags="-s -w"
TEST := `go list ./... | awk '!/cmd/ && !/static/ && !/template/'`

build:
	go build ${LDFLAGS} -o bin/web cmd/web/main.go

test:
	APP_ENV=test go test -v -p=1 -coverprofile=coverage.out ${TEST}
	go tool cover -html=coverage.out -o coverage.html

.PHONY: build
