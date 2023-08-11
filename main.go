package main

import (
	"log"
	"os"

	"github.com/pulkit-tanwar/dispense-quotes/lib/quoteslogic"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "dispenseQuotes"
	app.Usage = "Dispense a random programming quote"
	app.UsageText = "quotes dispense"
	app.Version = quoteslogic.Version

	app.Commands = []cli.Command{
		cli.Command{
			Name:      "dispense",
			Usage:     "dispense a single programming quote",
			UsageText: "quote dispense",
			Action:    DispenseQuote,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Print("Error occured: ", err)
		os.Exit(1)
	}
}

func DispenseQuote(c *cli.Context) error {
	dispenser, err := quoteslogic.LoadQuotesFromJSONFile("quotes.json")
	if err != nil {
		log.Print("Error occured while fetching quotes from the json file: ", err)
		os.Exit(1)
	}
	log.Print(dispenser.RandomQuote())
	return nil
}
