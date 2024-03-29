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

func TestTracksActive(t *testing.T) {
    url := "https://www.hackthebox.com/api/v4/tracks"
    responseCode, isValidJSON, err := checkResponse(session, url)

    if responseCode != 200 {
        t.Error("response code should be 200, got", responseCode)
    }

    if !isValidJSON {
        t.Error("response is not valid JSON")
    }

    _, err = session.TracksActive()

    if err != nil {
        t.Error("error creating struct: ", err)
    }
}

func TestTrack(t *testing.T) {
    trackIDs := []int{1, 2, 3}
    
    for _, trackID := range trackIDs {
        t.Run(fmt.Sprintf("id %d", trackID), func(t *testing.T) {
            trackIDString, _ := toPositiveIntString(trackID)
			url := "https://www.hackthebox.com/api/v4/tracks/" + trackIDString

            responseCode, isValidJSON, err := checkResponse(session, url)

            if responseCode != 200 {
                t.Error("response code should be 200, got", responseCode)
            }
        
            if !isValidJSON {
                t.Error("response is not valid JSON")
            }
        
            _, err = session.Track(trackID)
        
            if err != nil {
                t.Error("error creating struct: ", err)
            }
        })
    }
}