package htbgo

// Get a list of universities
type UniversitiesList struct {
	Message string                `json:"message"`
	Section string                `json:"section"`
	Data    *UniversitiesListData `json:"data"`
}

type UniversitiesListData struct {
	Page            int                         `json:"current_page"`
	University      *UniversitiesListUniversity `json:"data"`
	FirstPageURL    string                      `json:"first_page_url"`
	From            int                         `json:"from"`
	To              int                         `json:"to"`
	NextPageURL     string                      `json:"next_page_url"`
	LastPage        int                         `json:"last_page"`
	LastPageURL     string                      `json:"last_page_url"`
	Path            string                      `json:"path"`
	PerPage         int                         `json:"per_page"`
	PreviousPageURL string                      `json:"prev_page_url"`
	Links           *UniversitiesListLinks      `json:"links"`
	Total           int                         `json:"total"`
}

type UniversitiesListUniversity []struct {
	ID               int      `json:"id"`
	Name             string   `json:"name"`
	HasRandomLogo    int      `json:"has_auto_generated_logo"`
	CCA2             string   `json:"cca2"`
	CreatedAt        string   `json:"created_at"`
	UsersCount       int      `json:"users_count"`
	RespectedByCount int      `json:"respected_by_count"`
	Country          string   `json:"country"`
	UserAvatars      []string `json:"user_avatars"`
}

type UniversitiesListLinks []struct {
	URL      string `json:"url"`
	Label    string `json:"label"`
	IsActive bool   `json:"active"`
}

// get a profile of a university by id
type UniversityInfo struct {
	Message string              `json:"message"`
	Data    *UniversityInfoData `json:"data"`
}

type UniversityInfoData struct {
	ID              int                    `json:"id"`
	Name            string                 `json:"name"`
	Points          int                    `json:"points"`
	Motto           string                 `json:"motto"`
	Description     string                 `json:"description"`
	CountryName     string                 `json:"country_name"`
	CountryCode     string                 `json:"country_code"`
	HasRandomLogo   int                    `json:"has_auto_generated_logo"`
	JoinRequestSent bool                   `json:"join_request_sent"`
	IsRespected     bool                   `json:"is_respected"`
	URL             string                 `json:"url"`
	Twitter         string                 `json:"twitter"`
	Discord         string                 `json:"discord"`
	Logo            string                 `json:"logo_url"`
	Cover           string                 `json:"cover_image_url"`
	Active          int                    `json:"active"`
	Captain         *UniversityInfoCaptain `json:"captain"`
}

type UniversityInfoCaptain struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Thumb string `json:"avatar_thumb"`
}

// get stats of a university by id
type UniversityOwnsWeek struct {
	UserOwns      int    `json:"user_owns"`
	SystemOwns    int    `json:"system_owns"`
	ChallengeOwns int    `json:"challenge_owns"`
	FirstBloods   int    `json:"first_bloods"`
	Respects      int    `json:"respects"`
	WeekEndDate   string `json:"week_end_date"`
}
type UniversityOwnsWeeks struct {
	WeekOne    *UniversityOwnsWeek `json:"week1"`
	WeekTwo    *UniversityOwnsWeek `json:"week2"`
	WeekThree  *UniversityOwnsWeek `json:"week3"`
	WeekFour   *UniversityOwnsWeek `json:"week4"`
	WeekFive   *UniversityOwnsWeek `json:"week5"`
	WeekSix    *UniversityOwnsWeek `json:"week6"`
	WeekSeven  *UniversityOwnsWeek `json:"week7"`
	WeekEight  *UniversityOwnsWeek `json:"week8"`
	WeekNine   *UniversityOwnsWeek `json:"week9"`
	WeekTen    *UniversityOwnsWeek `json:"week10"`
	WeekEleven *UniversityOwnsWeek `json:"week11"`
	WeekTwelve *UniversityOwnsWeek `json:"week12"`
}
type UniversityOwnsInfo struct {
	Rank          int                  `json:"rank"`
	UserOwns      int                  `json:"user_owns"`
	SystemOwns    int                  `json:"system_owns"`
	Bloods        int                  `json:"first_bloods"`
	ChallengeOwns int                  `json:"challenge_owns"`
	Respects      int                  `json:"respects"`
	Weekly        *UniversityOwnsWeeks `json:"weekly"`
}

// List University Members
type UniversityMembersList []struct {
	ID          int                              `json:"id"`
	Name        string                           `json:"name"`
	Avatar      string                           `json:"avatar"`
	Rank        interface{}                      `json:"rank"`
	Points      int                              `json:"points"`
	RootOwns    int                              `json:"root_owns"`
	RootBloods  int                              `json:"root_bloods_count"`
	UserBloods  int                              `json:"user_bloods_count"`
	UserOwns    int                              `json:"user_owns"`
	RankText    string                           `json:"rank_text"`
	CountryName string                           `json:"country_name"`
	CountryCode string                           `json:"country_code"`
	Role        string                           `json:"role"`
	University  *UniversityMembersListUniversity `json:"university"`
	Public      int                              `json:"public"`
}

type UniversityMembersListUniversity struct {
	ID        int `json:"id"`
	CaptainID int `json:"captain_id"`
}

func (s *Session) UniversitiesByPage(page int) (universities *UniversitiesList, err error) {

	pageString, err := toPositiveIntString(page)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/university/all/list?page=" + pageString
	err = parseJSON(s, url, &universities)

	return
}

func (s *Session) UniversitiesSearch(query string, page int) (universities *UniversitiesList, err error) {

	pageString, err := toPositiveIntString(page)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/university/all/list?search=" + query + "&page=" + pageString
	err = parseJSON(s, url, &universities)

	return
}

func (s *Session) University(universityID int) (university *UniversityInfo, err error) {

	universityIDString, err := toPositiveIntString(universityID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/university/profile/" + universityIDString
	err = parseJSON(s, url, &university)

	return
}

func (s *Session) UniversityOwns(universityID int) (stats *UniversityOwnsInfo, err error) {

	universityIDString, err := toPositiveIntString(universityID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/university/stats/owns/" + universityIDString
	err = parseJSON(s, url, &stats)

	return
}

func (s *Session) UniversityMembers(universityID int) (members *UniversityMembersList, err error) {

	universityIDString, err := toPositiveIntString(universityID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/university/members/" + universityIDString
	err = parseJSON(s, url, &members)

	return
}
