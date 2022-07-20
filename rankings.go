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
		} `json:"data"`
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
