module api

go 1.12

replace common => ../common

replace account => ../account

replace shortener => ../shortener

require (
	account v0.0.0-00010101000000-000000000000
	github.com/gin-contrib/sse v0.0.0-20190301062529-5545eab6dad3 // indirect
	github.com/gin-gonic/gin v1.3.0
	github.com/go-ozzo/ozzo-validation v3.5.0+incompatible
	github.com/micro/go-web v1.0.0 // indirect
	gopkg.in/go-playground/validator.v8 v8.18.2 // indirect
	shortener v0.0.0-00010101000000-000000000000
)
