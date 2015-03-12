package ncaafb

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type API struct {
	apiKey     string
	production bool
}

func NewAPI(apiKey string, production bool) *API {
	return &API{
		apiKey:     apiKey,
		production: production,
	}
}

type AccessLevelType string

const (
	AccessLevelTrial      = AccessLevelType("t")
	AccessLevelProduction = AccessLevelType("p")
)

type ScheduleType string

const (
	ScheduleRegular    = ScheduleType("reg")
	SchedulePostSeason = ScheduleType("pst")
)

var ScheduleAll = []ScheduleType{
	ScheduleRegular,
	SchedulePostSeason,
}

type DivisionType string

const (
	DivisionFBS   = DivisionType("FBS")
	DivisionFCS   = DivisionType("FCS")
	DivisionD2    = DivisionType("D2")
	DivisionD3    = DivisionType("D3")
	DivisionNAIA  = DivisionType("NAIA")
	DivisionUSCAA = DivisionType("USCAA")
)

var DivisionAll = []DivisionType{
	DivisionFBS,
	DivisionFCS,
	DivisionD2,
	DivisionD3,
	DivisionNAIA,
	DivisionUSCAA,
}

func (a *API) baseEndpoint() string {
	var accessLevel AccessLevelType
	if a.production {
		accessLevel = AccessLevelProduction
	} else {
		accessLevel = AccessLevelTrial
	}
	return fmt.Sprintf("https://api.sportsdatallc.org/ncaafb-%s1", string(accessLevel))
}

func (a *API) divisionEndpoint(divisionType DivisionType) (*url.URL, error) {
	endpoint := fmt.Sprintf("%s/teams/%s/hierarchy.xml", a.baseEndpoint(), string(divisionType))
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	q := u.Query()
	q.Set("api_key", a.apiKey)
	u.RawQuery = q.Encode()
	return u, nil
}

func (a *API) scheduleEndpoint(year string, scheduleType ScheduleType) (*url.URL, error) {
	endpoint := fmt.Sprintf("%s/%s/%s/schedule.xml", a.baseEndpoint(), year, string(scheduleType))
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	q := u.Query()
	q.Set("api_key", a.apiKey)
	u.RawQuery = q.Encode()
	return u, nil
}

func (a *API) boxscoreEndpoint(year string, scheduleType ScheduleType, week, awayTeamId, homeTeamId string) (*url.URL, error) {
	//http(s)://api.sportsdatallc.org/ncaafb-[access_level][version]/[year]/[ncaafb_season]/[ncaafb_season_week]/[away_team]/[home_team]/boxscore.[format]?api_key=[your_api_key]
	endpoint := fmt.Sprintf("%s/%s/%s/%s/%s/%s/boxscore.xml", a.baseEndpoint(), year, scheduleType, week, awayTeamId, homeTeamId)
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	q := u.Query()
	q.Set("api_key", a.apiKey)
	u.RawQuery = q.Encode()
	return u, nil
}

func (a *API) Division(divisionType DivisionType) (*Division, error) {
	u, err := a.divisionEndpoint(divisionType)
	if err != nil {
		return nil, err
	}
	endpoint := u.String()
	resp, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("API Status Returned Code %d.\nRequest: %+v\nResponse: %+v\n", resp.StatusCode, resp.Request, resp))
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	division := new(Division)
	err = xml.Unmarshal(body, division)
	return division, err
}

func (a *API) AllDivisions() ([]*Division, error) {
	divisions := make([]*Division, 0)
	for i, divisionType := range DivisionAll {
		if i > 0 {
			time.Sleep(1 * time.Second)
		}
		division, err := a.Division(divisionType)
		if err != nil {
			return nil, err
		}
		divisions = append(divisions, division)
	}
	return divisions, nil
}

func (a *API) Schedule(year string, scheduleType ScheduleType) (*Schedule, error) {
	u, err := a.scheduleEndpoint(year, scheduleType)
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("API Status Returned Code %d.\nRequest: %+v\nResponse: %+v\n", resp.StatusCode, resp.Request, resp))
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	season := new(Season)
	err = xml.Unmarshal(body, season)
	if err != nil {
		return nil, err
	}
	schedule := new(Schedule)
	schedule.Year = year
	schedule.ScheduleType = scheduleType
	schedule.Season = season
	return schedule, nil
}

func (a *API) AllSchedules(years []string) ([]*Schedule, error) {
	schedules := make([]*Schedule, 0)
	for i, year := range years {
		for j, scheduleType := range ScheduleAll {
			if i > 0 || j > 0 {
				time.Sleep(1 * time.Second)
			}
			schedule, err := a.Schedule(year, scheduleType)
			if err != nil {
				return nil, err
			}
			schedules = append(schedules, schedule)
		}
	}
	return schedules, nil
}

func (a *API) Boxscore(year string, scheduleType ScheduleType, week, awayTeamId, homeTeamId string) (*Boxscore, error) {
	u, err := a.boxscoreEndpoint(year, scheduleType, week, awayTeamId, homeTeamId)
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("API Status Returned Code %d.\nRequest: %+v\nResponse: %+v\n", resp.StatusCode, resp.Request, resp))
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	boxscore := new(Boxscore)
	err = xml.Unmarshal(body, boxscore)
	if err != nil {
		return nil, err
	}
	boxscore.Year = year
	boxscore.ScheduleType = scheduleType
	boxscore.Week = week
	return boxscore, nil
}

func (a *API) ScheduleBoxscores(schedule *Schedule, games []*Game) ([]*Boxscore, error) {
	boxscores := make([]*Boxscore, 0)
	sleep := false
	for _, g := range games {
		for _, w := range schedule.Season.Weeks {
			if sleep {
				time.Sleep(1 * time.Second)
			}
			fmt.Printf("Getting boxscore for %s, %s, %s, %s, %s", schedule.Year, schedule.ScheduleType, w.Week, g.AwayTeamId, g.HomeTeamId)
			boxscore, err := a.Boxscore(schedule.Year, schedule.ScheduleType, w.Week, g.AwayTeamId, g.HomeTeamId)
			if err != nil {
				return nil, err
			}
			boxscores = append(boxscores, boxscore)
			sleep = true
		}
	}
	return boxscores, nil
}
