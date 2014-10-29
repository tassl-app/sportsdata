package ncaawb

import (
	"encoding/xml"
	"testing"
)

const leagueDivisionData = `
<league xmlns="http://feed.elasticstats.com/schema/basketball/ncaam/hierarchy-v2.0.xsd" id="cd4268ee-07aa-4c4d-a435-ec44ad2c76cb" name="NCAA MEN" alias="NCAAM">
	<division id="18b713d6-561d-4aab-8986-175c4da04aa8" name="ACCA" alias="ACCA">
		<conference id="d69f2770-9310-45f7-9465-a6c1fe2567cc" name="Independents (ACCA)" alias="ACCA-IND">
			<team id="cd34248e-6f7d-4e0a-b694-567699bd7917" name="Tigers" market="Champion Baptist" alias="CBAP">
				<venue id="c06cdbce-91ba-4306-b31c-97df5cbf2515" name="Jon M. Huntsman Center" capacity="15000" address="1825 East South Campus Dr" city="Salt Lake City" state="UT" zip="84112" country="USA"/>
			</team>
		</conference>
	</division>
	<division id="556da150-35b7-433c-ae6d-dd702c543cf3" name="NAIA" alias="NAIA">
		<conference id="02cdb06e-b644-48e8-a5c9-dd235b888f7a" name="Kansas Collegiate Athletic Conference" alias="KCAC">
			<team id="ea22a80e-0194-4bda-99d1-a32f3545fffc" name="Bluejays" market="Tabor College" alias="TAB"></team>
		</conference>
		<conference id="3e37a5b4-29da-488f-8176-3b33121c036d" name="Midwest Collegiate Conference" alias="MCC">
			<team id="fac4a71e-40a9-47a6-b8d9-61a50d3c6edc" name="Hawks" market="Viterbo" alias="VIT"></team>
		</conference>
	</division>
</league>
`

const leagueScheduleData = `
<league xmlns="http://feed.elasticstats.com/schema/basketball/schedule-v2.0.xsd" id="36e93ef4-8270-429c-be2d-bcd108b09507" name="NCAA MEN" alias="NCAAM">
	<season-schedule id="562c84a7-b3eb-4b95-8435-6e3e1624e007" year="2012" type="REG">
		<games>
			<game id="04d68600-024d-4f46-84aa-257da2f59127" status="scheduled" coverage="full" home_team="f861db3e-c1fc-4e90-9f17-db0d5f0f3e8b" away_team="35422c09-b48a-4a85-b99e-a2b06badd15e" scheduled="2012-11-09T14:22:00+00:00">
				<home name="Green Wave" alias="TULN" id="f861db3e-c1fc-4e90-9f17-db0d5f0f3e8b"></home>
				<away name="Yellow Jackets" alias="GT" id="35422c09-b48a-4a85-b99e-a2b06badd15e"></away>
			</game>
			<game id="04f5b010-4d33-4374-b270-17cfeae6da64" status="scheduled" coverage="full" home_team="98076615-ab08-4e9f-88ef-bab6702fd66b" away_team="ed08d6a7-580a-4d94-b4cc-4718be73cd10" scheduled="2012-11-09T14:22:00+00:00">
				<home name="Zips" alias="AKR" id="98076615-ab08-4e9f-88ef-bab6702fd66b"></home>
				<away name="Chanticleers" alias="CCAR" id="ed08d6a7-580a-4d94-b4cc-4718be73cd10"></away>
			</game>
		</games>
	</season-schedule>
</league>
`

func TestLeagueDivision(t *testing.T) {
	v := new(League)
	err := xml.Unmarshal([]byte(leagueDivisionData), v)
	if err != nil {
		t.Errorf("Could not unmarshal xml. Error: %s\n", err.Error())
		return
	}
	expectedLeagueId := "cd4268ee-07aa-4c4d-a435-ec44ad2c76cb"
	if v.Id != expectedLeagueId {
		t.Errorf("Expected league id %s, found %s\n", expectedLeagueId, v.Id)
		return
	}
	divisions := v.Division
	if len(divisions) != 2 {
		t.Errorf("Expected %d divisions, found %d\n", 2, len(divisions))
		t.Errorf("XML rendered as %+v\n", v)
		return
	}
	conferences := divisions[1].Conference
	if len(conferences) != 2 {
		t.Errorf("Expected %d conferences, found %d\n", 2, len(conferences))
		t.Errorf("XML rendered as %+v\n", v)
		return
	}
	teams := conferences[0].Team
	if len(teams) != 1 {
		t.Errorf("Expected %d teams, found %d\n", 1, len(teams))
		t.Errorf("XML rendered as %+v\n", v)
		return
	}
	team := teams[0]
	expectedTeamId := "ea22a80e-0194-4bda-99d1-a32f3545fffc"
	if team.Id != expectedTeamId {
		t.Errorf("Expected team id to be %s, found %s\n", expectedTeamId, team.Id)
		t.Errorf("XML rendered as %+v\n", v)
		return
	}
}

func TestLeagueSchedule(t *testing.T) {
	v := new(League)
	err := xml.Unmarshal([]byte(leagueScheduleData), v)
	if err != nil {
		t.Errorf("Could not unmarshal xml. Error: %s\n", err.Error())
		return
	}
	expectedLeagueId := "36e93ef4-8270-429c-be2d-bcd108b09507"
	if v.Id != expectedLeagueId {
		t.Errorf("Expected league id %s, found %s\n", expectedLeagueId, v.Id)
		return
	}
	expectedSeasonScheduleId := "562c84a7-b3eb-4b95-8435-6e3e1624e007"
	if v.SeasonSchedule.Id != expectedSeasonScheduleId {
		t.Errorf("Expected schedule id %s, found %s\n", expectedSeasonScheduleId, v.SeasonSchedule.Id)
		return
	}
	games := v.SeasonSchedule.Games.Game
	if len(games) != 2 {
		t.Errorf("Expected %d games, found %d\n", 2, len(games))
		return
	}
	game := games[0]
	expectedGameId := "04d68600-024d-4f46-84aa-257da2f59127"
	if game.Id != expectedGameId {
		t.Errorf("Expected game id %s, found %s\n", expectedGameId, game.Id)
		return
	}
	homeTeam := game.HomeTeam
	expectedHomeTeamId := "f861db3e-c1fc-4e90-9f17-db0d5f0f3e8b"
	if homeTeam.Id != expectedHomeTeamId {
		t.Errorf("Expected home team id %s, found %s\n", expectedGameId, homeTeam.Id)
		return
	}
	awayTeam := game.AwayTeam
	expectedAwayTeamId := "35422c09-b48a-4a85-b99e-a2b06badd15e"
	if awayTeam.Id != expectedAwayTeamId {
		t.Errorf("Expected away team id %s, found %s\n", expectedGameId, awayTeam.Id)
		return
	}
}
