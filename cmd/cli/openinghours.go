package cli

import (
	"log"

	"go-bolaget/pkg/bolaget"

	"github.com/urfave/cli"
)

func OpeningHours(c *cli.Context) error {
	storeId := c.Int("storeid")
	daysAhead := c.Int("daysahead")
	apiKey := c.String("apiKey")

	openingHour, err := bolaget.GetStoreOpeningHours(storeId, daysAhead, apiKey)
	if err != nil {
		log.Fatal(err)
	}

	println(openingHour.FullDescription)
	return nil
}
