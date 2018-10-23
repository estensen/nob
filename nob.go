package main

import (
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "nob"
	app.Usage = "get explanation for words"

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
