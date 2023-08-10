package main

import (
	"log"
	"os"

	. "github.com/pulkit-tanwar/dispense-quotes/lib/quotes"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "dispenseQuotes"
	app.Usage = "Dispense a random programming quote"
	app.UsageText = "quotes dispense"
	app.Version = Version

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
	q := NewQuote("Science is what we understand well enough to explain to a computer, Art is all the rest.	", "Donald Knuth")
	log.Print(q)
	return nil
}
