package htbgo

// Bearer Connection Status
// https://www.hackthebox.com/api/v4/user/connection/status

type ConnectionStatusInfo struct {
	Status     string      `json:"status"`
	Connection interface{} `json:"connection"`
}

// Bearer Active Machine
// https://www.hackthebox.com/api/v4/machine/active

type ActiveMachineInfo struct {
	Info *ActiveMachineInfoItem `json:"info"`
}

type ActiveMachineInfoItem struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Avatar     string `json:"avatar"`
	ExpiresAt  string `json:"expires_at"`
	IsVoting   bool   `json:"voting"`
	HasVoted   bool   `json:"voted"`
	IsSpawning bool   `json:"isSpawning"`
	Type       string `json:"type"`
	LabServer  string `json:"lab_server"`
}

// Bearer Followers
// https://www.hackthebox.com/api/v4/user/followers

type FollowersList struct {
	Info []*FollowersListItem `json:"info"`
}

type FollowersListItem struct {
	ID int `json:"id"`
}

// Bearer Profile
// https://www.hackthebox.com/api/v4/user/info

type ProfileInfo struct {
	Info *ProfileInfoItem `json:"info"`
}

type ProfileInfoItem struct {
	ID                        int                  `json:"id"`
	Name                      string               `json:"name"`
	Email                     string               `json:"email"`
	Timezone                  string               `json:"timezone"`
	IsVIP                     bool                 `json:"isVip"`
	IsModerator               bool                 `json:"isModerator"`
	IsBGModerator             bool                 `json:"isBGModerator"`
	IsChatBanned              bool                 `json:"isChatBanned"`
	IsDedicatedVIP            bool                 `json:"isDedicatedVip"`
	CanAccessVIP              bool                 `json:"canAccessVIP"`
	CanAccessDediLab          bool                 `json:"canAccessDedilab"`
	IsServerVIP               bool                 `json:"isServerVIP"`
	ServerID                  int                  `json:"server_id"`
	Avatar                    string               `json:"avatar"`
	BetaTester                int                  `json:"beta_tester"`
	RankID                    int                  `json:"rank_id"`
	IsOnboardComplete         bool                 `json:"onboarding_completed"`
	IsOnboardTutorialComplete int                  `json:"onboarding_tutorial_complete"`
	Verified                  bool                 `json:"verified"`
	CanDeleteAvatar           bool                 `json:"can_delete_avatar"`
	Team                      *ProfileInfoItemTeam `json:"team"`
	University                interface{}          `json:"university"`
	Identifier                string               `json:"identifier"`
	HasTeamInvitation         bool                 `json:"hasTeamInvitation"`
	HasTwoFAEnabled           bool                 `json:"TwoFaEnabled"`
	HasAppTokens              bool                 `json:"hasAppTokens"`
	SubscriptionPlan          string               `json:"subscription_plan"`
	DunningExists             bool                 `json:"dunning_exists"`
}

type ProfileInfoItemTeam struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	AvatarThumb string `json:"avatar_thumb_url"`
}

// Bearer Settings
// https://www.hackthebox.com/api/v4/user/settings

type SettingsInfo struct {
	Email           string `json:"email"`
	Notifications   int    `json:"email_notifications"`
	Public          int    `json:"public"`
	NameChangeDelay int    `json:"name_change_delay"`
	HideMachineTags int    `json:"hide_machine_tags"`
}

// Bearer Subscriptions
// https://www.hackthebox.com/api/v4/user/subscriptions

type SubscriptionsList struct {
	Subscriptions []*SubscriptionsListItem `json:"subscriptions"`
}

type SubscriptionsListItem struct {
	Name        string        `json:"name"`
	EndsAt      string        `json:"ends_at"`
	RenewsAt    string        `json:"renews_at"`
	RecurlyPlan string        `json:"recurly_plan"`
	PausedAt    string        `json:"paused_at"`
	CreatedAt   string        `json:"created_at"`
	ResumeAt    string        `json:"resume_at"`
	StripePlan  string        `json:"stripe_plan"`
	Items       []interface{} `json:"items"`
}

// Bearer Subscriptions Balance
// https://www.hackthebox.com/api/v4/user/subscriptions/balance

type SubscriptionsBalanceInfo struct {
	Balances        *SubscriptionsBalanceInfoBalances `json:"balances"`
	DefaultCurrency string                            `json:"default_currency"`
}

type SubscriptionsBalanceInfoBalances struct {
	USD int `json:"USD"`
	GBP int `json:"GBP"`
	EUR int `json:"EUR"`
}

// Bearer Recurly URL
// https://www.hackthebox.com/api/v4/user/subscriptions/manage/recurly

