package htbgo

type MachineMatrixGroup struct {
	Enum   float64 `json:"enum"`
	Real   float64 `json:"real"`
	CVE    float64 `json:"cve"`
	Custom float64 `json:"custom"`
	CTF    float64 `json:"ctf"`
}

type MachineCreator struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Avatar      string `json:"avatar"`
	IsRespected bool   `json:"isRespected"`
}

type MachineStatus struct {
	IsSpawned   bool   `json:"isSpawned"`
	IsSpawning  bool   `json:"isSpawning"`
	IsActive    bool   `json:"isActive"`
	PlayerCount int    `json:"active_player_count"`
	ExpiresAt   string `json:"expires_at"`
}

type MachineBlood struct {
	User      *MachineBloodUser `json:"user"`
	CreatedAt string            `json:"created_at"`
	BloodDiff string            `json:"blood_difference"`
}

type MachineBloodUser struct {
	Name   string `json:"name"`
	ID     int    `json:"id"`
	Avatar string `json:"avatar"`
}

type MachineTag struct {
	ID         int `json:"id"`
	CategoryID int `json:"tag_category_id"`
}

// Get Machine Matrix
// https://www.hackthebox.com/api/v4/machine/graph/matrix/{machineID}
type MachineMatrixInfo struct {
	Info *MachineMatrixInfoItem `json:"info"`
}

type MachineMatrixInfoItem struct {
	Aggregate *MachineMatrixGroup `json:"aggregate"`
	Creator   *MachineMatrixGroup `json:"maker"`
	User      *MachineMatrixGroup `json:"user"`
}

// Get the Information of the machine ( doesn't include the first blood information)
//  https://www.hackthebox.com/api/v4/machine/info/{machineID}
type MachineInfo struct {
	Info *MachineInfoItem `json:"info"`
}

type MachineInfoItem struct {
	ID              int             `json:"id"`
	Name            string          `json:"name"`
	OS              string          `json:"os"`
	Active          int             `json:"active"`
	Retired         int             `json:"retired"`
	IP              string          `json:"ip"`
	Points          int             `json:"points"`
	StaticPoints    int             `json:"static_points"`
	Release         string          `json:"release"`
	UserOwns        int             `json:"user_owns_count"`
	RootOwns        int             `json:"root_owns_count"`
	IsFree          bool            `json:"free"`
	HasOwnedUser    bool            `json:"authUserInUserOwns"`
	HasOwnedRoot    bool            `json:"authUserInRootOwns"`
	HasReviewed     bool            `json:"authUserHasReviewed"`
	Stars           string          `json:"stars"`
	DifficultyAvg   int             `json:"difficulty"`
	Avatar          string          `json:"avatar"`
	DifficultyStats *Difficulties   `json:"feedbackForChart"`
	Difficulty      string          `json:"difficultyText"`
	IsCompleted     bool            `json:"isCompleted"`
	LastResetTime   string          `json:"last_reset_time"`
	PlayInfo        *MachineStatus  `json:"playInfo"`
	CreatorOne      *MachineCreator `json:"maker"`
	CreatorTwo      *MachineCreator `json:"maker2"`
	Recommended     int             `json:"recommended"`
	ReviewsCount    int             `json:"reviews_count"`
	SPFlag          int             `json:"sp_flag"`
}

// Get the Profile of the machine (includes first blood information)
// https://www.hackthebox.com/api/v4/machine/profile/{machineID}
type MachineProfileInfo struct {
	Info *MachineProfileInfoItem `json:"info"`
}

