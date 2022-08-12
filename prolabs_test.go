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

func TestProLabsActive(t *testing.T) {
    url := "https://www.hackthebox.com/api/v4/prolabs"
    responseCode, isValidJSON, err := checkResponse(session, url)

    if responseCode != 200 {
        t.Error("response code should be 200, got", responseCode)
    }

    if !isValidJSON {
        t.Error("response is not valid JSON")
    }

    _, err = session.ProLabsActive()

    if err != nil {
        t.Error("error creating struct: ", err)
    }
}

func TestProLab(t *testing.T) {
    proLabIDs := []int{1, 3, 5}
    
    for _, proLabID := range proLabIDs {
        t.Run(fmt.Sprintf("id %d", proLabID), func(t *testing.T) {
            proLabIDString, _ := toPositiveIntString(proLabID)
			url := "https://www.hackthebox.com/api/v4/prolab/" + proLabIDString + "/info"

            responseCode, isValidJSON, err := checkResponse(session, url)

            if responseCode != 200 {
                t.Error("response code should be 200, got", responseCode)
            }
        
            if !isValidJSON {
                t.Error("response is not valid JSON")
            }
        
            _, err = session.ProLab(proLabID)
        
            if err != nil {
                t.Error("error creating struct: ", err)
            }
        })
    }
}

func TestProLabOverview(t *testing.T) {
    proLabIDs := []int{1, 3, 5}
    
    for _, proLabID := range proLabIDs {
        t.Run(fmt.Sprintf("id %d", proLabID), func(t *testing.T) {
            proLabIDString, _ := toPositiveIntString(proLabID)
			url := "https://www.hackthebox.com/api/v4/prolab/" + proLabIDString + "/overview"

            responseCode, isValidJSON, err := checkResponse(session, url)

            if responseCode != 200 {
                t.Error("response code should be 200, got", responseCode)
            }
        
            if !isValidJSON {
                t.Error("response is not valid JSON")
            }
        
            _, err = session.ProLabOverview(proLabID)
        
            if err != nil {
                t.Error("error creating struct: ", err)
            }
        })
    }
}

func TestProLabMachines(t *testing.T) {
    proLabIDs := []int{1, 3, 5}
    
    for _, proLabID := range proLabIDs {
        t.Run(fmt.Sprintf("id %d", proLabID), func(t *testing.T) {
            proLabIDString, _ := toPositiveIntString(proLabID)
			url := "https://www.hackthebox.com/api/v4/prolab/" + proLabIDString + "/machines"

            responseCode, isValidJSON, err := checkResponse(session, url)

            if responseCode != 200 {
                t.Error("response code should be 200, got", responseCode)
            }
        
            if !isValidJSON {
                t.Error("response is not valid JSON")
            }
        
            _, err = session.ProLabMachines(proLabID)
        
            if err != nil {
                t.Error("error creating struct: ", err)
            }
        })
    }
}

func TestProLabFlags(t *testing.T) {
    proLabIDs := []int{1, 3, 5}
    
    for _, proLabID := range proLabIDs {
        t.Run(fmt.Sprintf("id %d", proLabID), func(t *testing.T) {
            proLabIDString, _ := toPositiveIntString(proLabID)
			url := "https://www.hackthebox.com/api/v4/prolab/" + proLabIDString + "/flags"

            responseCode, isValidJSON, err := checkResponse(session, url)

            if responseCode != 200 {
                t.Error("response code should be 200, got", responseCode)
            }
        
            if !isValidJSON {
                t.Error("response is not valid JSON")
            }
        
            _, err = session.ProLabFlags(proLabID)
        
            if err != nil {
                t.Error("error creating struct: ", err)
            }
        })
    }
}

func TestProLabFAQ(t *testing.T) {
    proLabIDs := []int{1, 3, 5}
    
    for _, proLabID := range proLabIDs {
        t.Run(fmt.Sprintf("id %d", proLabID), func(t *testing.T) {
            proLabIDString, _ := toPositiveIntString(proLabID)
			url := "https://www.hackthebox.com/api/v4/prolab/" + proLabIDString + "/faq"

            responseCode, isValidJSON, err := checkResponse(session, url)

            if responseCode != 200 {
                t.Error("response code should be 200, got", responseCode)
            }
        
            if !isValidJSON {
                t.Error("response is not valid JSON")
            }
        
            _, err = session.ProLabFAQ(proLabID)
        
            if err != nil {
                t.Error("error creating struct: ", err)
            }
        })
    }
}

func TestProLabSelectedReviews(t *testing.T) {
    proLabIDs := []int{1, 3, 5}
    
    for _, proLabID := range proLabIDs {
        t.Run(fmt.Sprintf("id %d", proLabID), func(t *testing.T) {
            proLabIDString, _ := toPositiveIntString(proLabID)
			url := "https://www.hackthebox.com/api/v4/prolab/" + proLabIDString + "/reviews_overview"

            responseCode, isValidJSON, err := checkResponse(session, url)

            if responseCode != 200 {
                t.Error("response code should be 200, got", responseCode)
            }
        
            if !isValidJSON {
                t.Error("response is not valid JSON")
            }
        
            _, err = session.ProLabSelectedReviews(proLabID)
        
            if err != nil {
                t.Error("error creating struct: ", err)
            }
        })
    }
}

func TestProLabPaginatedReviews(t *testing.T) {
	testCases := []struct {
        proLabID  int
        page  int
    }{
        {1, 5},
        {3, 7},
        {5, 2},
    }
    
    for _, tc := range testCases {
        t.Run(fmt.Sprintf("id %d page %d", tc.proLabID, tc.page), func(t *testing.T) {
            proLabIDString, _ := toPositiveIntString(tc.proLabID)
			pageString, _ := toPositiveIntString(tc.page)
			url := "https://www.hackthebox.com/api/v4/prolab/" + proLabIDString + "/reviews?page=" + pageString

            responseCode, isValidJSON, err := checkResponse(session, url)

            if responseCode != 200 {
                t.Error("response code should be 200, got", responseCode)
            }
        
            if !isValidJSON {
                t.Error("response is not valid JSON")
            }
        
            _, err = session.ProLabPaginatedReviews(tc.proLabID, tc.page)
        
            if err != nil {
                t.Error("error creating struct: ", err)
            }
        })
    }
}