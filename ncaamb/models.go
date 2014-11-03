package ncaamb

import (
	"github.com/tassl/sportsdata"
)

type Team struct {
	Id     string            `xml:"id,attr"`
	Name   string            `xml:"name,attr"`
	Market string            `xml:"market,attr"`
	Alias  string            `xml:"alias,attr"`
	Venue  *sportsdata.Venue `xml:"venue"`
}

type Conference struct {
	Id    string `xml:"id,attr"`
	Name  string `xml:"name,attr"`
	Alias string `xml:"alias,attr"`
	Teams []Team `xml:"team"`
}

type Division struct {
	Id         string       `xml:"id,attr"`
	Name       string       `xml:"name,attr"`
	Alias      string       `xml:"alias,attr"`
	Conference []Conference `xml:"conference"`
}

type HomeTeam struct {
	Id    string `xml:"id,attr"`
	Name  string `xml:"name,attr"`
	Alias string `xml:"alias,attr"`
}

type AwayTeam struct {
	Id    string `xml:"id,attr"`
	Name  string `xml:"name,attr"`
	Alias string `xml:"alias,attr"`
}

type Game struct {
	Id         string   `xml:"id,attr"`
	Status     string   `xml:"status,attr"`
	Converage  string   `xml:"coverage,attr"`
	HomeTeamId string   `xml:"home_team,attr"`
	AwayTeamId string   `xml:"away_team,attr"`
	Scheduled  string   `xml:"scheduled,attr"`
	HomeTeam   HomeTeam `xml:"home"`
	AwayTeam   AwayTeam `xml:"away"`
}

type Games struct {
	Game []Game `xml:"game"`
}

type SeasonSchedule struct {
	Id         string `xml:"id,attr"`
	Year       string `xml:"year,attr"`
	SeasonType string `xml:"type,attr"`
	Games      Games  `xml:"games"`
}

type League struct {
	XMLNS          string         `xml:"xmlns,attr"`
	Id             string         `xml:"id,attr"`
	Name           string         `xml:"name,attr"`
	Alias          string         `xml:"alias,attr"`
	Division       []Division     `xml:"division"`
	SeasonSchedule SeasonSchedule `xml:"season-schedule"`
}

type Schedule struct {
	Season       string
	ScheduleType ScheduleType
	League       *League
}

func (s *Schedule) Venues() []*sportsdata.Venue {
	venues := make([]*sportsdata.Venue, 0)
	for _, division := range s.League.Division {
		for _, conference := range division.Conference {
			for _, team := range conference.Teams {
				if team.Venue != nil {
					venues = append(venues, team.Venue)
				}
			}
		}
	}
	return venues
}
