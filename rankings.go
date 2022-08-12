package htbgo

// RankBracket contains information on a ranking bracket
type RankBracket struct {
	Rank           interface{}    `json:"rank"`
	Points         int    `json:"points"`
	PointsToGo     int    `json:"points_for_next_bracket"`
	CurrentBracket string `json:"current_bracket"`
	NextBracket    string `json:"next_bracket"`
}

// RankDuring contains information on a rank during a period of time
type RankDuring struct {
	Rank      interface{}   `json:"rank"` // int or string "unranked"
	Date      string        `json:"date"`
	RankChart []interface{} `json:"rank_chart_data"`
}

// RankInCountryBestDuringInfo contains country best rank stats in period of time
type RankInCountryBestDuringInfo struct {
	Status bool        `json:"status"`
	Data   *RankDuring `json:"data"`
}

// RankInCountryDuringInfo contains country rank stats in period of time
type RankInCountryDuringInfo struct {
	Status bool                         `json:"status"`
	Data   *RankInCountryDuringInfoItem `json:"data"`
}

// RankInCountryDuringInfoItem contains information of country rank in a period of time
type RankInCountryDuringInfoItem struct {
	PointsDiff   int                                 `json:"points_diff"`
	PointsGrowth string                              `json:"points_growth"`
	RanksDiff    int                                 `json:"ranks_diff"`
	Chart        []int                               `json:"chart_data"`
	Country      *RankInCountryDuringInfoItemCountry `json:"country"`
}

// RankInCountryDuringInfoItemCountry contains information about a country in a country rank period of time
type RankInCountryDuringInfoItemCountry struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// RankBracketInCountryInfo contains information on a country rank bracket
type RankBracketInCountryInfo struct {
	Status bool         `json:"status"`
	Data   *RankBracket `json:"data"`
}

// RankOfTeamBestDuringInfo contains team best rank stats in period of time
type RankOfTeamBestDuringInfo struct {
	Status bool        `json:"status"`
	Data   *RankDuring `json:"data"`
}

// RankOfTeamDuringInfo contains team rank stats in period of time
type RankOfTeamDuringInfo struct {
	Status bool                      `json:"status"`
	Data   *RankOfTeamDuringInfoItem `json:"data"`
}

// RankOfTeamDuringInfoItem contains information of team rank in a period of time
type RankOfTeamDuringInfoItem struct {
	PointsDiff   int                           `json:"points_diff"`
	PointsGrowth string                        `json:"points_growth"`
	RanksDiff    int                           `json:"ranks_diff"`
	Chart        []int                         `json:"chart_data"`
	Team         *RankOfTeamDuringInfoItemTeam `json:"team"`
}

// RankOfTeamDuringInfoItemTeam contains information about a team in a team rank period of time
type RankOfTeamDuringInfoItemTeam struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar_url"`
	Thumb  string `json:"avatar_thumb_url"`
}

// RankBracketOfTeamInfo  contains information on a team rank bracket
type RankBracketOfTeamInfo struct {
	Status bool         `json:"status"`
	Data   *RankBracket `json:"data"`
}

// RankBestDuringInfo contains user best rank stats in period of time
type RankBestDuringInfo struct {
	Status bool        `json:"status"`
	Data   *RankDuring `json:"data"`
}

// RankDuringInfo contains user rank stats in period of time
type RankDuringInfo struct {
	Status bool                `json:"status"`
	Data   *RankDuringInfoItem `json:"data"`
}

// RankDuringInfoItem contains information of user rank in a period of time
type RankDuringInfoItem struct {
	PointsDiff   int                     `json:"points_diff"`
	PointsGrowth interface{}             `json:"points_growth"` // string or int 0
	RanksDiff    int                     `json:"ranks_diff"`
	Chart        []int                   `json:"chart_data"`
	User         *RankDuringInfoItemUser `json:"user"`
}

// RankDuringInfoItemUser contains information about a user in a user rank period of time
type RankDuringInfoItemUser struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Thumb  string `json:"avatar_thumb"`
}

// RankBracketInfo contains information for a users bracket rank
type RankBracketInfo struct {
	Status bool         `json:"status"`
	Data   *RankBracket `json:"data"`
}

