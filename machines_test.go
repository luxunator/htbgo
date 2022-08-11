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

func TestMachineMatrix(t *testing.T) {
    machineIDs := []int{483, 452, 1}
    
    for _, machineID := range machineIDs {
        t.Run(fmt.Sprintf("id %d", machineID), func(t *testing.T) {
            machineIDString, _ := toPositiveIntString(machineID)
			url := "https://www.hackthebox.com/api/v4/machine/graph/matrix/" + machineIDString

            responseCode, isValidJSON, err := checkResponse(session, url)

            if responseCode != 200 {
                t.Error("response code should be 200, got", responseCode)
            }
        
            if !isValidJSON {
                t.Error("response is not valid JSON")
            }
        
            _, err = session.MachineMatrix(machineID)
        
            if err != nil {
                t.Error("error creating struct: ", err)
            }
        })
    }
}

func TestMachine(t *testing.T) {
    machineIDs := []int{483, 452, 1}
    
    for _, machineID := range machineIDs {
        t.Run(fmt.Sprintf("id %d", machineID), func(t *testing.T) {
            machineIDString, _ := toPositiveIntString(machineID)
			url := "https://www.hackthebox.com/api/v4/machine/info/" + machineIDString

            responseCode, isValidJSON, err := checkResponse(session, url)

            if responseCode != 200 {
                t.Error("response code should be 200, got", responseCode)
            }
        
            if !isValidJSON {
                t.Error("response is not valid JSON")
            }
        
            _, err = session.Machine(machineID)
        
            if err != nil {
                t.Error("error creating struct: ", err)
            }
        })
    }
}

func TestMachineProfile(t *testing.T) {
    machineIDs := []int{483, 452, 1}
    
    for _, machineID := range machineIDs {
        t.Run(fmt.Sprintf("id %d", machineID), func(t *testing.T) {
            machineIDString, _ := toPositiveIntString(machineID)
			url := "https://www.hackthebox.com/api/v4/machine/profile/" + machineIDString

            responseCode, isValidJSON, err := checkResponse(session, url)

            if responseCode != 200 {
                t.Error("response code should be 200, got", responseCode)
            }
        
            if !isValidJSON {
                t.Error("response is not valid JSON")
            }
        
            _, err = session.MachineProfile(machineID)
        
            if err != nil {
                t.Error("error creating struct: ", err)
            }
        })
    }
}

func TestMachinesActive(t *testing.T) {
    url := "https://www.hackthebox.com/api/v4/machine/list"
    responseCode, isValidJSON, err := checkResponse(session, url)

    if responseCode != 200 {
        t.Error("response code should be 200, got", responseCode)
    }

    if !isValidJSON {
        t.Error("response is not valid JSON")
    }

    _, err = session.MachinesActive()

    if err != nil {
        t.Error("error creating struct: ", err)
    }
}

func TestMachinesRetired(t *testing.T) {
    url := "https://www.hackthebox.com/api/v4/machine/list/retired"
    responseCode, isValidJSON, err := checkResponse(session, url)

    if responseCode != 200 {
        t.Error("response code should be 200, got", responseCode)
    }

    if !isValidJSON {
        t.Error("response is not valid JSON")
    }

    _, err = session.MachinesRetired()

    if err != nil {
        t.Error("error creating struct: ", err)
    }
}

func TestMachineTopUsers(t *testing.T) {
    machineIDs := []int{483, 452, 1}
    
    for _, machineID := range machineIDs {
        t.Run(fmt.Sprintf("id %d", machineID), func(t *testing.T) {
            machineIDString, _ := toPositiveIntString(machineID)
			url := "https://www.hackthebox.com/api/v4/machine/owns/top/" + machineIDString

            responseCode, isValidJSON, err := checkResponse(session, url)

            if responseCode != 200 {
                t.Error("response code should be 200, got", responseCode)
            }
        
            if !isValidJSON {
                t.Error("response is not valid JSON")
            }
        
            _, err = session.MachineTopUsers(machineID)
        
            if err != nil {
                t.Error("error creating struct: ", err)
            }
        })
    }
}

func TestMachinesToRelease(t *testing.T) {
    url := "https://www.hackthebox.com/api/v4/machine/unreleased"
    responseCode, isValidJSON, err := checkResponse(session, url)

    if responseCode != 200 {
        t.Error("response code should be 200, got", responseCode)
    }

    if !isValidJSON {
        t.Error("response is not valid JSON")
    }

    _, err = session.MachinesToRelease()

    if err != nil {
        t.Error("error creating struct: ", err)
    }
}

func TestMachinesTodo(t *testing.T) {
    url := "https://www.hackthebox.com/api/v4/machine/list/todo"
    responseCode, isValidJSON, err := checkResponse(session, url)

    if responseCode != 200 {
        t.Error("response code should be 200, got", responseCode)
    }

    if !isValidJSON {
        t.Error("response is not valid JSON")
    }

    _, err = session.MachinesTodo()

    if err != nil {
        t.Error("error creating struct: ", err)
    }
}

func TestMachineTags(t *testing.T) {
    machineIDs := []int{3, 2, 1}
    
    for _, machineID := range machineIDs {
        t.Run(fmt.Sprintf("id %d", machineID), func(t *testing.T) {
            machineIDString, _ := toPositiveIntString(machineID)
			url := "https://www.hackthebox.com/api/v4/machine/tags/" + machineIDString

            responseCode, isValidJSON, err := checkResponse(session, url)

            if responseCode != 200 {
                t.Error("response code should be 200, got", responseCode)
            }
        
            if !isValidJSON {
                t.Error("response is not valid JSON")
            }
        
            _, err = session.MachineTags(machineID)
        
            if err != nil {
                t.Error("error creating struct: ", err)
            }
        })
    }
}

func TestMachineChangelog(t *testing.T) {
    machineIDs := []int{483, 452, 1}
    
    for _, machineID := range machineIDs {
        t.Run(fmt.Sprintf("id %d", machineID), func(t *testing.T) {
            machineIDString, _ := toPositiveIntString(machineID)
			url := "https://www.hackthebox.com/api/v4/machine/changelog/" + machineIDString

            responseCode, isValidJSON, err := checkResponse(session, url)

            if responseCode != 200 {
                t.Error("response code should be 200, got", responseCode)
            }
        
            if !isValidJSON {
                t.Error("response is not valid JSON")
            }
        
            _, err = session.MachineChangelog(machineID)
        
            if err != nil {
                t.Error("error creating struct: ", err)
            }
        })
    }
}

func TestMachineReviews(t *testing.T) {
    machineIDs := []int{3, 2, 1}
    
    for _, machineID := range machineIDs {
        t.Run(fmt.Sprintf("id %d", machineID), func(t *testing.T) {
            machineIDString, _ := toPositiveIntString(machineID)
			url := "https://www.hackthebox.com/api/v4/machine/reviews/" + machineIDString

            responseCode, isValidJSON, err := checkResponse(session, url)

            if responseCode != 200 {
                t.Error("response code should be 200, got", responseCode)
            }
        
            if !isValidJSON {
                t.Error("response is not valid JSON")
            }
        
            _, err = session.MachineReviews(machineID)
        
            if err != nil {
                t.Error("error creating struct: ", err)
            }
        })
    }
}