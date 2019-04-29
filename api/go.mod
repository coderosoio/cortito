module api

go 1.12

replace common => ../common

replace account => ../account

replace shortener => ../shortener

require (
	account v0.0.0-00010101000000-000000000000
	dmitri.shuralyov.com/html/belt v0.0.0-20180602232347-f7d459c86be0 // indirect
	dmitri.shuralyov.com/state v0.0.0-20180228185332-28bcc343414c // indirect
	github.com/99designs/gqlgen v0.7.2 // indirect
	github.com/brianvoe/gofakeit v3.17.0+incompatible
	github.com/gin-contrib/cors v0.0.0-20190424000812-bd1331c62cae
	github.com/gin-gonic/gin v1.3.0
	github.com/go-ozzo/ozzo-validation v3.5.0+incompatible
	github.com/micro/cli v0.1.0
	github.com/micro/go-micro v1.1.0
	github.com/micro/go-web v1.0.0
	github.com/micro/micro v1.1.1 // indirect
	github.com/neelance/astrewrite v0.0.0-20160511093645-99348263ae86 // indirect
	github.com/neelance/sourcemap v0.0.0-20151028013722-8c68805598ab // indirect
	github.com/shurcooL/component v0.0.0-20170202220835-f88ec8f54cc4 // indirect
	github.com/shurcooL/events v0.0.0-20181021180414-410e4ca65f48 // indirect
	github.com/shurcooL/github_flavored_markdown v0.0.0-20181002035957-2122de532470 // indirect
	github.com/shurcooL/gopherjslib v0.0.0-20160914041154-feb6d3990c2c // indirect
	github.com/shurcooL/home v0.0.0-20190204141146-5c8ae21d4240 // indirect
	github.com/shurcooL/httperror v0.0.0-20170206035902-86b7830d14cc // indirect
	github.com/shurcooL/httpgzip v0.0.0-20180522190206-b1c53ac65af9 // indirect
	github.com/shurcooL/users v0.0.0-20180125191416-49c67e49c537 // indirect
	github.com/sourcegraph/annotate v0.0.0-20160123013949-f4cad6c6324d // indirect
	github.com/sourcegraph/syntaxhighlight v0.0.0-20170531221838-bd320f5d308e // indirect
	github.com/stretchr/testify v1.3.0
	k8s.io/utils v0.0.0-20190204185745-a326ccf4f02b // indirect
	shortener v0.0.0-00010101000000-000000000000
	sourcegraph.com/sourcegraph/go-diff v0.5.0 // indirect
)
