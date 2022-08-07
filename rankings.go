package htbgo

// RankBracket
type RankBracket struct {
	Rank           int    `json:"rank"`
	Points         int    `json:"points"`
	PointsToGo     int    `json:"points_for_next_bracket"`
	CurrentBracket string `json:"current_bracket"`
	NextBracket    string `json:"next_bracket"`
}

// RankDuring
type RankDuring struct {
	Rank      int    `json:"rank"`
	Date      string `json:"date"`
	RankChart []int  `json:"rank_chart_data"`
}

// Get the user's best rank within the period
// Week - W 	Month - M 		Year -Y
// https://www.hackthebox.com/api/v4/rankings/country/best?period={period}

// RankInCountryBestDuringInfo
type RankInCountryBestDuringInfo struct {
	Status bool        `json:"status"`
	Data   *RankDuring `json:"data"`
}

// Get the overview of the user's rank in the country
// Week - W 	Month - M 		Year -Y
// https://www.hackthebox.com/api/v4/rankings/country/overview?period=1Y

// RankInCountryDuringInfo
type RankInCountryDuringInfo struct {
	Status bool                         `json:"status"`
	Data   *RankInCountryDuringInfoItem `json:"data"`
}

// RankInCountryDuringInfoItem
type RankInCountryDuringInfoItem struct {
	PointsDiff   int                                 `json:"points_diff"`
	PointsGrowth string                              `json:"points_growth"`
	RanksDiff    int                                 `json:"ranks_diff"`
	Chart        []int                               `json:"chart_data"`
	Country      *RankInCountryDuringInfoItemCountry `json:"country"`
}

// RankInCountryDuringInfoItemCountry
type RankInCountryDuringInfoItemCountry struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// Get the current user's rank bracket in their country
// https://www.hackthebox.com/api/v4/rankings/country/ranking_bracket

// RankBracketInCountryInfo
type RankBracketInCountryInfo struct {
	Status bool         `json:"status"`
	Data   *RankBracket `json:"data"`
}

// Get the user's current team their best rank within a period
// Week - W 	Month - M 		Year -Y
// https://www.hackthebox.com/api/v4/rankings/team/best?period=1Y

// RankOfTeamBestDuringInfo
type RankOfTeamBestDuringInfo struct {
	Status bool        `json:"status"`
	Data   *RankDuring `json:"data"`
}

// Get the user's current team their rank overview
// Week -W 		Month - M 		Year - Y
// https://www.hackthebox.com/api/v4/rankings/team/overview?period=1Y

// RankOfTeamDuringInfo
type RankOfTeamDuringInfo struct {
	Status bool                      `json:"status"`
	Data   *RankOfTeamDuringInfoItem `json:"data"`
}

// RankOfTeamDuringInfoItem
type RankOfTeamDuringInfoItem struct {
	PointsDiff   int                           `json:"points_diff"`
	PointsGrowth string                        `json:"points_growth"`
	RanksDiff    int                           `json:"ranks_diff"`
	Chart        []int                         `json:"chart_data"`
	Team         *RankOfTeamDuringInfoItemTeam `json:"team"`
}

// RankOfTeamDuringInfoItemTeam
type RankOfTeamDuringInfoItemTeam struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar_url"`
	Thumb  string `json:"avatar_thumb_url"`
}

// Get the user's current team their rank bracket
// https://www.hackthebox.com/api/v4/rankings/country/ranking_bracket

// RankBracketOfTeamInfo
type RankBracketOfTeamInfo struct {
	Status bool         `json:"status"`
	Data   *RankBracket `json:"data"`
}

// Get the user's best rank in the general leaderboard for a certain period.
// https://www.hackthebox.com/api/v4/rankings/user/best?period={period}&vip={numberValue}

// RankBestDuringInfo
type RankBestDuringInfo struct {
	Status bool        `json:"status"`
	Data   *RankDuring `json:"data"`
}

