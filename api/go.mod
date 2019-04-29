module api

go 1.12

replace common => ../common

replace account => ../account

replace shortener => ../shortener

require (
	account v0.0.0-00010101000000-000000000000
	common v0.0.0-00010101000000-000000000000
	github.com/99designs/gqlgen v0.7.2 // indirect
	github.com/brianvoe/gofakeit v3.17.0+incompatible
	github.com/gin-contrib/cors v0.0.0-20190424000812-bd1331c62cae
	github.com/gin-gonic/gin v1.3.0
	github.com/go-ozzo/ozzo-validation v3.5.0+incompatible
	github.com/micro/cli v0.1.0
	github.com/micro/go-micro v1.1.0
	github.com/micro/go-web v1.0.0
	github.com/micro/micro v1.1.1 // indirect
	github.com/shurcooL/home v0.0.0-20190204141146-5c8ae21d4240 // indirect
	github.com/stretchr/testify v1.3.0
	k8s.io/utils v0.0.0-20190204185745-a326ccf4f02b // indirect
	shortener v0.0.0-00010101000000-000000000000
)
