package htbgo

import (
	"strconv"
	"testing"
)

func TestTracks(t *testing.T) {
	t.Run("Retrieves the list of active tracks", func(t *testing.T) {
		url := "https://www.hackthebox.com/api/v4/tracks"
		responseCode, isValidJSON, err := checkResponse(session, url)

		if responseCode != 200 {
			t.Error("response code should be 200, got", responseCode)
		}

		if !isValidJSON {
			t.Error("response is not valid JSON")
		}
		if err != nil {
			t.Error("An unexpected error occured")
		}

		_, err = session.TracksActive()

		if err != nil {
			t.Error("error creating struct: ", err)
		}
	})
	t.Run("Retrieves the information of a track", func(t *testing.T) {
		trackID := 2
		url := "https://www.hackthebox.com/api/v4/tracks/" + strconv.Itoa(trackID)
		responseCode, isValidJSON, err := checkResponse(session, url)

		if responseCode != 200 {
			t.Error("response code should be 200, got", responseCode)
		}

		if !isValidJSON {
			t.Error("response is not valid JSON")
		}
		if err != nil {
			t.Error("An unexpected error occured")
		}

		_, err = session.Track(trackID)

		if err != nil {
			t.Error("error creating struct: ", err)
		}
	})
}
