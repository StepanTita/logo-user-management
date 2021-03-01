package main

import (
	"github.com/logo-user-management/app"
	"github.com/logo-user-management/config"
	"os"
)

func main() {
	cfg := config.New(os.Getenv("CONFIG"))

	if err := app.New(cfg).Run(); err != nil {
		panic(err)
	}
}