// Get the user's rank overview in the general ledearboard for a certain period.
// https://www.hackthebox.com/api/v4/rankings/user/overview?period=1Y&vip=0

// RankDuringInfo
type RankDuringInfo struct {
	Status bool                `json:"status"`
	Data   *RankDuringInfoItem `json:"data"`
}

// RankDuringInfoItem
type RankDuringInfoItem struct {
	PointsDiff   int                     `json:"points_diff"`
	PointsGrowth interface{}             `json:"points_growth"`
	RanksDiff    int                     `json:"ranks_diff"`
	Chart        []int                   `json:"chart_data"`
	User         *RankDuringInfoItemUser `json:"user"`
}

// RankDuringInfoItemUser
type RankDuringInfoItemUser struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Thumb  string `json:"avatar_thumb"`
}

// Get the user's  current rank bracket in the general leaderboard
// https://www.hackthebox.com/api/v4/rankings/user/ranking_bracket?vip=0

// RankBracketInfo
type RankBracketInfo struct {
	Status bool         `json:"status"`
	Data   *RankBracket `json:"data"`
}

// Get the Rankings of the countries
// https://www.hackthebox.com/api/v4/rankings/countries

// RanksOfCountriesList
type RanksOfCountriesList struct {
	Status bool                        `json:"status"`
	Data   []*RanksOfCountriesListItem `json:"data"`
}

// RanksOfCountriesListItem
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

// Get the rankings of the teams
// https://www.hackthebox.com/api/v4/rankings/teams

// RanksOfTeamsList
type RanksOfTeamsList struct {
	Status bool                    `json:"status"`
	Data   []*RanksOfTeamsListItem `json:"data"`
}

// RanksOfTeamsListItem
type RanksOfTeamsListItem struct {
	Rank            int    `json:"rank"`
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
	Thumb           string `json:"avatar_thumb_url"`
	Country         string `json:"country"`
	RanksDiff       int    `json:"ranks_diff"`
}

// Get the rankings of the users
// https://www.hackthebox.com/api/v4/rankings/users?vip=1

// RanksOfUsersList
type RanksOfUsersList struct {
	Status bool                    `json:"status"`
	Data   []*RanksOfUsersListItem `json:"data"`
}

// RanksOfUsersListItem
type RanksOfUsersListItem struct {
	Rank            int    `json:"rank"`
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
	Thumb           string `json:"avatar_thumb"`
	Country         string `json:"country"`
	Level           string `json:"level"`
	RanksDiff       int    `json:"ranks_diff"`
}

// RanksOfUniversitiesList
type RanksOfUniversitiesList struct {
	Status bool                           `json:"status"`
	Data   []*RanksOfUniversitiesListItem `json:"data"`
}

// RanksOfUniversitiesListItem
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

// RanksOfCountries returns a list of all the countries and their ranking information
func (s *Session) RanksOfCountries() (ranks *RanksOfCountriesList, err error) {

	var url string = "https://www.hackthebox.com/api/v4/rankings/countries"
	err = parseJSON(s, url, &ranks)

	return
}

// RanksOfTeams returns a list of all the teams and their ranking information
func (s *Session) RanksOfTeams() (ranks *RanksOfTeamsList, err error) {

	var url string = "https://www.hackthebox.com/api/v4/rankings/teams"
	err = parseJSON(s, url, &ranks)

	return
}

// RanksOfUsers returns a list of all the users and their ranking information
func (s *Session) RanksOfUsers(wantVIP bool) (ranks *RanksOfUsersList, err error) {

	var url string = "https://www.hackthebox.com/api/v4/rankings/users?vip=" + stringFromVIP(wantVIP)
	err = parseJSON(s, url, &ranks)

	return
}

// RanksOfUniversities returns a list of all the universities and their ranking information
func (s *Session) RanksOfUniversities() (universities *RanksOfUniversitiesList, err error) {

	var url string = "https://www.hackthebox.com/api/v4/rankings/universities"
	err = parseJSON(s, url, &universities)

	return
}
