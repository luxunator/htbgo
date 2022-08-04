package htbgo

type RankBracket struct {
	Rank           int    `json:"rank"`
	Points         int    `json:"points"`
	PointsToGo     int    `json:"points_for_next_bracket"`
	CurrentBracket string `json:"current_bracket"`
	NextBracket    string `json:"next_bracket"`
}

type RankDuring struct {
	Rank          int    `json:"rank"`
	Date          string `json:"date"`
	RankChart     []int  `json:"rank_chart_data"`
}

// Get the user's best rank within the period
// Week - W 	Month - M 		Year -Y
// https://www.hackthebox.com/api/v4/rankings/country/best?period={period}

type RankInCountryBestDuringInfo struct {
	Status bool       `json:"status"`
	Data   RankDuring `json:"data"`
}

// Get the overview of the user's rank in the country
// Week - W 	Month - M 		Year -Y
// https://www.hackthebox.com/api/v4/rankings/country/overview?period=1Y

type RankInCountryDuringInfo struct {
	Status bool `json:"status"`
	Data   struct {
		PointsDiff   int    `json:"points_diff"`
		PointsGrowth string `json:"points_growth"`
		RanksDiff    int    `json:"ranks_diff"`
		Chart        []int  `json:"chart_data"`
		Country      struct {
			Code string `json:"code"`
			Name string `json:"name"`
		} `json:"country"`
	} `json:"data"`
}

// Get the current user's rank bracket in their country
// https://www.hackthebox.com/api/v4/rankings/country/ranking_bracket

type RankBracketInCountryInfo struct {
	Status bool        `json:"status"`
	Data   RankBracket `json:"data"`
}

// Get the user's current team their best rank within a period
// Week - W 	Month - M 		Year -Y
// https://www.hackthebox.com/api/v4/rankings/team/best?period=1Y

type RankOfTeamBestDuringInfo struct {
	Status bool       `json:"status"`
	Data   RankDuring `json:"data"`
}

// Get the user's current team their rank overview
// Week -W 		Month - M 		Year - Y
// https://www.hackthebox.com/api/v4/rankings/team/overview?period=1Y

