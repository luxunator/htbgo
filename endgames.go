package htbgo

// EndgameCreator contains information on an endgame creator
type EndgameCreator struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Thumb string `json:"avatar_thumb"`
}

// EndgameAvailability contains information on the availability of an endgame
type EndgameAvailability struct {
	Available bool   `json:"available"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
}

// EndgamesActiveList contains a list of active endgames
type EndgamesActiveList struct {
	Status bool                      `json:"status"`
	Data   []*EndgamesActiveListItem `json:"data"`
}

// EndgamesActiveListItem contains information on an endgame within an active endgame list
type EndgamesActiveListItem struct {
	ID            int                  `json:"id"`
	Name          string               `json:"name"`
	Avatar        string               `json:"avatar_url"`
	Cover         string               `json:"cover_image_url"`
	IsRetired     bool                 `json:"retired"`
	IsVip         bool                 `json:"vip"`
	Creators      []*EndgameCreator    `json:"creators"`
	MachinesCount int                  `json:"endgame_machines_count"`
	FlagsCount    int                  `json:"endgame_flags_count"`
	Availability  *EndgameAvailability `json:"user_availability"`
	IsNew         bool                 `json:"new"`
}

// EndgameInfo contains an endgame
type EndgameInfo struct {
	Status bool             `json:"status"`
	Data   *EndgameInfoItem `json:"data"`
}

// EndgameInfoItem contains information on an endgame
type EndgameInfoItem struct {
	ID               int                  `json:"id"`
	Name             string               `json:"name"`
	Avatar           string               `json:"avatar_url"`
	Cover            string               `json:"cover_image_url"`
	IsRetired        bool                 `json:"retired"`
	IsVip            bool                 `json:"vip"`
	Creators         []*EndgameCreator    `json:"creators"`
	Points           interface{}          `json:"points"` // string or int 0
	PlayersCompleted int                  `json:"players_completed"`
	ResetVotes       int                  `json:"endgame_reset_votes"`
	LastReset        string               `json:"most_recent_reset"`
	EntryPoints      []string             `json:"entry_points"`
	Video            string               `json:"video_url"`
	Description      string               `json:"description"`
	CompletionIcon   string               `json:"completion_icon"`
	CompletionText   string               `json:"completion_text"`
	HasFinished      bool                 `json:"has_user_finished"`
	Availability     *EndgameAvailability `json:"user_availability"`
}

// EndgameFlagsList contains a list of endgame flags
type EndgameFlagsList struct {
	Status bool                    `json:"status"`
	Data   []*EndgameFlagsListItem `json:"data"`
}

// EndgameFlagsListItem contains information on an endgame flag
type EndgameFlagsListItem struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Points int    `json:"points"`
	Owned  bool   `json:"owned"`
}

// EndgameMachinesList contains a list of machines in an endame
type EndgameMachinesList struct {
	Status bool                       `json:"status"`
	Data   []*EndgameMachinesListItem `json:"data"`
}

// EndgameMachinesListItem contains information on a machine within an endgame
type EndgameMachinesListItem struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	OS    string `json:"os"`
	Thumb string `json:"avatar_thumb_url"`
}

// EndgamesActive returns a list of the currently active endgames
func (s *Session) EndgamesActive() (endgames *EndgamesActiveList, err error) {

	var url string = "https://www.hackthebox.com/api/v4/endgames"
	err = parseJSON(s, url, &endgames)

	return
}

// Endgame returns the information of an endgame, given an endgame ID
func (s *Session) Endgame(endgameID int) (endgame *EndgameInfo, err error) {

	endgameIDString, err := toPositiveIntString(endgameID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/endgame/" + endgameIDString
	err = parseJSON(s, url, &endgame)

	return
}

// EndgameFlags returns a list of information on the flags within an endgame, given an endgame ID
func (s *Session) EndgameFlags(endgameID int) (flags *EndgameFlagsList, err error) {

	endgameIDString, err := toPositiveIntString(endgameID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/endgame/" + endgameIDString + "/flags"
	err = parseJSON(s, url, &flags)

	return
}

// EndgameMachines returns a list of the machines that an endgame contains, given an endgame ID
func (s *Session) EndgameMachines(endgameID int) (machines *EndgameMachinesList, err error) {

	endgameIDString, err := toPositiveIntString(endgameID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/endgame/" + endgameIDString + "/machines"
	err = parseJSON(s, url, &machines)

	return
}
