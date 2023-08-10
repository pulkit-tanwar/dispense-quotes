package main

import (
	"log"
	"os"

	"github.com/pulkit-tanwar/dispense-quotes/lib/quotes"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "dispenseQuotes"
	app.Usage = "Dispense a random programming quote"
	app.UsageText = "quotes dispense"
	app.Version = quotes.Version

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
	log.Print("TODO: Dispense Quotes")
	return nil
}
