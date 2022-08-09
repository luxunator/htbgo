package htbgo

// TrackCreator contains information on a track creator
type TrackCreator struct {
	Type   string `json:"type"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

// TracksActiveList contains a list of available tracks and its information
type TracksActiveList []struct {
	ID         int           `json:"id"`
	Name       string        `json:"name"`
	Creator    *TrackCreator `json:"creator"`
	IsOfficial bool          `json:"official"`
	StaffPick  int           `json:"staff_pick"`
	Difficulty string        `json:"difficulty"`
	Cover      string        `json:"cover_image"`
	Likes      int           `json:"likes"`
}

// TrackInfo contains information on track
type TrackInfo struct {
	ID                   int              `json:"id"`
	Name                 string           `json:"name"`
	Description          string           `json:"description"`
	Difficulty           string           `json:"difficulty"`
	Creator              *TrackCreator    `json:"creator"`
	IsOfficial           bool             `json:"official"`
	StaffPick            int              `json:"staff_pick"`
	Items                []*TrackInfoItem `json:"items"`
	Cover                string           `json:"cover_image"`
	Likes                int              `json:"likes"`
	HasLiked             bool             `json:"liked"`
	IsEnrolled           bool             `json:"enrolled"`
	HasCompletionMessage bool             `json:"has_completion_message"`
	CompletionURL        interface{}      `json:"completion_url"`     // null or string "test"
	CompletionMessage    interface{}      `json:"completion_message"` // null or string "test"
	CompletionCTA        interface{}      `json:"completion_cta"`
}

// TrackInfoItem contains information on items in a track
type TrackInfoItem struct {
	ID              int           `json:"id"`
	Type            string        `json:"type"`
	Name            string        `json:"name"`
	Difficulty      string        `json:"difficulty"`
	DifficultyStats *Difficulties `json:"difficulty_ratings"`
	Avatar          interface{}   `json:"avatar"`   // null or string "test"
	OS              interface{}   `json:"os"`       // null or string "test"
	Category        interface{}   `json:"category"` // null or string "test"
	IsComplete      bool          `json:"complete"`
}

// TracksActive returns a list of the currently active tracks
func (s *Session) TracksActive() (tracks *TracksActiveList, err error) {

	var url string = "https://www.hackthebox.com/api/v4/tracks"
	err = parseJSON(s, url, &tracks)

	return
}

// Track returns the information of a track, given an track ID
func (s *Session) Track(trackID int) (info *TrackInfo, err error) {

	trackIDString, err := toPositiveIntString(trackID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/tracks/" + trackIDString
	err = parseJSON(s, url, &info)

	return
}