type SubscriptionsRecurlyInfo struct {
	Message string `json:"message"`
	URL     string `json:"url"`
}

// Bearer Enrolled Tracks
// https://www.hackthebox.com/api/v4/user/tracks

type EnrolledTracksList []struct {
	ID       int `json:"id"`
	Complete int `json:"complete"`
}

// User Profile with Bearer Relationship
// https://www.hackthebox.com/api/v4/user/profile/basic/{userID}

type UserRelationshipInfo struct {
	Profile *UserRelationshipInfoItem `json:"profile"`
}

type UserRelationshipInfoItem struct {
	ID              int                                 `json:"id"`
	SSOID           int                                 `json:"sso_id"`
	Name            string                              `json:"name"`
	SystemOwns      int                                 `json:"system_owns"`
	UserOwns        int                                 `json:"user_owns"`
	UserBloods      int                                 `json:"user_bloods"`
	SystemBloods    int                                 `json:"system_bloods"`
	Team            *UserRelationshipInfoItemTeam       `json:"team"`
	Respects        int                                 `json:"respects"`
	Rank            string                              `json:"rank"`
	RankID          int                                 `json:"rank_id"`
	RankProgress    int                                 `json:"current_rank_progress"`
	NextRank        string                              `json:"next_rank"`
	NextRankPoints  float64                             `json:"next_rank_points"`
	RankOwnership   string                              `json:"rank_ownership"`
	RankRequirement int                                 `json:"rank_requirement"`
	Ranking         int                                 `json:"ranking"`
	Avatar          string                              `json:"avatar"`
	Timezone        string                              `json:"timezone"`
	IsVIP           bool                                `json:"isVip"`
	IsDedicatedVip  bool                                `json:"isDedicatedVip"`
	IsPublic        bool                                `json:"public"`
	CountryName     string                              `json:"country_name"`
	CountryCode     string                              `json:"country_code"`
	Points          int                                 `json:"points"`
	University      *UserRelationshipInfoItemUniversity `json:"university"`
	UniversityName  string                              `json:"university_name"`
	Description     string                              `json:"description"`
	Github          string                              `json:"github"`
	LinkedIn        string                              `json:"linkedin"`
	Twitter         string                              `json:"twitter"`
	Website         string                              `json:"website"`
	IsRespected     bool                                `json:"isRespected"`
	IsFollowed      bool                                `json:"isFollowed"`
}

type UserRelationshipInfoItemTeam struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Ranking int    `json:"ranking"`
	Avatar  string `json:"avatar"`
}

type UserRelationshipInfoItemUniversity struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	LogoThumb string `json:"logo_thumb_url"`
	Rank      int    `json:"rank"`
}

// User Machines Progress
// https://www.hackthebox.com/api/v4/profile/progress/machines/os/{userID}

type UserMachinesList struct {
	Profile *UserMachinesListItem `json:"profile"`
}

type UserMachinesListItem struct {
	OperatingSystems []*UserMachinesListItemOperatingSystem `json:"operating_systems"`
}

type UserMachinesListItemOperatingSystem struct {
	Name          string  `json:"name"`
	Completion    float64 `json:"completion_percentage"`
	OwnedMachines int     `json:"owned_machines"`
	TotalMachines int     `json:"total_machines"`
}

// User Challenges Progress
// https://www.hackthebox.com/api/v4/profile/progress/challenges/{userID}

type UserChallengesList struct {
	Profile *UserChallengesListItem `json:"profile"`
}

type UserChallengesListItem struct {
	Owns       *UserChallengesListItemOwns       `json:"challenge_owns"`
	Categories []*UserChallengesListItemCategory `json:"challenge_categories"`
}

type UserChallengesListItemOwns struct {
	Solved int `json:"solved"`
	Total  int `json:"total"`
}

type UserChallengesListItemCategory struct {
	Name          string  `json:"name"`
	OwnedFlags    int     `json:"owned_flags"`
	TotalFlags    int     `json:"total_flags"`
	Completion    float64 `json:"completion_percentage"`
	UserSolvedAvg float64 `json:"avg_user_solved"`
}

// User Endgames Progress
// https://www.hackthebox.com/api/v4/profile/progress/endgame/{userID}

type UserEndgamesList struct {
	Profile *UserEndgamesListItem `json:"profile"`
}

type UserEndgamesListItem struct {
	Endgames []*UserEndgamesListItemEndgame `json:"endgames"`
}

type UserEndgamesListItemEndgame struct {
	Name       string  `json:"name"`
	Completion float64 `json:"completion_percentage"`
	OwnedFlags int     `json:"owned_flags"`
	TotalFlags int     `json:"total_flags"`
}

