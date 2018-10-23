package main

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/urfave/cli"
	"log"
	"os"
)

const ENDPOINT = "https://ordbok.uib.no/perl/ordbok.cgi?OPP="

func getExplanation(word string) {

	url := ENDPOINT + word
	doc, err := htmlquery.LoadURL(url)
	if err != nil {
		fmt.Println(err)
	}

	for _, n := range htmlquery.Find(doc, `//*[@id="byttutBM"]/tbody/tr[2]/td[2]/div/div[2]`) {
		a := htmlquery.FindOne(n, "//span[2]")
		fmt.Printf("%s\n", htmlquery.InnerText(a))

		// Example path:
		// 2888"]/div[2]/span[6]/div[1]
		list := htmlquery.Find(n, "//span")
		for _, element := range list[4:8] {
			b := htmlquery.Find(element, "//div")

			// TODO: Filter "kompakt" or "utvidet"
			for _, c := range b {
				fmt.Printf("%s\n", htmlquery.InnerText(c))
			}
		}
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "nob"
	app.Usage = "get explanation for words"

	app.Action = func(c *cli.Context) error {
		word := c.Args().Get(0)
		getExplanation(word)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
