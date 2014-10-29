package ncaamb

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
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

func (a *API) baseEndpoint() string {
	var accessLevel AccessLevelType
	if a.production {
		accessLevel = AccessLevelProduction
	} else {
		accessLevel = AccessLevelTrial
	}
	return fmt.Sprintf("https://api.sportsdatallc.org/ncaamb-%s3", string(accessLevel))
}

func (a *API) divisionEndpoint() string {
	return fmt.Sprintf("%s/league/hierarchy.xml?api_key=%s", a.baseEndpoint(), a.apiKey)
}

func (a *API) scheduleEndpoint(season string, scheduleType ScheduleType) string {
	return fmt.Sprintf("%s/games/%s/%s/schedule.xml?api_key=%s", a.baseEndpoint(), season, string(scheduleType), a.apiKey)
}

func (a *API) Division() (*League, error) {
	resp, err := http.Get(a.divisionEndpoint())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	//fmt.Print(string(body))
	if err != nil {
		return nil, err
	}
	league := new(League)
	err = xml.Unmarshal(body, league)
	return league, err
}

func (a *API) Schedule(season string, scheduleType ScheduleType) (*League, error) {
	resp, err := http.Get(a.scheduleEndpoint(season, scheduleType))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	//fmt.Print(string(body))
	if err != nil {
		return nil, err
	}
	league := new(League)
	err = xml.Unmarshal(body, league)
	return league, err
}
