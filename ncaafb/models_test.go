package ncaafb

import (
	"encoding/xml"
	"testing"
	"time"
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

const boxscoreData = `
<game xmlns="http://feed.elasticstats.com/schema/ncaafb/boxscore-v1.0.xsd" id="e5896e5f-3779-4726-bee9-512d9d0746b2" scheduled="2014-09-18T23:30:00+00:00" home="KST" away="AUB" status="closed" quarter="4" clock=":00" completed="2014-09-19T02:50:26+00:00">
  <team id="KST" name="Wildcats" market="Kansas State" remaining_challenges="2" remaining_timeouts="2">
    <scoring points="14">
      <quarter number="1" points="0"/>
      <quarter number="2" points="7"/>
      <quarter number="3" points="0"/>
      <quarter number="4" points="7"/>
    </scoring>
  </team>
  <team id="AUB" name="Tigers" market="Auburn" remaining_challenges="2" remaining_timeouts="2">
    <scoring points="20">
      <quarter number="1" points="3"/>
      <quarter number="2" points="7"/>
      <quarter number="3" points="0"/>
      <quarter number="4" points="10"/>
    </scoring>
  </team>
  <scoring_drives>
    <drive sequence="1" clock="13:07" quarter="1" team="AUB">
      <score id="a14cf3cc-2985-4f2a-a1f2-05fdf06635e5" type="fieldgoal" clock="11:19" quarter="1" points="3" team="AUB">
        <game-score>
          <team id="KST" points="0"/>
          <team id="AUB" points="3"/>
        </game-score>
        <summary>
          <![CDATA[38-D.Carlson 34 yards Field Goal is Good.]]>
        </summary>
        <links>
          <link rel="summary" href="/2014/REG/4/AUB/KST/plays/a14cf3cc-2985-4f2a-a1f2-05fdf06635e5.xml" type="application/xml"/>
        </links>
      </score>
    </drive>
    <drive sequence="2" clock="07:48" quarter="2" team="KST">
      <score id="afc02847-c2d7-457f-8e41-960da56681da" type="touchdown" clock="05:00" quarter="2" points="6" team="KST">
        <game-score>
          <team id="KST" points="6"/>
          <team id="AUB" points="3"/>
        </game-score>
        <summary>
          <![CDATA[20-D.Robinson runs 3 yards for a touchdown.]]>
        </summary>
        <links>
          <link rel="summary" href="/2014/REG/4/AUB/KST/plays/afc02847-c2d7-457f-8e41-960da56681da.xml" type="application/xml"/>
        </links>
      </score>
      <score id="5074c28c-8d77-4126-a877-92ea8b67ed4a" type="extrapoint" clock="04:56" quarter="2" points="1" team="KST">
        <game-score>
          <team id="KST" points="7"/>
          <team id="AUB" points="3"/>
        </game-score>
        <summary>
          <![CDATA[3-J.Cantele extra point is good.]]>
        </summary>
        <links>
          <link rel="summary" href="/2014/REG/4/AUB/KST/plays/5074c28c-8d77-4126-a877-92ea8b67ed4a.xml" type="application/xml"/>
        </links>
      </score>
    </drive>
    <drive sequence="3" clock="04:56" quarter="2" team="AUB">
      <score id="37baf44b-3593-45b0-8e75-3d0bd93e471c" type="touchdown" clock="01:45" quarter="2" points="6" team="AUB">
        <game-score>
          <team id="KST" points="7"/>
          <team id="AUB" points="9"/>
        </game-score>
        <summary>
          <![CDATA[14-N.Marshall complete to 5-R.Louis. 5-R.Louis runs 40 yards for a touchdown.]]>
        </summary>
        <links>
          <link rel="summary" href="/2014/REG/4/AUB/KST/plays/37baf44b-3593-45b0-8e75-3d0bd93e471c.xml" type="application/xml"/>
        </links>
      </score>
      <score id="bc359877-6cb5-4117-9c69-c99a644c3ded" type="extrapoint" clock="01:34" quarter="2" points="1" team="AUB">
        <game-score>
          <team id="KST" points="7"/>
          <team id="AUB" points="10"/>
        </game-score>
        <summary>
          <![CDATA[38-D.Carlson extra point is good.]]>
        </summary>
        <links>
          <link rel="summary" href="/2014/REG/4/AUB/KST/plays/bc359877-6cb5-4117-9c69-c99a644c3ded.xml" type="application/xml"/>
        </links>
      </score>
    </drive>
    <drive sequence="4" clock="04:44" quarter="3" team="AUB">
      <score id="71eba433-fce2-4e04-b559-f1d5ec8d8b54" type="touchdown" clock="14:16" quarter="4" points="6" team="AUB">
        <game-score>
          <team id="KST" points="7"/>
          <team id="AUB" points="16"/>
        </game-score>
        <summary>
          <![CDATA[14-N.Marshall complete to 1-D.Williams. 1-D.Williams runs 9 yards for a touchdown.]]>
        </summary>
        <links>
          <link rel="summary" href="/2014/REG/4/AUB/KST/plays/71eba433-fce2-4e04-b559-f1d5ec8d8b54.xml" type="application/xml"/>
        </links>
      </score>
      <score id="1992ebbd-299b-45fa-b3a9-ef0bfd67d3b8" type="extrapoint" clock="14:10" quarter="4" points="1" team="AUB">
        <game-score>
          <team id="KST" points="7"/>
          <team id="AUB" points="17"/>
        </game-score>
        <summary>
          <![CDATA[38-D.Carlson extra point is good.]]>
        </summary>
        <links>
          <link rel="summary" href="/2014/REG/4/AUB/KST/plays/1992ebbd-299b-45fa-b3a9-ef0bfd67d3b8.xml" type="application/xml"/>
        </links>
      </score>
    </drive>
    <drive sequence="5" clock="12:16" quarter="4" team="AUB">
      <score id="9fffb648-b574-4e98-96de-487d6eaa735f" type="fieldgoal" clock="06:30" quarter="4" points="3" team="AUB">
        <game-score>
          <team id="KST" points="7"/>
          <team id="AUB" points="20"/>
        </game-score>
        <summary>
          <![CDATA[38-D.Carlson 25 yards Field Goal is Good.]]>
        </summary>
        <links>
          <link rel="summary" href="/2014/REG/4/AUB/KST/plays/9fffb648-b574-4e98-96de-487d6eaa735f.xml" type="application/xml"/>
        </links>
      </score>
    </drive>
    <drive sequence="6" clock="06:28" quarter="4" team="KST">
      <score id="b6c3c540-19e4-4498-8a5e-1b2563f509c8" type="touchdown" clock="03:53" quarter="4" points="6" team="KST">
        <game-score>
          <team id="KST" points="13"/>
          <team id="AUB" points="20"/>
        </game-score>
        <summary>
          <![CDATA[24-C.Jones runs 1 yard for a touchdown.]]>
        </summary>
        <links>
          <link rel="summary" href="/2014/REG/4/AUB/KST/plays/b6c3c540-19e4-4498-8a5e-1b2563f509c8.xml" type="application/xml"/>
        </links>
      </score>
      <score id="9e9b1900-81a1-4154-b6b7-cf31baf1b7bf" type="extrapoint" clock="03:49" quarter="4" points="1" team="KST">
        <game-score>
          <team id="KST" points="14"/>
          <team id="AUB" points="20"/>
        </game-score>
        <summary>
          <![CDATA[16-M.McCrane extra point is good.]]>
        </summary>
        <links>
          <link rel="summary" href="/2014/REG/4/AUB/KST/plays/9e9b1900-81a1-4154-b6b7-cf31baf1b7bf.xml" type="application/xml"/>
        </links>
      </score>
    </drive>
  </scoring_drives>
</game>
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
	allTeams := v.Teams()
	if len(allTeams) != 6 {
		t.Errorf("Expected %d teams, found %d\n", 6, len(allTeams))
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
	gameTime, err := game.FormattedScheduled()
	if err != nil {
		t.Error(err.Error())
		return
	}
	expectedTime := time.Date(2014, 8, 23, 19, 30, 0, 0, time.UTC)
	if !expectedTime.Equal(gameTime) {
		t.Errorf("Expected time %v, found %v\n", expectedTime, gameTime)
		return
	}
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

func TestBoxscore(t *testing.T) {
	v := new(Boxscore)
	err := xml.Unmarshal([]byte(boxscoreData), v)
	if err != nil {
		t.Error(err.Error())
		return
	}
	gameTime, err := v.FormattedScheduled()
	if err != nil {
		t.Error(err.Error())
		return
	}
	expectedGameTime := time.Date(2014, 9, 18, 23, 30, 0, 0, time.UTC)
	if !expectedGameTime.Equal(gameTime) {
		t.Errorf("Expected game time %v, found %v\n", expectedGameTime, gameTime)
		return
	}
	teams := v.Teams
	if len(teams) != 2 {
		t.Errorf("Expected %d teams, found %d\n", 2, len(teams))
		return
	}
	expectedHomeTeamScore := int64(14)
	homeTeamScore, err := v.HomeTeamScore()
	if err != nil {
		t.Error(err.Error())
		return
	}
	if homeTeamScore != expectedHomeTeamScore {
		t.Errorf("Expected home score of %d, found %d\n", expectedHomeTeamScore, homeTeamScore)
		return
	}
	expectedAwayTeamScore := int64(20)
	awayTeamScore, err := v.AwayTeamScore()
	if err != nil {
		t.Error(err.Error())
		return
	}
	if awayTeamScore != expectedAwayTeamScore {
		t.Errorf("Expected away score of %d, found %d\n", expectedAwayTeamScore, awayTeamScore)
		return
	}
}