// RanksOfCountriesList contains a list of country ranks
type RanksOfCountriesList struct {
	Status bool                        `json:"status"`
	Data   []*RanksOfCountriesListItem `json:"data"`
}

// RanksOfCountriesListItem contains information on a countries rank
type RanksOfCountriesListItem struct {
	Rank          int    `json:"rank"`
	Country       string `json:"country"`
	Members       int    `json:"members"`
	Points        int    `json:"points"`
	RootOwns      int    `json:"root_owns"`
	UserOwns      int    `json:"user_owns"`
	ChallengeOwns int    `json:"challenge_owns"`
	RootBloods    int    `json:"root_bloods"`
	UserBloods    int    `json:"user_bloods"`
	Fortress      int    `json:"fortress"`
	Endgame       int    `json:"endgame"`
	Name          string `json:"name"`
	RanksDiff     int    `json:"ranks_diff"`
}

// RanksOfTeamsList contains a list of team ranks
type RanksOfTeamsList struct {
	Status bool                    `json:"status"`
	Data   []*RanksOfTeamsListItem `json:"data"`
}

// RanksOfTeamsListItem contains information on a teams rank
type RanksOfTeamsListItem struct {
	Rank            int         `json:"rank"`
	Points          int         `json:"points"`
	RootOwns        int         `json:"root_owns"`
	UserOwns        int         `json:"user_owns"`
	ChallengeOwns   int         `json:"challenge_owns"`
	RootBloods      int         `json:"root_bloods"`
	UserBloods      int         `json:"user_bloods"`
	ChallengeBloods int         `json:"challenge_bloods"`
	Fortress        int         `json:"fortress"`
	Endgame         int         `json:"endgame"`
	ID              int         `json:"id"`
	Name            string      `json:"name"`
	Thumb           string      `json:"avatar_thumb_url"`
	Country         string      `json:"country"`
	RanksDiff       int         `json:"ranks_diff"`
}

// RanksOfUsersList contains a list of user ranks
type RanksOfUsersList struct {
	Status bool                    `json:"status"`
	Data   []*RanksOfUsersListItem `json:"data"`
}

// RanksOfUsersListItem contains information on a users rank
type RanksOfUsersListItem struct {
	Rank            int         `json:"rank"`
	Points          int         `json:"points"`
	RootOwns        int         `json:"root_owns"`
	UserOwns        int         `json:"user_owns"`
	ChallengeOwns   int         `json:"challenge_owns"`
	RootBloods      int         `json:"root_bloods"`
	UserBloods      int         `json:"user_bloods"`
	ChallengeBloods int         `json:"challenge_bloods"`
	Fortress        int         `json:"fortress"`
	Endgame         int         `json:"endgame"`
	ID              int         `json:"id"`
	Name            string      `json:"name"`
	Thumb           string      `json:"avatar_thumb"`
	Country         string `json:"country"`
	Level           string      `json:"level"`
	RanksDiff       int `json:"ranks_diff"`
}

// RanksOfUniversitiesList contains a list of university ranks
type RanksOfUniversitiesList struct {
	Status bool                           `json:"status"`
	Data   []*RanksOfUniversitiesListItem `json:"data"`
}

// RanksOfUniversitiesListItem contains information on a universities rank
type RanksOfUniversitiesListItem struct {
	Rank            int    `json:"rank"`
	Students        int    `json:"students"`
	Points          int    `json:"points"`
	RootOwns        int    `json:"root_owns"`
	UserOwns        int    `json:"user_owns"`
	ChallengeOwns   int    `json:"challenge_owns"`
	RootBloods      int    `json:"root_bloods"`
	UserBloods      int    `json:"user_bloods"`
	ChallengeBloods int    `json:"challenge_bloods"`
	Fortress        int    `json:"fortress"`
	Endgame         int    `json:"endgame"`
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Country         string `json:"country"`
	RanksDiff       int    `json:"ranks_diff"`
}

// RankInCountryBestDuring returns the Token Users countries best rank within a period of time, given a duration
func (s *Session) RankInCountryBestDuring(period Duration) (rank *RankInCountryBestDuringInfo, err error) {

	var url string = "https://www.hackthebox.com/api/v4/rankings/country/best?period=" + string(period)
	err = parseJSON(s, url, &rank)

	return
}

