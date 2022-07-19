package htbgo

import "strconv"

// Get a list of universities
type Universities struct {
	Message string `json:"message"`
	Section string `json:"section"`
	Data    struct {
		University []struct {
			ID                   int      `json:"id"`
			Name                 string   `json:"name"`
			HasAutoGeneratedLogo int      `json:"has_auto_generated_logo"`
			CCA2                 string   `json:"cca2"`
			CreatedAt            string   `json:"created_at"`
			UsersCount           int      `json:"users_count"`
			RespectedByCount     int      `json:"respected_by_count"`
			Country              string   `json:"country"`
			UsersAvatartThumb    []string `json:"user_avatars"`
		} `json:"data"`
		NextPage     string `json:"next_page_url"`
		Path         string `json:"path"`
		PerPage      int    `json:"per_page"`
		PreviousPage string `json:"prev_page_url"`
		To           int    `json:"to"`
		Total        int    `json:"total"`
	}
}

// get a profile of a university by id
type University struct {
	Message string `json:"message"`
	Data    struct {
		ID                   int    `json:"id"`
		Name                 string `json:"name"`
		Points               int    `json:"points"`
		Motto                string `json:"motto"`
		Description          string `json:"description"`
		CountryName          string `json:"country_name"`
		CountryCode          string `json:"country_code"`
		HasAutoGeneratedLogo int    `json:"has_auto_generated_logo"`
		JoinRequestSent      bool   `json:"join_request_sent"`
		IsRespected          bool   `json:"is_respected"`
		URL                  string `json:"url"`
		Twitter              string `json:"twitter"`
		Discord              string `json:"discord"`
		LogoURL              string `json:"logo_url"`
		CoverImageURL        string `json:"cover_image_url"`
		Active               int    `json:"active"`
		Captain              struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			AvatarThumb string `json:"avatar_thumb"`
		} `json:"captain"`
	} `json:"data"`
}

// get stats of a university by id
// fixing weekly object later
type week struct {
	UserOwns      int    `json:"user_owns"`
	SystemOwns    int    `json:"system_owns"`
	ChallengeOwns int    `json:"challenge_owns"`
	FirstBloods   int    `json:"first_bloods"`
	Respects      int    `json:"respects"`
	WeekEndDate   string `json:"week_end_date"`
}
type Weeks struct {
	Week1  week `json:"week1"`
	Week2  week `json:"week2"`
	Week3  week `json:"week3"`
	Week4  week `json:"week4"`
	Week5  week `json:"week5"`
	Week6  week `json:"week6"`
	Week7  week `json:"week7"`
	Week8  week `json:"week8"`
	Week9  week `json:"week9"`
	Week10 week `json:"week10"`
	Week11 week `json:"week11"`
	Week12 week `json:"week12"`
}
type UniversityStats struct {
	Rank          int   `json:"rank"`
	UserOwns      int   `json:"user_owns"`
	SystemOwns    int   `json:"system_owns"`
	FirstBloods   int   `json:"first_bloods"`
	ChallengeOwns int   `json:"challenge_owns"`
	Respects      int   `json:"respects"`
	Weeks         Weeks `json:"weekly"`
}

// List University Members
// will fix Rank later because rank is int+string
type UniversityMembers []struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	// Rank            int    `json:"rank"`
	Points          int    `json:"points"`
	RootOwns        int    `json:"root_owns"`
	RootBloodsCount int    `json:"root_bloods_count"`
	UserBloodsCount int    `json:"user_bloods_count"`
	UserOwns        int    `json:"user_owns"`
	RankText        string `json:"rank_text"`
	CountryName     string `json:"country_name"`
	CountryCode     string `json:"country_code"`
	Role            string `json:"role"`
	University      struct {
		ID        int `json:"id"`
		CaptainID int `json:"captain_id"`
	} `json:"university"`
	Public int `json:"public"`
}

// will fix the query later
func (s *Session) Universities(page string, query string) (universities Universities) {
	if page == "" {
		page = "1"
	}
	var url string = "https://www.hackthebox.com/api/v4/university/all/list?page=" + page + "&search=" + query
	parseJSON(s, url, &universities)

	return
}

func (s *Session) University(id int) (university University) {
	var url string = "https://www.hackthebox.com/api/v4/university/profile/" + strconv.Itoa(id)
	parseJSON(s, url, &university)

	return
}

func (s *Session) UniversityStats(id int) (stats UniversityStats) {
	var url string = "https://www.hackthebox.com/api/v4/university/stats/owns/" + strconv.Itoa(id)
	parseJSON(s, url, &stats)

	return
}

func (s *Session) UniversityMembers(id int) (members UniversityMembers) {
	var url string = "https://www.hackthebox.com/api/v4/university/members/" + strconv.Itoa(id)
	parseJSON(s, url, &members)

	return
}
