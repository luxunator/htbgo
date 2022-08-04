package htbgo

// Get Team Profile
// https://www.hackthebox.com/api/v4/team/info/{teamid}

type TeamInfo struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Points          int    `json:"points"`
	Motto           string `json:"motto"`
	Description     string `json:"description"`
	CountryName     string `json:"country_name"`
	CountryCode     string `json:"country_code"`
	Avatar          string `json:"avatar_url"`
	Cover           string `json:"cover_image_url"`
	Twitter         string `json:"twitter"`
	Facebook        string `json:"facebook"`
	Discord         string `json:"discord"`
	IsPublic        bool   `json:"public"`
	CanDeleteAvatar bool   `json:"can_delete_avatar"`
	Captain         struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Thumb string `json:"avatar_thumb"`
	} `json:"captain"`
	IsRespected     bool `json:"is_respected"`
	JoinRequestSent bool `json:"join_request_sent"`
}

// Get Team Activity
// https://www.hackthebox.com/api/v4/team/activity/{teamid}

type TeamActivityList []struct {
	User struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Public int    `json:"public"`
		Thumb  string `json:"avatar_thumb"`
	} `json:"user"`
	Date              string `json:"date"`
	DateDiff          string `json:"date_diff"`
	Type              string `json:"type"`
	Blood             bool   `json:"first_blood"`
	ObjectType        string `json:"object_type"`
	ID                int    `json:"id"`
	Name              string `json:"name"`
	Points            int    `json:"points"`
	ChallengeCategory string `json:"challenge_category"`
	MachineAvatar     string `json:"machine_avatar"`
	FlagTitle         string `json:"flag_title"`
}

// Get Team Owns
// https://www.hackthebox.com/api/v4/team/stats/owns/{teamid}

type TeamOwnsWeek struct {
	UserOwns      int    `json:"user_owns"`
	SystemOwns    int    `json:"system_owns"`
	ChallengeOwns int    `json:"challenge_owns"`
	Bloods        int    `json:"first_bloods"`
	Respects      int    `json:"respects"`
	WeekEndDate   string `json:"week_end_date"`
}

type TeamOwnsByWeekMap struct {
	Rank          int                     `json:"rank"`
	UserOwns      int                     `json:"user_owns"`
	SystemOwns    int                     `json:"system_owns"`
	Bloods        int                     `json:"first_bloods"`
	ChallengeOwns int                     `json:"challenge_owns"`
	Respects      int                     `json:"respects"`
	Weekly        map[string]TeamOwnsWeek `json:"weekly"`
}

// Get Team Graph for Duration
// https://www.hackthebox.com/api/v4/team/graph/2102?duration={duration}

type TeamStatsDuringInfo struct {
	Status bool `json:"status"`
	Data   struct {
		Points  []int `json:"points"`
		Rank    []int `json:"rank"`
		Respect []int `json:"respect"`
	} `json:"data"`
}

// Get Team Ownage by Attack Path
// https://www.hackthebox.com/api/v4/team/chart/machines/attack/{teamid}

type TeamOwnsByPathMap struct {
	MachineOwns struct {
		Solved int `json:"solved"`
		Total  int `json:"total"`
	} `json:"machine_owns"`
	AttackPaths map[string]struct {
		Name           string  `json:"name"`
		Solved         int     `json:"solved"`
		Total          int     `json:"total"`
		TeamsSolvedAvg float64 `json:"avg_teams_solved"`
	} `json:"machine_attack_paths"`
}

// List Team Members
// https://www.hackthebox.com/api/v4/team/members/{teamid}

type TeamMembersList []struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Avatar      string      `json:"avatar"`
	Rank        interface{} `json:"rank"`
	Points      int         `json:"points"`
	RootOwns    int         `json:"root_owns"`
	RootBloods  int         `json:"root_bloods_count"`
	UserBloods  int         `json:"user_bloods_count"`
	UserOwns    int         `json:"user_owns"`
	RankText    string      `json:"rank_text"`
	CountryName string      `json:"country_name"`
	CountryCode string      `json:"country_code"`
	Role        string      `json:"role"`
	Team        struct {
		ID        int `json:"id"`
		CaptainID int `json:"captain_id"`
	} `json:"team"`
	Public int `json:"public"`
}

