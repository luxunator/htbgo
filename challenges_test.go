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


func TestChallengesSuggested(t *testing.T) {
    url := "https://www.hackthebox.com/api/v4/challenge/recommended"
    responseCode, isValidJSON, err := checkResponse(session, url)

    if responseCode != 200 {
        t.Error("response code should be 200, got", responseCode)
    }

    if !isValidJSON {
        t.Error("response is not valid JSON")
    }

    _, err = session.ChallengesSuggested()

    if err != nil {
        t.Error("error creating struct: ", err)
    }
}

func TestChallengesRetiredSuggested(t *testing.T) {
    url := "https://www.hackthebox.com/api/v4/challenge/recommended/retired"
    responseCode, isValidJSON, err := checkResponse(session, url)

    if responseCode != 200 {
        t.Error("response code should be 200, got", responseCode)
    }

    if !isValidJSON {
        t.Error("response is not valid JSON")
    }

    _, err = session.ChallengesRetiredSuggested()

    if err != nil {
        t.Error("error creating struct: ", err)
    }
}

func TestChallengeCategories(t *testing.T) {
    url := "https://www.hackthebox.com/api/v4/challenge/categories/list"
    responseCode, isValidJSON, err := checkResponse(session, url)

    if responseCode != 200 {
        t.Error("response code should be 200, got", responseCode)
    }

    if !isValidJSON {
        t.Error("response is not valid JSON")
    }

    _, err = session.ChallengeCategories()

    if err != nil {
        t.Error("error creating struct: ", err)
    }
}

func TestChallengesActive(t *testing.T) {
    url := "https://www.hackthebox.com/api/v4/challenge/list"
    responseCode, isValidJSON, err := checkResponse(session, url)

    if responseCode != 200 {
        t.Error("response code should be 200, got", responseCode)
    }

    if !isValidJSON {
        t.Error("response is not valid JSON")
    }

    _, err = session.ChallengesActive()

    if err != nil {
        t.Error("error creating struct: ", err)
    }
}

func TestChallengesRetired(t *testing.T) {
    url := "https://www.hackthebox.com/api/v4/challenge/list/retired"
    responseCode, isValidJSON, err := checkResponse(session, url)

    if responseCode != 200 {
        t.Error("response code should be 200, got", responseCode)
    }

    if !isValidJSON {
        t.Error("response is not valid JSON")
    }

    _, err = session.ChallengesRetired()

    if err != nil {
        t.Error("error creating struct: ", err)
    }
}

func TestChallenge(t *testing.T) {
    challengeIDs := []int{390, 4, 2}

    for _, challengeID := range challengeIDs {
        t.Run(fmt.Sprintf("id %d", challengeID), func(t *testing.T) {
            challengeIDString, _ := toPositiveIntString(challengeID)
            url := "https://www.hackthebox.com/api/v4/challenge/info/" + challengeIDString

            responseCode, isValidJSON, err := checkResponse(session, url)

            if responseCode != 200 {
                t.Error("response code should be 200, got", responseCode)
            }
        
            if !isValidJSON {
                t.Error("response is not valid JSON")
            }
        
            _, err = session.Challenge(challengeID)
        
            if err != nil {
                t.Error("error creating struct: ", err)
            }
        })
    }
}

func TestChallengeActivty(t *testing.T) {
    challengeIDs := []int{390, 4, 2}
    
    for _, challengeID := range challengeIDs {
        t.Run(fmt.Sprintf("id %d", challengeID), func(t *testing.T) {
            challengeIDString, _ := toPositiveIntString(challengeID)
            url := "https://www.hackthebox.com/api/v4/challenge/activity/" + challengeIDString

            responseCode, isValidJSON, err := checkResponse(session, url)

            if responseCode != 200 {
                t.Error("response code should be 200, got", responseCode)
            }
        
            if !isValidJSON {
                t.Error("response is not valid JSON")
            }
        
            _, err = session.ChallengeActivity(challengeID)
        
            if err != nil {
                t.Error("error creating struct: ", err)
            }
        })
    }
}