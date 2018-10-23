package main

import (
	"fmt"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var endpoint = "https://ordbok.uib.no/perl/ordbok.cgi?OPP="

func getExplanation(word string) string {
	url := endpoint + word
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	return string(body)
}

func main() {
	app := cli.NewApp()
	app.Name = "nob"
	app.Usage = "get explanation for words"

	app.Action = func(c *cli.Context) error {
		word := c.Args().Get(0)
		explanation := getExplanation(word)
		fmt.Println(explanation)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
