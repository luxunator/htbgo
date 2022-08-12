package htbgo

import (
    "encoding/json"
    "io/ioutil"
    "net/http"
    "testing"
	"fmt"
    "log"
    "os"
)

func checkResponse(s *Session, url string) (responseCode int, isValidJSON bool, err error) {
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return
    }

    req.Header = s.Headers
    resp, err := s.Client.Do(req)
    if err != nil {
        return
    }

    responseCode = resp.StatusCode

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return
    }
    
    isValidJSON = json.Valid([]byte(body))

    return
}

var token string;
var session *Session;

func init() {
    token = os.Getenv("HTB_TOKEN")
    if token == ""{
        log.Fatal("empty HTB_TOKEN environment variable")
    }
    session, _ = New(token)
}

func TestRankInCountryBestDuring(t *testing.T) {
    durations := []Duration{Year, HalfYear, Week}
    
    for _, duration := range durations {
        t.Run(fmt.Sprintf("duration %s", duration), func(t *testing.T) {
			url := "https://www.hackthebox.com/api/v4/rankings/country/best?period=" + string(duration)

            responseCode, isValidJSON, err := checkResponse(session, url)

            if responseCode != 200 {
                t.Error("response code should be 200, got", responseCode)
            }
        
            if !isValidJSON {
                t.Error("response is not valid JSON")
            }
        
            _, err = session.RankInCountryBestDuring(duration)
        
            if err != nil {
                t.Error("error creating struct: ", err)
            }
        })
    }
}

func TestRankInCountryDuring(t *testing.T) {
    durations := []Duration{Year, HalfYear, Week}
    
    for _, duration := range durations {
        t.Run(fmt.Sprintf("duration %s", duration), func(t *testing.T) {
			url := "https://www.hackthebox.com/api/v4/rankings/country/overview?period=" + string(duration)

            responseCode, isValidJSON, err := checkResponse(session, url)

            if responseCode != 200 {
                t.Error("response code should be 200, got", responseCode)
            }
        
            if !isValidJSON {
                t.Error("response is not valid JSON")
            }
        
            _, err = session.RankInCountryDuring(duration)
        
            if err != nil {
                t.Error("error creating struct: ", err)
            }
        })
    }
}

func TestRankBracketInCountry(t *testing.T) {
    url := "https://www.hackthebox.com/api/v4/rankings/country/ranking_bracket"
    responseCode, isValidJSON, err := checkResponse(session, url)

    if responseCode != 200 {
        t.Error("response code should be 200, got", responseCode)
    }

    if !isValidJSON {
        t.Error("response is not valid JSON")
    }

    _, err = session.RankBracketInCountry()

    if err != nil {
        t.Error("error creating struct: ", err)
    }
}

func TestRankOfTeamBestDuring(t *testing.T) {
    durations := []Duration{Year, HalfYear, Week}
    
    for _, duration := range durations {
        t.Run(fmt.Sprintf("duration %s", duration), func(t *testing.T) {
			url := "https://www.hackthebox.com/api/v4/rankings/team/best?period=" + string(duration)

            responseCode, isValidJSON, err := checkResponse(session, url)

            if responseCode != 200 {
                t.Error("response code should be 200, got", responseCode)
            }
        
            if !isValidJSON {
                t.Error("response is not valid JSON")
            }
        
            _, err = session.RankOfTeamBestDuring(duration)
        
            if err != nil {
                t.Error("error creating struct: ", err)
            }
        })
    }
}

func TestRankOfTeamDuring(t *testing.T) {
    durations := []Duration{Year, HalfYear, Week}
    
    for _, duration := range durations {
        t.Run(fmt.Sprintf("duration %s", duration), func(t *testing.T) {
			url := "https://www.hackthebox.com/api/v4/rankings/team/overview?period=" + string(duration)

            responseCode, isValidJSON, err := checkResponse(session, url)

            if responseCode != 200 {
                t.Error("response code should be 200, got", responseCode)
            }
        
            if !isValidJSON {
                t.Error("response is not valid JSON")
            }
        
            _, err = session.RankOfTeamDuring(duration)
        
            if err != nil {
                t.Error("error creating struct: ", err)
            }
        })
    }
}

func TestRankBracketOfTeam(t *testing.T) {
    url := "https://www.hackthebox.com/api/v4/rankings/team/ranking_bracket"
    responseCode, isValidJSON, err := checkResponse(session, url)

    if responseCode != 200 {
        t.Error("response code should be 200, got", responseCode)
    }

    if !isValidJSON {
        t.Error("response is not valid JSON")
    }

    _, err = session.RankBracketOfTeam()

    if err != nil {
        t.Error("error creating struct: ", err)
    }
}

