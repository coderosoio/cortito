package handler

import (
	"time"

	"github.com/brianvoe/gofakeit"
)

func init() {
	gofakeit.Seed(time.Now().UnixNano())
}
