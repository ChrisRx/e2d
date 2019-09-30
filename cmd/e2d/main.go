package main

import (
	"github.com/criticalstack/e2d/cmd/e2d/app"
	"github.com/criticalstack/e2d/pkg/log"
)

func main() {
	if err := app.NewRootCmd().Execute(); err != nil {
		log.Fatal(err)
	}
}
