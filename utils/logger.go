package utils

import (
	"log"

	"go.uber.org/zap"
)

func InitLogger() {
	c := zap.NewProductionConfig()
	l, err := c.Build()
	if err != nil {
		log.Fatal(err)
	}
	zap.ReplaceGlobals(l)
}
