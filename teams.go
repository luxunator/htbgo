package htbgo

// TeamInfo contains information about a team
type TeamInfo struct {
	ID              int              `json:"id"`
	Name            string           `json:"name"`
	Points          int              `json:"points"`
	Motto           string           `json:"motto"`
	Description     string           `json:"description"`
	CountryName     string           `json:"country_name"`
	CountryCode     string           `json:"country_code"`
	Avatar          string           `json:"avatar_url"`
	Cover           string           `json:"cover_image_url"`
	Twitter         string           `json:"twitter"`
	Facebook        string           `json:"facebook"`
	Discord         string           `json:"discord"`
	IsPublic        bool             `json:"public"`
	CanDeleteAvatar bool             `json:"can_delete_avatar"`
	Captain         *TeamInfoCaptain `json:"captain"`
	IsRespected     bool             `json:"is_respected"`
	JoinRequestSent bool             `json:"join_request_sent"`
}

// TeamInfoCaptain contains information about a teams captain
type TeamInfoCaptain struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Thumb string `json:"avatar_thumb"`
}

// TeamActivityList contains a list of team activity and its information
type TeamActivityList []struct {
	User              *TeamActivityListUser `json:"user"`
	Date              string                `json:"date"`
	DateDiff          string                `json:"date_diff"`
	Type              string                `json:"type"`
	Blood             bool                  `json:"first_blood"`
	ObjectType        string                `json:"object_type"`
	ID                int                   `json:"id"`
	Name              string                `json:"name"`
	Points            int                   `json:"points"`
	ChallengeCategory string                `json:"challenge_category"`
	MachineAvatar     string                `json:"machine_avatar"`
	FlagTitle         string                `json:"flag_title"`
}

// TeamActivityListUser contains information on the user of a team activity
type TeamActivityListUser struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Public int    `json:"public"`
	Thumb  string `json:"avatar_thumb"`
}

// TeamOwnsByWeekMap contains information about team owns by the week
type TeamOwnsByWeekMap struct {
	Rank          int                               `json:"rank"`
	UserOwns      int                               `json:"user_owns"`
	SystemOwns    int                               `json:"system_owns"`
	Bloods        int                               `json:"first_bloods"`
	ChallengeOwns int                               `json:"challenge_owns"`
	Respects      int                               `json:"respects"`
	Weekly        map[string]*TeamOwnsByWeekMapWeek `json:"weekly"`
}

// TeamOwnsByWeekMapWeek contains information about team owns in a week
type TeamOwnsByWeekMapWeek struct {
	UserOwns      int    `json:"user_owns"`
	SystemOwns    int    `json:"system_owns"`
	ChallengeOwns int    `json:"challenge_owns"`
	Bloods        int    `json:"first_bloods"`
	Respects      int    `json:"respects"`
	WeekEndDate   string `json:"week_end_date"`
}

// TeamStatsDuringInfo contains team stats during a period of time
type TeamStatsDuringInfo struct {
	Status bool                     `json:"status"`
	Data   *TeamStatsDuringInfoItem `json:"data"`
}

// TeamStatsDuringInfoItem contains information on a teams stats during a period of time
type TeamStatsDuringInfoItem struct {
	Points  []int `json:"points"`
	Rank    []int `json:"rank"`
	Respect []int `json:"respect"`
}

// TeamOwnsByPathMap contains team owns by the attack path
type TeamOwnsByPathMap struct {
	MachineOwns *TeamOwnsByPathMapMachineOwns           `json:"machine_owns"`
	AttackPaths map[string]*TeamOwnsByPathMapAttackPath `json:"machine_attack_paths"`
}

// TeamOwnsByPathMapMachineOwns contains total information of team machine owns with attack path
type TeamOwnsByPathMapMachineOwns struct {
	Solved int `json:"solved"`
	Total  int `json:"total"`
}

// TeamOwnsByPathMapAttackPath contains information about team owns in an attack path
type TeamOwnsByPathMapAttackPath struct {
	Name           string  `json:"name"`
	Solved         int     `json:"solved"`
	Total          int     `json:"total"`
	TeamsSolvedAvg float64 `json:"avg_teams_solved"`
}

