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

func TestConnectionStatus(t *testing.T) {
    url := "https://www.hackthebox.com/api/v4/user/connection/status"
    responseCode, isValidJSON, err := checkResponse(session, url)

    if responseCode != 200 {
        t.Error("response code should be 200, got", responseCode)
    }

    if !isValidJSON {
        t.Error("response is not valid JSON")
    }

    _, err = session.ConnectionStatus()

    if err != nil {
        t.Error("error creating struct: ", err)
    }
}

func TestActiveMachine(t *testing.T) {
    url := "https://www.hackthebox.com/api/v4/machine/active"
    responseCode, isValidJSON, err := checkResponse(session, url)

    if responseCode != 200 {
        t.Error("response code should be 200, got", responseCode)
    }

    if !isValidJSON {
        t.Error("response is not valid JSON")
    }

    _, err = session.ActiveMachine()

    if err != nil {
        t.Error("error creating struct: ", err)
    }
}

func TestFollowers(t *testing.T) {
    url := "https://www.hackthebox.com/api/v4/user/followers"
    responseCode, isValidJSON, err := checkResponse(session, url)

    if responseCode != 200 {
        t.Error("response code should be 200, got", responseCode)
    }

    if !isValidJSON {
        t.Error("response is not valid JSON")
    }

    _, err = session.Followers()

    if err != nil {
        t.Error("error creating struct: ", err)
    }
}

func TestProfile(t *testing.T) {
    url := "https://www.hackthebox.com/api/v4/user/info"
    responseCode, isValidJSON, err := checkResponse(session, url)

    if responseCode != 200 {
        t.Error("response code should be 200, got", responseCode)
    }

    if !isValidJSON {
        t.Error("response is not valid JSON")
    }

    _, err = session.Profile()

    if err != nil {
        t.Error("error creating struct: ", err)
    }
}

func TestSettings(t *testing.T) {
    url := "https://www.hackthebox.com/api/v4/user/settings"
    responseCode, isValidJSON, err := checkResponse(session, url)

    if responseCode != 200 {
        t.Error("response code should be 200, got", responseCode)
    }

    if !isValidJSON {
        t.Error("response is not valid JSON")
    }

    _, err = session.Settings()

    if err != nil {
        t.Error("error creating struct: ", err)
    }
}

func TestSubscriptions(t *testing.T) {
    url := "https://www.hackthebox.com/api/v4/user/subscriptions"
    responseCode, isValidJSON, err := checkResponse(session, url)

    if responseCode != 200 {
        t.Error("response code should be 200, got", responseCode)
    }

    if !isValidJSON {
        t.Error("response is not valid JSON")
    }

    _, err = session.Subscriptions()

    if err != nil {
        t.Error("error creating struct: ", err)
    }
}

func TestSubscriptionsBalance(t *testing.T) {
    url := "https://www.hackthebox.com/api/v4/user/subscriptions/balance"
    responseCode, isValidJSON, err := checkResponse(session, url)

    if responseCode != 200 {
        t.Error("response code should be 200, got", responseCode)
    }

    if !isValidJSON {
        t.Error("response is not valid JSON")
    }

    _, err = session.SubscriptionsBalance()

    if err != nil {
        t.Error("error creating struct: ", err)
    }
}

func TestSubscriptionsRecurly(t *testing.T) {
    url := "https://www.hackthebox.com/api/v4/user/subscriptions/manage/recurly"
    responseCode, isValidJSON, err := checkResponse(session, url)

    if responseCode != 200 {
        t.Error("response code should be 200, got", responseCode)
    }

    if !isValidJSON {
        t.Error("response is not valid JSON")
    }

    _, err = session.SubscriptionsRecurly()

    if err != nil {
        t.Error("error creating struct: ", err)
    }
}

func TestEnrolledTracks(t *testing.T) {
    url := "https://www.hackthebox.com/api/v4/user/tracks"
    responseCode, isValidJSON, err := checkResponse(session, url)

    if responseCode != 200 {
        t.Error("response code should be 200, got", responseCode)
    }

    if !isValidJSON {
        t.Error("response is not valid JSON")
    }

    _, err = session.EnrolledTracks()

    if err != nil {
        t.Error("error creating struct: ", err)
    }
}

func TestUserRelationship(t *testing.T) {
    userIDs := []int{1063066, 13489, 15059}
    
    for _, userID := range userIDs {
        t.Run(fmt.Sprintf("id %d", userID), func(t *testing.T) {
            userIDString, _ := toPositiveIntString(userID)
			url := "https://www.hackthebox.com/api/v4/user/profile/basic/" + userIDString

            responseCode, isValidJSON, err := checkResponse(session, url)

            if responseCode != 200 {
                t.Error("response code should be 200, got", responseCode)
            }
        
            if !isValidJSON {
                t.Error("response is not valid JSON")
            }
        
            _, err = session.UserRelationship(userID)
        
            if err != nil {
                t.Error("error creating struct: ", err)
            }
        })
    }
}