// List Team Invitations (only if team captain)
// https://www.hackthebox.com/api/v4/team/invitations/{teamid}

type TeamInvitationsList struct {
	Headers  struct{} `json:"headers"`
	Original []struct {
		ID          int `json:"id"`
		UserID      int `json:"user_id"`
		TeamID      int `json:"team_id"`
		UserRequest int `json:"user_request"`
		User        struct {
			ID       int         `json:"id"`
			Name     string      `json:"name"`
			Points   int         `json:"points"`
			UserOwns int         `json:"user_owns_count"`
			RootOwns int         `json:"root_owns_count"`
			Respects int         `json:"respected_by_count"`
			Thumb    string      `json:"avatar_thumb"`
			RankName string      `json:"rank_name"`
			Ranking  interface{} `json:"ranking"`
		} `json:"user"`
	} `json:"original"`
	Exception bool `json:"exception"`
}

// Global Rank History for Bearer's Team
// https://www.hackthebox.com/api/v4/rankings/team/best?period={duration}

type TeamRankingsInfo struct {
	Status bool `json:"status"`
	Data   struct {
		Rank      int    `json:"rank"`
		Date      string `json:"date"`
		RankChart []int  `json:"rank_chart_data"`
	} `json:"data"`
}

// Global Points History for Bearer's Team
// https://www.hackthebox.com/api/v4/rankings/team/overview?period={duration}

type TeamPointsInfo struct {
	Status bool `json:"status"`
	Data   struct {
		PointsDiff   int    `json:"points_diff"`
		PointsGrowth string `json:"points_growth"`
		RanksDiff    int    `json:"ranks_diff"`
		Chart        []int  `json:"chart_data"`
		Team         struct {
			ID     int    `json:"id"`
			Name   string `json:"name"`
			Avatar string `json:"avatar_url"`
			Thumb  string `json:"avatar_thumb_url"`
		} `json:"team"`
	} `json:"data"`
}

func (s *Session) Team(teamID int) (info *TeamInfo, err error) {

	teamIDString, err := toPositiveIntString(teamID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/team/info/" + teamIDString
	err = parseJSON(s, url, &info)

	return
}

func (s *Session) TeamActivity(teamID int) (activities *TeamActivityList, err error) {

	teamIDString, err := toPositiveIntString(teamID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/team/activity/" + teamIDString
	err = parseJSON(s, url, &activities)

	return
}

func (s *Session) TeamOwnsByWeek(teamID int) (owns *TeamOwnsByWeekMap, err error) {

	teamIDString, err := toPositiveIntString(teamID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/team/stats/owns/" + teamIDString
	err = parseJSON(s, url, &owns)

	return
}

func (s *Session) TeamStatsDuring(teamID int, duration Duration) (graph *TeamStatsDuringInfo, err error) {

	teamIDString, err := toPositiveIntString(teamID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/team/graph/" + teamIDString + "?duration=" + string(duration)
	err = parseJSON(s, url, &graph)

	return
}

func (s *Session) TeamOwnsByPath(teamID int) (path *TeamOwnsByPathMap, err error) {

	teamIDString, err := toPositiveIntString(teamID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/team/chart/machines/attack/" + teamIDString
	err = parseJSON(s, url, &path)

	return
}

func (s *Session) TeamMembers(teamID int) (members *TeamMembersList, err error) {

	teamIDString, err := toPositiveIntString(teamID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/team/members/" + teamIDString
	err = parseJSON(s, url, &members)

	return
}

func (s *Session) TeamInvitations(teamID int) (invitations *TeamInvitationsList, err error) {

	teamIDString, err := toPositiveIntString(teamID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/team/invitations/" + teamIDString
	err = parseJSON(s, url, &invitations)

	return
}

func (s *Session) TeamRankings(period Duration) (rankings *TeamRankingsInfo, err error) {

	var url string = "https://www.hackthebox.com/api/v4/rankings/team/best?period=" + string(period)
	err = parseJSON(s, url, &rankings)

	return
}

func (s *Session) TeamPoints(period Duration) (points *TeamPointsInfo, err error) {

	var url string = "https://www.hackthebox.com/api/v4/rankings/team/overview?period=" + string(period)
	err = parseJSON(s, url, &points)

	return
}
