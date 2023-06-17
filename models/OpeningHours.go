package models

type BolagetOpeningHour struct {
	IsClosed         bool
	IsDeviant        bool
	ShortDescription string
	FullDescription  string
}

type SystembolagetOpeningHour struct {
	Date     string
	OpenFrom string
	OpenTo   string
	Reason   string
}
