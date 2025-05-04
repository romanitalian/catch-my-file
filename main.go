package main

import (
	"github.com/romanitalian/catch-my-file/internal/app"
	"github.com/romanitalian/catch-my-file/internal/config"
)

func main() {
	cfg := config.New()
	application := app.New(cfg)
	application.Run()
}
