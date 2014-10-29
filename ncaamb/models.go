package ncaamb

type Venue struct {
	Id       string `xml:"id,attr"`
	Name     string `xml:"name,attr"`
	Capacity string `xml:"capacity,attr"`
	Address  string `xml:"address,attr"`
	City     string `xml:"city,attr"`
	State    string `xml:"state,attr"`
	Zip      string `xml:"zip,attr"`
	Country  string `xml:"country,attr"`
}

type Team struct {
	Id     string `xml:"id,attr"`
	Name   string `xml:"name,attr"`
	Market string `xml:"market,attr"`
	Alias  string `xml:"alias,attr"`
	Venue  *Venue `xml:"venue,attr"`
}

type Conference struct {
	Id    string `xml:"id,attr"`
	Name  string `xml:"name,attr"`
	Alias string `xml:"alias,attr"`
	Team  []Team `xml:"team"`
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