func TestUserMachines(t *testing.T) {
    userIDs := []int{1063066, 13489, 15059}
    
    for _, userID := range userIDs {
        t.Run(fmt.Sprintf("id %d", userID), func(t *testing.T) {
            userIDString, _ := toPositiveIntString(userID)
			url := "https://www.hackthebox.com/api/v4/profile/progress/machines/os/" + userIDString

            responseCode, isValidJSON, err := checkResponse(session, url)

            if responseCode != 200 {
                t.Error("response code should be 200, got", responseCode)
            }
        
            if !isValidJSON {
                t.Error("response is not valid JSON")
            }
        
            _, err = session.UserMachines(userID)
        
            if err != nil {
                t.Error("error creating struct: ", err)
            }
        })
    }
}

func TestUserChallenges(t *testing.T) {
    userIDs := []int{1063066, 13489, 15059}
    
    for _, userID := range userIDs {
        t.Run(fmt.Sprintf("id %d", userID), func(t *testing.T) {
            userIDString, _ := toPositiveIntString(userID)
			url := "https://www.hackthebox.com/api/v4/profile/progress/challenges/" + userIDString

            responseCode, isValidJSON, err := checkResponse(session, url)

            if responseCode != 200 {
                t.Error("response code should be 200, got", responseCode)
            }
        
            if !isValidJSON {
                t.Error("response is not valid JSON")
            }
        
            _, err = session.UserChallenges(userID)
        
            if err != nil {
                t.Error("error creating struct: ", err)
            }
        })
    }
}

func TestUserEndgames(t *testing.T) {
    userIDs := []int{1063066, 13489, 15059}
    
    for _, userID := range userIDs {
        t.Run(fmt.Sprintf("id %d", userID), func(t *testing.T) {
            userIDString, _ := toPositiveIntString(userID)
			url := "https://www.hackthebox.com/api/v4/profile/progress/endgame/" + userIDString

            responseCode, isValidJSON, err := checkResponse(session, url)

            if responseCode != 200 {
                t.Error("response code should be 200, got", responseCode)
            }
        
            if !isValidJSON {
                t.Error("response is not valid JSON")
            }
        
            _, err = session.UserEndgames(userID)
        
            if err != nil {
                t.Error("error creating struct: ", err)
            }
        })
    }
}

func TestUserFortresses(t *testing.T) {
    userIDs := []int{1063066, 13489, 15059}
    
    for _, userID := range userIDs {
        t.Run(fmt.Sprintf("id %d", userID), func(t *testing.T) {
            userIDString, _ := toPositiveIntString(userID)
			url := "https://www.hackthebox.com/api/v4/profile/progress/fortress/" + userIDString

            responseCode, isValidJSON, err := checkResponse(session, url)

            if responseCode != 200 {
                t.Error("response code should be 200, got", responseCode)
            }
        
            if !isValidJSON {
                t.Error("response is not valid JSON")
            }
        
            _, err = session.UserFortresses(userID)
        
            if err != nil {
                t.Error("error creating struct: ", err)
            }
        })
    }
}

func TestUserProLabs(t *testing.T) {
    userIDs := []int{1063066, 13489, 15059}
    
    for _, userID := range userIDs {
        t.Run(fmt.Sprintf("id %d", userID), func(t *testing.T) {
            userIDString, _ := toPositiveIntString(userID)
			url := "https://www.hackthebox.com/api/v4/profile/progress/prolab/" + userIDString

            responseCode, isValidJSON, err := checkResponse(session, url)

            if responseCode != 200 {
                t.Error("response code should be 200, got", responseCode)
            }
        
            if !isValidJSON {
                t.Error("response is not valid JSON")
            }
        
            _, err = session.UserProLabs(userID)
        
            if err != nil {
                t.Error("error creating struct: ", err)
            }
        })
    }
}

func TestUserActivity(t *testing.T) {
    userIDs := []int{1063066, 13489, 15059}
    
    for _, userID := range userIDs {
        t.Run(fmt.Sprintf("id %d", userID), func(t *testing.T) {
            userIDString, _ := toPositiveIntString(userID)
			url := "https://www.hackthebox.com/api/v4/profile/activity/" + userIDString

            responseCode, isValidJSON, err := checkResponse(session, url)

            if responseCode != 200 {
                t.Error("response code should be 200, got", responseCode)
            }
        
            if !isValidJSON {
                t.Error("response is not valid JSON")
            }
        
            _, err = session.UserActivity(userID)
        
            if err != nil {
                t.Error("error creating struct: ", err)
            }
        })
    }
}

