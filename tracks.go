package htbgo

// List Available Tracks
// https://www.hackthebox.com/api/v4/tracks

type TracksList []struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Creator struct {
		Type   string `json:"type"`
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Avatar string `json:"avatar"`
	} `json:"creator"`
	Official   bool   `json:"official"`
	StaffPick  int    `json:"staff_pick"`
	Difficulty string `json:"difficulty"`
	CoverImage string `json:"cover_image"`
	Likes      int    `json:"likes"`
}

// Track Profile
// https://www.hackthebox.com/api/v4/tracks/{trackID}

type TrackProfile struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Difficulty  string `json:"difficulty"`
	Creator     struct {
		Type   string `json:"type"`
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Avatar string `json:"avatar"`
	} `json:"creator"`
	Official  bool `json:"official"`
	StaffPick int  `json:"staff_pick"`
	Items     []struct {
		ID                int    `json:"id"`
		Type              string `json:"type"`
		Name              string `json:"name"`
		Difficulty        string `json:"difficulty"`
		DifficultyRatings struct {
			CakeDifficulty      int `json:"counterCake"`
			VeryEasyDifficulty  int `json:"counterVeryEasy"`
			EasyDifficulty      int `json:"counterEasy"`
			TooEasyDifficulty   int `json:"counterTooEasy"`
			MediumDifficulty    int `json:"counterMedium"`
			BitHardDifficulty   int `json:"counterBitHard"`
			HardDifficulty      int `json:"counterHard"`
			TooHardDifficulty   int `json:"counterTooHard"`
			ExtraHardDifficulty int `json:"counterExHard"`
			BrainFuckDifficulty int `json:"counterBrainFuck"`
		} `json:"difficulty_ratings"`
		Avatar   string `json:"avatar"`
		OS       string `json:"os"`
		Category string `json:"category"`
		Complete bool   `json:"complete"`
	}
	CoverImage           string `json:"cover_image"`
	Likes                int    `json:"likes"`
	Liked                bool   `json:"liked"`
	Enrolled             bool   `json:"enrolled"`
	HasCompletionMessage bool   `json:"has_completion_message"`
	CompletionURL        string `json:"completion_url"`
	CompletionMessage    string `json:"completion_message"`
	// CompletionCTA `json:"completion_cta"` TODO
}

func (s *Session) ActiveTracks() (tracks TracksList) {

	var url string = "https://www.hackthebox.com/api/v4/tracks"
	parseJSON(s, url, &tracks)

	return
}

func (s *Session) TrackInfo(trackID string) (info TrackProfile) {

	var url string = "https://www.hackthebox.com/api/v4/tracks/" + trackID
	parseJSON(s, url, &info)

	return
}
