package cli

import (
	"errors"
	"go-bolaget/pkg/bolaget"
	"log"
	"strconv"

	"github.com/urfave/cli"
)

func Search(c *cli.Context) error {
	if len(c.Args()) > 0 {
		return errors.New("no arguments is expected, use flags")
	}

	if c.String("query") == "" {
		return errors.New("'query' flag cannot be empty")
	}

	stores, err := bolaget.SearchStore(c.String("query"), c.String("apikey"))
	if err != nil {
		log.Fatal(err)
	}

	if len(stores) == 0 {
		println("found no store.")
		return nil
	}

	hits := len(stores)
	store := stores[0]
	if hits > 0 {
		println(store.Name + " - " + store.SiteId + " (hits: " + strconv.Itoa(hits) + ")")
	} else {
		println(store.Name + " - " + store.SiteId)
	}

	for i, openingHour := range store.OpeningHours {
		println(openingHour.FullDescription)

		if i == 2 {
			break
		}
	}

	return nil
}
