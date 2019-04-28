module web

go 1.12

replace common => ../common

replace account => ../account

replace shortener => ../shortener

require (
	account v0.0.0-00010101000000-000000000000
	common v0.0.0-00010101000000-000000000000
	github.com/foolin/gin-template v0.0.0-20190415034731-41efedfb393b
	github.com/gin-contrib/sessions v0.0.0-20190226023029-1532893d996f
	github.com/gin-gonic/gin v1.3.0
	shortener v0.0.0-00010101000000-000000000000
)
