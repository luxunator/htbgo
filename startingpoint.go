package htbgo

// StartingPointCreator contains information about a starting points creator
type StartingPointCreator struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Avatar      string `json:"avatar"`
	IsRespected bool   `json:"isRespected"`
}

// StartingPointMachinesList contains a list of starting point machines
type StartingPointMachinesList struct {
	Info []*StartingPointMachinesListItem `json:"info"`
}

// StartingPointMachinesListItem contains information about a starting point machine
type StartingPointMachinesListItem struct {
	ID              int                                    `json:"id"`
	Name            string                                 `json:"name"`
	OS              string                                 `json:"os"`
	Points          int                                    `json:"points"`
	StaticPoints    int                                    `json:"static_points"`
	Release         string                                 `json:"release"`
	UserOwns        int                                    `json:"user_owns_count"`
	RootOwns        int                                    `json:"root_owns_count"`
	InUserOwns      bool                                   `json:"authUserInUserOwns"`
	InRootOwns      bool                                   `json:"authUserInRootOwns"`
	IsTodo          bool                                   `json:"isTodo"`
	HasReviewed     bool                                   `json:"authUserHasReviewed"`
	Stars           float64                                `json:"stars"`
	DifficultyAvg   int                                    `json:"difficulty"`
	DifficultyStats *Difficulties                          `json:"feedbackForChart"`
	Avatar          string                                 `json:"avatar"`
	Difficulty      string                                 `json:"difficultyText"`
	PlayInfo        *StartingPointMachinesListItemPlayInfo `json:"playInfo"`
	IsFree          bool                                   `json:"free"`
	CreatorOne      *StartingPointCreator                  `json:"maker"`
	CreatorTwo      *StartingPointCreator                  `json:"maker2"`
	Recommended     int                                    `json:"recommended"`
	SPFlag          int                                    `json:"sp_flag"`
	EasyMonth       int                                    `json:"easy_month"`
	HasRootFlagOnly bool                                   `json:"root_flag_only"`
}

// StartingPointMachinesListItemPlayInfo contains information about players on a starting point machine
type StartingPointMachinesListItemPlayInfo struct {
	IsSpawned   bool   `json:"isSpawned"`
	IsSpawning  bool   `json:"isSpawning"`
	IsActive    bool   `json:"isActive"`
	PlayerCount int    `json:"active_player_count"`
	ExpiresAt   string `json:"expires_at"`
}

// StartingPointMachines returns a list of the current starting point machines
func (s *Session) StartingPointMachines() (machines *StartingPointMachinesList, err error) {

	var url string = "https://www.hackthebox.com/api/v4/sp/machines"
	err = parseJSON(s, url, &machines)

	return
}
