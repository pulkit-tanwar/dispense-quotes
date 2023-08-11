package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/pulkit-tanwar/dispense-quotes/lib/config"
	"github.com/pulkit-tanwar/dispense-quotes/lib/constants"
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
		{
			Name:      "dispense",
			Usage:     "dispense a single programming quote",
			UsageText: "quote dispense",
			Action:    DispenseQuote,
		},
		{
			Name:      "serve",
			Usage:     "Run and API Server to dispense a programming quotes.",
			UsageText: "dispense serve -e ENV -p PORT --host HOST --api-path API_PATH",
			Action:    RunHTTPServer,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "env,e",
					Value:  constants.DefaultEnv,
					Usage:  "environment (dev | test | stage | prod)",
					EnvVar: "ENV",
				},
				cli.StringFlag{
					Name:   "host",
					Value:  constants.DefaultHost,
					Usage:  "host to listen on",
					EnvVar: "HOST",
				},
				cli.IntFlag{
					Name:   "port,p",
					Value:  constants.DefaultPort,
					Usage:  "port to listen on",
					EnvVar: "PORT",
				},
				cli.StringFlag{
					Name:   "api-path",
					Value:  constants.DefaultAPIPath,
					Usage:  "url path prefix for mounting API router",
					EnvVar: "API_PATH",
				},
			},
		},
	}

	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

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

func RunHTTPServer(c *cli.Context) error {
	cfg := config.NewConfig(c.String("env"), c.String("host"), c.Int("port"), c.String("api-path"))

	log.Infof("ENV: %s", cfg.Env)
	log.Infof("HOST: %s", cfg.Host)
	log.Infof("PORT: %d", cfg.Port)
	log.Infof("API_PATH: %s", cfg.APIPath)
	return nil
}
