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

func TestEndgamesActive(t *testing.T) {
    url := "https://www.hackthebox.com/api/v4/endgames"
    responseCode, isValidJSON, err := checkResponse(session, url)

    if responseCode != 200 {
        t.Error("response code should be 200, got", responseCode)
    }

    if !isValidJSON {
        t.Error("response is not valid JSON")
    }

    _, err = session.EndgamesActive()

    if err != nil {
        t.Error("error creating struct: ", err)
    }
}

func TestEndgame(t *testing.T) {
    endgameIDs := []int{1, 4, 7}
    
    for _, endgameID := range endgameIDs {
        t.Run(fmt.Sprintf("id %d", endgameID), func(t *testing.T) {
            endgameIDString, _ := toPositiveIntString(endgameID)
			url := "https://www.hackthebox.com/api/v4/endgame/" + endgameIDString

            responseCode, isValidJSON, err := checkResponse(session, url)

            if responseCode != 200 {
                t.Error("response code should be 200, got", responseCode)
            }
        
            if !isValidJSON {
                t.Error("response is not valid JSON")
            }
        
            _, err = session.Endgame(endgameID)
        
            if err != nil {
                t.Error("error creating struct: ", err)
            }
        })
    }
}

func TestEndgameFlags(t *testing.T) {
    endgameIDs := []int{1, 4, 7}

    for _, endgameID := range endgameIDs {
        t.Run(fmt.Sprintf("id %d", endgameID), func(t *testing.T) {
            endgameIDString, _ := toPositiveIntString(endgameID)
			url := "https://www.hackthebox.com/api/v4/endgame/" + endgameIDString + "/flags"

            responseCode, isValidJSON, err := checkResponse(session, url)

            if responseCode != 200 {
                t.Error("response code should be 200, got", responseCode)
            }
        
            if !isValidJSON {
                t.Error("response is not valid JSON")
            }
        
            _, err = session.EndgameFlags(endgameID)
        
            if err != nil {
                t.Error("error creating struct: ", err)
            }
        })
    }
}

func TestEndgameMachines(t *testing.T) {
    endgameIDs := []int{1, 4, 7}

    for _, endgameID := range endgameIDs {
        t.Run(fmt.Sprintf("id %d", endgameID), func(t *testing.T) {
            endgameIDString, _ := toPositiveIntString(endgameID)
			url := "https://www.hackthebox.com/api/v4/endgame/" + endgameIDString + "/machines"

            responseCode, isValidJSON, err := checkResponse(session, url)

            if responseCode != 200 {
                t.Error("response code should be 200, got", responseCode)
            }
        
            if !isValidJSON {
                t.Error("response is not valid JSON")
            }
        
            _, err = session.EndgameMachines(endgameID)
        
            if err != nil {
                t.Error("error creating struct: ", err)
            }
        })
    }
}