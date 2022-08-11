package htbgo

import (
	"strconv"
	"testing"
)

func TestMachines(t *testing.T) {
	t.Run("Retrieves the matrix of a machine", func(t *testing.T) {
		id := 298
		url := "https://www.hackthebox.com/api/v4/machine/graph/matrix/" + strconv.Itoa(id)
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

		_, err = session.MachineMatrix(id)

		if err != nil {
			t.Error("error creating struct: ", err)
		}

	})
	t.Run("Retrieves the information of a machine", func(t *testing.T) {
		id := 298
		url := "https://www.hackthebox.com/api/v4/machine/graph/matrix/" + strconv.Itoa(id)
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

		_, err = session.Machine(id)

		if err != nil {
			t.Error("error creating struct: ", err)
		}

	})

	t.Run("Retrieves the profile of a machine", func(t *testing.T) {
		id := 298
		url := "https://www.hackthebox.com/api/v4/machine/profile/" + strconv.Itoa(id)
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

		_, err = session.MachineProfile(id)

		if err != nil {
			t.Error("error creating struct: ", err)
		}
	})
	t.Run("Retrieves list of active machines", func(t *testing.T) {

		url := "https://www.hackthebox.com/api/v4/machine/list"
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

		_, err = session.MachinesRetired()

		if err != nil {
			t.Error("error creating struct: ", err)
		}
	})
	t.Run("Retrieves the top users in a machine", func(t *testing.T) {
		machineID := 298
		url := "https://www.hackthebox.com/api/v4/machine/owns/top/" + strconv.Itoa(machineID)
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

		_, err = session.MachineTopUsers(machineID)

		if err != nil {
			t.Error("error creating struct: ", err)
		}
	})
	t.Run("Retrieves list of scheduled machines to release", func(t *testing.T) {
		url := "https://www.hackthebox.com/api/v4/machine/list/todo"
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

		_, err = session.MachinesTodo()

		if err != nil {
			t.Error("error creating struct: ", err)
		}
	})
	// Cause issues because golang automatically sort the hashmap
	t.Run("Retrieves list of tags of a machine", func(t *testing.T) {
		machineID := 298
		url := "https://www.hackthebox.com/api/v4/machine/tags/" + strconv.Itoa(machineID)
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

		_, err = session.MachineTags(machineID)

		if err != nil {
			t.Error("error creating struct: ", err)
		}

	})
	t.Run("Retrieves the changelogs of a machine", func(t *testing.T) {
		machineID := 298
		url := "https://www.hackthebox.com/api/v4/machine/tags/" + strconv.Itoa(machineID)
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

		_, err = session.MachineChangelog(machineID)

		if err != nil {
			t.Error("error creating struct: ", err)
		}
	})

	// Issues with unicode characters in the response of the http request
	t.Run("Retrieves the reviews of a machine", func(t *testing.T) {
		machineID := 298
		url := "https://www.hackthebox.com/api/v4/machine/reviews/" + strconv.Itoa(machineID)
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

		_, err = session.MachineReviews(machineID)

		if err != nil {
			t.Error("error creating struct: ", err)
		}
	})
}
