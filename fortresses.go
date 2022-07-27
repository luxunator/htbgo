package htbgo

// Active Fortesses List
// https://www.hackthebox.com/api/v4/fortresses

type ActiveFortressesList struct {
	Status bool `json:"status"`
	Data   map[string]struct {
		ID            int    `json:"id"`
		Name          string `json:"name"`
		Image         string `json:"image"`
		CoverImageUrl string `json:"cover_image_url"`
		New           bool   `json:"new"`
		NumberOfFlags int    `json:"number_of_flags"`
	} `json:"data"`
}

// Fortress Profile
// https://www.hackthebox.com/api/v4/fortress/{fortressID}

type FortessProfile struct {
	Status bool `json:"enum"`
	Data   struct {
		ID            int    `json:"id"`
		Name          string `json:"name"`
		IP            string `json:"ip"`
		Image         string `json:"image"`
		CoverImageUrl string `json:"cover_image_url"`
		Company       struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
			URL         string `json:"url"`
			Image       string `json:"image"`
		} `json:"company"`
		ResetVotes           int    `json:"reset_votes"`
		Description          string `json:"description"`
		HasCompletionMessage bool   `json:"has_completion_message"`
		CompletionMessage    string `json:"completion_message"`
		ProgressPercent      int    `json:"progress_percent"`
		PlayersCompleted     int    `json:"players_completed"`
		Points               string `json:"points"`
		UserAvailability     struct {
			Available bool   `json:"available"`
			Code      int    `json:"code"`
			Message   string `json:"message"`
		} `json:"user_availability"`
		Flags []struct {
			ID     int    `json:"id"`
			Title  string `json:"title"`
			Points int    `json:"points"`
			Owned  bool   `json:"owned"`
		} `json:"flags"`
	} `json:"data"`
}

// Fortress Flag List
// https://www.hackthebox.com/api/v4/fortress/{fortressID}/flags

type FortressFlagList struct {
	Status bool `json:"status"`
	Data   []struct {
		ID     int    `json:"id"`
		Title  string `json:"title"`
		Points int    `json:"points"`
		Owned  bool   `json:"owned"`
	} `json:"data"`
}

func (s *Session) ActiveFortresses() (fortresses ActiveFortressesList) {

	var url string = "https://www.hackthebox.com/api/v4/fortresses"
	parseJSON(s, url, &fortresses)

	return
}

func (s *Session) FortressInfo(fortressID int) (info FortessProfile, err error) {

	fortressIDString, err := toPositiveIntString(fortressID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/fortress/" + fortressIDString
	parseJSON(s, url, &info)

	return
}

func (s *Session) FortressFlags(fortressID int) (flags FortressFlagList, err error) {

	fortressIDString, err := toPositiveIntString(fortressID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/fortress/" + fortressIDString + "/flags"
	parseJSON(s, url, &flags)

	return
}
