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

func TestTeam(t *testing.T) {
    teamIDs := []int{3804, 3421, 2978}
    
    for _, teamID := range teamIDs {
        t.Run(fmt.Sprintf("id %d", teamID), func(t *testing.T) {
            teamIDString, _ := toPositiveIntString(teamID)
			url := "https://www.hackthebox.com/api/v4/team/info/" + teamIDString

            responseCode, isValidJSON, err := checkResponse(session, url)

            if responseCode != 200 {
                t.Error("response code should be 200, got", responseCode)
            }
        
            if !isValidJSON {
                t.Error("response is not valid JSON")
            }
        
            _, err = session.Team(teamID)
        
            if err != nil {
                t.Error("error creating struct: ", err)
            }
        })
    }
}

func TestActivity(t *testing.T) {
    teamIDs := []int{3804, 3421, 2978}
    
    for _, teamID := range teamIDs {
        t.Run(fmt.Sprintf("id %d", teamID), func(t *testing.T) {
            teamIDString, _ := toPositiveIntString(teamID)
			url := "https://www.hackthebox.com/api/v4/team/activity/" + teamIDString

            responseCode, isValidJSON, err := checkResponse(session, url)

            if responseCode != 200 {
                t.Error("response code should be 200, got", responseCode)
            }
        
            if !isValidJSON {
                t.Error("response is not valid JSON")
            }
        
            _, err = session.TeamActivity(teamID)
        
            if err != nil {
                t.Error("error creating struct: ", err)
            }
        })
    }
}

func TestTeamOwnsByWeek(t *testing.T) {
    teamIDs := []int{3804, 3421, 2978}
    
    for _, teamID := range teamIDs {
        t.Run(fmt.Sprintf("id %d", teamID), func(t *testing.T) {
            teamIDString, _ := toPositiveIntString(teamID)
			url := "https://www.hackthebox.com/api/v4/team/stats/owns/" + teamIDString

            responseCode, isValidJSON, err := checkResponse(session, url)

            if responseCode != 200 {
                t.Error("response code should be 200, got", responseCode)
            }
        
            if !isValidJSON {
                t.Error("response is not valid JSON")
            }
        
            _, err = session.TeamOwnsByWeek(teamID)
        
            if err != nil {
                t.Error("error creating struct: ", err)
            }
        })
    }
}

func TestTeamStatsDuring(t *testing.T) {
	testCases := []struct {
        teamID  int
        duration  Duration
    }{
        {3804, Year},
        {3421, Week},
        {2978, HalfYear},
    }
    
    for _, tc := range testCases {
        t.Run(fmt.Sprintf("id %d duration %s", tc.teamID, tc.duration), func(t *testing.T) {
            teamIDString, _ := toPositiveIntString(tc.teamID)
			url := "https://www.hackthebox.com/api/v4/team/graph/" + teamIDString + "?duration=" + string(tc.duration)

            responseCode, isValidJSON, err := checkResponse(session, url)

            if responseCode != 200 {
                t.Error("response code should be 200, got", responseCode)
            }
        
            if !isValidJSON {
                t.Error("response is not valid JSON")
            }
        
            _, err = session.TeamStatsDuring(tc.teamID, tc.duration)
        
            if err != nil {
                t.Error("error creating struct: ", err)
            }
        })
    }
}

func TestTeamOwnsByPath(t *testing.T) {
    teamIDs := []int{3804, 3421, 2978}
    
    for _, teamID := range teamIDs {
        t.Run(fmt.Sprintf("id %d", teamID), func(t *testing.T) {
            teamIDString, _ := toPositiveIntString(teamID)
			url := "https://www.hackthebox.com/api/v4/team/chart/machines/attack/" + teamIDString

            responseCode, isValidJSON, err := checkResponse(session, url)

            if responseCode != 200 {
                t.Error("response code should be 200, got", responseCode)
            }
        
            if !isValidJSON {
                t.Error("response is not valid JSON")
            }
        
            _, err = session.TeamOwnsByPath(teamID)
        
            if err != nil {
                t.Error("error creating struct: ", err)
            }
        })
    }
}

func TestTeamMembers(t *testing.T) {
    teamIDs := []int{3804, 3421, 2978}
    
    for _, teamID := range teamIDs {
        t.Run(fmt.Sprintf("id %d", teamID), func(t *testing.T) {
            teamIDString, _ := toPositiveIntString(teamID)
			url := "https://www.hackthebox.com/api/v4/team/members/" + teamIDString

            responseCode, isValidJSON, err := checkResponse(session, url)

            if responseCode != 200 {
                t.Error("response code should be 200, got", responseCode)
            }
        
            if !isValidJSON {
                t.Error("response is not valid JSON")
            }
        
            _, err = session.TeamMembers(teamID)
        
            if err != nil {
                t.Error("error creating struct: ", err)
            }
        })
    }
}

func TestTeamInvitations(t *testing.T) {
    teamIDs := []int{3804, 3421, 2978}
    
    for _, teamID := range teamIDs {
        t.Run(fmt.Sprintf("id %d", teamID), func(t *testing.T) {
            teamIDString, _ := toPositiveIntString(teamID)
			url := "https://www.hackthebox.com/api/v4/team/invitations/" + teamIDString

            responseCode, isValidJSON, err := checkResponse(session, url)

            if responseCode != 200 {
                t.Error("response code should be 200, got", responseCode)
            }
        
            if !isValidJSON {
                t.Error("response is not valid JSON")
            }
        
            _, err = session.TeamInvitations(teamID)
        
            if err != nil {
                t.Error("error creating struct: ", err)
            }
        })
    }
}