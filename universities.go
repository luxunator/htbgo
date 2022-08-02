package htbgo

// Get a list of universities
type UniversitiesList struct {
	Message string `json:"message"`
	Section string `json:"section"`
	Data    struct {
		Page       int `json:"current_page"`
		University []struct {
			ID                   int      `json:"id"`
			Name                 string   `json:"name"`
			HasAutoGeneratedLogo int      `json:"has_auto_generated_logo"`
			CCA2                 string   `json:"cca2"`
			CreatedAt            string   `json:"created_at"`
			UsersCount           int      `json:"users_count"`
			RespectedByCount     int      `json:"respected_by_count"`
			Country              string   `json:"country"`
			UserAvatars          []string `json:"user_avatars"`
		} `json:"data"`
		FirstPageURL    string `json:"first_page_url"`
		From            int    `json:"from"`
		To              int    `json:"to"`
		NextPageURL     string `json:"next_page_url"`
		LastPage        int    `json:"last_page"`
		LastPageURL     string `json:"last_page_url"`
		Path            string `json:"path"`
		PerPage         int    `json:"per_page"`
		PreviousPageURL string `json:"prev_page_url"`
		Links           []struct {
			URL    string `json:"url"`
			Label  string `json:"label"`
			Active bool   `json:"active"`
		} `json:"links"`
		Total int `json:"total"`
	} `json:"data"`
}

// get a profile of a university by id
type UniversityInfo struct {
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
type OwnWeek struct {
	UserOwns      int    `json:"user_owns"`
	SystemOwns    int    `json:"system_owns"`
	ChallengeOwns int    `json:"challenge_owns"`
	FirstBloods   int    `json:"first_bloods"`
	Respects      int    `json:"respects"`
	WeekEndDate   string `json:"week_end_date"`
}
type Weeks struct {
	Week1  OwnWeek `json:"week1"`
	Week2  OwnWeek `json:"week2"`
	Week3  OwnWeek `json:"week3"`
	Week4  OwnWeek `json:"week4"`
	Week5  OwnWeek `json:"week5"`
	Week6  OwnWeek `json:"week6"`
	Week7  OwnWeek `json:"week7"`
	Week8  OwnWeek `json:"week8"`
	Week9  OwnWeek `json:"week9"`
	Week10 OwnWeek `json:"week10"`
	Week11 OwnWeek `json:"week11"`
	Week12 OwnWeek `json:"week12"`
}
type UniversityOwnsInfo struct {
	Rank          int   `json:"rank"`
	UserOwns      int   `json:"user_owns"`
	SystemOwns    int   `json:"system_owns"`
	FirstBloods   int   `json:"first_bloods"`
	ChallengeOwns int   `json:"challenge_owns"`
	Respects      int   `json:"respects"`
	Weeks         Weeks `json:"weekly"`
}

// List University Members
type UniversityMembersList []struct {
	ID              int         `json:"id"`
	Name            string      `json:"name"`
	Avatar          string      `json:"avatar"`
	Rank            interface{} `json:"rank"`
	Points          int         `json:"points"`
	RootOwns        int         `json:"root_owns"`
	RootBloodsCount int         `json:"root_bloods_count"`
	UserBloodsCount int         `json:"user_bloods_count"`
	UserOwns        int         `json:"user_owns"`
	RankText        string      `json:"rank_text"`
	CountryName     string      `json:"country_name"`
	CountryCode     string      `json:"country_code"`
	Role            string      `json:"role"`
	University      struct {
		ID        int `json:"id"`
		CaptainID int `json:"captain_id"`
	} `json:"university"`
	Public int `json:"public"`
}

func (s *Session) UniversitiesByPage(page int) (universities UniversitiesList, err error) {

	pageString, err := toPositiveIntString(page)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/university/all/list?page=" + pageString
	err = parseJSON(s, url, &universities)

	return
}

func (s *Session) UniversitiesSearch(query string, page int) (universities UniversitiesList, err error) {

	pageString, err := toPositiveIntString(page)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/university/all/list?search=" + query + "&page=" + pageString
	err = parseJSON(s, url, &universities)

	return
}

func (s *Session) University(universityID int) (university UniversityInfo, err error) {

	universityIDString, err := toPositiveIntString(universityID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/university/profile/" + universityIDString
	err = parseJSON(s, url, &university)

	return
}

func (s *Session) UniversityOwns(universityID int) (stats UniversityOwnsInfo, err error) {

	universityIDString, err := toPositiveIntString(universityID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/university/stats/owns/" + universityIDString
	err = parseJSON(s, url, &stats)

	return
}

func (s *Session) UniversityMembers(universityID int) (members UniversityMembersList, err error) {

	universityIDString, err := toPositiveIntString(universityID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/university/members/" + universityIDString
	err = parseJSON(s, url, &members)

	return
}
