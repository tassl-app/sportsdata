package ncaawb

import (
	"github.com/tassl/sportsdata"
	"time"
)

type Team struct {
	Id           string            `xml:"id,attr"`
	ConferenceId string            `xml:"-"`
	Name         string            `xml:"name,attr"`
	Market       string            `xml:"market,attr"`
	Alias        string            `xml:"alias,attr"`
	Venue        *sportsdata.Venue `xml:"venue"`
}

type Conference struct {
	Id    string  `xml:"id,attr"`
	Name  string  `xml:"name,attr"`
	Alias string  `xml:"alias,attr"`
	Teams []*Team `xml:"team"`
}

type Division struct {
	Id          string        `xml:"id,attr"`
	Name        string        `xml:"name,attr"`
	Alias       string        `xml:"alias,attr"`
	Conferences []*Conference `xml:"conference"`
}

type HomeTeam struct {
	Id    string `xml:"id,attr"`
	Name  string `xml:"name,attr"`
	Alias string `xml:"alias,attr"`
}

func (t *HomeTeam) Team() *Team {
	return &Team{
		Id:    t.Id,
		Name:  t.Name,
		Alias: t.Alias,
	}
}

type AwayTeam struct {
	Id    string `xml:"id,attr"`
	Name  string `xml:"name,attr"`
	Alias string `xml:"alias,attr"`
}

func (t *AwayTeam) Team() *Team {
	return &Team{
		Id:    t.Id,
		Name:  t.Name,
		Alias: t.Alias,
	}
}

type Game struct {
	Id         string    `xml:"id,attr"`
	Status     string    `xml:"status,attr"`
	Coverage   string    `xml:"coverage,attr"`
	HomeTeamId string    `xml:"home_team,attr"`
	AwayTeamId string    `xml:"away_team,attr"`
	Scheduled  string    `xml:"scheduled,attr"`
	HomeTeam   *HomeTeam `xml:"home"`
	AwayTeam   *AwayTeam `xml:"away"`
}

func (g *Game) FormattedScheduled() (time.Time, error) {
	return time.Parse(sportsdata.SportsDataTimeFormat, g.Scheduled)
}

type Games struct {
	Games []*Game `xml:"game"`
}

type SeasonSchedule struct {
	Id         string `xml:"id,attr"`
	Year       string `xml:"year,attr"`
	SeasonType string `xml:"type,attr"`
	Games      Games  `xml:"games"`
}

type League struct {
	XMLNS          string          `xml:"xmlns,attr"`
	Id             string          `xml:"id,attr"`
	Name           string          `xml:"name,attr"`
	Alias          string          `xml:"alias,attr"`
	Divisions      []*Division     `xml:"division"`
	SeasonSchedule *SeasonSchedule `xml:"season-schedule"`
}

func (l *League) Teams() []*Team {
	teams := make([]*Team, 0)
	for _, division := range l.Divisions {
		for _, conference := range division.Conferences {
			for _, team := range conference.Teams {
				team.ConferenceId = conference.Id
				teams = append(teams, team)
			}
		}
	}
	return teams
}

type Schedule struct {
	Season       string
	ScheduleType ScheduleType
	League       *League
}

func (s *Schedule) Venues() []*sportsdata.Venue {
	venues := make([]*sportsdata.Venue, 0)
	for _, division := range s.League.Divisions {
		for _, conference := range division.Conferences {
			for _, team := range conference.Teams {
				if team.Venue != nil {
					venues = append(venues, team.Venue)
				}
			}
		}
	}
	return venues
}
