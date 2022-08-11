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

func TestUniversitiesByPage(t *testing.T) {
    pages := []int{1, 200, 1000}
    
    for _, page := range pages {
        t.Run(fmt.Sprintf("id %d", page), func(t *testing.T) {
            pageString, _ := toPositiveIntString(page)
			url := "https://www.hackthebox.com/api/v4/university/all/list?page=" + pageString

            responseCode, isValidJSON, err := checkResponse(session, url)

            if responseCode != 200 {
                t.Error("response code should be 200, got", responseCode)
            }
        
            if !isValidJSON {
                t.Error("response is not valid JSON")
            }
        
            _, err = session.UniversitiesByPage(page)
        
            if err != nil {
                t.Error("error creating struct: ", err)
            }
        })
    }
}

func TestUniversitiesSearch(t *testing.T) {
	testCases := []struct {
        search  string
        page  int
    }{
        {"K", 1},
        {"Jdd", 10},
        {"L", 3},
    }
    
    for _, tc := range testCases {
        t.Run(fmt.Sprintf("search %s page %d", tc.search, tc.page), func(t *testing.T) {
			pageString, err := toPositiveIntString(tc.page)
			url := "https://www.hackthebox.com/api/v4/university/all/list?search=" + tc.search + "&page=" + pageString

            responseCode, isValidJSON, err := checkResponse(session, url)

            if responseCode != 200 {
                t.Error("response code should be 200, got", responseCode)
            }
        
            if !isValidJSON {
                t.Error("response is not valid JSON")
            }
        
            _, err = session.UniversitiesSearch(tc.search, tc.page)
        
            if err != nil {
                t.Error("error creating struct: ", err)
            }
        })
    }
}

func TestUniversity(t *testing.T) {
    universityIDs := []int{24, 498, 844}
    
    for _, universityID := range universityIDs {
        t.Run(fmt.Sprintf("id %d", universityID), func(t *testing.T) {
            universityIDString, _ := toPositiveIntString(universityID)
			url := "https://www.hackthebox.com/api/v4/university/profile/" + universityIDString

            responseCode, isValidJSON, err := checkResponse(session, url)

            if responseCode != 200 {
                t.Error("response code should be 200, got", responseCode)
            }
        
            if !isValidJSON {
                t.Error("response is not valid JSON")
            }
        
            _, err = session.University(universityID)
        
            if err != nil {
                t.Error("error creating struct: ", err)
            }
        })
    }
}

func TestUniversityOwns(t *testing.T) {
    universityIDs := []int{24, 498, 844}
    
    for _, universityID := range universityIDs {
        t.Run(fmt.Sprintf("id %d", universityID), func(t *testing.T) {
            universityIDString, _ := toPositiveIntString(universityID)
			url := "https://www.hackthebox.com/api/v4/university/stats/owns/" + universityIDString

            responseCode, isValidJSON, err := checkResponse(session, url)

            if responseCode != 200 {
                t.Error("response code should be 200, got", responseCode)
            }
        
            if !isValidJSON {
                t.Error("response is not valid JSON")
            }
        
            _, err = session.UniversityOwns(universityID)
        
            if err != nil {
                t.Error("error creating struct: ", err)
            }
        })
    }
}

func TestUniversityMembers(t *testing.T) {
    universityIDs := []int{24, 498, 844}
    
    for _, universityID := range universityIDs {
        t.Run(fmt.Sprintf("id %d", universityID), func(t *testing.T) {
            universityIDString, _ := toPositiveIntString(universityID)
			url := "https://www.hackthebox.com/api/v4/university/members/" + universityIDString

            responseCode, isValidJSON, err := checkResponse(session, url)

            if responseCode != 200 {
                t.Error("response code should be 200, got", responseCode)
            }
        
            if !isValidJSON {
                t.Error("response is not valid JSON")
            }
        
            _, err = session.UniversityMembers(universityID)
        
            if err != nil {
                t.Error("error creating struct: ", err)
            }
        })
    }
}