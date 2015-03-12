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

func (g *Game) ParseScheduled(t time.Time) string {
	return t.Format(sportsdata.SportsDataTimeFormat)
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

func (s *Schedule) Games() []*Game {
	return s.League.SeasonSchedule.Games.Games
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

func (s *Schedule) FilterBoxscores(l []*Boxscore) []*Boxscore {
	filtered := make([]*Boxscore, 0)
	for _, b := range l {
		for _, g := range s.Games() {
			if g.Id == b.Id {
				filtered = append(filtered, b)
				break
			}
		}
	}
	return filtered
}

type Boxscore struct {
	XMLNS       string          `xml:"xmlns,attr"`
	Id          string          `xml:"id,attr"`
	Status      string          `xml:"status,attr"`
	Coverage    string          `xml:"coverage,attr"`
	HomeTeamId  string          `xml:"home_team,attr"`
	AwayTeamId  string          `xml:"away_team,attr"`
	Scheduled   string          `xml:"scheduled,attr"`
	Attendance  int64           `xml:"attendance,attr"`
	LeadChanges int64           `xml:"lead_chages,attr"`
	TimesTied   int64           `xml:"times_tied,attr"`
	Half        int64           `xml:"half"`
	Teams       []*BoxscoreTeam `xml:"team"`
}

func (b *Boxscore) FormattedScheduled() (time.Time, error) {
	return time.Parse(sportsdata.SportsDataTimeFormat, b.Scheduled)
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
	return homeTeam.Points, nil
}

func (b *Boxscore) AwayTeamScore() (int64, error) {
	awawyTeam := b.AwayTeam()
	if awawyTeam == nil {
		return 0, sportsdata.ErrScoreNotFound
	}
	return awawyTeam.Points, nil
}

type BoxscoreTeam struct {
	Name            string           `xml:"name,attr"`
	Market          string           `xml:"market,attr"`
	Id              string           `xml:"id,attr"`
	Points          int64            `xml:"points,attr"`
	Rank            int64            `xml:"rank,attr"`
	BoxscoreScoring *BoxscoreScoring `xml:"scoring"`
	Leaders         *BoxscoreLeader  `xml:"leaders"`
}

type BoxscoreScoring struct {
	Halves []*BoxcoreScoringHalf `xml:"half"`
}

type BoxcoreScoringHalf struct {
	Number   int64 `xml:"number,attr"`
	Sequence int64 `xml:"sequence,attr"`
	Points   int64 `xml:"points,attr"`
}

type BoxscoreLeader struct {
	Points *BoxscoreLeaderPoint `xml:"points"`
	// TODO
	// BoxscoreLeaderRebound
	// BoxscoreLeaderAssist
}

type BoxscoreLeaderPoint struct {
	Player *BoxscoreLeaderPointPlayer `xml:"player"`
}

type BoxscoreLeaderPointPlayer struct {
	FullName    string                               `xml:"full_name"`
	Position    string                               `xml:"position"`
	JersyNumber string                               `xml:"jersey_number"`
	Id          string                               `xml:"id"`
	Statistics  *BoxscoreLeaderPointPlayerStatistics `xml:"statistics"`
}

type BoxscoreLeaderPointPlayerStatistics struct {
	Minutes              string `xml:"minutes"`
	FieldGoalsMade       string `xml:"field_goals_made"`
	FieldGoalsAtt        string `xml:"field_goals_att"`
	ThreePointsMade      string `xml:"three_points_made"`
	ThreePointsAttempted string `xml:"three_points_att"`
	ThreePointsPercent   string `xml:"three_points_pct"`
	TwoPointsMade        string `xml:"two_points_made"`
	TwoPointsAttempted   string `xml:"two_points_attempted"`
	TwoPointsPercent     string `xml:"two_points_pct"`
	FreeThrowsMade       string `xml:"free_throws_made"`
	FreeThrowsAttempted  string `xml:"free_throws_att"`
	FreeThrowsPercent    string `xml:"free_throws_pct"`
	OffensiveRebounds    string `xml:"offensive_rebounds"`
	DefensiveRebounds    string `xml:"defensive_rebounds"`
	Rebounds             string `xml:"rebounds"`
	Assists              string `xml:"assists"`
	Turnovers            string `xml:"turnovers"`
	Steals               string `xml:"steals"`
	Blocks               string `xml:"blocks"`
	AssistsTurnoverRatio string `xml:"assists_turnover_ratio"`
	PersonalFouls        string `xml:"personal_fouls"`
	TechFouls            string `xml:"tech_fouls"`
	Points               string `xml:"points"`
}

// TODO
type BoxscoreLeaderRebound struct{}

// TODO
type BoxscoreLeaderAssist struct{}
