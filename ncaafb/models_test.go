package ncaafb

import (
	"encoding/xml"
	"testing"
)

const divisionConferenceData = `
<division xmlns="http://feed.elasticstats.com/schema/ncaafb/hierarchy-v1.0.xsd" id="FBS" name="I-A">
	<conference id="ACC" name="ACC">
		<subdivision id="ACC-ATLANTIC" name="ATLANTIC">
			<team id="BC" name="Eagles" market="Boston College" coverage="full"/>
			<team id="CLE" name="Tigers" market="Clemson" coverage="full"/>
		</subdivision>
		<subdivision id="ACC-COASTAL" name="COASTAL">
			<team id="DUK" name="Blue Devils" market="Duke" coverage="full"/>
			<team id="GT" name="Yellow Jackets" market="Georgia Tech" coverage="full"/>
		</subdivision>
	</conference>
	<conference id="AAC" name="American Athletic">
		<team id="CIN" name="Bearcats" market="Cincinnati" coverage="full"/>
		<team id="UCONN" name="Huskies" market="Connecticut" coverage="full"/>
	</conference>
</division>
`

const seasonData = `
<season xmlns="http://feed.elasticstats.com/schema/ncaafb/schedule-v1.0.xsd" season="2014" type="REG">
	<week week="1">
		<game id="92044ce9-3698-443d-88a9-47967462dd61" scheduled="2014-08-23T19:30:00+00:00" coverage="full" home_rotation="" away_rotation="" home="EW" away="SHS" status="closed">
			<venue id="61b61700-a5e3-4f72-9cce-e0e6c3e652fa" country="USA" name="Roos Field" city="Cheney" state="WA" capacity="8600" surface="artificial" type="outdoor" zip="99004" address="1136 Washington St."/>
			<weather temperature="69" condition="Sunny" humidity="37">
				<wind speed="12" direction="NE"/>
			</weather>
			<broadcast network="ESPN" satellite="206" internet="WatchESPN" cable=""/>
			<links>
				<link rel="statistics" href="/2014/REG/1/SHS/EW/statistics.xml" type="application/xml"/>
				<link rel="summary" href="/2014/REG/1/SHS/EW/summary.xml" type="application/xml"/>
				<link rel="pbp" href="/2014/REG/1/SHS/EW/pbp.xml" type="application/xml"/>
				<link rel="boxscore" href="/2014/REG/1/SHS/EW/boxscore.xml" type="application/xml"/>
				<link rel="roster" href="/2014/REG/1/SHS/EW/roster.xml" type="application/xml"/>
			</links>
		</game>
		<game id="9d403f64-4b0c-4b3d-a882-beb507522004" scheduled="2014-08-27T23:00:00+00:00" coverage="full" home_rotation="" away_rotation="" home="GST" away="ACU" status="closed">
			<venue id="1167273b-94e9-4a51-ac65-2585f6da22b2" country="USA" name="Georgia Dome" city="Atlanta" state="GA" capacity="74228" surface="artificial" type="dome" zip="30313" address="1 Georgia Dome Drive Northwest"/>
			<weather temperature="88" condition="Sunny" humidity="36">
				<wind speed="6" direction="NE"/>
			</weather>
			<broadcast network="ESPNU" satellite="208" internet="WatchESPN" cable=""/>
			<links>
				<link rel="statistics" href="/2014/REG/1/ACU/GST/statistics.xml" type="application/xml"/>
				<link rel="summary" href="/2014/REG/1/ACU/GST/summary.xml" type="application/xml"/>
				<link rel="pbp" href="/2014/REG/1/ACU/GST/pbp.xml" type="application/xml"/>
				<link rel="boxscore" href="/2014/REG/1/ACU/GST/boxscore.xml" type="application/xml"/>
				<link rel="roster" href="/2014/REG/1/ACU/GST/roster.xml" type="application/xml"/>
			</links>
		</game>
	</week>
	<week week="2">
		<game id="b02a3bee-7afc-4c02-a352-a76baff1c26e" scheduled="2014-09-05T00:00:00+00:00" coverage="full" home_rotation="" away_rotation="" home="UTSA" away="ARI" status="closed">
			<venue id="0f0b1691-1f01-41f4-a34d-4f34ac0dc413" country="USA" name="Alamodome" city="San Antonio" state="TX" capacity="65000" surface="artificial" type="dome" zip="78203" address="100 Montana Street"/>
			<weather temperature="82" condition="Partly Cloudy " humidity="72">
				<wind speed="8" direction="E"/>
			</weather>
			<broadcast network="Fox Sports 1" satellite="219" internet="" cable=""/>
			<links>
				<link rel="statistics" href="/2014/REG/2/ARI/UTSA/statistics.xml" type="application/xml"/>
				<link rel="summary" href="/2014/REG/2/ARI/UTSA/summary.xml" type="application/xml"/>
				<link rel="pbp" href="/2014/REG/2/ARI/UTSA/pbp.xml" type="application/xml"/>
				<link rel="boxscore" href="/2014/REG/2/ARI/UTSA/boxscore.xml" type="application/xml"/>
				<link rel="roster" href="/2014/REG/2/ARI/UTSA/roster.xml" type="application/xml"/>
			</links>
		</game>
	</week>
</season>
`

