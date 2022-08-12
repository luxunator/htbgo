// Package htbgo contains functionality to interact with the Hack The Box v4 API.
package htbgo

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const VERSION string = "0.1.0"

const (
	// Duration constants are used when duration is needed for obtaining information from within a specified period of time
	Year        Duration = "1Y"
	HalfYear    Duration = "6M"
	QuarterYear Duration = "3M"
	Month       Duration = "1M"
	Week        Duration = "1W"
)

var (
	// Err contants are general errors that could be encountered
	ErrIntNotPositive = errors.New("value: integer not positive")
)

// A Session containter holds the information needed to create a valid session with the API
type Session struct {
	Token   string
	Client  *http.Client
	Headers http.Header
}

// A Duration holds a representation of a period of time
type Duration string

// Difficulties is a container for difficulty types defined by the API
type Difficulties struct {
	Cake      int `json:"counterCake"`
	VeryEasy  int `json:"counterVeryEasy"`
	Easy      int `json:"counterEasy"`
	TooEasy   int `json:"counterTooEasy"`
	Medium    int `json:"counterMedium"`
	BitHard   int `json:"counterBitHard"`
	Hard      int `json:"counterHard"`
	TooHard   int `json:"counterTooHard"`
	ExtraHard int `json:"counterExHard"`
	BrainFuck int `json:"counterBrainFuck"`
}

// New is a function that will create a session to interact with the API when given a valid token
func New(token string) (s *Session, err error) {

	s = &Session{
		Token:  token,
		Client: &http.Client{Timeout: (15 * time.Second)},
		Headers: http.Header{
			"User-Agent":    {"htbgo v" + VERSION},
			"Authorization": {"Bearer " + token},
		},
	}

	return
}

func parseJSON(s *Session, url string, schema interface{}) (err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	req.Header = s.Headers

	resp, err := s.Client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(body), &schema)

	return
}

func stringFromVIP(isVIP bool) (stringValue string) {
	stringValue = "0"

	if isVIP {
		stringValue = "1"
	}

	return
}

func toPositiveIntString(num int) (positiveIntString string, err error) {
	if num < 1 {
		return "", ErrIntNotPositive
	}

	positiveIntString = strconv.Itoa(num)

	return
}
