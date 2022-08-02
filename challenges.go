package htbgo

// Challenge Recommendation Cards
// https://www.hackthebox.com/api/v4/challenge/recommended

type ChallengesSuggestedInfo struct {
	State   []string `json:"state"`
	CardOne struct {
		ID           int    `json:"id"`
		Name         string `json:"name"`
		UrlName      string `json:"url_name"`
		Difficulty   string `json:"difficulty"`
		CategoryName string `json:"category_name"`
		Retired      int    `json:"retired"`
		ReleaseDate  string `json:"release_date"`
	} `json:"card1"`
	CardTwo struct {
		ID           int    `json:"id"`
		Name         string `json:"name"`
		UrlName      string `json:"url_name"`
		Difficulty   string `json:"difficulty"`
		CategoryName string `json:"category_name"`
		Retired      int    `json:"retired"`
		ReleaseDate  string `json:"release_date"`
	} `json:"card2"`
}

// Challenge Reccomendation Cards Retired
// https://www.hackthebox.com/api/v4/challenge/recommended/retired

type ChallengesRetiredSuggestedInfo struct {
	CardOne struct {
		ID           int    `json:"id"`
		Name         string `json:"name"`
		UrlName      string `json:"url_name"`
		Difficulty   string `json:"difficulty"`
		CategoryName string `json:"category_name"`
		Retired      int    `json:"retired"`
		ReleaseDate  string `json:"release_date"`
	} `json:"card1"`
	CardTwo struct {
		ID           int    `json:"id"`
		Name         string `json:"name"`
		UrlName      string `json:"url_name"`
		Difficulty   string `json:"difficulty"`
		CategoryName string `json:"category_name"`
		Retired      int    `json:"retired"`
		ReleaseDate  string `json:"release_date"`
	} `json:"card2"`
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

type ChallengesActiveList struct {
	Challenges []struct {
		ID              int    `json:"id"`
		Name            string `json:"name"`
		URLName         string `json:"url_name"`
		Retired         int    `json:"retired"`
		Difficulty      string `json:"difficulty"`
		AvgDifficulty   int    `json:"avg_difficulty"`
		Points          string `json:"points"`
		StaticPoints    string `json:"static_points"`
		DifficultyChart struct {
			CakeDifficulty      int `json:"counterCake"`
			VeryEasyDifficulty  int `json:"counterVeryEasy"`
			EasyDifficulty      int `json:"counterEasy"`
			TooEasyDifficulty   int `json:"counterTooEasy"`
			MediumDifficulty    int `json:"counterMedium"`
			BitHardDifficulty   int `json:"counterBitHard"`
			HardDifficulty      int `json:"counterHard"`
			TooHardDifficulty   int `json:"counterTooHard"`
			ExtraHardDifficulty int `json:"counterExHard"`
			BrainFuckDifficulty int `json:"counterBrainFuck"`
		} `json:"difficulty_chart"`
		DifficultyChartArr  []int  `json:"difficulty_chart_arr"`
		Solves              int    `json:"solves"`
		Likes               int    `json:"likes"`
		Dislikes            int    `json:"dislikes"`
		ReleaseDate         string `json:"release_date"`
		IsCompleted         bool   `json:"isCompleted"`
		ChallengeCategoryID int    `json:"challenge_category_id"`
		LikeByAuthUser      bool   `json:"likeByAuthUser"`
		DislikeByAuthUser   bool   `json:"dislikeByAuthUser"`
		AuthUserSolve       bool   `json:"authUserSolve"`
		IsActive            bool   `json:"isActive"`
		IsTodo              bool   `json:"isTodo"`
		Recommended         int    `json:"recommended"`
	} `json:"challenges"`
}

// List Retired Challenges
// https://www.hackthebox.com/api/v4/challenge/list/retired

type ChallengesRetiredList struct {
	Challenges []struct {
		ID              int    `json:"id"`
		Name            string `json:"name"`
		URLName         string `json:"url_name"`
		Retired         int    `json:"retired"`
		Difficulty      string `json:"difficulty"`
		AvgDifficulty   int    `json:"avg_difficulty"`
		Points          int    `json:"points"`
		StaticPoints    string `json:"static_points"`
		DifficultyChart struct {
			CakeDifficulty      int `json:"counterCake"`
			VeryEasyDifficulty  int `json:"counterVeryEasy"`
			EasyDifficulty      int `json:"counterEasy"`
			TooEasyDifficulty   int `json:"counterTooEasy"`
			MediumDifficulty    int `json:"counterMedium"`
			BitHardDifficulty   int `json:"counterBitHard"`
			HardDifficulty      int `json:"counterHard"`
			TooHardDifficulty   int `json:"counterTooHard"`
			ExtraHardDifficulty int `json:"counterExHard"`
			BrainFuckDifficulty int `json:"counterBrainFuck"`
		} `json:"difficulty_chart"`
		DifficultyChartArr  []int  `json:"difficulty_chart_arr"`
		Solves              int    `json:"solves"`
		Likes               int    `json:"likes"`
		Dislikes            int    `json:"dislikes"`
		ReleaseDate         string `json:"release_date"`
		IsCompleted         bool   `json:"isCompleted"`
		ChallengeCategoryID int    `json:"challenge_category_id"`
		LikeByAuthUser      bool   `json:"likeByAuthUser"`
		DislikeByAuthUser   bool   `json:"dislikeByAuthUser"`
		AuthUserSolve       bool   `json:"authUserSolve"`
		IsActive            bool   `json:"isActive"`
		IsTodo              bool   `json:"isTodo"`
		Recommended         int    `json:"recommended"`
	} `json:"challenges"`
}

// Challenge Profile
// https://www.hackthebox.com/api/v4/challenge/info/{challengeID}

type ChallengeInfo struct {
	Challenge struct {
		ID              int    `json:"id"`
		Name            string `json:"name"`
		Retired         int    `json:"retired"`
		Difficulty      string `json:"difficulty"`
		Points          string `json:"points"`
		DifficultyChart struct {
			CakeDifficulty      int `json:"counterCake"`
			VeryEasyDifficulty  int `json:"counterVeryEasy"`
			EasyDifficulty      int `json:"counterEasy"`
			TooEasyDifficulty   int `json:"counterTooEasy"`
			MediumDifficulty    int `json:"counterMedium"`
			BitHardDifficulty   int `json:"counterBitHard"`
			HardDifficulty      int `json:"counterHard"`
			TooHardDifficulty   int `json:"counterTooHard"`
			ExtraHardDifficulty int `json:"counterExHard"`
			BrainFuckDifficulty int `json:"counterBrainFuck"`
		} `json:"difficulty_chart"`
		Solves               int    `json:"solves"`
		AuthUserSolveTime    string `json:"authUserSolveTime"`
		Likes                int    `json:"likes"`
		Dislikes             int    `json:"dislikes"`
		Description          string `json:"description"`
		CategoryName         string `json:"category_name"`
		FirstBloodUserID     int    `json:"first_blood_user_id"`
		FirstBloodUser       string `json:"first_blood_user"`
		FirstBloodTime       string `json:"first_blood_time"`
		FirstBloodUserAvatar string `json:"first_blood_user_avatar"`
		CreatorID            int    `json:"creator_id"`
		CreatorName          string `json:"creator_name"`
		CreatorAvatar        string `json:"creator_avatar"`
		IsRespected          bool   `json:"isRespected"`
		CreatorTwoID         int    `json:"creator2_id"`
		CreatorTwoName       string `json:"creator2_name"`
		CreatorTwoAvatar     string `json:"creator2_avatar"`
		IsRespectedTwo       bool   `json:"isRespected2"`
		Download             bool   `json:"download"`
		SHA256               string `json:"sha256"`
		Docker               bool   `json:"docker"`
		DockerIP             string `json:"docker_ip"`
		DockerPort           int    `json:"docker_port"`
		ReleaseDate          string `json:"release_date"`
		Released             int    `json:"released"`
		AuthUserSolve        bool   `json:"authUserSolve"`
		LikeByAuthUser       bool   `json:"likeByAuthUser"`
		DislikeByAuthUser    bool   `json:"dislikeByAuthUser"`
		IsTodo               bool   `json:"isTodo"`
		Recommended          int    `json:"recommended"`
		AuthUserHasReviewed  bool   `json:"authUserHasReviewed"`
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

func (s *Session) ChallengesActive() (challenges ChallengesActiveList, err error) {

	var url string = "https://www.hackthebox.com/api/v4/challenge/list"
	err = parseJSON(s, url, &challenges)

	return
}

func (s *Session) ChallengesRetired() (challenges ChallengesRetiredList, err error) {

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