package htbgo

// ConnectionStatusInfo contains information on the connection status of the Token User
type ConnectionStatusInfo struct {
	Status     string      `json:"status"`
	Connection interface{} `json:"connection"`
}

// ActiveMachineInfo contains Token User active machine
type ActiveMachineInfo struct {
	Info *ActiveMachineInfoItem `json:"info"`
}

// ActiveMachineInfoItem information on the Token Users active machine
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

// FollowersList contains a list of the Token Users followers
type FollowersList struct {
	Info []*FollowersListItem `json:"info"`
}

// FollowersListItem contains information on a follower of the Token User
type FollowersListItem struct {
	ID int `json:"id"`
}

// ProfileInfo contains the Token User profile
type ProfileInfo struct {
	Info *ProfileInfoItem `json:"info"`
}

// ProfileInfoItem contains the information of the Token User
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

// ProfileInfoItemTeam contains the team information of a Token User
type ProfileInfoItemTeam struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	AvatarThumb string `json:"avatar_thumb_url"`
}

// SettingsInfo contains the settings information for the Token User
type SettingsInfo struct {
	Email           string `json:"email"`
	Notifications   int    `json:"email_notifications"`
	Public          int    `json:"public"`
	NameChangeDelay int    `json:"name_change_delay"`
	HideMachineTags int    `json:"hide_machine_tags"`
}

// SubscriptionsList contains a list of the Token Users subscriptions
type SubscriptionsList struct {
	Subscriptions []*SubscriptionsListItem `json:"subscriptions"`
}

// SubscriptionsListItem contains the information of a Token Users subscription
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

// SubscriptionsBalanceInfo contains Token User balance and currency
type SubscriptionsBalanceInfo struct {
	Balances        *SubscriptionsBalanceInfoBalances `json:"balances"`
	DefaultCurrency string                            `json:"default_currency"`
}

// SubscriptionsBalanceInfoBalances contains information on the Token Users balance
type SubscriptionsBalanceInfoBalances struct {
	USD int `json:"USD"`
	GBP int `json:"GBP"`
	EUR int `json:"EUR"`
}

// SubscriptionsRecurlyInfo contains the Token Users subscription recurly information
type SubscriptionsRecurlyInfo struct {
	Message string `json:"message"`
	URL     string `json:"url"`
}

// EnrolledTracksList contains a list of the Token Users enrolled tracks
type EnrolledTracksList []struct {
	ID       int `json:"id"`
	Complete int `json:"complete"`
}

// UserRelationshipInfo contains user and Token User relationship
type UserRelationshipInfo struct {
	Profile *UserRelationshipInfoItem `json:"profile"`
}

// UserRelationshipInfoItem contains information on a user and the relationship to the Token User
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

// UserRelationshipInfoItemTeam contains information of the team of a user in a Token User relationship
type UserRelationshipInfoItemTeam struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Ranking int    `json:"ranking"`
	Avatar  string `json:"avatar"`
}

// UserRelationshipInfoItemUniversity contains information of the university of a user in a Token User relationship
type UserRelationshipInfoItemUniversity struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	LogoThumb string `json:"logo_thumb_url"`
	Rank      int    `json:"rank"`
}

// UserMachinesList contains a users machine progress
type UserMachinesList struct {
	Profile *UserMachinesListItem `json:"profile"`
}

// UserMachinesListItem contains a list of a users machine progress by operating system
type UserMachinesListItem struct {
	OperatingSystems []*UserMachinesListItemOperatingSystem `json:"operating_systems"`
}

// UserMachinesListItemOperatingSystem contains information on a users machine progress of an operating system
type UserMachinesListItemOperatingSystem struct {
	Name          string  `json:"name"`
	Completion    float64 `json:"completion_percentage"`
	OwnedMachines int     `json:"owned_machines"`
	TotalMachines int     `json:"total_machines"`
}

// UserChallengesList contains a users challenge progress
type UserChallengesList struct {
	Profile *UserChallengesListItem `json:"profile"`
}

// UserChallengesListItem contains a list of users challenge progress by categories
type UserChallengesListItem struct {
	Owns       *UserChallengesListItemOwns       `json:"challenge_owns"`
	Categories []*UserChallengesListItemCategory `json:"challenge_categories"`
}

// UserChallengesListItemOwns contains information on solves in users challenge progress
type UserChallengesListItemOwns struct {
	Solved int `json:"solved"`
	Total  int `json:"total"`
}

