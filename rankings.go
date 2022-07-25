package htbgo

// Get the user's best rank within the period
// Week - W 	Month - M 		Year -Y
// https://www.hackthebox.com/api/v4/rankings/country/best?period={period}

type UserRankInCountry struct {
	Status bool `json:"status"`
	Data   struct {
		Rank          int    `json:"rank"`
		Date          string `json:"date"`
		RankChartData []int  `json:"rank_chart_data"`
	} `json:"data"`
}

// Get the overview of the user's rank in the country
// Week - W 	Month - M 		Year -Y
// https://www.hackthebox.com/api/v4/rankings/country/overview?period=1Y

type UserRankOverviewInCountry struct {
	Status bool `json:"status"`
	Data   struct {
		PointsDiff   int    `json:"points_diff"`
		PointsGrowth string `json:"points_growth"`
		RanksDiff    int    `json:"ranks_diff"`
		ChartData    []int  `json:"chart_data"`
		Country      struct {
			Code string `json:"code"`
			Name string `json:"name"`
		} `json:"country"`
	} `json:"data"`
}

// Get the current user's rank bracket in their country
// https://www.hackthebox.com/api/v4/rankings/country/ranking_bracket

type UserRankBracketInCountry struct {
	Status bool `json:"status"`
	Data   struct {
		Rank                 int    `json:"rank"`
		Points               int    `json:"points"`
		PointsForNextBracket int    `json:"points_for_next_bracket"`
		CurrentBracket       string `json:"current_bracket"`
		NextBracket          string `json:"next_bracket"`
	} `json:"data"`
}

// Get the user's current team their best rank within a period
// Week - W 	Month - M 		Year -Y
// https://www.hackthebox.com/api/v4/rankings/team/best?period=1Y

type TeamRank struct {
	Status bool `json:"status"`
	Data   struct {
		Rank          int    `json:"rank"`
		Date          string `json:"date"`
		RankChartData []int  `json:"rank_chart_data"`
	} `json:"data"`
}

// Get the user's current team their rank overview
// Week -W 		Month - M 		Year - Y
// https://www.hackthebox.com/api/v4/rankings/team/overview?period=1Y

type TeamRankOverview struct {
	Status bool `json:"status"`
	Data   struct {
		PointsDiff   int    `json:"points_diff"`
		PointsGrowth string `json:"points_growth"`
		RanksDiff    int    `json:"ranks_diff"`
		ChartData    []int  `json:"chart_data"`
		Team         struct {
			ID             int    `json:"id"`
			Name           string `json:"name"`
			AvatarURL      string `json:"avatar_url"`
			AvatarThumbURL string `json:"avatar_thumb_url"`
		} `json:"team"`
	} `json:"data"`
}

// Get the user's current team their rank bracket
// https://www.hackthebox.com/api/v4/rankings/country/ranking_bracket

type TeamRankBracket struct {
	Status bool `json:"status"`
	Data   struct {
		Rank                 int    `json:"rank"`
		Points               int    `json:"points"`
		PointsForNextBracket int    `json:"points_for_next_bracket"`
		CurrentBracket       string `json:"current_bracket"`
		NextBracket          string `json:"next_bracket"`
	} `json:"data"`
}

// Get the user's best rank in the general leaderboard for a certain period.
// https://www.hackthebox.com/api/v4/rankings/user/best?period={period}&vip={numberValue}

type UserRank struct {
	Status bool `json:"status"`
	Data   struct {
		Rank          interface{} `json:"rank"`
		Date          string      `json:"date"`
		RankChartData []int       `json:"rank_chart_data"`
	} `json:"data"`
}

// Get the user's rank overview in the general ledearboard for a certain period.
// https://www.hackthebox.com/api/v4/rankings/user/overview?period=1Y&vip=0

type UserOverview struct {
	Status bool `json:"status"`
	Data   struct {
		PointsDiff   int         `json:"points_diff"`
		PointsGrowth interface{} `json:"points_growth"`
		RanksDiff    int         `json:"ranks_diff"`
		ChartData    []int       `json:"chart_data"`
		User         struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Avatar      string `json:"avatar"`
			AvatarThumb string `json:"avatar_thumb"`
		} `json:"user"`
	} `json:"data"`
}

// Get the user's  current rank bracket in the general leaderboard
// https://www.hackthebox.com/api/v4/rankings/user/ranking_bracket?vip=0

