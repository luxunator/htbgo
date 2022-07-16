package htbgo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

const VERSION string = "0.1.0"

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

func parseJSON(s *Session, url string, location interface{}) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header = s.Headers

	resp, err := s.Client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal([]byte(body), &location)
	if err != nil {
		panic(err)
	}

	return
}
