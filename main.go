package main

import (
	"github.com/jmscatena/ecert-back-go/database"
	"github.com/jmscatena/ecert-back-go/server"
)

func main() {
	database.Init()
	r := server.NewServer()
	r.Run()
}