// User Fortress Progress
// https://www.hackthebox.com/api/v4/profile/progress/fortress/{userID}

type UserFortressesList struct {
	Profile *UserFortressesListItem `json:"profile"`
}

type UserFortressesListItem struct {
	Fortresses []*UserFortressesListItemFortress `json:"fortresses"`
}

type UserFortressesListItemFortress struct {
	Name       string  `json:"name"`
	Avatar     string  `json:"avatar"`
	Completion float64 `json:"completion_percentage"`
	OwnedFlags int     `json:"owned_flags"`
	TotalFlags int     `json:"total_flags"`
}

// User ProLabs Progress
// https://www.hackthebox.com/api/v4/profile/progress/prolab/{userID}

type UserProLabsList struct {
	Profile *UserProLabsListItem `json:"profile"`
}

type UserProLabsListItem struct {
	ProLabs []*UserProLabsListItemProLab `json:"prolabs"`
}

type UserProLabsListItemProLab struct {
	Name          string  `json:"name"`
	Completion    float64 `json:"completion_percentage"`
	OwnedFlags    int     `json:"owned_flags"`
	TotalFlags    int     `json:"total_flags"`
	TotalMachines int     `json:"total_machines"`
	RatingsAvg    float64 `json:"average_ratings"`
}

// User Activity
// https://www.hackthebox.com/api/v4/profile/activity/{userID}

type UserActivityList struct {
	Profile *UserActivityListItem `json:"profile"`
}

type UserActivityListItem struct {
	Activity []*UserActivityListItemActivity `json:"activity"`
}

type UserActivityListItemActivity struct {
	Date              string `json:"date"`
	DateDiff          string `json:"date_diff"`
	ObjectType        string `json:"object_type"`
	Type              string `json:"type"`
	Blood             bool   `json:"first_blood"`
	ID                int    `json:"id"`
	Name              string `json:"name"`
	Points            int    `json:"points"`
	MachineAvatar     string `json:"machine_avatar"`
	ChallengeCategory string `json:"challenge_category"`
	FlagTitle         string `json:"flag_title"`
}

// User Bloods
// https://www.hackthebox.com/api/v4/profile/bloods/{userID}

type UserBloodsList struct {
	Profile *UserBloodsListItem `json:"profile"`
}

type UserBloodsListItem struct {
	Bloods *UserBloodsListItemBloods `json:"bloods"`
}

type UserBloodsListItemBloods struct {
	Machines   []*UserBloodsListItemBloodsMachine   `json:"machines"`
	Challenges []*UserBloodsListItemBloodsChallenge `json:"challenges"`
}

type UserBloodsListItemBloodsMachine struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Avatar        string `json:"avatar"`
	OS            string `json:"os"`
	Difficulty    string `json:"difficulty_text"`
	UserBlood     bool   `json:"user_blood"`
	UserBloodDiff string `json:"user_blood_difference"`
	RootBlood     bool   `json:"root_blood"`
	RootBloodDiff string `json:"root_blood_difference"`
}

type UserBloodsListItemBloodsChallenge struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Category   string `json:"category_name"`
	Difficulty string `json:"difficulty_test"`
	BloodDiff  string `json:"blood_difference"`
}

// User Submissions
// https://www.hackthebox.com/api/v4/profile/content/{userID}

type UserSubmissionsList struct {
	Profile *UserSubmissionsListItem `json:"profile"`
}

type UserSubmissionsListItem struct {
	Content *UserSubmissionsListItemContent `json:"content"`
}

type UserSubmissionsListItemContent struct {
	Machines   []*UserSubmissionsListItemContentMachine   `json:"machines"`
	Challenges []*UserSubmissionsListItemContentChallenge `json:"challenges"`
	Writeups   []*UserSubmissionsListItemContentWriteup   `json:"writeups"`
}

type UserSubmissionsListItemContentMachine struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	OS         string `json:"os"`
	Avatar     string `json:"machine_avatar"`
	Difficulty string `json:"difficulty"`
	Rating     string `json:"rating"`
	UserOwns   int    `json:"user_owns"`
	SystemOwns int    `json:"system_owns"`
}

type UserSubmissionsListItemContentChallenge struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Likes    int    `json:"likes"`
	Dislikes int    `json:"dislikes"`
}

type UserSubmissionsListItemContentWriteup struct {
	ID            int    `json:"id"`
	MachineID     int    `json:"machine_id"`
	MachineAvatar string `json:"machine_avatar"`
	MachineName   string `json:"machine_name"`
	URL           string `json:"url"`
	Likes         int    `json:"likes"`
	Dislikes      int    `json:"dislikes"`
	Type          string `json:"type"`
}

