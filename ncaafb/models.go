package ncaafb

import (
	"github.com/tassl-app/sportsdata"
	"time"
)

type Team struct {
	Id            string `xml:"id,attr"`
	SubdivisionId string `xml:"-"`
	ConferenceId  string `xml:"-"`
	Name          string `xml:"name,attr"`
	Market        string `xml:"market,attr"`
	Coverage      string `xml:"coverage,attr"`
}

type Subdivision struct {
	Id    string  `xml:"id,attr"`
	Name  string  `xml:"name,attr"`
	Teams []*Team `xml:"team"`
}

type Conference struct {
	Id           string         `xml:"id,attr"`
	Name         string         `xml:"name,attr"`
	Subdivisions []*Subdivision `xml:"subdivision"`
	Teams        []*Team        `xml:"team"`
}

type Division struct {
	XMLNS       string        `xml:"xmlns,attr"`
	Id          string        `xml:"id,attr"`
	Name        string        `xml:"name,attr"`
	Conferences []*Conference `xml:"conference"`
}

func (d *Division) Teams() []*Team {
	teams := make([]*Team, 0)
	for _, conference := range d.Conferences {
		for _, subdivision := range conference.Subdivisions {
			for _, team := range subdivision.Teams {
				team.SubdivisionId = subdivision.Id
				teams = append(teams, team)
			}
		}
		for _, team := range conference.Teams {
			team.ConferenceId = conference.Id
			teams = append(teams, team)
		}
	}
	return teams
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

func (g *Game) FormattedScheduled() (time.Time, error) {
	return time.Parse(sportsdata.SportsDataTimeFormat, g.Scheduled)
}

func (g *Game) ParseScheduled(t time.Time) string {
	return t.Format(sportsdata.SportsDataTimeFormat)
}

type Week struct {
	Week  string  `xml:"week,attr"`
	Games []*Game `xml:"game"`
}

type Season struct {
	XMLNS      string  `xml:"xmlns,attr"`
	Season     string  `xml:"season,attr"`
	SeasonType string  `xml:"type,attr"`
	Weeks      []*Week `xml:"week"`
}

func (s *Season) Games() []*Game {
	games := make([]*Game, 0)
	for _, w := range s.Weeks {
		for _, g := range w.Games {
			games = append(games, g)
		}
	}
	return games
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

func (s *Schedule) Games() []*Game {
	games := make([]*Game, 0)
	for _, w := range s.Season.Weeks {
		for _, g := range w.Games {
			games = append(games, g)
		}
	}
	return games
}

func (s *Schedule) FilterGames(l []*Game) []*Game {
	filtered := make([]*Game, 0)
	for _, g := range l {
		for _, sg := range s.Games() {
			if sg.Id == g.Id {
				filtered = append(filtered, g)
				break
			}
		}
	}
	return filtered
}

type Boxscore struct {
	Year          string                 `xml:"-"`
	ScheduleType  ScheduleType           `xml:"-"`
	Week          string                 `xml:"-"`
	XMLNS         string                 `xml:"xmlns,attr"`
	Id            string                 `xml:"id,attr"`
	Scheduled     string                 `xml:"scheduled,attr"`
	HomeTeamId    string                 `xml:"home,attr"`
	AwayTeamId    string                 `xml:"away,attr"`
	Status        string                 `xml:"status,attr"`
	Quarter       string                 `xml:"quarter,attr"`
	Clock         string                 `xml:"clock,attr"`
	Completed     string                 `xml:"completed,attr"`
	Teams         []*BoxscoreTeam        `xml:"team"`
	ScoringDrives *BoxscoreScoringDrives `xml:"scoring_drives"`
}

func (b *Boxscore) FormattedScheduled() (time.Time, error) {
	return time.Parse(sportsdata.SportsDataTimeFormat, b.Scheduled)
}

func (b *Boxscore) FormattedCompleted() (time.Time, error) {
	return time.Parse(sportsdata.SportsDataTimeFormat, b.Completed)
}

func (b *Boxscore) HomeTeam() *BoxscoreTeam {
	for _, t := range b.Teams {
		if t.Id == b.HomeTeamId {
			return t
		}
	}
	return nil
}

func (b *Boxscore) AwayTeam() *BoxscoreTeam {
	for _, t := range b.Teams {
		if t.Id == b.AwayTeamId {
			return t
		}
	}
	return nil
}

func (b *Boxscore) HomeTeamScore() (int64, error) {
	homeTeam := b.HomeTeam()
	if homeTeam == nil {
		return 0, sportsdata.ErrScoreNotFound
	}
	return homeTeam.Points()
}

func (b *Boxscore) AwayTeamScore() (int64, error) {
	awawyTeam := b.AwayTeam()
	if awawyTeam == nil {
		return 0, sportsdata.ErrScoreNotFound
	}
	return awawyTeam.Points()
}

type BoxscoreTeam struct {
	Id                  string               `xml:"id,attr"`
	Name                string               `xml:"name,attr"`
	Market              string               `xml:"market,attr"`
	RemainingChallenges int64                `xml:"remaining_challenges,attr"`
	RemainingTimeouts   int64                `xml:"remaining_timeouts,attr"`
	Scoring             *BoxscoreTeamScoring `xml:"scoring"`
}

func (t *BoxscoreTeam) Points() (int64, error) {
	if t.Scoring == nil {
		return 0, sportsdata.ErrScoreNotFound
	}
	return t.Scoring.Points, nil
}

type BoxscoreTeamScoring struct {
	Points  int64                         `xml:"points,attr"`
	Quarter []*BoxscoreTeamScoringQuarter `xml:"quarter"`
}

type BoxscoreTeamScoringQuarter struct {
	Number int64 `xml:"number,attr"`
	Points int64 `xml:"points,attr"`
}

type BoxscoreScoringDrives struct {
	Drives []*BoxscoreScoringDrive `xml:"drive"`
}

type BoxscoreScoringDrive struct {
	Sequence string                       `xml:"sequence,attr"`
	Clock    string                       `xml:"clock,attr"`
	Quarter  string                       `xml:"quarter,attr"`
	Team     string                       `xml:"team,attr"`
	Scores   []*BoxscoreScoringDriveScore `xml:"score"`
}

type BoxscoreSummary struct {
	Data string `xml:",chardata"`
}

type BoxscoreScoringDriveScore struct {
	Id        string                         `xml:"id,attr"`
	Type      string                         `xml:"type,attr"`
	Clock     string                         `xml:"clock,attr"`
	Quarter   string                         `xml:"quarter,attr"`
	Points    int64                          `xml:"points,attr"`
	Team      string                         `xml:"team,attr"`
	GameScore *BoxscoreScoringDriveGameScore `xml:"game-score"`
	Summary   *BoxscoreSummary               `xml:"summary"`
	Links     *Links                         `xml:"links"`
}

type BoxscoreScoringDriveGameScore struct {
	Teams []*BoxscoreScoringDriveGameScoreTeam `xml:"team"`
}

type BoxscoreScoringDriveGameScoreTeam struct {
	Id     string `xml:"id,attr"`
	Points int64  `xml:"points,attr"`
}