func TestUserBloods(t *testing.T) {
    userIDs := []int{1063066, 13489, 15059}
    
    for _, userID := range userIDs {
        t.Run(fmt.Sprintf("id %d", userID), func(t *testing.T) {
            userIDString, _ := toPositiveIntString(userID)
			url := "https://www.hackthebox.com/api/v4/profile/bloods/" + userIDString

            responseCode, isValidJSON, err := checkResponse(session, url)

            if responseCode != 200 {
                t.Error("response code should be 200, got", responseCode)
            }
        
            if !isValidJSON {
                t.Error("response is not valid JSON")
            }
        
            _, err = session.UserBloods(userID)
        
            if err != nil {
                t.Error("error creating struct: ", err)
            }
        })
    }
}

func TestUserSubmissions(t *testing.T) {
    userIDs := []int{1063066, 13489, 15059}
    
    for _, userID := range userIDs {
        t.Run(fmt.Sprintf("id %d", userID), func(t *testing.T) {
            userIDString, _ := toPositiveIntString(userID)
			url := "https://www.hackthebox.com/api/v4/profile/content/" + userIDString

            responseCode, isValidJSON, err := checkResponse(session, url)

            if responseCode != 200 {
                t.Error("response code should be 200, got", responseCode)
            }
        
            if !isValidJSON {
                t.Error("response is not valid JSON")
            }
        
            _, err = session.UserSubmissions(userID)
        
            if err != nil {
                t.Error("error creating struct: ", err)
            }
        })
    }
}

func TestUserAchievements(t *testing.T) {
	testCases := []struct {
        userID  int
        duration  Duration
    }{
        {1063066, Year},
        {13489, Week},
        {15059, HalfYear},
    }
    
    for _, tc := range testCases {
        t.Run(fmt.Sprintf("id %d duration %s", tc.userID, tc.duration), func(t *testing.T) {
            userIDString, _ := toPositiveIntString(tc.userID)
			url := "https://www.hackthebox.com/api/v4/profile/graph/" + string(tc.duration) + "/" + userIDString

            responseCode, isValidJSON, err := checkResponse(session, url)

            if responseCode != 200 {
                t.Error("response code should be 200, got", responseCode)
            }
        
            if !isValidJSON {
                t.Error("response is not valid JSON")
            }
        
            _, err = session.UserAchievements(tc.userID, tc.duration)
        
            if err != nil {
                t.Error("error creating struct: ", err)
            }
        })
    }
}

func TestUserOwnsByPath(t *testing.T) {
    userIDs := []int{1063066, 13489, 15059}
    
    for _, userID := range userIDs {
        t.Run(fmt.Sprintf("id %d", userID), func(t *testing.T) {
            userIDString, _ := toPositiveIntString(userID)
			url := "https://www.hackthebox.com/api/v4/profile/chart/machines/attack/" + userIDString

            responseCode, isValidJSON, err := checkResponse(session, url)

            if responseCode != 200 {
                t.Error("response code should be 200, got", responseCode)
            }
        
            if !isValidJSON {
                t.Error("response is not valid JSON")
            }
        
            _, err = session.UserOwnsByPath(userID)
        
            if err != nil {
                t.Error("error creating struct: ", err)
            }
        })
    }
}

func TestUser(t *testing.T) {
    userIDs := []int{1063066, 13489, 15059}
    
    for _, userID := range userIDs {
        t.Run(fmt.Sprintf("id %d", userID), func(t *testing.T) {
            userIDString, _ := toPositiveIntString(userID)
			url := "https://www.hackthebox.com/api/v4/profile/" + userIDString

            responseCode, isValidJSON, err := checkResponse(session, url)

            if responseCode != 200 {
                t.Error("response code should be 200, got", responseCode)
            }
        
            if !isValidJSON {
                t.Error("response is not valid JSON")
            }
        
            _, err = session.User(userID)
        
            if err != nil {
                t.Error("error creating struct: ", err)
            }
        })
    }
}

func TestUserBadges(t *testing.T) {
    userIDs := []int{1063066, 13489, 15059}
    
    for _, userID := range userIDs {
        t.Run(fmt.Sprintf("id %d", userID), func(t *testing.T) {
            userIDString, _ := toPositiveIntString(userID)
			url := "https://www.hackthebox.com/api/v4/profile/badges/" + userIDString

            responseCode, isValidJSON, err := checkResponse(session, url)

            if responseCode != 200 {
                t.Error("response code should be 200, got", responseCode)
            }
        
            if !isValidJSON {
                t.Error("response is not valid JSON")
            }
        
            _, err = session.UserBadges(userID)
        
            if err != nil {
                t.Error("error creating struct: ", err)
            }
        })
    }
}