// User Achievements
// https://www.hackthebox.com/api/v4/profile/graph/{duration}/{userID}

type UserAchievementsList struct {
	Profile *UserAchievementsListItem `json:"profile"`
}

type UserAchievementsListItem struct {
	Graph *UserAchievementsListItemGraph `json:"graphData"`
}

type UserAchievementsListItemGraph struct {
	UserOwns      []int `json:"user_owns"`
	SystemOwns    []int `json:"system_owns"`
	ChallengeOwns []int `json:"challenge_owns"`
	Bloods        []int `json:"first_bloods"`
	Respects      []int `json:"respects"`
}

// User Machine Owns By Attack Path
// https://www.hackthebox.com/api/v4/profile/chart/machines/attack/{userID}

type UserOwnsByPathMap struct {
	Profile *UserOwnsByPathMapItem `json:"profile"`
}

type UserOwnsByPathMapItem struct {
	MachineOwns        *UserOwnsByPathMapItemMachineOwns                  `json:"machine_owns"`
	MachineAttackPaths map[string]*UserOwnsByPathMapItemMachineAttackPath `json:"machine_attack_paths"`
}

type UserOwnsByPathMapItemMachineOwns struct {
	Solved int `json:"solved"`
	Total  int `json:"total"`
}

type UserOwnsByPathMapItemMachineAttackPath struct {
	Solved        int     `json:"solved"`
	Total         int     `json:"total"`
	UserSolvedAvg float64 `json:"avg_user_solved"`
}

// User Profile
// https://www.hackthebox.com/api/v4/profile/{userID}

type UserInfo struct {
	Profile *UserInfoItem `json:"profile"`
}

type UserInfoItem struct {
	ID              int               `json:"id"`
	SSOID           int               `json:"sso_id"`
	Name            string            `json:"name"`
	SystemOwns      int               `json:"system_owns"`
	UserOwns        int               `json:"user_owns"`
	UserBloods      int               `json:"user_bloods"`
	SystemBloods    int               `json:"system_bloods"`
	Team            *UserInfoItemTeam `json:"team"`
	Respects        int               `json:"respects"`
	Rank            string            `json:"rank"`
	RankID          int               `json:"rank_id"`
	RankProgress    int               `json:"current_rank_progress"`
	NextRank        string            `json:"next_rank"`
	NextRankPoints  float64           `json:"next_rank_points"`
	RankOwnership   interface{}       `json:"rank_ownership"`
	RankRequirement int               `json:"rank_requirement"`
	Ranking         int               `json:"ranking"`
	Avatar          string            `json:"avatar"`
	Timezone        string            `json:"timezone"`
	Points          int               `json:"points"`
	CountryName     string            `json:"country_name"`
	CountryCode     string            `json:"country_code"`
	UniversityName  string            `json:"university_name"`
	Description     string            `json:"description"`
	Github          string            `json:"github"`
	LinkedIn        string            `json:"linkedin"`
	Twitter         string            `json:"twitter"`
	Website         string            `json:"website"`
}

type UserInfoItemTeam struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Ranking int    `json:"ranking"`
	Avatar  string `json:"avatar"`
}

// User Badges
// https://www.hackthebox.com/api/v4/profile/badges/{userID}

type UserBadgesList struct {
	Badges []*UserBadgesListItem `json:"badges"`
}

type UserBadgesListItem struct {
	ID           int                      `json:"id"`
	Name         string                   `json:"name"`
	Description  string                   `json:"description_en"`
	Color        string                   `json:"color"`
	Icon         string                   `json:"icon"`
	CategoryID   int                      `json:"badge_category_id"`
	BadgableType string                   `json:"badgable_type"`
	BadgableID   int                      `json:"badgable_id"`
	Percentage   float64                  `json:"percentage"`
	Pivot        *UserBadgesListItemPivot `json:"pivot"`
}