// UserChallengesListItemCategory contains information on challenge sovles in users challenge progress in a category
type UserChallengesListItemCategory struct {
	Name          string  `json:"name"`
	OwnedFlags    int     `json:"owned_flags"`
	TotalFlags    int     `json:"total_flags"`
	Completion    float64 `json:"completion_percentage"`
	UserSolvedAvg float64 `json:"avg_user_solved"`
}

// UserEndgamesList contains a users endgame progress
type UserEndgamesList struct {
	Profile *UserEndgamesListItem `json:"profile"`
}

// UserEndgamesListItem contains a list of users endgame progress
type UserEndgamesListItem struct {
	Endgames []*UserEndgamesListItemEndgame `json:"endgames"`
}

// UserEndgamesListItemEndgame contains information on a users progress of an endgame
type UserEndgamesListItemEndgame struct {
	Name       string  `json:"name"`
	Completion float64 `json:"completion_percentage"`
	OwnedFlags int     `json:"owned_flags"`
	TotalFlags int     `json:"total_flags"`
}

// UserFortressesList contains a users fortress progress
type UserFortressesList struct {
	Profile *UserFortressesListItem `json:"profile"`
}

// UserFortressesListItem contains a list of users fortress progress
type UserFortressesListItem struct {
	Fortresses []*UserFortressesListItemFortress `json:"fortresses"`
}

// UserFortressesListItemFortress contains information on a users progress of a fortress
type UserFortressesListItemFortress struct {
	Name       string  `json:"name"`
	Avatar     string  `json:"avatar"`
	Completion float64 `json:"completion_percentage"`
	OwnedFlags int     `json:"owned_flags"`
	TotalFlags int     `json:"total_flags"`
}

// UserProLabsList contains a users prolab progress
type UserProLabsList struct {
	Profile *UserProLabsListItem `json:"profile"`
}

// UserProLabsListItem contains a list of users prolab progress
type UserProLabsListItem struct {
	ProLabs []*UserProLabsListItemProLab `json:"prolabs"`
}

// UserProLabsListItemProLab contains information on a users progress of a prolab
type UserProLabsListItemProLab struct {
	Name          string  `json:"name"`
	Completion    float64 `json:"completion_percentage"`
	OwnedFlags    int     `json:"owned_flags"`
	TotalFlags    int     `json:"total_flags"`
	TotalMachines int     `json:"total_machines"`
	RatingsAvg    float64 `json:"average_ratings"`
}

// UserActivityList contains a users activity
type UserActivityList struct {
	Profile *UserActivityListItem `json:"profile"`
}

// UserActivityListItem contains a list of a users activity
type UserActivityListItem struct {
	Activity []*UserActivityListItemActivity `json:"activity"`
}

// UserActivityListItemActivity contains information of an activity of a user
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

// UserBloodsList contains a users bloods
type UserBloodsList struct {
	Profile *UserBloodsListItem `json:"profile"`
}

// UserBloodsListItem contains a users machine and challenge bloods
type UserBloodsListItem struct {
	Bloods *UserBloodsListItemBloods `json:"bloods"`
}

// UserBloodsListItemBloods contains lists of machine and challenge bloods of a user
type UserBloodsListItemBloods struct {
	Machines   []*UserBloodsListItemBloodsMachine   `json:"machines"`
	Challenges []*UserBloodsListItemBloodsChallenge `json:"challenges"`
}

// UserBloodsListItemBloodsMachine contains information on a users machine blood
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

// UserBloodsListItemBloodsChallenge contains information on a users challenge blood
type UserBloodsListItemBloodsChallenge struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Category   string `json:"category_name"`
	Difficulty string `json:"difficulty_test"`
	BloodDiff  string `json:"blood_difference"`
}

// UserSubmissionsList contains user submissions
type UserSubmissionsList struct {
	Profile *UserSubmissionsListItem `json:"profile"`
}

// UserSubmissionsListItem contains a users machine, challenge, and writeup submissions
type UserSubmissionsListItem struct {
	Content *UserSubmissionsListItemContent `json:"content"`
}

// UserSubmissionsListItemContent contains lists of machine, challenge, and writeup submissions of a user
type UserSubmissionsListItemContent struct {
	Machines   []*UserSubmissionsListItemContentMachine   `json:"machines"`
	Challenges []*UserSubmissionsListItemContentChallenge `json:"challenges"`
	Writeups   []*UserSubmissionsListItemContentWriteup   `json:"writeups"`
}

// UserSubmissionsListItemContentMachine contains information of a users machine submission
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

// UserSubmissionsListItemContentChallenge contains information of a users challenge submission
type UserSubmissionsListItemContentChallenge struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Likes    int    `json:"likes"`
	Dislikes int    `json:"dislikes"`
}

