package htbgo

import (
	"testing"
)

func TestRankings(t *testing.T) {
	t.Run("Retrieves the user's best rank during a period of time in its country", func(t *testing.T) {
		period := Duration("1W")
		url := "https://www.hackthebox.com/api/v4/rankings/country/best?period=" + string(period)
		responseCode, isValidJSON, err := checkResponse(session, url)

		if responseCode != 200 {
			t.Error("response code should be 200, got", responseCode)
		}

		if !isValidJSON {
			t.Error("response is not valid JSON")
		}
		if err != nil {
			t.Error("An unexpected error occured")
		}

		_, err = session.RankInCountryBestDuring(period)

		if err != nil {
			t.Error("error creating struct: ", err)
		}
	})
	t.Run("Retrieves the user's rank during a period of time in its country", func(t *testing.T) {
		period := Duration("1W")
		url := "https://www.hackthebox.com/api/v4/rankings/country/overview?period=" + string(period)
		responseCode, isValidJSON, err := checkResponse(session, url)

		if responseCode != 200 {
			t.Error("response code should be 200, got", responseCode)
		}

		if !isValidJSON {
			t.Error("response is not valid JSON")
		}
		if err != nil {
			t.Error("An unexpected error occured")
		}

		_, err = session.RankInCountryDuring(period)

		if err != nil {
			t.Error("error creating struct: ", err)
		}
	})
	t.Run("Retrieves the user's rank bracket in its country", func(t *testing.T) {

		url := "https://www.hackthebox.com/api/v4/rankings/country/ranking_bracket"
		responseCode, isValidJSON, err := checkResponse(session, url)

		if responseCode != 200 {
			t.Error("response code should be 200, got", responseCode)
		}

		if !isValidJSON {
			t.Error("response is not valid JSON")
		}
		if err != nil {
			t.Error("An unexpected error occured")
		}

		_, err = session.RankBracketInCountry()

		if err != nil {
			t.Error("error creating struct: ", err)
		}
	})
	t.Run("Retrieves the user's current team's best rank during a period of time", func(t *testing.T) {
		period := Duration("1W")
		url := "https://www.hackthebox.com/api/v4/rankings/team/best?period=" + string(period)
		responseCode, isValidJSON, err := checkResponse(session, url)

		if responseCode != 200 {
			t.Error("response code should be 200, got", responseCode)
		}

		if !isValidJSON {
			t.Error("response is not valid JSON")
		}
		if err != nil {
			t.Error("An unexpected error occured")
		}

		_, err = session.RankOfTeamBestDuring(period)

		if err != nil {
			t.Error("error creating struct: ", err)
		}
	})
	t.Run("Retrieves the user's current team's rank during a period of time", func(t *testing.T) {

		period := Duration("1W")
		url := "https://www.hackthebox.com/api/v4/rankings/team/overview?period=" + string(period)
		responseCode, isValidJSON, err := checkResponse(session, url)

		if responseCode != 200 {
			t.Error("response code should be 200, got", responseCode)
		}

		if !isValidJSON {
			t.Error("response is not valid JSON")
		}
		if err != nil {
			t.Error("An unexpected error occured")
		}

		_, err = session.RankOfTeamDuring(period)

		if err != nil {
			t.Error("error creating struct: ", err)
		}
	})
	t.Run("Retrieves the user's current team's bracket in the leaderboards", func(t *testing.T) {
		url := "https://www.hackthebox.com/api/v4/rankings/team/ranking_bracket"
		responseCode, isValidJSON, err := checkResponse(session, url)

		if responseCode != 200 {
			t.Error("response code should be 200, got", responseCode)
		}

		if !isValidJSON {
			t.Error("response is not valid JSON")
		}
		if err != nil {
			t.Error("An unexpected error occured")
		}

		_, err = session.RankBracketOfTeam()

		if err != nil {
			t.Error("error creating struct: ", err)
		}
	})
	t.Run("Retrieves the user's best rank during a period of time", func(t *testing.T) {
		period := Duration("1W")
		wantVIP := true
		url := "https://www.hackthebox.com/api/v4/rankings/user/best?period=" + string(period) + "&vip=" + stringFromVIP(wantVIP)
		responseCode, isValidJSON, err := checkResponse(session, url)

		if responseCode != 200 {
			t.Error("response code should be 200, got", responseCode)
		}

		if !isValidJSON {
			t.Error("response is not valid JSON")
		}
		if err != nil {
			t.Error("An unexpected error occured")
		}

		_, err = session.RankBestDuring(period, wantVIP)

		if err != nil {
			t.Error("error creating struct: ", err)
		}
	})
	t.Run("Retrieve the country leaderboard", func(t *testing.T) {
		url := "https://www.hackthebox.com/api/v4/rankings/countries"
		responseCode, isValidJSON, err := checkResponse(session, url)

		if responseCode != 200 {
			t.Error("response code should be 200, got", responseCode)
		}

		if !isValidJSON {
			t.Error("response is not valid JSON")
		}
		if err != nil {
			t.Error("An unexpected error occured")
		}

		_, err = session.RanksOfCountries()

		if err != nil {
			t.Error("error creating struct: ", err)
		}
	})
	t.Run("Retrieves the team leaaderboard", func(t *testing.T) {
		url := "https://www.hackthebox.com/api/v4/rankings/teams"
		responseCode, isValidJSON, err := checkResponse(session, url)

		if responseCode != 200 {
			t.Error("response code should be 200, got", responseCode)
		}

		if !isValidJSON {
			t.Error("response is not valid JSON")
		}
		if err != nil {
			t.Error("An unexpected error occured")
		}

		_, err = session.RanksOfTeams()

		if err != nil {
			t.Error("error creating struct: ", err)
		}
	})
	t.Run("Retrieves the user leaderboard", func(t *testing.T) {
		wantVIP := true
		url := "https://www.hackthebox.com/api/v4/rankings/users?vip=" + stringFromVIP(wantVIP)
		responseCode, isValidJSON, err := checkResponse(session, url)

		if responseCode != 200 {
			t.Error("response code should be 200, got", responseCode)
		}

		if !isValidJSON {
			t.Error("response is not valid JSON")
		}
		if err != nil {
			t.Error("An unexpected error occured")
		}

		_, err = session.RanksOfUsers(wantVIP)

		if err != nil {
			t.Error("error creating struct: ", err)
		}
	})
	t.Run("Retrieves the university leaderboard", func(t *testing.T) {
		url := "https://www.hackthebox.com/api/v4/rankings/universities"
		responseCode, isValidJSON, err := checkResponse(session, url)

		if responseCode != 200 {
			t.Error("response code should be 200, got", responseCode)
		}

		if !isValidJSON {
			t.Error("response is not valid JSON")
		}
		if err != nil {
			t.Error("An unexpected error occured")
		}

		_, err = session.RanksOfUniversities()

		if err != nil {
			t.Error("error creating struct: ", err)
		}
	})

}
