module web

go 1.12

require (
	common v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.3.0
	shortener v0.0.0-00010101000000-000000000000
)

replace shortener => ../shortener

replace common => ../common
