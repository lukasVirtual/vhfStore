package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"www.github.com/vhfStore/internals/handler"
	"www.github.com/vhfStore/internals/router"
	application "www.github.com/vhfStore/internals/usecases/Application"
	"www.github.com/vhfStore/utils"
)

func init() {
	log.SetOutput(os.Stdout)
	// log.SetFormatter(&log.JSONFormatter{})

	log.SetLevel(log.DebugLevel)
}

func main() {
	utils.GetSystemInformation()
	r := router.New()
	a := application.New()

	h := handler.New(a)
	h.Register(r)

	if err := r.Listen(":2173"); err != nil {
		log.Fatal(err)
	}
}