type MachineProfileInfoItem struct {
	ID              int             `json:"id"`
	Name            string          `json:"name"`
	OS              string          `json:"os"`
	Active          int             `json:"active"`
	Retired         int             `json:"retired"`
	IP              string          `json:"ip"`
	Points          int             `json:"points"`
	StaticPoints    int             `json:"static_points"`
	Release         string          `json:"release"`
	UserOwns        int             `json:"user_owns_count"`
	RootOwns        int             `json:"root_owns_count"`
	IsFree          bool            `json:"free"`
	HasOwnedUser    bool            `json:"authUserInUserOwns"`
	HasOwnedRoot    bool            `json:"authUserInRootOwns"`
	HasReviewed     bool            `json:"authUserHasReviewed"`
	Stars           string          `json:"stars"`
	DifficultyAvg   int             `json:"difficulty"`
	Avatar          string          `json:"avatar"`
	DifficultyStats *Difficulties   `json:"feedbackForChart"`
	Difficulty      string          `json:"difficultyText"`
	IsCompleted     bool            `json:"isCompleted"`
	LastResetTime   string          `json:"last_reset_time"`
	PlayInfo        *MachineStatus  `json:"playInfo"`
	CreatorOne      *MachineCreator `json:"maker"`
	CreatorTwo      *MachineCreator `json:"maker2"`
	UserTime        string          `json:"authUserFirstUserTime"`
	RootTime        string          `json:"authUserFirstRootTime"`
	UserBlood       *MachineBlood   `json:"userBlood"`
	UserBloodAvatar string          `json:"userBloodAvatar"`
	RootBlood       *MachineBlood   `json:"rootBlood"`
	RootBloodAvatar string          `json:"rootBloodAvatar"`
	UserBloodTime   string          `json:"firstUserBloodTime"`
	RootBloodTime   string          `json:"firstRootBloodTime"`
	Recommended     int             `json:"recommended"`
}

//	Get Active Machines
// https://www.hackthebox.com/api/v4/machine/list
type MachinesActiveList struct {
	Info []*MachinesActiveListItem `json:"info"`
}

type MachinesActiveListItem struct {
	ID              int             `json:"id"`
	Name            string          `json:"name"`
	OS              string          `json:"os"`
	Points          int             `json:"points"`
	StaticPoints    int             `json:"static_points"`
	Release         string          `json:"release"`
	UserOwns        int             `json:"user_owns_count"`
	RootOwns        int             `json:"root_owns_count"`
	HasOwnedUser    bool            `json:"authUserInUserOwns"`
	HasOwnedRoot    bool            `json:"authUserInRootOwns"`
	IsTodo          bool            `json:"isTodo"`
	HasReviewed     bool            `json:"authUserHasReviewed"`
	Stars           string          `json:"stars"`
	DifficultyAvg   int             `json:"difficulty"`
	DifficultyStats *Difficulties   `json:"feedbackForChart"`
	Avatar          string          `json:"avatar"`
	Difficulty      string          `json:"difficultyText"`
	IsCompleted     bool            `json:"isCompleted"`
	LastResetTime   string          `json:"last_reset_time"`
	PlayInfo        *MachineStatus  `json:"playInfo"`
	IsFree          bool            `json:"free"`
	CreatorOne      *MachineCreator `json:"maker"`
	CreatorTwo      *MachineCreator `json:"maker2"`
	Recommended     int             `json:"recommended"`
	SPFlag          int             `json:"sp_flag"`
	EasyMonth       int             `json:"easy_month"`
	IP              string          `json:"ip"`
}

// Get Retired Machines
// https://www.hackthebox.com/api/v4/machine/list/retired
type MachinesRetiredList struct {
	Info []*MachinesRetiredListItem `json:"info"`
}

