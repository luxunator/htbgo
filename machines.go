package htbgo

// Get Machine Matrix
// https://www.hackthebox.com/api/v4/machine/graph/matrix/{machineID}
type MachineMatrixInformation struct {
	Info struct {
		Aggregate struct {
			Enum   float64 `json:"enum"`
			Real   float64 `json:"real"`
			CVE    float64 `json:"cve"`
			Custom float64 `json:"custom"`
			CTF    float64 `json:"ctf"`
		}
		Maker struct {
			Enum   float64 `json:"enum"`
			Real   float64 `json:"real"`
			CVE    float64 `json:"cve"`
			Custom float64 `json:"custom"`
			CTF    float64 `json:"ctf"`
		}
		User struct {
			Enum   float64 `json:"enum"`
			Real   float64 `json:"real"`
			CVE    float64 `json:"cve"`
			Custom float64 `json:"custom"`
			CTF    float64 `json:"ctf"`
		}
	}
}

// Get the Information of the machine ( doesn't include the first blood information)
//  https://www.hackthebox.com/api/v4/machine/info/{machineID}
type MachineInformation struct {
	Info struct {
		ID                  int    `json:"id"`
		Name                string `json:"name"`
		OS                  string `json:"os"`
		Active              int    `json:"active"`
		Retired             int    `json:"retired"`
		IP                  string `json:"ip"`
		Points              int    `json:"points"`
		StaticPoints        int    `json:"static_points"`
		Release             string `json:"release"`
		UserOwnsCount       int    `json:"user_owns_count"`
		RootOwnsCount       int    `json:"root_owns_count"`
		Free                bool   `json:"free"`
		AuthUserInUserOwns  bool   `json:"authUserInUserOwns"`
		AuthUserInRootOwns  bool   `json:"authUserInRootOwns"`
		AuthUserHasReviewed bool   `json:"authUserHasReviewed"`
		Stars               string `json:"stars"`
		Difficulty          int    `json:"difficulty"`
		Avatar              string `json:"avatar"`
		FeedbackChart       struct {
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
		DifficultyText string `json:"difficultyText"`
		IsCompleted    bool   `json:"isCompleted"`
		LastResetTime  string `json:"last_reset_time"`
		PlayInfo       struct {
			IsSpawned         bool   `json:"isSpawned"`
			IsSpawning        bool   `json:"isSpawning"`
			IsActive          bool   `json:"isActive"`
			ActivePlayerCount int    `json:"active_player_count"`
			ExpiresAt         string `json:"expires_at"`
		} `json:"playInfo"`
		MakerPrimary struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Avatar      string `json:"avatar"`
			IsRespected bool   `json:"isRespected"`
		} `json:"maker"`

		MakerSecondary struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Avatar      string `json:"avatar"`
			IsRespected bool   `json:"isRespected"`
		} `json:"maker2"`
		Recommended  int `json:"recommended"`
		ReviewsCount int `json:"reviews_count"`
		SPFlag       int `json:"sp_flag"`
	}
}

