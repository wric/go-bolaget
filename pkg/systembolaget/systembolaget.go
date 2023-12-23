package systembolaget

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

const (
	API_ROOT = "https://api-extern.systembolaget.se/site/V2"
)

type SystembolagetOpeningHour struct {
	Date     string `json:"date"`
	OpenFrom string `json:"openFrom"`
	OpenTo   string `json:"openTo"`
	Reason   string `json:"reason"`
}

type SystembolagetSite struct {
	SiteId         string                     `json:"siteId"`
	Alias          string                     `json:"alias"`
	StreetAddress  string                     `json:"streetAddress"`
	DisplayName    string                     `json:"displayName"`
	City           string                     `json:"city"`
	County         string                     `json:"county"`
	Postalcode     string                     `json:"postalcode"`
	IsTastingStore bool                       `json:"isTastingStore"`
	IsAgent        bool                       `json:"isAgent"`
	IsOpen         bool                       `json:"isOpen"`
	IsBlocked      bool                       `json:"isBlocked"`
	OpeningHours   []SystembolagetOpeningHour `json:"openingHours"`
	Position       Position                   `json:"position"`
}

type SystembolagetStore struct {
	SiteId                     string                     `json:"siteId"`
	Alias                      string                     `json:"alias"`
	IsActive                   bool                       `json:"isActive"`
	IsBlocked                  bool                       `json:"isBlocked"`
	IsOpen                     bool                       `json:"isOpen"`
	IsBlockedByOrderLimit      bool                       `json:"isBlockedByOrderLimit"`
	MaxOrdersPerDay            int                        `json:"maxOrdersPerDay"`
	OrdersToday                int                        `json:"ordersToday"`
	Address                    string                     `json:"address"`
	PostalCode                 string                     `json:"postalCode"`
	City                       string                     `json:"city"`
	Phone                      string                     `json:"phone"`
	County                     string                     `json:"county"`
	IsFullAssortmentOrderStore bool                       `json:"isFullAssortmentOrderStore"`
	IsTastingStore             bool                       `json:"isTastingStore"`
	Position                   Position                   `json:"position"`
	OpeningHours               []SystembolagetOpeningHour `json:"openingHours"`
	ParentSiteId               string                     `json:"parentSiteId"`
	SearchArea                 string                     `json:"searchArea"`
}

type Position struct {
	Latitude  float64 `json:"Latitude"`
	Longitude float64 `json:"Longitude"`
}

type SystembolagetSearch struct {
	SiteViewModel []SystembolagetSite `json:"siteViewModel"`
}

type ApiError struct {
	Error []string `json:"error"`
}

func ApiGet(path string, apiKey string) ([]byte, error) {
	if apiKey == "" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Panic("No api key specified and failt to load .env file. Is it missing?")
		}

		apiKey = os.Getenv("API_KEY")
	}

	req, err := http.NewRequest("GET", API_ROOT+path, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Ocp-Apim-Subscription-Key", apiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		var apiError ApiError
		err = json.Unmarshal(body, &apiError)

		if err != nil {
			return nil, err
		}

		return nil, errors.New("api error: " + strings.Join(apiError.Error, ", "))
	}

	return body, nil
}