type MachinesRetiredListItem struct {
	ID              int             `json:"id"`
	Name            string          `json:"name"`
	OS              string          `json:"os"`
	Points          int             `json:"points"`
	StaticPoints    int             `json:"static_points"`
	Release         string          `json:"release"`
	UserOwns        int             `json:"user_owns_count"`
	RootOwns        int             `json:"root_owns_count"`
	HasOwnedUser    bool            `json:"authUserInUserOwns"`
	HasOwnedRoot    bool            `json:"authUserInRootOwns"`
	IsTodo          bool            `json:"isTodo"`
	HasReviewed     bool            `json:"authUserHasReviewed"`
	Stars           string          `json:"stars"`
	DifficultyAvg   int             `json:"difficulty"`
	DifficultyStats *Difficulties   `json:"feedbackForChart"`
	Avatar          string          `json:"avatar"`
	Difficulty      string          `json:"difficultyText"`
	IsCompleted     bool            `json:"isCompleted"`
	LastResetTime   string          `json:"last_reset_time"`
	PlayInfo        *MachineStatus  `json:"playInfo"`
	IsFree          bool            `json:"free"`
	CreatorOne      *MachineCreator `json:"maker"`
	CreatorTwo      *MachineCreator `json:"maker2"`
	Recommended     int             `json:"recommended"`
	SPFlag          int             `json:"sp_flag"`
	EasyMonth       int             `json:"easy_month"`
	IP              string          `json:"ip"`
	Tags            []*MachineTag   `json:"tags"`
}

// Get the top 25 users in a machine
// https://www.hackthebox.com/api/v4/machine/owns/top/{machineID}
type MachineTopsUsersList struct {
	Info []*MachineTopsUsersListItem `json:"info"`
}

type MachineTopsUsersListItem struct {
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
}

// Get the list of scheduled machines to release
// https://www.hackthebox.com/api/v4/machine/unreleased
type MachinesToReleaseList struct {
	Data []*MachinesToReleaseListItem `json:"data"`
}

type MachinesToReleaseListItem struct {
	ID            int                                 `json:"id"`
	Name          string                              `json:"name"`
	OS            string                              `json:"os"`
	Avatar        string                              `json:"avatar"`
	Release       string                              `json:"release"`
	DifficultyAvg int                                 `json:"difficulty"`
	Difficulty    string                              `json:"difficulty_text"`
	CreatorOne    []*MachinesToReleaseListItemCreator `json:"firstCreator"`
	CreatorTwo    []*MachinesToReleaseListItemCreator `json:"coCreators"`
	Retiring      *MachinesToReleaseListItemRetiring  `json:"retiring"`
}

type MachinesToReleaseListItemCreator struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

type MachinesToReleaseListItemRetiring struct {
	Difficulty string `json:"difficulty_text"`
	Avatar     string `json:"avatar"`
	OS         string `json:"os"`
	Name       string `json:"name"`
	ID         int    `json:"id"`
}

// Get the list of machines using the todo endpoint
// https://www.hackthebox.com/api/v4/machine/list/todo

type MachinesTodoList struct {
	Info []*MachinesTodoListItem `json:"info"`
}

type MachinesTodoListItem struct {
	ID              int             `json:"id"`
	Name            string          `json:"name"`
	OS              string          `json:"os"`
	Points          int             `json:"points"`
	StaticPoints    int             `json:"static_points"`
	Release         string          `json:"release"`
	UserOwns        int             `json:"user_owns_count"`
	RootOwns        int             `json:"root_owns_count"`
	HasOwnedUser    bool            `json:"authUserInUserOwns"`
	HasOwnedRoot    bool            `json:"authUserInRootOwns"`
	IsTodo          bool            `json:"isTodo"`
	HasReviewed     bool            `json:"authUserHasReviewed"`
	Stars           string          `json:"stars"`
	DifficultyAvg   int             `json:"difficulty"`
	DifficultyStats *Difficulties   `json:"feedbackForChart"`
	Avatar          string          `json:"avatar"`
	Difficulty      string          `json:"difficultyText"`
	PlayInfo        *MachineStatus  `json:"playInfo"`
	IsFree          bool            `json:"free"`
	CreatorOne      *MachineCreator `json:"maker"`
	CreatorTwo      *MachineCreator `json:"maker2"`
	Recommended     int             `json:"recommended"`
	SPFlag          int             `json:"sp_flag"`
	EasyMonth       int             `json:"easy_month"`
	IP              string          `json:"ip"`
	Tags            []*MachineTag   `json:"tags"`
}

// Get a retired machine their tags
// https://www.hackthebox.com/api/v4/machine/tags/{machineID}

