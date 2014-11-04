package ncaamb

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
	ScheduleRegular              = ScheduleType("reg")
	ScheduleConferenceTournament = ScheduleType("ct")
	SchedulePostSeason           = ScheduleType("pst")
)

var ScheduleAll = []ScheduleType{
	ScheduleRegular,
	ScheduleConferenceTournament,
	SchedulePostSeason,
}

func (a *API) baseEndpoint() string {
	var accessLevel AccessLevelType
	if a.production {
		accessLevel = AccessLevelProduction
	} else {
		accessLevel = AccessLevelTrial
	}
	return fmt.Sprintf("https://api.sportsdatallc.org/ncaamb-%s3", string(accessLevel))
}

func (a *API) divisionEndpoint() (*url.URL, error) {
	endpoint := fmt.Sprintf("%s/league/hierarchy.xml", a.baseEndpoint())
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	q := u.Query()
	q.Set("api_key", a.apiKey)
	u.RawQuery = q.Encode()
	return u, nil
}

func (a *API) scheduleEndpoint(season string, scheduleType ScheduleType) (*url.URL, error) {
	endpoint := fmt.Sprintf("%s/games/%s/%s/schedule.xml?", a.baseEndpoint(), season, string(scheduleType))
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	q := u.Query()
	q.Set("api_key", a.apiKey)
	u.RawQuery = q.Encode()
	return u, nil
}

func (a *API) League() (*League, error) {
	endpoint, err := a.divisionEndpoint()
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(endpoint.String())
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
	league := new(League)
	err = xml.Unmarshal(body, league)
	return league, err
}

func (a *API) Schedule(season string, scheduleType ScheduleType) (*Schedule, error) {
	endpoint, err := a.scheduleEndpoint(season, scheduleType)
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(endpoint.String())
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
	league := new(League)
	err = xml.Unmarshal(body, league)
	if err != nil {
		return nil, err
	}
	schedule := new(Schedule)
	schedule.Season = season
	schedule.ScheduleType = scheduleType
	schedule.League = league
	return schedule, nil
}

func (a *API) AllSchedules(seasons []string) ([]*Schedule, error) {
	schedules := make([]*Schedule, 0)
	for i, season := range seasons {
		for j, scheduleType := range ScheduleAll {
			if i > 0 || j > 0 {
				time.Sleep(1 * time.Second)
			}
			schedule, err := a.Schedule(season, scheduleType)
			if err != nil {
				return nil, err
			}
			schedules = append(schedules, schedule)
		}
	}
	return schedules, nil
}
