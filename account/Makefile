LDFLAGS := -ldflags="-s -w"
TEST := `go list ./... | awk '!/cmd/ && !/model/ && !/proto/'`

build: proto
	go build ${LDFLAGS} -o bin/account cmd/account/main.go

run:
	dogo

proto:
	protoc --proto_path=. \
				 --micro_out=. \
				 --go_out=. \
				 proto/**/*.proto

test:
	APP_ENV=test go test -v -p=1 -coverprofile=coverage.out ${TEST}
	go tool cover -html=coverage.out -o coverage.html

.PHONY: build proto