type UserRankBracket struct {
	Status bool `json:"status"`
	Data   struct {
		Rank                 interface{} `json:"rank"`
		Points               int         `json:"points"`
		PointsForNextBracket int         `json:"points_for_next_bracket"`
		CurrentBracket       string      `json:"current_bracket"`
		NextBracket          string      `json:"next_bracket"`
	} `json:"data"`
}

// Get the Rankings of the countries
// https://www.hackthebox.com/api/v4/rankings/countries

type CountryRanks struct {
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

type TeamRanks struct {
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
		AvatarThumbURL  string `json:"avatar_thumb_url"`
		Country         string `json:"country"`
		RanksDiff       int    `json:"ranks_diff"`
	} `json:"data"`
}

// Get the rankings of the users
// https://www.hackthebox.com/api/v4/rankings/users?vip=1

type UserRanks struct {
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
		AvatarThumb     string `json:"avatar_thumb"`
		Country         string `json:"country"`
		Level           string `json:"level"`
		RanksDiff       int    `json:"ranks_diff"`
	} `json:"data"`
}

type UniversitiesRankingList struct {
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

func (s *Session) UserRankInCountry(period string) (rank UserRankInCountry) {

	var url string = "https://www.hackthebox.com/api/v4/rankings/country/best?period=" + period
	parseJSON(s, url, &rank)

	return
}

func (s *Session) UserRankOverviewInCountry(period string) (rankOverview UserRankOverviewInCountry) {

	var url string = "https://www.hackthebox.com/api/v4/rankings/country/overview?period=" + period
	parseJSON(s, url, &rankOverview)

	return
}

func (s *Session) UserRankBracketInCountry() (rankBracket UserRankBracketInCountry) {

	var url string = "https://www.hackthebox.com/api/v4/rankings/country/ranking_bracket"
	parseJSON(s, url, &rankBracket)

	return
}

func (s *Session) TeamRank(period string) (teamRank TeamRank) {
	var url string = "https://www.hackthebox.com/api/v4/rankings/team/best?period=" + period
	parseJSON(s, url, &teamRank)
	return
}

func (s *Session) TeamRankOverview(period string) (rankOverview TeamRankOverview) {
	var url string = "https://www.hackthebox.com/api/v4/rankings/team/overview?period=" + period
	parseJSON(s, url, &rankOverview)
	return
}

func (s *Session) TeamRankBracket() (rankBracket TeamRankBracket) {
	var url string = "https://www.hackthebox.com/api/v4/rankings/team/ranking_bracket"
	parseJSON(s, url, &rankBracket)
	return
}

func (s *Session) UserRank(period string, wantVIP bool) (rank UserRank) {

	var url string = "https://www.hackthebox.com/api/v4/rankings/user/best?period=" + period + "&vip=" + stringFromVIP(wantVIP)
	parseJSON(s, url, &rank)

	return
}

func (s *Session) UserOverview(period string, wantVIP bool) (rankOverview UserOverview) {

	var url string = "https://www.hackthebox.com/api/v4/rankings/user/overview?period=" + period + "&vip=" + stringFromVIP(wantVIP)
	parseJSON(s, url, &rankOverview)

	return
}

func (s *Session) UserRankBracket(wantVIP bool) (rankBracket UserRankBracket) {

	var url string = "https://www.hackthebox.com/api/v4/rankings/user/ranking_bracket?vip=" + stringFromVIP(wantVIP)
	parseJSON(s, url, &rankBracket)

	return
}

func (s *Session) CountryRanks() (ranks CountryRanks) {

	var url string = "https://www.hackthebox.com/api/v4/rankings/countries"
	parseJSON(s, url, &ranks)

	return
}

func (s *Session) TeamRanks() (ranks TeamRanks) {

	var url string = "https://www.hackthebox.com/api/v4/rankings/teams"
	parseJSON(s, url, &ranks)

	return
}

func (s *Session) UserRanks(wantVIP bool) (ranks UserRanks) {

	var url string = "https://www.hackthebox.com/api/v4/rankings/users?vip=" + stringFromVIP(wantVIP)
	parseJSON(s, url, &ranks)

	return
}

func (s *Session) UniversitiesRanking() (universities UniversitiesRankingList) {

	var url string = "https://www.hackthebox.com/api/v4/rankings/universities"
	parseJSON(s, url, &universities)

	return
}