// Get the Profile of the machine (includes first blood information)
// https://www.hackthebox.com/api/v4/machine/profile/{machineID}
type MachineProfile struct {
	Info struct {
		ID                  int    `json:"id"`
		Name                string `json:"name"`
		OS                  string `json:"os"`
		Active              int    `json:"active"`
		Retired             int    `json:"retired"`
		IP                  string `json:"ip"`
		Points              int    `json:"points"`
		StaticPoints        int    `json:"static_points"`
		Release             string `json:"release"`
		UserOwnsCount       int    `json:"user_owns_count"`
		RootOwnsCount       int    `json:"root_owns_count"`
		Free                bool   `json:"free"`
		AuthUserInUserOwns  bool   `json:"authUserInUserOwns"`
		AuthUserInRootOwns  bool   `json:"authUserInRootOwns"`
		AuthUserHasReviewed bool   `json:"authUserHasReviewed"`
		Stars               string `json:"stars"`
		Difficulty          int    `json:"difficulty"`
		Avatar              string `json:"avatar"`
		FeedbackChart       struct {
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
		DifficultyText string `json:"difficultyText"`
		IsCompleted    bool   `json:"isCompleted"`
		LastResetTime  string `json:"last_reset_time"`
		PlayInfo       struct {
			IsSpawned         bool   `json:"isSpawned"`
			IsSpawning        bool   `json:"isSpawning"`
			IsActive          bool   `json:"isActive"`
			ActivePlayerCount int    `json:"active_player_count"`
			ExpiresAt         string `json:"expires_at"`
		} `json:"playInfo"`

		MakerPrimary struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Avatar      string `json:"avatar"`
			IsRespected bool   `json:"isRespected"`
		} `json:"maker"`

		MakerSecondary struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Avatar      string `json:"avatar"`
			IsRespected bool   `json:"isRespected"`
		} `json:"maker2"`

		AuthUserFirstUserTime string `json:"authUserFirstUserTime"`
		AuthUserFirstRootTime string `json:"authUserFirstRootTime"`
		UserBlood             struct {
			User struct {
				Name   string `json:"name"`
				ID     int    `json:"id"`
				Avatar string `json:"avatar"`
			} `json:"user"`

			CreatedAt       string `json:"created_at"`
			BloodDifference string `json:"blood_difference"`
		} `json:"userBlood"`
		UserBloodAvatar string `json:"userBloodAvatar"`
		RootBlood       struct {
			User struct {
				Name   string `json:"name"`
				ID     int    `json:"id"`
				Avatar string `json:"avatar"`
			} `json:"user"`

			CreatedAt       string `json:"created_at"`
			BloodDifference string `json:"blood_difference"`
		} `json:"rootBlood"`

		RootBloodAvatar    string `json:"rootBloodAvatar"`
		FirstUserBloodTime string `json:"firstUserBloodTime"`
		FirstRootBloodTime string `json:"firstRootBloodTime"`
		Recommended        int    `json:"recommended"`
	}
}

//	Get Active Machines
// https://www.hackthebox.com/api/v4/machine/list
type ActiveMachines struct {
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
		FeedbackChart       struct {
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
		IsCompleted    bool   `json:"isCompleted"`
		LastResetTime  string `json:"last_reset_time"`
		PlayInfo       struct {
			IsSpawned         bool   `json:"isSpawned"`
			IsSpawning        bool   `json:"isSpawning"`
			IsActive          bool   `json:"isActive"`
			ActivePlayerCount int    `json:"active_player_count"`
			ExpiresAt         string `json:"expires_at"`
		} `json:"playInfo"`

		Free         bool `json:"free"`
		MakerPrimary struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Avatar      string `json:"avatar"`
			IsRespected bool   `json:"isRespected"`
		} `json:"maker"`
		MakerSecondary struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Avatar      string `json:"avatar"`
			IsRespected bool   `json:"isRespected"`
		} `json:"maker2"`
		Recommended int    `json:"recommended"`
		SPFlag      int    `json:"sp_flag"`
		EasyMonth   int    `json:"easy_month"`
		IP          string `json:"ip"`
	}
}

// Get Retired Machines
// https://www.hackthebox.com/api/v4/machine/list/retired
type RetiredMachines struct {
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
		FeedbackChart       struct {
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
		IsCompleted    bool   `json:"isCompleted"`
		LastResetTime  string `json:"last_reset_time"`
		PlayInfo       struct {
			IsSpawned         bool   `json:"isSpawned"`
			IsSpawning        bool   `json:"isSpawning"`
			IsActive          bool   `json:"isActive"`
			ActivePlayerCount int    `json:"active_player_count"`
			ExpiresAt         string `json:"expires_at"`
		} `json:"playInfo"`

		Free         bool `json:"free"`
		MakerPrimary struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Avatar      string `json:"avatar"`
			IsRespected bool   `json:"isRespected"`
		} `json:"maker"`
		MakerSecondary struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Avatar      string `json:"avatar"`
			IsRespected bool   `json:"isRespected"`
		} `json:"maker2"`
		Recommended int    `json:"recommended"`
		SPFlag      int    `json:"sp_flag"`
		EasyMonth   int    `json:"easy_month"`
		IP          string `json:"ip"`
		Tags        []struct {
			ID            int `json:"id"`
			TagCategoryID int `json:"tag_category_id"`
		} `json:"tags"`
	}
}