type UserBadgesListItemPivot struct {
	UserID    int    `json:"user_id"`
	BadgeID   int    `json:"badge_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (s *Session) ConnectionStatus() (status *ConnectionStatusInfo, err error) {

	var url string = "https://www.hackthebox.com/api/v4/user/connection/status"
	err = parseJSON(s, url, &status)

	return
}

func (s *Session) ActiveMachine() (active *ActiveMachineInfo, err error) {

	var url string = "https://www.hackthebox.com/api/v4/machine/active"
	err = parseJSON(s, url, &active)

	return
}

func (s *Session) Followers() (followers *FollowersList, err error) {

	var url string = "https://www.hackthebox.com/api/v4/user/followers"
	err = parseJSON(s, url, &followers)

	return
}

func (s *Session) Profile() (profile *ProfileInfo, err error) {

	var url string = "https://www.hackthebox.com/api/v4/user/info"
	err = parseJSON(s, url, &profile)

	return
}

func (s *Session) Settings() (settings *SettingsInfo, err error) {

	var url string = "https://www.hackthebox.com/api/v4/user/settings"
	err = parseJSON(s, url, &settings)

	return
}

func (s *Session) Subscriptions() (subscriptions *SubscriptionsList, err error) {

	var url string = "https://www.hackthebox.com/api/v4/user/subscriptions"
	err = parseJSON(s, url, &subscriptions)

	return
}

func (s *Session) SubscriptionsBalance() (balance *SubscriptionsBalanceInfo, err error) {

	var url string = "https://www.hackthebox.com/api/v4/user/subscriptions/balance"
	err = parseJSON(s, url, &balance)

	return
}

func (s *Session) SubscriptionsRecurly() (recurly *SubscriptionsRecurlyInfo, err error) {

	var url string = "https://www.hackthebox.com/api/v4/user/subscriptions/manage/recurly"
	err = parseJSON(s, url, &recurly)

	return
}

func (s *Session) EnrolledTracks() (tracks *EnrolledTracksList, err error) {

	var url string = "https://www.hackthebox.com/api/v4/user/tracks"
	err = parseJSON(s, url, &tracks)

	return
}

func (s *Session) UserRelationship(userID int) (relationship *UserRelationshipInfo, err error) {

	userIDString, err := toPositiveIntString(userID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/user/profile/basic/" + userIDString
	err = parseJSON(s, url, &relationship)

	return
}

func (s *Session) UserMachines(userID int) (machines *UserMachinesList, err error) {

	userIDString, err := toPositiveIntString(userID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/profile/progress/machines/os/" + userIDString
	err = parseJSON(s, url, &machines)

	return
}

func (s *Session) UserChallenges(userID int) (challenges *UserChallengesList, err error) {

	userIDString, err := toPositiveIntString(userID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/profile/progress/challenges/" + userIDString
	err = parseJSON(s, url, &challenges)

	return
}

func (s *Session) UserEndgames(userID int) (endgames *UserEndgamesList, err error) {

	userIDString, err := toPositiveIntString(userID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/profile/progress/endgame/" + userIDString
	err = parseJSON(s, url, &endgames)

	return
}

func (s *Session) UserFortresses(userID int) (fortresses *UserFortressesList, err error) {

	userIDString, err := toPositiveIntString(userID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/profile/progress/fortress/" + userIDString
	err = parseJSON(s, url, &fortresses)

	return
}

func (s *Session) UserProLabs(userID int) (proLabs *UserProLabsList, err error) {

	userIDString, err := toPositiveIntString(userID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/profile/progress/prolab/" + userIDString
	err = parseJSON(s, url, &proLabs)

	return
}

func (s *Session) UserActivity(userID int) (activities *UserActivityList, err error) {

	userIDString, err := toPositiveIntString(userID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/profile/activity/" + userIDString
	err = parseJSON(s, url, &activities)

	return
}

func (s *Session) UserBloods(userID int) (bloods *UserBloodsList, err error) {

	userIDString, err := toPositiveIntString(userID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/profile/bloods/" + userIDString
	err = parseJSON(s, url, &bloods)

	return
}

func (s *Session) UserSubmissions(userID int) (submissions *UserSubmissionsList, err error) {

	userIDString, err := toPositiveIntString(userID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/profile/content/" + userIDString
	err = parseJSON(s, url, &submissions)

	return
}

func (s *Session) UserAchievements(userID int, duration Duration) (acheivements *UserAchievementsList, err error) {

	userIDString, err := toPositiveIntString(userID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/profile/graph/" + string(duration) + "/" + userIDString
	err = parseJSON(s, url, &acheivements)

	return
}

func (s *Session) UserOwnsByPath(userID int) (path *UserOwnsByPathMap, err error) {

	userIDString, err := toPositiveIntString(userID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/profile/chart/machines/attack/" + userIDString
	err = parseJSON(s, url, &path)

	return
}

func (s *Session) User(userID int) (profile *UserInfo, err error) {

	userIDString, err := toPositiveIntString(userID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/profile/" + userIDString
	err = parseJSON(s, url, &profile)

	return
}

func (s *Session) UserBadges(userID int) (badges *UserBadgesList, err error) {

	userIDString, err := toPositiveIntString(userID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/profile/badges/" + userIDString
	err = parseJSON(s, url, &badges)

	return
}
