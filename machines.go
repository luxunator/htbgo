package htbgo

type maker struct {
	ID                  int    `json:"id"`
	Name                string `json:"name"`
	Avatar              string `json:"avatar"`
	AuthUserSentRespect bool   `json:"isRespected"`
}
type user struct {
	Name   string `json:"name"`
	ID     int    `json:"id"`
	Avatar string `json:"avatar"`
}

type firstBlood struct {
	User               user   `json:"user"`
	MachineCreatedAt   string `json:"created_at"`
	FirstBloodDuration string `json:"blood_difference"`
}

type machineInfo struct {
	ID                  int    `json:"id"`
	Name                string `json:"name"`
	OS                  string `json:"os"`
	IsTodo              bool   `json:"isTodo"`
	Active              int    `json:"active"`
	Retired             int    `json:"retired"`
	IP                  string `json:"ip"`
	Points              int    `json:"points"`
	StaticPoints        int    `json:"static_points"`
	Release             string `json:"release"`
	UserOwnCounts       int    `json:"user_owns_count"`
	RootOwnCounts       int    `json:"root_owns_count"`
	Free                bool   `json:"free"`
	AuthUserIfUserOwns  bool   `json:"authUserInUserOwns"`
	AuthUserIfRootOwns  bool   `json:"authUserInRootOwns"`
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
	}
}

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

// Get the Profile of the machine (includes first blood information)
// https://www.hackthebox.com/api/v4/machine/info/{machineID}
type MachineProfile struct {
	Info struct {
		machineInfo
		MakerPrimary          maker      `json:"maker"`
		MakerSecondary        maker      `json:"maker2"`
		AuthUserTimeOwnedUser string     `json:"authUserFirstUserTime"`
		AuthUserTimeOwnedRoot string     `json:"authUserFirstRootTime"`
		FirstBloodInUser      firstBlood `json:"userBlood"`
		UserFirstBloodAvatar  string     `json:"userBloodAvatar"`
		FirstBloodInRoot      firstBlood `json:"rootBlood"`
		RootFirstBloodAvatar  string     `json:"rootBloodAvatar"`
		UserFirstBloodTime    string     `json:"firstUserBloodTime"`
		RootFirstBloodTime    string     `json:"firstRootBloodTime"`
		Recommended           int        `json:"recommended"`
	}
}

//	Get Active Machines
// https://www.hackthebox.com/api/v4/machine/list
type ActiveMachines struct {
	Info []machineInfo
}

// Get Retired Machines
// https://www.hackthebox.com/api/v4/machine/list/retired
type RetiredMachines struct {
	Info []machineInfo
}

func (s *Session) MachineMatrix(machineID string) (matrix MachineMatrixInformation) {
	var url string = "https://www.hackthebox.com/api/v4/machine/graph/matrix/" + machineID
	parseJSON(s, url, &matrix)
	return
}

func (s *Session) MachineProfile(machineID string) (machineProfile MachineProfile) {
	var url string = "https://www.hackthebox.com/api/v4/machine/profile/" + machineID
	parseJSON(s, url, &machineProfile)
	return
}
func (s *Session) ActiveMachines() (activeMachines ActiveMachines) {
	var url string = "https://www.hackthebox.com/api/v4/machine/list"
	parseJSON(s, url, &activeMachines)
	return
}
func (s *Session) RetiredMachines() (retiredMachines RetiredMachines) {
	var url string = "https://www.hackthebox.com/api/v4/machine/list/retired"
	parseJSON(s, url, &retiredMachines)
	return
}