// RankInCountryDuring returns the Token Users countries rank info within a period of time, given a duration
func (s *Session) RankInCountryDuring(period Duration) (rank *RankInCountryDuringInfo, err error) {

	var url string = "https://www.hackthebox.com/api/v4/rankings/country/overview?period=" + string(period)
	err = parseJSON(s, url, &rank)

	return
}

// RankBracketInCountry returns the Token Users countries rank bracket information
func (s *Session) RankBracketInCountry() (bracket *RankBracketInCountryInfo, err error) {

	var url string = "https://www.hackthebox.com/api/v4/rankings/country/ranking_bracket"
	err = parseJSON(s, url, &bracket)

	return
}

// RankOfTeamBestDuring returns the Token Users teams best rank within a period of time, given a duration
func (s *Session) RankOfTeamBestDuring(period Duration) (rank *RankOfTeamBestDuringInfo, err error) {

	var url string = "https://www.hackthebox.com/api/v4/rankings/team/best?period=" + string(period)
	err = parseJSON(s, url, &rank)

	return
}

// RankOfTeamDuring returns the Token Users team rank info within a period of time, given a duration
func (s *Session) RankOfTeamDuring(period Duration) (info *RankOfTeamDuringInfo, err error) {

	var url string = "https://www.hackthebox.com/api/v4/rankings/team/overview?period=" + string(period)
	err = parseJSON(s, url, &info)

	return
}

// RankBracketOfTeam returns the Token Users teams rank bracket information
func (s *Session) RankBracketOfTeam() (bracket *RankBracketOfTeamInfo, err error) {

	var url string = "https://www.hackthebox.com/api/v4/rankings/team/ranking_bracket"
	err = parseJSON(s, url, &bracket)

	return
}

// RankBestDuring returns the Token Users best rank within a period of time, given a duration and indication of VIP info wanting to be returned
func (s *Session) RankBestDuring(period Duration, wantVIP bool) (rank *RankBestDuringInfo, err error) {

	var url string = "https://www.hackthebox.com/api/v4/rankings/user/best?period=" + string(period) + "&vip=" + stringFromVIP(wantVIP)
	err = parseJSON(s, url, &rank)

	return
}

// RankDuring returns the Token Users rank info within a period of time, given a duration
func (s *Session) RankDuring(period Duration, wantVIP bool) (rank *RankDuringInfo, err error) {

	var url string = "https://www.hackthebox.com/api/v4/rankings/user/overview?period=" + string(period) + "&vip=" + stringFromVIP(wantVIP)
	err = parseJSON(s, url, &rank)

	return
}

// RankBracket returns the Token Users rank bracket information
func (s *Session) RankBracket(wantVIP bool) (bracket *RankBracketInfo, err error) {

	var url string = "https://www.hackthebox.com/api/v4/rankings/user/ranking_bracket?vip=" + stringFromVIP(wantVIP)
	err = parseJSON(s, url, &bracket)

	return
}

// RanksOfCountries returns a list of top countries and their ranking information
func (s *Session) RanksOfCountries() (ranks *RanksOfCountriesList, err error) {

	var url string = "https://www.hackthebox.com/api/v4/rankings/countries"
	err = parseJSON(s, url, &ranks)

	return
}

// RanksOfTeams returns a list of top teams and their ranking information
func (s *Session) RanksOfTeams() (ranks *RanksOfTeamsList, err error) {

	var url string = "https://www.hackthebox.com/api/v4/rankings/teams"
	err = parseJSON(s, url, &ranks)

	return
}

// RanksOfUsers returns a list of top users and their ranking information
func (s *Session) RanksOfUsers(wantVIP bool) (ranks *RanksOfUsersList, err error) {

	var url string = "https://www.hackthebox.com/api/v4/rankings/users?vip=" + stringFromVIP(wantVIP)
	err = parseJSON(s, url, &ranks)

	return
}

// RanksOfUniversities returns a list of top universities and their ranking information
func (s *Session) RanksOfUniversities() (universities *RanksOfUniversitiesList, err error) {

	var url string = "https://www.hackthebox.com/api/v4/rankings/universities"
	err = parseJSON(s, url, &universities)

	return
}