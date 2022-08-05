package htbgo

type EndgameCreator struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Thumb string `json:"avatar_thumb"`
}

type EndgameAvailability struct {
	Available bool   `json:"available"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
}

// Endgames List
// https://www.hackthebox.com/api/v4/endgames

type EndgamesActiveList struct {
	Status bool                      `json:"status"`
	Data   []*EndgamesActiveListItem `json:"data"`
}

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

// Endgame Profile
// https://www.hackthebox.com/api/v4/endgame/{endgameID}

type EndgameInfo struct {
	Status bool             `json:"status"`
	Data   *EndgameInfoItem `json:"data"`
}

type EndgameInfoItem struct {
	ID               int                  `json:"id"`
	Name             string               `json:"name"`
	Avatar           string               `json:"avatar_url"`
	Cover            string               `json:"cover_image_url"`
	IsRetired        bool                 `json:"retired"`
	IsVip            bool                 `json:"vip"`
	Creators         []*EndgameCreator    `json:"creators"`
	Points           int                  `json:"points"`
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

// Endgame Flag List
// https://www.hackthebox.com/api/v4/endgame/{endgameID}/flags

type EndgameFlagsList struct {
	Status bool                    `json:"status"`
	Data   []*EndgameFlagsListItem `json:"data"`
}

type EndgameFlagsListItem struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Points int    `json:"points"`
	Owned  bool   `json:"owned"`
}

// Endgame Machine List
// https://www.hackthebox.com/api/v4/endgame/{endgameID}}/machines

type EndgameMachinesList struct {
	Status bool                       `json:"status"`
	Data   []*EndgameMachinesListItem `json:"data"`
}

type EndgameMachinesListItem struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	OS    string `json:"os"`
	Thumb string `json:"avatar_thumb_url"`
}

func (s *Session) EndgamesActive() (endgames *EndgamesActiveList, err error) {

	var url string = "https://www.hackthebox.com/api/v4/endgames"
	err = parseJSON(s, url, &endgames)

	return
}

func (s *Session) Endgame(endgameID int) (endgame *EndgameInfo, err error) {

	endgameIDString, err := toPositiveIntString(endgameID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/endgame/" + endgameIDString
	err = parseJSON(s, url, &endgame)

	return
}

func (s *Session) EndgameFlags(endgameID int) (flags *EndgameFlagsList, err error) {

	endgameIDString, err := toPositiveIntString(endgameID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/endgame/" + endgameIDString + "/flags"
	err = parseJSON(s, url, &flags)

	return
}

func (s *Session) EndgameMachines(endgameID int) (machines *EndgameMachinesList, err error) {

	endgameIDString, err := toPositiveIntString(endgameID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/endgame/" + endgameIDString + "/machines"
	err = parseJSON(s, url, &machines)

	return
}
