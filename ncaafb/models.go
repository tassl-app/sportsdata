package ncaafb

import (
	"github.com/tassl/sportsdata"
)

type Team struct {
	Id       string `xml:"id,attr"`
	Name     string `xml:"name,attr"`
	Market   string `xml:"market,attr"`
	Coverage string `xml:"coverage,attr"`
}

type Subdivision struct {
	Id    string `xml:"id,attr"`
	Name  string `xml:"name,attr"`
	Teams []Team `xml:"team"`
}

type Conference struct {
	Id           string        `xml:"id,attr"`
	Name         string        `xml:"name,attr"`
	Subdivisions []Subdivision `xml:"subdivision"`
	Teams        []Team        `xml:"team"`
}

type Division struct {
	XMLNS       string       `xml:"xmlns,attr"`
	Id          string       `xml:"id,attr"`
	Name        string       `xml:"name,attr"`
	Conferences []Conference `xml:"conference"`
}

type Link struct {
	Rel      string `xml:"rel,attr"`
	Href     string `xml:"href,attr"`
	LinkType string `xml:"link,attr"`
}

type Links struct {
	Links []Link `xml:"link"`
}

type Broadcast struct {
	Network   string `xml:"network,attr"`
	Satellite string `xml:"satellite,attr"`
	Internet  string `xml:"internet,attr"`
	Cable     string `xml:"cable,attr"`
}

type Wind struct {
	Speed     string `xml:"speed,attr"`
	Direction string `xml:"direction,attr"`
}

type Weather struct {
	Temperature string `xml:"temperature,attr"`
	Condition   string `xml:"condition,attr"`
	Humidity    string `xml:"humidty,attr"`
	Wind        Wind   `xml:"wind"`
}

type Venue struct {
	Id        string `xml:"id,attr"`
	Name      string `xml:"name,attr"`
	Address   string `xml:"address,attr"`
	City      string `xml:"city,attr"`
	State     string `xml:"state,attr"`
	Zip       string `xml:"zip,attr"`
	Country   string `xml:"country,attr"`
	Capacity  string `xml:"capacity,attr"`
	Surface   string `xml:"surface,attr"`
	VenueType string `xml:"type,attr"`
}

type Game struct {
	Id           string            `xml:"id,attr"`
	Scheduled    string            `xml:"scheduled,attr"`
	Coverage     string            `xml:"coverage,attr"`
	HomeRotation string            `xml:"home_rotation,attr"`
	AwayRotation string            `xml:"away_rotation,attr"`
	HomeTeamId   string            `xml:"home,attr"`
	AwayTeamId   string            `xml:"away,attr"`
	Status       string            `xml:"status,attr"`
	Venue        *sportsdata.Venue `xml:"venue"`
	Broadcast    *Broadcast        `xml:"broadcast"`
	Links        Links             `xml:"links"`
}

type Week struct {
	Week  string `xml:"week,attr"`
	Games []Game `xml:"game"`
}

type Season struct {
	XMLNS      string `xml:"xmlns,attr"`
	Season     string `xml:"season,attr"`
	SeasonType string `xml:"type,attr"`
	Weeks      []Week `xml:"week"`
}

type Schedule struct {
	Year         string
	ScheduleType ScheduleType
	Season       *Season
}

func (s *Schedule) Venues() []*sportsdata.Venue {
	venues := make([]*sportsdata.Venue, 0)
	for _, week := range s.Season.Weeks {
		for _, game := range week.Games {
			if game.Venue != nil {
				venues = append(venues, game.Venue)
			}
		}
	}
	return venues
}
