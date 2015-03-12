package sportsdata

import (
	"errors"
)

const SportsDataTimeFormat = "2006-01-02T15:04:05-07:00"

var ErrScoreNotFound = errors.New("Score not found")

type Venue struct {
	Id        string `xml:"id,attr"`
	Name      string `xml:"name,attr"`
	Address   string `xml:"address,attr"`
	City      string `xml:"city,attr"`
	State     string `xml:"state,attr"`
	Zip       string `xml:"zip,attr"`
	Country   string `xml:"country,attr"`
	Capacity  int64  `xml:"capacity,attr"`
	Surface   string `xml:"surface,attr"`
	VenueType string `xml:"type,attr"`
}