type RankOfTeamDuringInfo struct {
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

// Get the user's current team their rank bracket
// https://www.hackthebox.com/api/v4/rankings/country/ranking_bracket

type RankBracketOfTeamInfo struct {
	Status bool        `json:"status"`
	Data   RankBracket `json:"data"`
}

// Get the user's best rank in the general leaderboard for a certain period.
// https://www.hackthebox.com/api/v4/rankings/user/best?period={period}&vip={numberValue}

type RankBestDuringInfo struct {
	Status bool       `json:"status"`
	Data   RankDuring `json:"data"`
}

// Get the user's rank overview in the general ledearboard for a certain period.
// https://www.hackthebox.com/api/v4/rankings/user/overview?period=1Y&vip=0

type RankDuringInfo struct {
	Status bool `json:"status"`
	Data   struct {
		PointsDiff   int         `json:"points_diff"`
		PointsGrowth interface{} `json:"points_growth"`
		RanksDiff    int         `json:"ranks_diff"`
		Chart        []int       `json:"chart_data"`
		User         struct {
			ID     int    `json:"id"`
			Name   string `json:"name"`
			Avatar string `json:"avatar"`
			Thumb  string `json:"avatar_thumb"`
		} `json:"user"`
	} `json:"data"`
}

// Get the user's  current rank bracket in the general leaderboard
// https://www.hackthebox.com/api/v4/rankings/user/ranking_bracket?vip=0

type RankBracketInfo struct {
	Status bool        `json:"status"`
	Data   RankBracket `json:"data"`
}

// Get the Rankings of the countries
// https://www.hackthebox.com/api/v4/rankings/countries

type RanksOfCountriesList struct {
	Status bool `json:"status"`
	Data   []struct {
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
	} `json:"data"`
}

// Get the rankings of the teams
// https://www.hackthebox.com/api/v4/rankings/teams

type RanksOfTeamsList struct {
	Status bool `json:"status"`
	Data   []struct {
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
	} `json:"data"`
}

// Get the rankings of the users
// https://www.hackthebox.com/api/v4/rankings/users?vip=1

type RanksOfUsersList struct {
	Status bool `json:"status"`
	Data   []struct {
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
	} `json:"data"`
}

type RanksOfUniversitiesList struct {
	Status bool `json:"status"`
	Data   []struct {
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
	} `json:"data"`
}

func (s *Session) RankInCountryBestDuring(period Duration) (rank RankInCountryBestDuringInfo, err error) {

	var url string = "https://www.hackthebox.com/api/v4/rankings/country/best?period=" + string(period)
	err = parseJSON(s, url, &rank)

	return
}

func (s *Session) RankInCountryDuring(period Duration) (rank RankInCountryDuringInfo, err error) {

	var url string = "https://www.hackthebox.com/api/v4/rankings/country/overview?period=" + string(period)
	err = parseJSON(s, url, &rank)

	return
}

func (s *Session) RankBracketInCountry() (bracket RankBracketInCountryInfo, err error) {

	var url string = "https://www.hackthebox.com/api/v4/rankings/country/ranking_bracket"
	err = parseJSON(s, url, &bracket)

	return
}

func (s *Session) RankOfTeamBestDuring(period Duration) (rank RankOfTeamBestDuringInfo, err error) {

	var url string = "https://www.hackthebox.com/api/v4/rankings/team/best?period=" + string(period)
	err = parseJSON(s, url, &rank)

	return
}

func (s *Session) RankOfTeamDuring(period Duration) (info RankOfTeamDuringInfo, err error) {

	var url string = "https://www.hackthebox.com/api/v4/rankings/team/overview?period=" + string(period)
	err = parseJSON(s, url, &info)

	return
}

func (s *Session) RankBracketOfTeam() (bracket RankBracketOfTeamInfo, err error) {

	var url string = "https://www.hackthebox.com/api/v4/rankings/team/ranking_bracket"
	err = parseJSON(s, url, &bracket)

	return
}

func (s *Session) RankBestDuring(period Duration, wantVIP bool) (rank RankBestDuringInfo, err error) {

	var url string = "https://www.hackthebox.com/api/v4/rankings/user/best?period=" + string(period) + "&vip=" + stringFromVIP(wantVIP)
	err = parseJSON(s, url, &rank)

	return
}

func (s *Session) RankDuring(period Duration, wantVIP bool) (rank RankDuringInfo, err error) {

	var url string = "https://www.hackthebox.com/api/v4/rankings/user/overview?period=" + string(period) + "&vip=" + stringFromVIP(wantVIP)
	err = parseJSON(s, url, &rank)

	return
}

func (s *Session) RankBracket(wantVIP bool) (bracket RankBracketInfo, err error) {

	var url string = "https://www.hackthebox.com/api/v4/rankings/user/ranking_bracket?vip=" + stringFromVIP(wantVIP)
	err = parseJSON(s, url, &bracket)

	return
}

func (s *Session) RanksOfCountries() (ranks RanksOfCountriesList, err error) {

	var url string = "https://www.hackthebox.com/api/v4/rankings/countries"
	err = parseJSON(s, url, &ranks)

	return
}

func (s *Session) RanksOfTeams() (ranks RanksOfTeamsList, err error) {

	var url string = "https://www.hackthebox.com/api/v4/rankings/teams"
	err = parseJSON(s, url, &ranks)

	return
}

func (s *Session) RanksOfUsers(wantVIP bool) (ranks RanksOfUsersList, err error) {

	var url string = "https://www.hackthebox.com/api/v4/rankings/users?vip=" + stringFromVIP(wantVIP)
	err = parseJSON(s, url, &ranks)

	return
}

func (s *Session) RanksOfUniversities() (universities RanksOfUniversitiesList, err error) {

	var url string = "https://www.hackthebox.com/api/v4/rankings/universities"
	err = parseJSON(s, url, &universities)

	return
}
