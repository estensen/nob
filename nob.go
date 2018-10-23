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
		//*[@id="art22888"]/div[2]
		//*[@id="art22723"]/div[2]
		//*[@id="art17346"]/div[2]


		//*[@id="byttutBM"]/tbody


		//*[@id="byttutBM"]/tbody/tr[2]/td[2]


		a := htmlquery.FindOne(n, "//span[2]")
		fmt.Printf("%s\n", htmlquery.InnerText(a))

		//*[@id="art22892"]/div[2]/span[5]/div[1]/text()
		//*[@id="art22892"]/div[2]/span[5]/div[1]/div/span[1]/span/text()
		//*[@id="art22892"]/div[2]/span[5]/div[1]
		b := htmlquery.FindOne(n, "//span[6]/div")
		fmt.Printf("%s\n", htmlquery.InnerText(b))
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
