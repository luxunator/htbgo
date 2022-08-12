package htbgo

import (
    "encoding/json"
    "io/ioutil"
    "net/http"
    "testing"
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

func TestBadges(t *testing.T) {
    url := "https://www.hackthebox.com/api/v4/category/badges"
    responseCode, isValidJSON, err := checkResponse(session, url)

    if responseCode != 200 {
        t.Error("response code should be 200, got", responseCode)
    }

    if !isValidJSON {
        t.Error("response is not valid JSON")
    }

    _, err = session.Badges()

    if err != nil {
        t.Error("error creating struct: ", err)
    }
}

func TestReportBugAreas(t *testing.T) {
    url := "https://www.hackthebox.com/api/v4/user/feedback/bug/areas"
    responseCode, isValidJSON, err := checkResponse(session, url)

    if responseCode != 200 {
        t.Error("response code should be 200, got", responseCode)
    }

    if !isValidJSON {
        t.Error("response is not valid JSON")
    }

    _, err = session.ReportBugAreas()

    if err != nil {
        t.Error("error creating struct: ", err)
    }
}

func TestLabsStats(t *testing.T) {
    url := "https://www.hackthebox.com/api/v4/content/stats"
    responseCode, isValidJSON, err := checkResponse(session, url)

    if responseCode != 200 {
        t.Error("response code should be 200, got", responseCode)
    }

    if !isValidJSON {
        t.Error("response is not valid JSON")
    }

    _, err = session.LabsStats()

    if err != nil {
        t.Error("error creating struct: ", err)
    }
}

func TestChangelogs(t *testing.T) {
    url := "https://www.hackthebox.com/api/v4/changelogs"
    responseCode, isValidJSON, err := checkResponse(session, url)

    if responseCode != 200 {
        t.Error("response code should be 200, got", responseCode)
    }

    if !isValidJSON {
        t.Error("response is not valid JSON")
    }

    _, err = session.Changelogs()

    if err != nil {
        t.Error("error creating struct: ", err)
    }
}

func TestServers(t *testing.T) {
    url := "https://www.hackthebox.com/api/v4/lab/list"
    responseCode, isValidJSON, err := checkResponse(session, url)

    if responseCode != 200 {
        t.Error("response code should be 200, got", responseCode)
    }

    if !isValidJSON {
        t.Error("response is not valid JSON")
    }

    _, err = session.Servers()

    if err != nil {
        t.Error("error creating struct: ", err)
    }
}

func TestReportImprovementAreas(t *testing.T) {
    url := "https://www.hackthebox.com/api/v4/user/feedback/improvement/areas"
    responseCode, isValidJSON, err := checkResponse(session, url)

    if responseCode != 200 {
        t.Error("response code should be 200, got", responseCode)
    }

    if !isValidJSON {
        t.Error("response is not valid JSON")
    }

    _, err = session.ReportImprovementAreas()

    if err != nil {
        t.Error("error creating struct: ", err)
    }
}

func TestStats(t *testing.T) {
    url := "https://www.hackthebox.com/api/v4/user/dashboard"
    responseCode, isValidJSON, err := checkResponse(session, url)

    if responseCode != 200 {
        t.Error("response code should be 200, got", responseCode)
    }

    if !isValidJSON {
        t.Error("response is not valid JSON")
    }

    _, err = session.Stats()

    if err != nil {
        t.Error("error creating struct: ", err)
    }
}

func TestChangelogsSidebar(t *testing.T) {
    url := "https://www.hackthebox.com/api/v4/sidebar/changelog"
    responseCode, isValidJSON, err := checkResponse(session, url)

    if responseCode != 200 {
        t.Error("response code should be 200, got", responseCode)
    }

    if !isValidJSON {
        t.Error("response is not valid JSON")
    }

    _, err = session.ChangelogsSidebar()

    if err != nil {
        t.Error("error creating struct: ", err)
    }
}

func TestAnnouncement(t *testing.T) {
    url := "https://www.hackthebox.com/api/v4/sidebar/announcement"
    responseCode, isValidJSON, err := checkResponse(session, url)

    if responseCode != 200 {
        t.Error("response code should be 200, got", responseCode)
    }

    if !isValidJSON {
        t.Error("response is not valid JSON")
    }

    _, err = session.Announcement()

    if err != nil {
        t.Error("error creating struct: ", err)
    }
}