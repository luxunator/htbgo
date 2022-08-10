package htbgo

// ChallengeCard contains information of a challenge as a card
type ChallengeCard struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	UrlName      string `json:"url_name"`
	Difficulty   string `json:"difficulty"`
	CategoryName string `json:"category_name"`
	Retired      int    `json:"retired"`
	ReleaseDate  string `json:"release_date"`
}

// ChallengesSuggestedInfo contains information on suggested challenges
type ChallengesSuggestedInfo struct {
	State   []string       `json:"state"`
	CardOne *ChallengeCard `json:"card1"`
	CardTwo *ChallengeCard `json:"card2"`
}

// ChallengesRetiredSuggestedInfo contains information on suggested retired challenges
type ChallengesRetiredSuggestedInfo struct {
	CardOne *ChallengeCard `json:"card1"`
	CardTwo *ChallengeCard `json:"card2"`
}

// ChallengeCategoriesList contains a list of challenge categories
type ChallengeCategoriesList struct {
	Info []*ChallengeCategoriesListItem `json:"info"`
}

// ChallengeCategoriesListItem contains a challenge category
type ChallengeCategoriesListItem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// ChallengesList contains a list of challenges
type ChallengesList struct {
	Challenges []*ChallengeListItem `json:"challenges"`
}

// ChallengeListItem contains information on a challenge within a challenge list
type ChallengeListItem struct {
	ID                 int           `json:"id"`
	Name               string        `json:"name"`
	URLName            string        `json:"url_name"`
	Retired            int           `json:"retired"`
	Difficulty         string        `json:"difficulty"`
	DifficultyAvg      int           `json:"avg_difficulty"`
	Points             interface{}   `json:"points"`
	StaticPoints       string        `json:"static_points"`
	DifficultyStats    *Difficulties `json:"difficulty_chart"`
	DifficultyStatsArr []int         `json:"difficulty_chart_arr"`
	Solves             int           `json:"solves"`
	Likes              int           `json:"likes"`
	Dislikes           int           `json:"dislikes"`
	ReleaseDate        string        `json:"release_date"`
	IsCompleted        bool          `json:"isCompleted"`
	CategoryID         int           `json:"challenge_category_id"`
	IsLiked            bool          `json:"likeByAuthUser"`
	IsDisliked         bool          `json:"dislikeByAuthUser"`
	IsSolved           bool          `json:"authUserSolve"`
	IsActive           bool          `json:"isActive"`
	IsTodo             bool          `json:"isTodo"`
	Recommended        int           `json:"recommended"`
}

// ChallengeInfo contains a challenge
type ChallengeInfo struct {
	Challenge *ChallengeInfoItem `json:"challenge"`
}

// ChallengeInfoItem contains information on a challenge
type ChallengeInfoItem struct {
	ID                    int           `json:"id"`
	Name                  string        `json:"name"`
	Retired               int           `json:"retired"`
	Difficulty            string        `json:"difficulty"`
	Points                interface{}   `json:"points"` // string or int 0
	DifficultyStats       *Difficulties `json:"difficulty_chart"`
	Solves                int           `json:"solves"`
	SolveTime             string        `json:"authUserSolveTime"`
	Likes                 int           `json:"likes"`
	Dislikes              int           `json:"dislikes"`
	Description           string        `json:"description"`
	CategoryName          string        `json:"category_name"`
	BloodUserID           int           `json:"first_blood_user_id"`
	BloodUser             string        `json:"first_blood_user"`
	BloodTime             string        `json:"first_blood_time"`
	BloodUserAvatar       string        `json:"first_blood_user_avatar"`
	CreatorID             int           `json:"creator_id"`
	CreatorName           string        `json:"creator_name"`
	CreatorAvatar         string        `json:"creator_avatar"`
	CreatorIsRespected    bool          `json:"isRespected"`
	CreatorTwoID          int           `json:"creator2_id"`
	CreatorTwoName        string        `json:"creator2_name"`
	CreatorTwoAvatar      string        `json:"creator2_avatar"`
	CreatorTwoIsRespected string        `json:"isRespected2"`
	HasDownload           bool          `json:"download"`
	SHA256                string        `json:"sha256"`
	HasDocker             bool          `json:"docker"`
	DockerIP              string        `json:"docker_ip"`
	DockerPort            int           `json:"docker_port"`
	ReleaseDate           string        `json:"release_date"`
	Released              int           `json:"released"`
	IsSolved              bool          `json:"authUserSolve"`
	IsLiked               bool          `json:"likeByAuthUser"`
	IsDisliked            bool          `json:"dislikeByAuthUser"`
	IsTodo                bool          `json:"isTodo"`
	Recommended           int           `json:"recommended"`
	HasReviewed           bool          `json:"authUserHasReviewed"`
}

// ChallengeActivityList contains challenge activities
type ChallengeActivityList struct {
	Info *ChallengeActivityListItem `json:"info"`
}

// ChallengeActivityList contains a list of challenge activity
type ChallengeActivityListItem struct {
	Activity []*ChallengeActivityListItemActivity `json:"activity"`
}

// ChallengeActivityListItemActivity contains information on a challenge activity history item
type ChallengeActivityListItemActivity struct {
	CreatedAt  string `json:"created_at"`
	Date       string `json:"date"`
	DateDiff   string `json:"date_diff"`
	UserID     int    `json:"user_id"`
	UserName   string `json:"user_name"`
	UserAvatar string `json:"user_avatar"`
	Type       string `json:"type"`
}

// ChallengesSuggested returns the currently recommended challenges
func (s *Session) ChallengesSuggested() (cards *ChallengesSuggestedInfo, err error) {

	var url string = "https://www.hackthebox.com/api/v4/challenge/recommended"
	err = parseJSON(s, url, &cards)

	return
}

// ChallengesRetiredSuggested returns the currently recommended challenges that are retired
func (s *Session) ChallengesRetiredSuggested() (cards *ChallengesRetiredSuggestedInfo, err error) {

	var url string = "https://www.hackthebox.com/api/v4/challenge/recommended/retired"
	err = parseJSON(s, url, &cards)

	return
}

// ChallengeCategories returns a list of the current challenge categories
func (s *Session) ChallengeCategories() (categories *ChallengeCategoriesList, err error) {

	var url string = "https://www.hackthebox.com/api/v4/challenge/categories/list"
	err = parseJSON(s, url, &categories)

	return
}

// ChallengesActive returns a list of the currently active challenges
func (s *Session) ChallengesActive() (challenges *ChallengesList, err error) {

	var url string = "https://www.hackthebox.com/api/v4/challenge/list"
	err = parseJSON(s, url, &challenges)

	return
}

// ChallengesRetired returns a list of all of the retired challenges
func (s *Session) ChallengesRetired() (challenges *ChallengesList, err error) {

	var url string = "https://www.hackthebox.com/api/v4/challenge/list/retired"
	err = parseJSON(s, url, &challenges)

	return
}

// Challenge returns the information of a challenge, given a challenge ID
func (s *Session) Challenge(challengeID int) (info *ChallengeInfo, err error) {

	challengeIDString, err := toPositiveIntString(challengeID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/challenge/info/" + challengeIDString
	err = parseJSON(s, url, &info)

	return
}

// ChallengeActivity returns the recent activity of users on a challenge, given a challenge ID
func (s *Session) ChallengeActivity(challengeID int) (activities *ChallengeActivityList, err error) {

	challengeIDString, err := toPositiveIntString(challengeID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/challenge/activity/" + challengeIDString
	err = parseJSON(s, url, &activities)

	return
}
