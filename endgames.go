package htbgo

// Endgames List
// https://www.hackthebox.com/api/v4/endgames

type EndgamesList struct {
	Status bool `json:"status"`
	Data   []struct {
		ID            int    `json:"id"`
		Name          string `json:"name"`
		AvatarUrl     string `json:"avatar_url"`
		CoverImageUrl string `json:"cover_image_url"`
		Retired       bool   `json:"retired"`
		Vip           bool   `json:"vip"`
		Creators      []struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			AvatarThumb string `json:"avatar_thumb"`
		} `json:"creators"`
		EndgameMachinesCount int `json:"endgame_machines_count"`
		EndgameFlagsCount    int `json:"endgame_flags_count"`
		UserAvailability     struct {
			Available bool   `json:"available"`
			Code      int    `json:"code"`
			Message   string `json:"message"`
		} `json:"user_availability"`
		New bool `json:"new"`
	} `json:"data"`
}

// Endgame Profile
// https://www.hackthebox.com/api/v4/endgame/{endgameID}

type EndgameProfile struct {
	Status bool `json:"status"`
	Data   struct {
		ID            int    `json:"id"`
		Name          string `json:"name"`
		AvatarUrl     string `json:"avatar_url"`
		CoverImageUrl string `json:"cover_image_url"`
		Retired       bool   `json:"retired"`
		Vip           bool   `json:"vip"`
		Creators      []struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			AvatarThumb string `json:"avatar_thumb"`
		} `json:"creators"`
		Points            int      `json:"points"`
		PlayersCompleted  int      `json:"players_completed"`
		EndgameResetVotes int      `json:"endgame_reset_votes"`
		MostRecentReset   string   `json:"most_recent_reset"`
		EntryPoints       []string `json:"entry_points"`
		VideoURL          string   `json:"video_url"`
		Description       string   `json:"description"`
		CompletionIcon    string   `json:"completion_icon"`
		CompletionText    string   `json:"completion_text"`
		HasUserFinished   bool     `json:"has_user_finished"`
		UserAvailability  struct {
			Available bool   `json:"available"`
			Code      int    `json:"code"`
			Message   string `json:"message"`
		} `json:"user_availability"`
	}
}

// Endgame Flag List
// https://www.hackthebox.com/api/v4/endgame/{endgameID}/flags

type EndgameFlagsList struct {
	Status bool `json:"status"`
	Data   []struct {
		ID     int    `json:"id"`
		Title  string `json:"title"`
		Points int    `json:"points"`
		Owned  bool   `json:"owned"`
	} `json:"data"`
}

// Endgame Machine List
// https://www.hackthebox.com/api/v4/endgame/{endgameID}}/machines

type EndgameMachinesList struct {
	Status bool `json:"status"`
	Data   []struct {
		ID             int    `json:"id"`
		Name           string `json:"name"`
		OS             string `json:"os"`
		AvatarThumbURL string `json:"avatar_thumb_url"`
	} `json:"data"`
}

func (s *Session) ActiveEndgames() (endgames EndgamesList) {

	var url string = "https://www.hackthebox.com/api/v4/endgames"
	parseJSON(s, url, &endgames)

	return
}

func (s *Session) EndgameInfo(endgameID int) (endgame EndgameProfile, err error) {

	endgameIDString, err := toPositiveIntString(endgameID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/endgame/" + endgameIDString
	parseJSON(s, url, &endgame)

	return
}

func (s *Session) EndgameFlags(endgameID int) (flags EndgameFlagsList, err error) {

	endgameIDString, err := toPositiveIntString(endgameID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/endgame/" + endgameIDString + "/flags"
	parseJSON(s, url, &flags)

	return
}

func (s *Session) EndgameMachines(endgameID int) (machines EndgameMachinesList, err error) {

	endgameIDString, err := toPositiveIntString(endgameID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/endgame/" + endgameIDString + "/machines"
	parseJSON(s, url, &machines)

	return
}
