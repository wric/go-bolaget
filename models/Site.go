package models

type BolagetSite struct {
	siteId       string
	name         string
	address      string
	position     string
	openingHours []BolagetOpeningHour
}

type SystembolagetSite struct {
	siteId         string
	alias          string
	streetAddress  string
	displayName    string
	city           string
	county         string
	postalcode     string
	isTastingStore bool
	isAgent        bool
	isOpen         bool
	isBlocked      bool
	openingHours   []SystembolagetOpeningHour
	position       Position
}

type SystembolagetStore struct {
	siteId                     string
	alias                      string
	isActive                   bool
	isBlocked                  bool
	isOpen                     bool
	isBlockedByOrderLimit      bool
	maxOrdersPerDay            int
	ordersToday                int
	address                    string
	postalCode                 string
	city                       string
	phone                      string
	county                     string
	isFullAssortmentOrderStore bool
	isTastingStore             bool
	position                   Position
	openingHours               []SystembolagetOpeningHour
	parentSiteId               string
	searchArea                 string
}

type Position struct {
	latitude  float32
	longitude float32
}
