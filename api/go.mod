module api

go 1.12

replace common => ../common

replace account => ../account

replace shortener => ../shortener

require (
	account v0.0.0-00010101000000-000000000000
	common v0.0.0-00010101000000-000000000000
	github.com/brianvoe/gofakeit v3.17.0+incompatible
	github.com/gin-gonic/gin v1.3.0
	github.com/go-ozzo/ozzo-validation v3.5.0+incompatible
	github.com/micro/cli v0.1.0
	github.com/micro/go-micro v1.1.0
	github.com/micro/go-web v1.0.0
	github.com/stretchr/testify v1.3.0
	shortener v0.0.0-00010101000000-000000000000
)
