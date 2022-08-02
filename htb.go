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
	Year        Duration = "1Y"
	HalfYear    Duration = "6M"
	QuarterYear Duration = "3M"
	Month       Duration = "1M"
	Week        Duration = "1W"
)

var (
	ErrIntNotPositive = errors.New("value: integer not positive")
)

type Duration string

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