type MachineTagsMap struct {
	Info map[string]*MachineTagsMapItem `json:"info"`
}

type MachineTagsMapItem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Get the changelog of a machine
// https://www.hackthebox.com/api/v4/machine/changelog/{machineID}

type MachineChangelogList struct {
	Info []*MachineChangelogListItem `json:"info"`
}

type MachineChangelogListItem struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	MachineID   int    `json:"machine_id"`
	Type        string `json:"type"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Released    int    `json:"released"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// Get the reviews of a retired machine
// https://www.hackthebox.com/api/v4/machine/reviews/{machineID}

type MachineReviewsList struct {
	Message []*MachineReviewsListItem `json:"message"`
}

type MachineReviewsListItem struct {
	ID        int                         `json:"id"`
	UserID    int                         `json:"user_id"`
	MachineID int                         `json:"machine_id"`
	Stars     int                         `json:"stars"`
	Message   string                      `json:"message"`
	CreatedAt string                      `json:"created_at"`
	UpdatedAt string                      `json:"updated_at"`
	Title     string                      `json:"title"`
	Released  int                         `json:"released"`
	User      *MachineReviewsListItemUser `json:"user"`
}

type MachineReviewsListItemUser struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

func (s *Session) MachineMatrix(machineID int) (matrix *MachineMatrixInfo, err error) {

	machineIDString, err := toPositiveIntString(machineID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/machine/graph/matrix/" + machineIDString
	err = parseJSON(s, url, &matrix)

	return
}

func (s *Session) Machine(machineID int) (info *MachineInfo, err error) {

	machineIDString, err := toPositiveIntString(machineID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/machine/info/" + machineIDString
	err = parseJSON(s, url, &info)

	return
}

func (s *Session) MachineProfile(machineID int) (profile *MachineProfileInfo, err error) {

	machineIDString, err := toPositiveIntString(machineID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/machine/profile/" + machineIDString
	err = parseJSON(s, url, &profile)

	return
}

func (s *Session) MachinesActive() (activeMachines *MachinesActiveList, err error) {

	var url string = "https://www.hackthebox.com/api/v4/machine/list"
	err = parseJSON(s, url, &activeMachines)

	return
}

func (s *Session) MachinesRetired() (retiredMachines *MachinesRetiredList, err error) {

	var url string = "https://www.hackthebox.com/api/v4/machine/list/retired"
	err = parseJSON(s, url, &retiredMachines)

	return
}

func (s *Session) MachineTopUsers(machineID int) (topUsers *MachineTopsUsersList, err error) {

	machineIDString, err := toPositiveIntString(machineID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/machine/owns/top/" + machineIDString
	err = parseJSON(s, url, &topUsers)
	return
}

func (s *Session) MachinesToRelease() (scheduledMachines *MachinesToReleaseList, err error) {

	var url string = "https://www.hackthebox.com/api/v4/machine/unreleased"
	err = parseJSON(s, url, &scheduledMachines)

	return
}

func (s *Session) MachinesTodo() (machines *MachinesTodoList, err error) {

	var url string = "https://www.hackthebox.com/api/v4/machine/list/todo"
	err = parseJSON(s, url, &machines)

	return
}

func (s *Session) MachineTags(machineID int) (tags *MachineTagsMap, err error) {

	machineIDString, err := toPositiveIntString(machineID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/machine/tags/" + machineIDString
	err = parseJSON(s, url, &tags)

	return
}

func (s *Session) MachineChangelog(machineID int) (changelog *MachineChangelogList, err error) {

	machineIDString, err := toPositiveIntString(machineID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/machine/changelog/" + machineIDString
	err = parseJSON(s, url, &changelog)

	return
}

func (s *Session) MachineReviews(machineID int) (reviews *MachineReviewsList, err error) {

	machineIDString, err := toPositiveIntString(machineID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/machine/reviews/" + machineIDString
	err = parseJSON(s, url, &reviews)

	return
}