func TestRankBestDuring(t *testing.T) {
	testCases := []struct {
        duration  Duration
        wantVIP  bool
    }{
        {Year, false},
        {Week, false},
        {HalfYear, false},
    }
    
    for _, tc := range testCases {
        t.Run(fmt.Sprintf("duration %s wantVIP %t", tc.duration, tc.wantVIP), func(t *testing.T) {
			url := "https://www.hackthebox.com/api/v4/rankings/user/best?period=" + string(tc.duration) + "&vip=" + stringFromVIP(tc.wantVIP)

            responseCode, isValidJSON, err := checkResponse(session, url)

            if responseCode != 200 {
                t.Error("response code should be 200, got", responseCode)
            }
        
            if !isValidJSON {
                t.Error("response is not valid JSON")
            }
        
            _, err = session.RankBestDuring(tc.duration, tc.wantVIP)
        
            if err != nil {
                t.Error("error creating struct: ", err)
            }
        })
    }
}

func TestRankDuring(t *testing.T) {
	testCases := []struct {
        duration  Duration
        wantVIP  bool
    }{
        {Year, false},
        {Week, false},
        {HalfYear, false},
    }
    
    for _, tc := range testCases {
        t.Run(fmt.Sprintf("duration %s wantVIP %t", tc.duration, tc.wantVIP), func(t *testing.T) {
			url := "https://www.hackthebox.com/api/v4/rankings/user/overview?period=" + string(tc.duration) + "&vip=" + stringFromVIP(tc.wantVIP)

            responseCode, isValidJSON, err := checkResponse(session, url)

            if responseCode != 200 {
                t.Error("response code should be 200, got", responseCode)
            }
        
            if !isValidJSON {
                t.Error("response is not valid JSON")
            }
        
            _, err = session.RankDuring(tc.duration, tc.wantVIP)
        
            if err != nil {
                t.Error("error creating struct: ", err)
            }
        })
    }
}

func TestRankBracket(t *testing.T) {
	wantVIP := false
    url := "https://www.hackthebox.com/api/v4/rankings/user/ranking_bracket?vip=" + stringFromVIP(wantVIP)
    responseCode, isValidJSON, err := checkResponse(session, url)

    if responseCode != 200 {
        t.Error("response code should be 200, got", responseCode)
    }

    if !isValidJSON {
        t.Error("response is not valid JSON")
    }

    _, err = session.RankBracket(wantVIP)

    if err != nil {
        t.Error("error creating struct: ", err)
    }
}

func TestRanksOfCountries(t *testing.T) {
    url := "https://www.hackthebox.com/api/v4/rankings/countries"
    responseCode, isValidJSON, err := checkResponse(session, url)

    if responseCode != 200 {
        t.Error("response code should be 200, got", responseCode)
    }

    if !isValidJSON {
        t.Error("response is not valid JSON")
    }

    _, err = session.RanksOfCountries()

    if err != nil {
        t.Error("error creating struct: ", err)
    }
}

func TestRanksOfTeams(t *testing.T) {
    url := "https://www.hackthebox.com/api/v4/rankings/teams"
    responseCode, isValidJSON, err := checkResponse(session, url)

    if responseCode != 200 {
        t.Error("response code should be 200, got", responseCode)
    }

    if !isValidJSON {
        t.Error("response is not valid JSON")
    }

    _, err = session.RanksOfTeams()

    if err != nil {
        t.Error("error creating struct: ", err)
    }
}

func TestRanksOfUsers(t *testing.T) {
	wantVIP := false
    url := "https://www.hackthebox.com/api/v4/rankings/users?vip=" + stringFromVIP(wantVIP)
    responseCode, isValidJSON, err := checkResponse(session, url)

    if responseCode != 200 {
        t.Error("response code should be 200, got", responseCode)
    }

    if !isValidJSON {
        t.Error("response is not valid JSON")
    }

    _, err = session.RanksOfUsers(wantVIP)

    if err != nil {
        t.Error("error creating struct: ", err)
    }
}

func TestRanksOfUniversities(t *testing.T) {
    url := "https://www.hackthebox.com/api/v4/rankings/universities"
    responseCode, isValidJSON, err := checkResponse(session, url)

    if responseCode != 200 {
        t.Error("response code should be 200, got", responseCode)
    }

    if !isValidJSON {
        t.Error("response is not valid JSON")
    }

    _, err = session.RanksOfUniversities()

    if err != nil {
        t.Error("error creating struct: ", err)
    }
}