// UserSubmissionsListItemContentWriteup contains information of a users writeup submission
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

// UserAchievementsList contains user achievements
type UserAchievementsList struct {
	Profile *UserAchievementsListItem `json:"profile"`
}

// UserAchievementsListItem contains a users achievement graph
type UserAchievementsListItem struct {
	Graph *UserAchievementsListItemGraph `json:"graphData"`
}

// UserAchievementsListItemGraph contains lists of a users achievments
type UserAchievementsListItemGraph struct {
	UserOwns      []int `json:"user_owns"`
	SystemOwns    []int `json:"system_owns"`
	ChallengeOwns []int `json:"challenge_owns"`
	Bloods        []int `json:"first_bloods"`
	Respects      []int `json:"respects"`
}

// UserOwnsByPathMap contains user attack path owns
type UserOwnsByPathMap struct {
	Profile *UserOwnsByPathMapItem `json:"profile"`
}

// UserOwnsByPathMapItem contains user owns by the attack path
type UserOwnsByPathMapItem struct {
	MachineOwns        *UserOwnsByPathMapItemMachineOwns                  `json:"machine_owns"`
	MachineAttackPaths map[string]*UserOwnsByPathMapItemMachineAttackPath `json:"machine_attack_paths"`
}

// UserOwnsByPathMapItemMachineOwns contains total information of user machine owns with attack path
type UserOwnsByPathMapItemMachineOwns struct {
	Solved int `json:"solved"`
	Total  int `json:"total"`
}

// UserOwnsByPathMapItemMachineAttackPath contains information about user owns in an attack path
type UserOwnsByPathMapItemMachineAttackPath struct {
	Solved        int     `json:"solved"`
	Total         int     `json:"total"`
	UserSolvedAvg float64 `json:"avg_user_solved"`
}

// UserInfo contains a user
type UserInfo struct {
	Profile *UserInfoItem `json:"profile"`
}

// UserInfoItem contains a users information
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

// UserInfoItemTeam contains team information of a user
type UserInfoItemTeam struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Ranking int    `json:"ranking"`
	Avatar  string `json:"avatar"`
}

// UserBadgesList contains a list of badges of a user
type UserBadgesList struct {
	Badges []*UserBadgesListItem `json:"badges"`
}

// UserBadgesListItem contains information on a badge of a user
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

