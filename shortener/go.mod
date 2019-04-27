module shortener

go 1.12

replace common => ../common

require (
	common v0.0.0-00010101000000-000000000000
	github.com/armon/circbuf v0.0.0-20190214190532-5111143e8da2 // indirect
	github.com/go-ozzo/ozzo-validation v3.5.0+incompatible
	github.com/golang/protobuf v1.3.1
	github.com/jinzhu/configor v1.0.0 // indirect
	github.com/jinzhu/gorm v1.9.4
	github.com/mattn/go-colorable v0.1.1 // indirect
	github.com/micro/go-micro v1.1.0
	github.com/thoas/go-funk v0.4.0 // indirect
)