// Get the top 25 users in a machine
// https://www.hackthebox.com/api/v4/machine/owns/top/{machineID}
type TopUsers struct {
	Info []struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Avatar      string `json:"avatar"`
		RankID      int    `json:"rank_id"`
		RankText    string `json:"rank_text"`
		OwnDate     string `json:"own_date"`
		UserOwnDate string `json:"user_own_date"`
		UserOwnTime string `json:"user_own_time"`
		RootOwnTime string `json:"root_own_time"`
		IsUserBlood bool   `json:"is_user_blood"`
		IsRootBlood bool   `json:"is_root_blood"`
		Position    int    `json:"position"`
	} `json:"info"`
}

// Get the list of scheduled machines to release
// https://www.hackthebox.com/api/v4/machine/unreleased
type ScheduledMachines struct {
	Data []struct {
		ID             int    `json:"id"`
		Name           string `json:"name"`
		OS             string `json:"os"`
		Avatar         string `json:"avatar"`
		Release        string `json:"release"`
		Difficulty     int    `json:"difficulty"`
		DifficultyText string `json:"difficulty_text"`
		FirstCreator   []struct {
			ID     int    `json:"id"`
			Name   string `json:"name"`
			Avatar string `json:"avatar"`
		} `json:"firstCreator"`

		CoCreators []struct {
			ID     int    `json:"id"`
			Name   string `json:"name"`
			Avatar string `json:"avatar"`
		} `json:"coCreators"`

		Retiring struct {
			DifficultyText string `json:"difficulty_text"`
			Avatar         string `json:"avatar"`
			OS             string `json:"os"`
			Name           string `json:"name"`
			ID             int    `json:"id"`
		} `json:"retiring"`
	}
}

// Get the list of machines using the todo endpoint (STUPID API)
// https://www.hackthebox.com/api/v4/machine/list/todo

// type TodoMachines struct {
// 	Info []struct {
// 		ID                  int    `json:"id"`
// 		Name                string `json:"name"`
// 		OS                  string `json:"os"`
// 		Points              int    `json:"points"`
// 		StaticPoints        int    `json:"static_points"`
// 		Release             string `json:"release"`
// 		UserOwnCounts       int    `json:"user_owns_count"`
// 		RootOwnCounts       int    `json:"root_owns_count"`
// 		AuthUserIfUserOwns  bool   `json:"authUserInUserOwns"`
// 		AuthUserIfRootOwns  bool   `json:"authUserInRootOwns"`
// 		IsTodo              bool   `json:"isTodo"`
// 		AuthUserHasReviewed bool   `json:"authUserHasReviewed"`
// 		Stars               string `json:"stars"`
// 		Difficulty          int    `json:"difficulty"`
// 		FeedbackChart       struct {
// 			CakeDifficulty      int `json:"counterCake"`
// 			VeryEasyDifficulty  int `json:"counterVeryEasy"`
// 			EasyDifficulty      int `json:"counterEasy"`
// 			TooEasyDifficulty   int `json:"counterTooEasy"`
// 			MediumDifficulty    int `json:"counterMedium"`
// 			BitHardDifficulty   int `json:"counterBitHard"`
// 			HardDifficulty      int `json:"counterHard"`
// 			TooHardDifficulty   int `json:"counterTooHard"`
// 			ExtraHardDifficulty int `json:"counterExHard"`
// 			BrainFuckDifficulty int `json:"counterBrainFuck"`
// 		} `json:"feedbackForChart"`

// 		Avatar         string `json:"avatar"`
// 		DifficultyText string `json:"difficultyText"`
// 		PlayInfo       struct {
// 			IsSpawned         bool   `json:"isSpawned"`
// 			IsSpawning        bool   `json:"isSpawning"`
// 			IsActive          bool   `json:"isActive"`
// 			ActivePlayerCount int    `json:"active_player_count"`
// 			ExpiresAt         string `json:"expires_at"`
// 		} `json:"playInfo"`

// 		Free         bool `json:"free"`
// 		MakerPrimary struct {
// 			ID                  int    `json:"id"`
// 			Name                string `json:"name"`
// 			Avatar              string `json:"avatar"`
// 			AuthUserSentRespect bool   `json:"isRespected"`
// 		} `json:"maker"`

// 		MakerSecondary struct {
// 			ID                  int    `json:"id"`
// 			Name                string `json:"name"`
// 			Avatar              string `json:"avatar"`
// 			AuthUserSentRespect bool   `json:"isRespected"`
// 		} `json:"maker2"`

