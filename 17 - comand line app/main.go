package main

import (
	"log"
	"os"

	"github.com/alexandersantosdev/comandlinego/app"
)

func main() {

	application := app.Generate()
	if err := application.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
