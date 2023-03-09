package main

import (
	"log"

	"www.github.com/vhfStore/internals/handler"
	"www.github.com/vhfStore/internals/router"
	application "www.github.com/vhfStore/internals/usecases/Application"
	"www.github.com/vhfStore/utils"
)

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
