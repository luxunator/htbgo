package htbgo

type FortessFlags []struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Points int    `json:"points"`
	Owned  bool   `json:"owned"`
}

// Active Fortesses List
// https://www.hackthebox.com/api/v4/fortresses

type FortressesActiveMap struct {
	Status bool                     `json:"status"`
	Data   *FortressesActiveMapData `json:"data"`
}

type FortressesActiveMapData map[string]struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Logo       string `json:"image"`
	CoverSmall string `json:"cover_image_url"`
	IsNew      bool   `json:"new"`
	Flags      int    `json:"number_of_flags"`
}

// Fortress Profile
// https://www.hackthebox.com/api/v4/fortress/{fortressID}

type FortessInfo struct {
	Status bool             `json:"enum"`
	Data   *FortessInfoData `json:"data"`
}

type FortessInfoData struct {
	ID                   int                      `json:"id"`
	Name                 string                   `json:"name"`
	IP                   string                   `json:"ip"`
	Logo                 string                   `json:"image"`
	CoverFull            string                   `json:"cover_image_url"`
	Company              *FortessInfoCompany      `json:"company"`
	ResetVotes           int                      `json:"reset_votes"`
	Description          string                   `json:"description"`
	HasCompletionMessage bool                     `json:"has_completion_message"`
	CompletionMessage    string                   `json:"completion_message"`
	Progress             int                      `json:"progress_percent"`
	PlayersCompleted     int                      `json:"players_completed"`
	Points               string                   `json:"points"`
	Availability         *FortessInfoAvailability `json:"user_availability"`
	Flags                *FortessFlags            `json:"flags"`
}

type FortessInfoCompany struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Image       string `json:"image"`
}

type FortessInfoAvailability struct {
	Available bool   `json:"available"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
}

// Fortress Flag List
// https://www.hackthebox.com/api/v4/fortress/{fortressID}/flags

type FortressFlagList struct {
	Status bool          `json:"status"`
	Data   *FortessFlags `json:"data"`
}

func (s *Session) FortressesActive() (fortresses *FortressesActiveMap, err error) {

	var url string = "https://www.hackthebox.com/api/v4/fortresses"
	err = parseJSON(s, url, &fortresses)

	return
}

func (s *Session) Fortress(fortressID int) (info *FortessInfo, err error) {

	fortressIDString, err := toPositiveIntString(fortressID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/fortress/" + fortressIDString
	err = parseJSON(s, url, &info)

	return
}

func (s *Session) FortressFlags(fortressID int) (flags *FortressFlagList, err error) {

	fortressIDString, err := toPositiveIntString(fortressID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/fortress/" + fortressIDString + "/flags"
	err = parseJSON(s, url, &flags)

	return
}