// 		Recommended int    `json:"recommended"`
// 		SPFlag      int    `json:"sp_flag"`
// 		EasyMonth   int    `json:"easy_month"`
// 		IP          string `json:"ip"`
// 		Tags        []struct {
// 			ID             int `json:"id"`
// 			TagCategoryIDs int `json:"tag_category_id"`
// 		} `json:"tags"`
// 	} `json:"info"`
// }

// Get a retired machine their tags
// https://www.hackthebox.com/api/v4/machine/tags/{machineID}

type RetiredMachineTags struct {
	Info map[string]struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"info"`
}

// Get the changelog of a machine
// https://www.hackthebox.com/api/v4/machine/changelog/{machineID}

type MachineChangelog struct {
	Info []struct {
		ID          int    `json:"id"`
		UserID      int    `json:"user_id"`
		MachineID   int    `json:"machine_id"`
		Type        string `json:"type"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Released    int    `json:"released"`
		CreatedAt   string `json:"created_at"`
		UpdatedAt   string `json:"updated_at"`
	} `json:"info"`
}

// Get the reviews of a retired machine
// https://www.hackthebox.com/api/v4/machine/reviews/{machineID}

type MachineReviews struct {
	Message []struct {
		ID        int    `json:"id"`
		UserID    int    `json:"user_id"`
		MachineID int    `json:"machine_id"`
		Stars     int    `json:"stars"`
		Message   string `json:"message"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
		Title     string `json:"title"`
		Released  int    `json:"released"`
		User      struct {
			ID     int    `json:"id"`
			Name   string `json:"name"`
			Avatar string `json:"avatar"`
		} `json:"user"`
	} `json:"message"`
}

func (s *Session) MachineMatrix(machineID int) (matrix MachineMatrixInformation, err error) {

	machineIDString, err := toPositiveIntString(machineID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/machine/graph/matrix/" + machineIDString
	err = parseJSON(s, url, &matrix)

	return
}

func (s *Session) MachineInformation(machineID int) (machineInformation MachineInformation, err error) {

	machineIDString, err := toPositiveIntString(machineID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/machine/info/" + machineIDString
	err = parseJSON(s, url, &machineInformation)

	return
}

func (s *Session) MachineProfile(machineID int) (machineProfile MachineProfile, err error) {

	machineIDString, err := toPositiveIntString(machineID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/machine/profile/" + machineIDString
	err = parseJSON(s, url, &machineProfile)

	return
}

func (s *Session) ActiveMachines() (activeMachines ActiveMachines, err error) {

	var url string = "https://www.hackthebox.com/api/v4/machine/list"
	err = parseJSON(s, url, &activeMachines)

	return
}

func (s *Session) RetiredMachines() (retiredMachines RetiredMachines, err error) {

	var url string = "https://www.hackthebox.com/api/v4/machine/list/retired"
	err = parseJSON(s, url, &retiredMachines)

	return
}

func (s *Session) MachineTopUsers(machineID int) (topUsers TopUsers, err error) {

	machineIDString, err := toPositiveIntString(machineID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/machine/owns/top/" + machineIDString
	err = parseJSON(s, url, &topUsers)
	return
}

func (s *Session) ScheduledMachines() (scheduledMachines ScheduledMachines, err error) {

	var url string = "https://www.hackthebox.com/api/v4/machine/unreleased"
	err = parseJSON(s, url, &scheduledMachines)

	return
}

// func (s *Session) TodoMachines() (todoMachines TodoMachines) {
// 	var url string = "https://www.hackthebox.com/api/v4/machine/list/todo"
// 	parseJSON(s, url, &todoMachines)
// 	return
// }

func (s *Session) RetiredMachineTags(machineID int) (tags RetiredMachineTags, err error) {

	machineIDString, err := toPositiveIntString(machineID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/machine/tags/" + machineIDString
	err = parseJSON(s, url, &tags)

	return
}

func (s *Session) MachineChangelog(machineID int) (changelog MachineChangelog, err error) {

	machineIDString, err := toPositiveIntString(machineID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/machine/changelog/" + machineIDString
	err = parseJSON(s, url, &changelog)

	return
}

func (s *Session) MachineReviews(machineID int) (reviews MachineReviews, err error) {

	machineIDString, err := toPositiveIntString(machineID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/machine/reviews/" + machineIDString
	err = parseJSON(s, url, &reviews)

	return
}
