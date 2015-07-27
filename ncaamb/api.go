package ncaamb

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

type API struct {
	apiKey     string
	production bool
	log        bool
}

func NewAPI(apiKey string, production, log bool) *API {
	return &API{
		apiKey:     apiKey,
		production: production,
		log:        log,
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
	endpoint := fmt.Sprintf("https://api.sportsdatallc.org/ncaamb-%s3", string(accessLevel))
	if a.log {
		log.Printf("base endpoint: %+v\n", endpoint)
	}
	return endpoint
}

func (a *API) boxscoreEndpoint(gameId string) (*url.URL, error) {
	endpoint := fmt.Sprintf("%s/games/%s/boxscore.xml", a.baseEndpoint(), gameId)
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	q := u.Query()
	q.Set("api_key", a.apiKey)
	u.RawQuery = q.Encode()
	if a.log {
		log.Printf("boxscore endpoint: %+v\n", u.String())
	}
	return u, nil
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
	if a.log {
		log.Printf("division endpoint: %+v\n", u.String())
	}
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
	if a.log {
		log.Printf("schedule endpoint: %+v\n", u.String())
	}
	return u, nil
}

func (a *API) League() (*League, error) {
	endpoint, err := a.divisionEndpoint()
	if err != nil {
		return nil, err
	}
	time.Sleep(1 * time.Second)
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
	time.Sleep(1 * time.Second)
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
	for _, season := range seasons {
		for _, scheduleType := range ScheduleAll {
			schedule, err := a.Schedule(season, scheduleType)
			if err != nil {
				return nil, err
			}
			schedules = append(schedules, schedule)
		}
	}
	return schedules, nil
}

func (a *API) Boxscore(gameId string) (*Boxscore, error) {
	endpoint, err := a.boxscoreEndpoint(gameId)
	if err != nil {
		return nil, err
	}
	time.Sleep(1 * time.Second)
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
	boxscore := new(Boxscore)
	err = xml.Unmarshal(body, boxscore)
	return boxscore, err
}

func (a *API) Boxscores(ids []string) ([]*Boxscore, error) {
	boxscores := make([]*Boxscore, 0)
	for _, id := range ids {
		fmt.Printf("Getting boxscore for %s\n", id)
		boxscore, err := a.Boxscore(id)
		if err != nil {
			return nil, err
		}
		boxscores = append(boxscores, boxscore)
	}
	return boxscores, nil
}
