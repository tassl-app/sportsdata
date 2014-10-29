package ncaafb

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
	ScheduleRegular    = ScheduleType("reg")
	SchedulePostSeason = ScheduleType("pst")
)

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
	return fmt.Sprintf("https://api.sportsdatallc.org/ncaafb-%s3", string(accessLevel))
}

func (a *API) divisionEndpoint(divisionType DivisionType) string {
	return fmt.Sprintf("%s/teams/%s/hierarchy.xml?api_key=%s", a.baseEndpoint(), string(divisionType), a.apiKey)
}

func (a *API) scheduleEndpoint(year string, scheduleType ScheduleType) string {
	return fmt.Sprintf("%s/%s/%s/schedule.xml?api_key=%s", a.baseEndpoint(), year, string(scheduleType), a.apiKey)
}

func (a *API) Division(divisionType DivisionType) (*Division, error) {
	resp, err := http.Get(a.divisionEndpoint(divisionType))
	if err != nil {
		return nil, err
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

func (a *API) Schedule(year string, scheduleType ScheduleType) (*Schedule, error) {
	resp, err := http.Get(a.scheduleEndpoint(year, scheduleType))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	//fmt.Print(string(body))
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