// TeamMembersList contains a list of team members and their information
type TeamMembersList []struct {
	ID          int                  `json:"id"`
	Name        string               `json:"name"`
	Avatar      string               `json:"avatar"`
	Rank        interface{}          `json:"rank"`
	Points      int                  `json:"points"`
	RootOwns    int                  `json:"root_owns"`
	RootBloods  int                  `json:"root_bloods_count"`
	UserBloods  int                  `json:"user_bloods_count"`
	UserOwns    int                  `json:"user_owns"`
	RankText    string               `json:"rank_text"`
	CountryName string               `json:"country_name"`
	CountryCode string               `json:"country_code"`
	Role        string               `json:"role"`
	Team        *TeamMembersListTeam `json:"team"`
	Public      int                  `json:"public"`
}

// TeamMembersListTeam contains information on the team in a list of team members
type TeamMembersListTeam struct {
	ID        int `json:"id"`
	CaptainID int `json:"captain_id"`
}

// TeamInvitationsList contains a list of team invitations
type TeamInvitationsList struct {
	Headers   struct{}                       `json:"headers"`
	Original  []*TeamInvitationsListOriginal `json:"original"`
	Exception bool                           `json:"exception"`
}

// TeamInvitationsListOriginal contains information on a team invitation
type TeamInvitationsListOriginal struct {
	ID          int                              `json:"id"`
	UserID      int                              `json:"user_id"`
	TeamID      int                              `json:"team_id"`
	UserRequest int                              `json:"user_request"`
	User        *TeamInvitationsListOriginalUser `json:"user"`
}

// TeamInvitationsListOriginalUser contains information on the user of a team invitation
type TeamInvitationsListOriginalUser struct {
	ID       int         `json:"id"`
	Name     string      `json:"name"`
	Points   int         `json:"points"`
	UserOwns int         `json:"user_owns_count"`
	RootOwns int         `json:"root_owns_count"`
	Respects int         `json:"respected_by_count"`
	Thumb    string      `json:"avatar_thumb"`
	RankName string      `json:"rank_name"`
	Ranking  interface{} `json:"ranking"`
}

// Team returns the information of a team, given a team ID
func (s *Session) Team(teamID int) (info *TeamInfo, err error) {

	teamIDString, err := toPositiveIntString(teamID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/team/info/" + teamIDString
	err = parseJSON(s, url, &info)

	return
}

// TeamActivity returns the recent activity of users in a team, given a team ID
func (s *Session) TeamActivity(teamID int) (activities *TeamActivityList, err error) {

	teamIDString, err := toPositiveIntString(teamID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/team/activity/" + teamIDString
	err = parseJSON(s, url, &activities)

	return
}

// TeamOwnsByWeek returns the owns information of a team by the week, given a team ID
func (s *Session) TeamOwnsByWeek(teamID int) (owns *TeamOwnsByWeekMap, err error) {

	teamIDString, err := toPositiveIntString(teamID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/team/stats/owns/" + teamIDString
	err = parseJSON(s, url, &owns)

	return
}

// TeamStatsDuring returns the statistics of a team within a period of time, given a team ID and duration
func (s *Session) TeamStatsDuring(teamID int, duration Duration) (graph *TeamStatsDuringInfo, err error) {

	teamIDString, err := toPositiveIntString(teamID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/team/graph/" + teamIDString + "?duration=" + string(duration)
	err = parseJSON(s, url, &graph)

	return
}

// TeamOwnsByPath returns the owns information of a team by attack path, given a team ID
func (s *Session) TeamOwnsByPath(teamID int) (path *TeamOwnsByPathMap, err error) {

	teamIDString, err := toPositiveIntString(teamID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/team/chart/machines/attack/" + teamIDString
	err = parseJSON(s, url, &path)

	return
}

// TeamMembers returns a list of a teams members, given a team ID
func (s *Session) TeamMembers(teamID int) (members *TeamMembersList, err error) {

	teamIDString, err := toPositiveIntString(teamID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/team/members/" + teamIDString
	err = parseJSON(s, url, &members)

	return
}

// TeamInvitations returns a list of invitations to a team, given a team ID (TOKEN USER NEEDS TO BE TEAM CAPTAIN)
func (s *Session) TeamInvitations(teamID int) (invitations *TeamInvitationsList, err error) {

	teamIDString, err := toPositiveIntString(teamID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/team/invitations/" + teamIDString
	err = parseJSON(s, url, &invitations)

	return
}