func TestDivisionConferences(t *testing.T) {
	v := new(Division)
	err := xml.Unmarshal([]byte(divisionConferenceData), v)
	if err != nil {
		t.Errorf("Error: %s\n", err.Error())
		return
	}
	expectedDivisionId := "FBS"
	if v.Id != expectedDivisionId {
		t.Errorf("Expected division id %s, found %s\n", expectedDivisionId)
		return
	}
	conferences := v.Conferences
	if len(conferences) != 2 {
		t.Errorf("Expected %d conferences, found %d\n", 2, len(conferences))
		return
	}
	subdivisions := conferences[0].Subdivisions
	if len(subdivisions) != 2 {
		t.Errorf("Expected %d subdivisions, found %d\n", 2, len(subdivisions))
		return
	}
	teams := subdivisions[0].Teams
	if len(teams) != 2 {
		t.Errorf("Expected %d teams, found %d\n", 2, len(teams))
		return
	}
	expectedTeamId := "BC"
	team := teams[0]
	if team.Id != expectedTeamId {
		t.Errorf("Expected team id %s, found %s\n", expectedTeamId, team.Id)
		return
	}
}

func TestSeasons(t *testing.T) {
	v := new(Season)
	err := xml.Unmarshal([]byte(seasonData), v)
	if err != nil {
		t.Error(err.Error())
		return
	}
	expectedSeason := "2014"
	if v.Season != expectedSeason {
		t.Errorf("Expected season %s, found %s\n", expectedSeason, v.Season)
		return
	}
	weeks := v.Weeks
	if len(weeks) != 2 {
		t.Errorf("Expected %d weeks, found %d\n", 2, len(weeks))
		return
	}
	games := weeks[0].Games
	if len(games) != 2 {
		t.Errorf("Expected %d games, found %d\n", 2, len(games))
		return
	}
	game := games[0]
	venue := game.Venue
	if venue == nil {
		t.Errorf("Venue not found\n")
		return
	}
	expectedVenueId := "61b61700-a5e3-4f72-9cce-e0e6c3e652fa"
	if expectedVenueId != venue.Id {
		t.Errorf("Expected venue id %s, found %s\n", expectedVenueId, venue.Id)
		return
	}
	links := game.Links.Links
	if len(links) != 5 {
		t.Errorf("Expected %d links, found %d\n", 5, len(links))
		return
	}
	schedule := new(Schedule)
	schedule.Season = v
	venues := schedule.Venues()
	if len(venues) != 3 {
		t.Errorf("Expected %d venues, found %d\n", 3, len(venues))
		return
	}
}
