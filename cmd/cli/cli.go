package cli

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func RunCLI() {
	app := cli.NewApp()
	app.Name = "Go Bolaget"
	app.Usage = "search bolaget store information"

	app.Commands = []cli.Command{
		{
			Name:        "search",
			HelpName:    "search",
			Action:      Search,
			ArgsUsage:   ` `,
			Usage:       `Give a query to search for a store`,
			Description: `Search for a store given a fuzzy query.`,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "query",
					Usage:    "query to search for.",
					Required: true,
				},
				&cli.StringFlag{
					Name:     "apikey",
					Usage:    "api key (if no env file)",
					Required: false,
				},
			},
		},
		{
			Name:        "openinghours",
			HelpName:    "openinghours",
			Action:      OpeningHours,
			ArgsUsage:   ` `,
			Usage:       `Provided the id of the store it returns the opening hours`,
			Description: `Get opening hours for a specific stor.`,
			Flags: []cli.Flag{
				&cli.IntFlag{
					Name:     "storeid",
					Usage:    "Id for the store to get.",
					Required: true,
				},
				&cli.IntFlag{
					Name:  "daysahead",
					Usage: "Days in future to get open hours (0 = today).",
					Value: 0,
				},
				&cli.StringFlag{
					Name:     "apikey",
					Usage:    "api key (if no env file)",
					Required: false,
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
