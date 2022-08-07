package htbgo

// FortessFlag
type FortessFlag struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Points int    `json:"points"`
	Owned  bool   `json:"owned"`
}

// Active Fortesses List
// https://www.hackthebox.com/api/v4/fortresses

// FortressesActiveMap
type FortressesActiveMap struct {
	Status bool                                `json:"status"`
	Data   map[string]*FortressesActiveMapItem `json:"data"`
}

// FortressesActiveMapItem
type FortressesActiveMapItem struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Logo       string `json:"image"`
	CoverSmall string `json:"cover_image_url"`
	IsNew      bool   `json:"new"`
	Flags      int    `json:"number_of_flags"`
}

// Fortress Profile
// https://www.hackthebox.com/api/v4/fortress/{fortressID}

// FortessInfo
type FortessInfo struct {
	Status bool             `json:"enum"`
	Data   *FortessInfoItem `json:"data"`
}

// FortessInfoItem
type FortessInfoItem struct {
	ID                   int                          `json:"id"`
	Name                 string                       `json:"name"`
	IP                   string                       `json:"ip"`
	Logo                 string                       `json:"image"`
	CoverFull            string                       `json:"cover_image_url"`
	Company              *FortessInfoItemCompany      `json:"company"`
	ResetVotes           int                          `json:"reset_votes"`
	Description          string                       `json:"description"`
	HasCompletionMessage bool                         `json:"has_completion_message"`
	CompletionMessage    string                       `json:"completion_message"`
	Progress             int                          `json:"progress_percent"`
	PlayersCompleted     int                          `json:"players_completed"`
	Points               string                       `json:"points"`
	Availability         *FortessInfoItemAvailability `json:"user_availability"`
	Flags                []*FortessFlag               `json:"flags"`
}

// FortessInfoItemCompany
type FortessInfoItemCompany struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Image       string `json:"image"`
}

// FortessInfoItemAvailability
type FortessInfoItemAvailability struct {
	Available bool   `json:"available"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
}

// Fortress Flag List
// https://www.hackthebox.com/api/v4/fortress/{fortressID}/flags

// FortressFlagList
type FortressFlagList struct {
	Status bool           `json:"status"`
	Data   []*FortessFlag `json:"data"`
}

// FortressesActive returns a list of the currently active fortresses
func (s *Session) FortressesActive() (fortresses *FortressesActiveMap, err error) {

	var url string = "https://www.hackthebox.com/api/v4/fortresses"
	err = parseJSON(s, url, &fortresses)

	return
}

// Fortress returns the information of a fortress, given a fortress ID
func (s *Session) Fortress(fortressID int) (info *FortessInfo, err error) {

	fortressIDString, err := toPositiveIntString(fortressID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/fortress/" + fortressIDString
	err = parseJSON(s, url, &info)

	return
}

// FortressFlags returns a list of information on the flags within a fortress, given a fortress ID
func (s *Session) FortressFlags(fortressID int) (flags *FortressFlagList, err error) {

	fortressIDString, err := toPositiveIntString(fortressID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/fortress/" + fortressIDString + "/flags"
	err = parseJSON(s, url, &flags)

	return
}