// UserBadgesListItemPivot contains pivot information on a badge of a user
type UserBadgesListItemPivot struct {
	UserID    int    `json:"user_id"`
	BadgeID   int    `json:"badge_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// ConnectionStatus returns the connection status of the Token User
func (s *Session) ConnectionStatus() (status *ConnectionStatusInfo, err error) {

	var url string = "https://www.hackthebox.com/api/v4/user/connection/status"
	err = parseJSON(s, url, &status)

	return
}

// ActiveMachine returns the active machine of the Token User
func (s *Session) ActiveMachine() (active *ActiveMachineInfo, err error) {

	var url string = "https://www.hackthebox.com/api/v4/machine/active"
	err = parseJSON(s, url, &active)

	return
}

// Followers returns a list of users following the Token User
func (s *Session) Followers() (followers *FollowersList, err error) {

	var url string = "https://www.hackthebox.com/api/v4/user/followers"
	err = parseJSON(s, url, &followers)

	return
}

// Profile returns the profile information of the Token User
func (s *Session) Profile() (profile *ProfileInfo, err error) {

	var url string = "https://www.hackthebox.com/api/v4/user/info"
	err = parseJSON(s, url, &profile)

	return
}

// Settings returns the setting information of the Token User
func (s *Session) Settings() (settings *SettingsInfo, err error) {

	var url string = "https://www.hackthebox.com/api/v4/user/settings"
	err = parseJSON(s, url, &settings)

	return
}

// Subscriptions returns the current subscriptions of the Token User
func (s *Session) Subscriptions() (subscriptions *SubscriptionsList, err error) {

	var url string = "https://www.hackthebox.com/api/v4/user/subscriptions"
	err = parseJSON(s, url, &subscriptions)

	return
}

// SubscriptionsBalance returns the current subscriptions balance of the Token User
func (s *Session) SubscriptionsBalance() (balance *SubscriptionsBalanceInfo, err error) {

	var url string = "https://www.hackthebox.com/api/v4/user/subscriptions/balance"
	err = parseJSON(s, url, &balance)

	return
}

// SubscriptionsRecurly returns the subscriptions recurly url of the Token User
func (s *Session) SubscriptionsRecurly() (recurly *SubscriptionsRecurlyInfo, err error) {

	var url string = "https://www.hackthebox.com/api/v4/user/subscriptions/manage/recurly"
	err = parseJSON(s, url, &recurly)

	return
}

// EnrolledTracks returns a list of tracks that the Token User is enrolled in
func (s *Session) EnrolledTracks() (tracks *EnrolledTracksList, err error) {

	var url string = "https://www.hackthebox.com/api/v4/user/tracks"
	err = parseJSON(s, url, &tracks)

	return
}

// UserRelationship returns the profile information and Token User relation of a user, given a user ID
func (s *Session) UserRelationship(userID int) (relationship *UserRelationshipInfo, err error) {

	userIDString, err := toPositiveIntString(userID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/user/profile/basic/" + userIDString
	err = parseJSON(s, url, &relationship)

	return
}

// UserMachines returns a list of information containing a users machines progess, given a user ID
func (s *Session) UserMachines(userID int) (machines *UserMachinesList, err error) {

	userIDString, err := toPositiveIntString(userID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/profile/progress/machines/os/" + userIDString
	err = parseJSON(s, url, &machines)

	return
}

// UserChallenges returns a list of information containing a users challenges progess, given a user ID
func (s *Session) UserChallenges(userID int) (challenges *UserChallengesList, err error) {

	userIDString, err := toPositiveIntString(userID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/profile/progress/challenges/" + userIDString
	err = parseJSON(s, url, &challenges)

	return
}

// UserEndgames  returns a list of information containing a users endgames progess, given a user ID
func (s *Session) UserEndgames(userID int) (endgames *UserEndgamesList, err error) {

	userIDString, err := toPositiveIntString(userID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/profile/progress/endgame/" + userIDString
	err = parseJSON(s, url, &endgames)

	return
}

// UserFortresses returns a list of information containing a users fortresses progess, given a user ID
func (s *Session) UserFortresses(userID int) (fortresses *UserFortressesList, err error) {

	userIDString, err := toPositiveIntString(userID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/profile/progress/fortress/" + userIDString
	err = parseJSON(s, url, &fortresses)

	return
}

// UserProLabs returns a list of information containing a users prolabs progess, given a user ID
func (s *Session) UserProLabs(userID int) (proLabs *UserProLabsList, err error) {

	userIDString, err := toPositiveIntString(userID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/profile/progress/prolab/" + userIDString
	err = parseJSON(s, url, &proLabs)

	return
}

// UserActivity returns the recent activity of a user, given a user ID
func (s *Session) UserActivity(userID int) (activities *UserActivityList, err error) {

	userIDString, err := toPositiveIntString(userID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/profile/activity/" + userIDString
	err = parseJSON(s, url, &activities)

	return
}

// UserBloods returns a users blood information, given a user ID
func (s *Session) UserBloods(userID int) (bloods *UserBloodsList, err error) {

	userIDString, err := toPositiveIntString(userID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/profile/bloods/" + userIDString
	err = parseJSON(s, url, &bloods)

	return
}

// UserSubmissions returns a users submissions, given a user ID
func (s *Session) UserSubmissions(userID int) (submissions *UserSubmissionsList, err error) {

	userIDString, err := toPositiveIntString(userID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/profile/content/" + userIDString
	err = parseJSON(s, url, &submissions)

	return
}

// UserAchievements returns a users achievements, given a user ID
func (s *Session) UserAchievements(userID int, duration Duration) (acheivements *UserAchievementsList, err error) {

	userIDString, err := toPositiveIntString(userID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/profile/graph/" + string(duration) + "/" + userIDString
	err = parseJSON(s, url, &acheivements)

	return
}

// UserOwnsByPath returns the owns information of a user by attack path, given a user ID
func (s *Session) UserOwnsByPath(userID int) (path *UserOwnsByPathMap, err error) {

	userIDString, err := toPositiveIntString(userID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/profile/chart/machines/attack/" + userIDString
	err = parseJSON(s, url, &path)

	return
}

// User returns the information of a user, given a user ID
func (s *Session) User(userID int) (profile *UserInfo, err error) {

	userIDString, err := toPositiveIntString(userID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/profile/" + userIDString
	err = parseJSON(s, url, &profile)

	return
}

// UserBadges returns a list of a users badges, given a user ID
func (s *Session) UserBadges(userID int) (badges *UserBadgesList, err error) {

	userIDString, err := toPositiveIntString(userID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/profile/badges/" + userIDString
	err = parseJSON(s, url, &badges)

	return
}
