package bolaget

import (
	"encoding/json"
	"errors"
	"go-bolaget/pkg/systembolaget"
	"strconv"
)

type BolagetOpeningHour struct {
	IsClosed         bool
	IsDeviant        bool
	ShortDescription string
	FullDescription  string
}

type BolagetSite struct {
	SiteId       string
	Name         string
	Address      string
	Position     string
	OpeningHours []BolagetOpeningHour
}

func notAgentNorBlocked(s *systembolaget.SystembolagetSite) bool {
	return !s.IsAgent && !s.IsBlocked
}

func SearchStore(query string, apiKey string) ([]BolagetSite, error) {
	body, err := systembolaget.ApiGet("/Search/Site?q="+query, apiKey)
	if err != nil {
		return nil, err
	}

	var searchResult systembolaget.SystembolagetSearch
	err = json.Unmarshal(body, &searchResult)
	if err != nil {
		return nil, err
	}

	filteredBolagetSites := []BolagetSite{}
	for _, site := range searchResult.SiteViewModel {
		if notAgentNorBlocked(&site) {
			bolagetSite := MapSite(&site)
			filteredBolagetSites = append(filteredBolagetSites, bolagetSite)
		}
	}

	return filteredBolagetSites, nil
}

func GetStoreOpeningHours(storeId int, daysAhead int, apikey string) (BolagetOpeningHour, error) {
	if daysAhead < 0 {
		return BolagetOpeningHour{}, errors.New("'daysAhead' must be >= 0")
	}

	body, err := systembolaget.ApiGet("/Store/"+strconv.Itoa(storeId), apikey)
	if err != nil {
		return BolagetOpeningHour{}, err
	}

	var site systembolaget.SystembolagetSite
	err = json.Unmarshal(body, &site)
	if err != nil {
		return BolagetOpeningHour{}, err
	}

	if len(site.OpeningHours) == 0 {
		return BolagetOpeningHour{}, errors.New("no opening hours returned")
	}

	if daysAhead >= len(site.OpeningHours) {
		return BolagetOpeningHour{}, errors.New("'daysAhead' out of range")
	}

	openingHour := site.OpeningHours[daysAhead]
	name := ParseSiteName(&site)
	bolagetOpeningHour := MapOpeningHour(&openingHour, name)

	return bolagetOpeningHour, nil
}
