package main

import (
	"log"
	"os"

	"github.com/var512/ddda-save-corrupter/internal/api"
	"github.com/var512/ddda-save-corrupter/internal/logger"
)

var AppVersion = "dev"

func main() {
	logger.Default = log.New(os.Stdout, "[api] ", log.LstdFlags)
	logger.Default.Println("============================================================")
	logger.Default.Println("== DDDA Save Corrupter")
	logger.Default.Println("============================================================")
	logger.Default.Println("== [Version] " + AppVersion)
	logger.Default.Println("============================================================")
	api.Start()
}
