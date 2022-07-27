package htbgo

// Starting Point Machines List
// https://www.hackthebox.com/api/v4/sp/machines

type StartingPointMachinesList struct {
	Info []struct {
		ID                  int    `json:"id"`
		Name                string `json:"name"`
		OS                  string `json:"os"`
		Points              int    `json:"points"`
		StaticPoints        int    `json:"static_points"`
		Release             string `json:"release"`
		UserOwnsCount       int    `json:"user_owns_count"`
		RootOwnsCount       int    `json:"root_owns_count"`
		AuthUserInUserOwns  bool   `json:"authUserInUserOwns"`
		AuthUserInRootOwns  bool   `json:"authUserInRootOwns"`
		IsTodo              bool   `json:"isTodo"`
		AuthUserHasReviewed bool   `json:"authUserHasReviewed"`
		Stars               string `json:"stars"`
		Difficulty          int    `json:"difficulty"`
		FeedbackForChart    struct {
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
		} `json:"feedbackForChart"`
		Avatar         string `json:"avatar"`
		DifficultyText string `json:"difficultyText"`
		PlayInfo       struct {
			IsSpawned         bool   `json:"isSpawned"`
			IsSpawning        bool   `json:"isSpawning"`
			IsActive          bool   `json:"isActive"`
			ActivePlayerCount int    `json:"active_player_count"`
			ExpiresAt         string `json:"expires_at"`
		} `json:"playInfo"`
		Free  bool `json:"free"`
		Maker struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Avatar      string `json:"avatar"`
			IsRespected bool   `json:"isRespected"`
		} `json:"maker"`
		MakerTwo struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Avatar      string `json:"avatar"`
			IsRespected bool   `json:"isRespected"`
		} `json:"maker2"`
		Recommended  int  `json:"recommended"`
		SPFlag       int  `json:"sp_flag"`
		EasyMonth    int  `json:"easy_month"`
		RootFlagOnly bool `json:"root_flag_only"`
	}
}

func (s *Session) StartingPointMachines() (machines StartingPointMachinesList, err error) {

	var url string = "https://www.hackthebox.com/api/v4/sp/machines"
	err = parseJSON(s, url, &machines)

	return
}
