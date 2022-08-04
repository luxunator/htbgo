package htbgo

type ChallengeCard struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	UrlName      string `json:"url_name"`
	Difficulty   string `json:"difficulty"`
	CategoryName string `json:"category_name"`
	Retired      int    `json:"retired"`
	ReleaseDate  string `json:"release_date"`
}

// Challenge Recommendation Cards
// https://www.hackthebox.com/api/v4/challenge/recommended

type ChallengesSuggestedInfo struct {
	State   []string      `json:"state"`
	CardOne ChallengeCard `json:"card1"`
	CardTwo ChallengeCard `json:"card2"`
}

// Challenge Reccomendation Cards Retired
// https://www.hackthebox.com/api/v4/challenge/recommended/retired

type ChallengesRetiredSuggestedInfo struct {
	CardOne ChallengeCard `json:"card1"`
	CardTwo ChallengeCard `json:"card2"`
}

/* Not Implemented

// Challenge Reviews
// https://www.hackthebox.com/api/v4/challenge/reviews/{challengeID}

type UserChallengeReviews struct {

}

// Bearer Challenge Reviews
// https://www.hackthebox.com/api/v4/challenge/reviews/user/{challengeID}

type BearerChallengeReviews struct {

}

*/

// List Challenge Categories
// https://www.hackthebox.com/api/v4/challenge/categories/list

type ChallengeCategoriesList struct {
	Info []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"info"`
}

// List Active Challenges
// https://www.hackthebox.com/api/v4/challenge/list

// List Retired Challenges
// https://www.hackthebox.com/api/v4/challenge/list/retired

type ChallengesList struct {
	Challenges []struct {
		ID                 int          `json:"id"`
		Name               string       `json:"name"`
		URLName            string       `json:"url_name"`
		Retired            int          `json:"retired"`
		Difficulty         string       `json:"difficulty"`
		DifficultyAvg      int          `json:"avg_difficulty"`
		Points             interface{}  `json:"points"`
		StaticPoints       string       `json:"static_points"`
		DifficultyStats    Difficulties `json:"difficulty_chart"`
		DifficultyStatsArr []int        `json:"difficulty_chart_arr"`
		Solves             int          `json:"solves"`
		Likes              int          `json:"likes"`
		Dislikes           int          `json:"dislikes"`
		ReleaseDate        string       `json:"release_date"`
		IsCompleted        bool         `json:"isCompleted"`
		CategoryID         int          `json:"challenge_category_id"`
		IsLiked            bool         `json:"likeByAuthUser"`
		IsDisliked         bool         `json:"dislikeByAuthUser"`
		IsSolved           bool         `json:"authUserSolve"`
		IsActive           bool         `json:"isActive"`
		IsTodo             bool         `json:"isTodo"`
		Recommended        int          `json:"recommended"`
	} `json:"challenges"`
}

// Challenge Profile
// https://www.hackthebox.com/api/v4/challenge/info/{challengeID}

type ChallengeInfo struct {
	Challenge struct {
		ID                    int          `json:"id"`
		Name                  string       `json:"name"`
		Retired               int          `json:"retired"`
		Difficulty            string       `json:"difficulty"`
		Points                string       `json:"points"`
		DifficultyStats       Difficulties `json:"difficulty_chart"`
		Solves                int          `json:"solves"`
		SolveTime             string       `json:"authUserSolveTime"`
		Likes                 int          `json:"likes"`
		Dislikes              int          `json:"dislikes"`
		Description           string       `json:"description"`
		CategoryName          string       `json:"category_name"`
		BloodUserID           int          `json:"first_blood_user_id"`
		BloodUser             string       `json:"first_blood_user"`
		BloodTime             string       `json:"first_blood_time"`
		BloodUserAvatar       string       `json:"first_blood_user_avatar"`
		CreatorID             int          `json:"creator_id"`
		CreatorName           string       `json:"creator_name"`
		CreatorAvatar         string       `json:"creator_avatar"`
		CreatorIsRespected    bool         `json:"isRespected"`
		CreatorTwoID          int          `json:"creator2_id"`
		CreatorTwoName        string       `json:"creator2_name"`
		CreatorTwoAvatar      string       `json:"creator2_avatar"`
		CreatorTwoIsRespected bool         `json:"isRespected2"`
		HasDownload           bool         `json:"download"`
		SHA256                string       `json:"sha256"`
		HasDocker             bool         `json:"docker"`
		DockerIP              string       `json:"docker_ip"`
		DockerPort            int          `json:"docker_port"`
		ReleaseDate           string       `json:"release_date"`
		Released              int          `json:"released"`
		IsSolved              bool         `json:"authUserSolve"`
		IsLiked               bool         `json:"likeByAuthUser"`
		IsDisliked            bool         `json:"dislikeByAuthUser"`
		IsTodo                bool         `json:"isTodo"`
		Recommended           int          `json:"recommended"`
		HasReviewed           bool         `json:"authUserHasReviewed"`
	} `json:"challenge"`
}

// Challenge Activity
// https://www.hackthebox.com/api/v4/challenge/activity/{challengeID}

type ChallengeActivityList struct {
	Info struct {
		Activity []struct {
			CreatedAt  string `json:"created_at"`
			Date       string `json:"date"`
			DateDiff   string `json:"date_diff"`
			UserID     int    `json:"user_id"`
			UserName   string `json:"user_name"`
			UserAvatar string `json:"user_avatar"`
			Type       string `json:"type"`
		} `json:"activity"`
	} `json:"info"`
}

/* Not Implemented

// Changelog
// https://www.hackthebox.com/api/v4/challenge/reviews/{challengeID}

type ChangeLogList struct {

}

*/

/* TODO

// Challenge Download (ZIP Blob)
// https://www.hackthebox.com/api/v4/challenge/download/{challengeID}

*/

func (s *Session) ChallengesSuggested() (cards ChallengesSuggestedInfo, err error) {

	var url string = "https://www.hackthebox.com/api/v4/challenge/recommended"
	err = parseJSON(s, url, &cards)

	return
}

func (s *Session) ChallengesRetiredSuggested() (cards ChallengesRetiredSuggestedInfo, err error) {

	var url string = "https://www.hackthebox.com/api/v4/challenge/recommended/retired"
	err = parseJSON(s, url, &cards)

	return
}

func (s *Session) ChallengeCategories() (categories ChallengeCategoriesList, err error) {

	var url string = "https://www.hackthebox.com/api/v4/challenge/categories/list"
	err = parseJSON(s, url, &categories)

	return
}

func (s *Session) ChallengesActive() (challenges ChallengesList, err error) {

	var url string = "https://www.hackthebox.com/api/v4/challenge/list"
	err = parseJSON(s, url, &challenges)

	return
}

func (s *Session) ChallengesRetired() (challenges ChallengesList, err error) {

	var url string = "https://www.hackthebox.com/api/v4/challenge/list/retired"
	err = parseJSON(s, url, &challenges)

	return
}

func (s *Session) Challenge(challengeID int) (info ChallengeInfo, err error) {

	challengeIDString, err := toPositiveIntString(challengeID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/challenge/info/" + challengeIDString
	err = parseJSON(s, url, &info)

	return
}

func (s *Session) ChallengeActivity(challengeID int) (activities ChallengeActivityList, err error) {

	challengeIDString, err := toPositiveIntString(challengeID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/challenge/activity/" + challengeIDString
	err = parseJSON(s, url, &activities)

	return
}
