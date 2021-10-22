package main

import (
	"github/Re44e/civoo/infrastructure/storage/pg"
)
import "github/Re44e/civoo/infrastructure"

func main() {
	pg.StartDb()
	server := infrastructure.NewServer()
	server.Run()
}