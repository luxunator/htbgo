package htbgo

// Bearer Connection Status
// https://www.hackthebox.com/api/v4/user/connection/status

type BearerConnectionStatus struct {
	Status string `json:"status"`
	Connection interface{} `json:"connection"`
}

// Bearer Active Machine
// https://www.hackthebox.com/api/v4/machine/active

type BearerActiveMachine struct {
	Info struct {
		ID int `json:"id"`
		Name string `json:"name"`
		Avatar string `json:"avatar"`
		ExpiresAt string `json:"expires_at"`
		Voting bool `json:"voting"`
		Voted bool `json:"voted"`
		IsSpawning bool `json:"isSpawning"`
		Type string `json:"type"`
		LabServer string `json:"lab_server"`
	} `json:"info"`
}

// Bearer Followers
// https://www.hackthebox.com/api/v4/user/followers

type BearerFollowers struct {
	Info []struct {
		ID int `json:"id"`
	} `json:"info"`
}

// Bearer Profile
// https://www.hackthebox.com/api/v4/user/info

type BearerProfile struct {
	Info struct {
		ID int `json:"id"`
		Name string `json:"name"`
		Email string `json:"email"`
		Timezone string `json:"timezone"`
		IsVIP bool `json:"isVip"`
		IsModerator bool `json:"isModerator"`
		IsBGModerator bool `json:"isBGModerator"`
		IsChatBanned bool `json:"isChatBanned"`
		IsDedicatedVIP bool `json:"isDedicatedVip"`
		CanAccessVIP bool `json:"canAccessVIP"`
		CanAccessDediLab bool `json:"canAccessDedilab"`
		IsServerVIP bool `json:"isServerVIP"`
		ServerID int `json:"server_id"`
		Avatar string `json:"avatar"`
		BetaTester int `json:"beta_tester"`
		RankID int `json:"rank_id"`
		OnboardingCompleted bool `json:"onboarding_completed"`
		OnboardingTutorialComplete int `json:"onboarding_tutorial_complete"`
		Verified bool `json:"verified"`
		CanDeleteAvatar bool `json:"can_delete_avatar"`
		Team struct {
			ID int `json:"id"`
			Name string `json:"name"`
			AvatarThumb string `json:"avatar_thumb_url"`
		} `json:"team"`
		// University `json:"university"` TODO
		Identifier string `json:"identifier"`
		HasTeamInvitation bool `json:"hasTeamInvitation"`
		TwoFAEnabled bool `json:"TwoFaEnabled"`
		HasAppTokens bool `json:"hasAppTokens"`
		SubscriptionPlan string `json:"subscription_plan"`
		DunningExists bool `json:"dunning_exists"`
	} `json:"info"`
}

// Bearer Settings
// https://www.hackthebox.com/api/v4/user/settings

type BearerSettings struct {
	Email string `json:"email"`
	EmailNotifications int `json:"email_notifications"`
	Public int `json:"public"`
	NameChangeDelay int `json:"name_change_delay"`
	HideMachineTags int `json:"hide_machine_tags"`
}

// Bearer Subscriptions
// https://www.hackthebox.com/api/v4/user/subscriptions

type BearerSubscriptions struct {
	Subscriptions []struct {
		Name string `json:"name"`
		EndsAt string `json:"ends_at"`
		RenewsAt string `json:"renews_at"`
		RecurlyPlan string `json:"recurly_plan"`
		PausedAt string `json:"paused_at"`
		CreatedAt string `json:"created_at"`
		ResumeAt string `json:"resume_at"`
		StripePlan string `json:"stripe_plan"`
		// Items [] `json:"items"` TODO
	} `json:"subscriptions"`
}

// Bearer Subscriptions Balance
// https://www.hackthebox.com/api/v4/user/subscriptions/balance

type BearerSubscriptionsBalance struct {
	Balances struct {
		USD int `json:"USD"`
		GBP int `json:"GBP"`
		EUR int `json:"EUR"`
	} `json:"balances"`

	DefaultCurrency string `json:"default_currency"`
}

// Bearer Recurly URL
// https://www.hackthebox.com/api/v4/user/subscriptions/manage/recurly

type BearerSubscriptionsRecurly struct {
	Message string `json:"message"`
	URL string `json:"url"`
}

// Bearer Enrolled Tracks
// https://www.hackthebox.com/api/v4/user/tracks

type BearerTracks []struct {
	ID int `json:"id"`
	Complete int `json:"complete"`
}

/* TODO

// Bearer Machine Submissions
// https://www.hackthebox.com/api/v4/user/submissions/machines

type BearerMachineSubmissions struct {

}

// Bearer Challenge Submissions
// https://www.hackthebox.com/api/v4/user/submissions/challenges

type BearerChallengeSubmissions struct {

} 

*/

// User Profile with Bearer Relationship
// https://www.hackthebox.com/api/v4/user/profile/basic/{userID}

type BearerUserRelationship struct {
	Profile struct {
		ID int `json:"id"`
		SSOID int `json:"sso_id"`
		Name string `json:"name"`
		SystemOwns int `json:"system_owns"`
		UserOwns int `json:"user_owns"`
		UserBloods int `json:"user_bloods"`
		SystemBloods int `json:"system_bloods"`
		Team struct {
			ID int `json:"id"`
			Name string `json:"name"`
			Ranking int `json:"ranking"`
			Avatar string `json:"avatar"`
		} `json:"team"`
		Respects int `json:"respects"`
		Rank string `json:"rank"`
		RankID int `json:"rank_id"`
		CurrentRankProgress int `json:"current_rank_progress"`
		NextRank string `json:"next_rank"`
		NextRankPoints float64 `json:"next_rank_points"`
		RankOwnership string `json:"rank_ownership"`
		RankRequirement int `json:"rank_requirement"`
		Ranking int `json:"ranking"`
		Avatar string `json:"avatar"`
		Timezone string `json:"timezone"`
		IsVIP bool `json:"isVip"`
		IsDedicatedVip bool `json:"isDedicatedVip"`
		Public bool `json:"public"`
		CountryName string `json:"country_name"`
		CountryCode string `json:"country_code"`
		Points int `json:"points"`
		University struct {
			ID int `json:"id"`
			Name string `json:"name"`
			LogoThumbURL string `json:"logo_thumb_url"`
			Rank int `json:"rank"`
		} `json:"university"`
		UniversityName string `json:"university_name"`
		Description string `json:"description"`
		Github string `json:"github"`
		LinkedIn string `json:"linkedin"`
		Twitter string `json:"twitter"`
		Website string `json:"website"`
		IsRespected bool `json:"isRespected"`
		IsFollowed bool `json:"isFollowed"`
	}
}

// User Machines Progress
// https://www.hackthebox.com/api/v4/profile/progress/machines/os/{userID}

type UserProgressMachines struct {
	Profile struct {
		OperatingSystems []struct {
			Name string `json:"name"`
			CompletionPrecentage float64 `json:"completion_percentage"`
			OwnedMachines int `json:"owned_machines"`
			TotalMachines int `json:"total_machines"`
		} `json:"operating_systems"`
	} `json:"profile"`
}

// User Challenges Progress
// https://www.hackthebox.com/api/v4/profile/progress/challenges/{userID}

type UserProgressChallenges struct {
	Profile struct {
		ChallengeOwns struct {
			Solved int `json:"solved"`
			Total int `json:"total"`
		} `json:"challenge_owns"`
		ChallengeCategories []struct {
			Name string `json:"name"`
			OwnedFlags int `json:"owned_flags"`
			TotalFlags int `json:"total_flags"`
			CompletionPrecentage float64 `json:"completion_percentage"`
			AvgUserSolved float64 `json:"avg_user_solved"`
		} `json:"challenge_categories"`
	} `json:"profile"`
}

// User Endgames Progress
// https://www.hackthebox.com/api/v4/profile/progress/endgame/{userID}

type UserProgressEndgames struct {
	Profile struct {
		Endgames []struct {
			Name string `json:"name"`
			CompletionPrecentage float64 `json:"completion_percentage"`
			OwnedFlags int `json:"owned_flags"`
			TotalFlags int `json:"total_flags"`
		} `json:"endgames"`
	} `json:"profile"`
}

// User Fortress Progress
// https://www.hackthebox.com/api/v4/profile/progress/fortress/{userID}

type UserProgressFortresses struct {
	Profile struct {
		Fortresses []struct {
			Name string `json:"name"`
			Avatar string `json:"avatar"`
			CompletionPrecentage float64 `json:"completion_percentage"`
			OwnedFlags int `json:"owned_flags"`
			TotalFlags int `json:"total_flags"`
		} `json:"fortresses"`
	} `json:"profile"`
}

// User ProLabs Progress
// https://www.hackthebox.com/api/v4/profile/progress/prolab/{userID}

type UserProgressProLabs struct {
	Profile struct {
		ProLabs []struct {
			Name string `json:"name"`
			CompletionPrecentage float64 `json:"completion_percentage"`
			OwnedFlags int `json:"owned_flags"`
			TotalFlags int `json:"total_flags"`
			TotalMachines int `json:"total_machines"`
			AverageRatings float64 `json:"average_ratings"`
		} `json:"prolabs"`
	} `json:"profile"`
}

// User Activity
// https://www.hackthebox.com/api/v4/profile/activity/{userID}

type UserAllActivities struct {
	Profile struct {
		Activity []struct {
			Date string `json:"date"`
			DateDiff string `json:"date_diff"`
			ObjectType string `json:"object_type"`
			Type string `json:"type"`
			FirstBlood bool `json:"first_blood"`
			ID int `json:"id"`
			Name string `json:"name"`
			Points int `json:"points"`
			MachineAvatar string `json:"machine_avatar"` // Machine
			ChallengeCategory string `json:"challenge_category"` // Challenge
			FlagTitle string `json:"flag_title"` // Fortress and Endgame
		} `json:"activity"`
	} `json:"profile"`
}

// User Bloods
// https://www.hackthebox.com/api/v4/profile/bloods/{userID}

type UserAllBloods struct {
	Profile struct {
		Bloods struct {
			Machines []struct {
				ID int `json:"id"`
				Name string `json:"name"`
				Avatar string `json:"avatar"`
				OS string `json:"os"`
				DifficultyText string `json:"difficulty_text"`
				UserBlood bool `json:"user_blood"`
				UserBloodDifference string `json:"user_blood_difference"`
				RootBlood bool `json:"root_blood"`
				RootBloodDifference string `json:"root_blood_difference"`
			} `json:"machines"`
			Challenges []struct {
				ID int `json:"id"`
				Name string `json:"name"`
				CategoryName string `json:"category_name"`
				DifficultyText string `json:"difficulty_test"`
				BloodDiffference string `json:"blood_difference"`
			} `json:"challenges"`
		} `json:"bloods"`
	} `json:"profile"`
}

// User Submissions
// https://www.hackthebox.com/api/v4/profile/content/{userID}

type UserAllSubmissions struct {
	Profile struct {
		Content struct {
			Machines []struct {
				ID int `json:"id"`
				Name string `json:"name"`
				OS string `json:"os"`
				MachineAvatar string `json:"machine_avatar"`
				Difficulty string `json:"difficulty"`
				Rating string `json:"rating"`
				UserOwns int `json:"user_owns"`
				SystemOwns int `json:"system_owns"`
			} `json:"machines"`
			Challenges []struct {
				ID int `json:"id"`
				Name string `json:"name"`
				Category string `json:"category"`
				Likes int `json:"likes"`
				Dislikes int `json:"dislikes"`
			} `json:"challenges"`
			Writeups []struct {
				ID int `json:"id"`
				MachineID int `json:"machine_id"`
				MachineAvatar string `json:"machine_avatar"`
				MachineName string `json:"machine_name"`
				URL string `json:"url"`
				Likes int `json:"likes"`
				Dislikes int `json:"dislikes"`
				Type string `json:"type"`
			} `json:"writeups"`
		} `json:"content"`
	} `json:"profile"`
}

// User Achievements
// https://www.hackthebox.com/api/v4/profile/graph/{duration}/{userID}

type UserAllAchievements struct {
	Profile struct {
		GraphData struct {
			UserOwns []int `json:"user_owns"`
			SystemOwns []int `json:"system_owns"`
			ChallengeOwns []int `json:"challenge_owns"`
			FirstBloods []int `json:"first_bloods"`
			Respects []int `json:"respects"`
		} `json:"graphData"`
	} `json:"profile"`
} 

// User Machine Owns By Attack Path
// https://www.hackthebox.com/api/v4/profile/chart/machines/attack/{userID}

type UserAllAttackPaths struct {
	Profile struct {
		MachineOwns struct {
			Solved int `json:"solved"`
			Total int `json:"total"`
		} `json:"machine_owns"`
		MachineAttackPaths interface{} `json:"machine_attack_paths"`
	} `json:"profile"`
}

// User Profile
// https://www.hackthebox.com/api/v4/profile/{userID}

type UserProfile struct {
	Profile struct {
		ID int `json:"id"`
		SSOID int `json:"sso_id"`
		Name string `json:"name"`
		SystemOwns int `json:"system_owns"`
		UserOwns int `json:"user_owns"`
		UserBloods int `json:"user_bloods"`
		SystemBloods int `json:"system_bloods"`
		Team struct {
			ID int `json:"id"`
			Name string `json:"name"`
			Ranking int `json:"ranking"`
			Avatar string `json:"avatar"`
		} `json:"team"`
		Respects int `json:"respects"`
		Rank string `json:"rank"`
		RankID int `json:"rank_id"`
		CurrentRankProgress int `json:"current_rank_progress"`
		NextRank string `json:"next_rank"`
		NextRankPoints float64 `json:"next_rank_points"`
		RankOwnership string `json:"rank_ownership"`
		RankRequirement int `json:"rank_requirement"`
		Ranking int `json:"ranking"`
		Avatar string `json:"avatar"`
		Timezone string `json:"timezone"`
		Points int `json:"points"`
		CountryName string `json:"country_name"`
		CountryCode string `json:"country_code"`
		UniversityName string `json:"university_name"`
		Description string `json:"description"`
		Github string `json:"github"`
		LinkedIn string `json:"linkedin"`
		Twitter string `json:"twitter"`
		Website string `json:"website"`
	}
}

// User Badges
// https://www.hackthebox.com/api/v4/profile/badges/{userID}

type UserAllBadges struct {
	Badges []struct {
		ID int `json:"id"`
		Name string `json:"name"`
		DescriptionEN string `json:"description_en"`
		Color string `json:"color"`
		Icon string `json:"icon"`
		BadgeCategoryID int `json:"badge_category_id"`
		BadgableType string `json:"badgable_type"`
		BadgableID int `json:"badgable_id"`
		Percentage float64 `json:"percentage"`
		Pivot struct {
			UserID int `json:"user_id"`
			BadgeID int `json:"badge_id"`
			CreatedAt string `json:"created_at"`
			UpdatedAt string `json:"updated_at"`
		} `json:"pivot"`
	} `json:"badges"`
}

func (s *Session) ConnectionStatus() (status BearerConnectionStatus) {

	var url string = "https://www.hackthebox.com/api/v4/user/connection/status"
	parseJSON(s, url, &status)

	return
}

func (s *Session) ActiveMachine() (active BearerActiveMachine) {

	var url string = "https://www.hackthebox.com/api/v4/machine/active"
	parseJSON(s, url, &active)

	return
}

func (s *Session) Followers() (followers BearerFollowers) {

	var url string = "https://www.hackthebox.com/api/v4/user/followers"
	parseJSON(s, url, &followers)

	return
}

func (s *Session) Profile() (profile BearerProfile) {

	var url string = "https://www.hackthebox.com/api/v4/user/info"
	parseJSON(s, url, &profile)

	return
}

func (s *Session) Settings() (settings BearerSettings) {

	var url string = "https://www.hackthebox.com/api/v4/user/settings"
	parseJSON(s, url, &settings)

	return
}

func (s *Session) Subscriptions() (subscriptions BearerSubscriptions) {

	var url string = "https://www.hackthebox.com/api/v4/user/subscriptions"
	parseJSON(s, url, &subscriptions)

	return
}

func (s *Session) SubscriptionsBalance() (balance BearerSubscriptionsBalance) {

	var url string = "https://www.hackthebox.com/api/v4/user/subscriptions/balance"
	parseJSON(s, url, &balance)

	return
}

func (s *Session) SubscriptionsRecurly() (recurly BearerSubscriptionsRecurly) {

	var url string = "https://www.hackthebox.com/api/v4/user/subscriptions/manage/recurly"
	parseJSON(s, url, &recurly)

	return
}

func (s *Session) Tracks() (tracks BearerTracks) {

	var url string = "https://www.hackthebox.com/api/v4/user/tracks"
	parseJSON(s, url, &tracks)

	return
}

/*
func (s *Session) MachineSubmissions() (submissions BearerMachineSubmissions) {

	var url string = "https://www.hackthebox.com/api/v4/team/info/" + teamID
	parseJSON(s, url, &team)

	return
}

func (s *Session) ChallengeSubmissions() (submissions BearerChallengeSubmissions) {

	var url string = "https://www.hackthebox.com/api/v4/team/info/" + teamID
	parseJSON(s, url, &team)

	return
}
*/

func (s *Session) UserRelationship(userID string) (relationship BearerUserRelationship) {

	var url string = "https://www.hackthebox.com/api/v4/user/profile/basic/" + userID
	parseJSON(s, url, &relationship)

	return
}

func (s *Session) UserMachines(userID string) (machines UserProgressMachines) {

	var url string = "https://www.hackthebox.com/api/v4/profile/progress/machines/os/" + userID
	parseJSON(s, url, &machines)

	return
}

func (s *Session) UserChallenges(userID string) (challenges UserProgressChallenges) {

	var url string = "https://www.hackthebox.com/api/v4/profile/progress/challenges/" + userID
	parseJSON(s, url, &challenges)

	return
}

func (s *Session) UserEndgames(userID string) (endgames UserProgressEndgames) {

	var url string = "https://www.hackthebox.com/api/v4/profile/progress/endgame/" + userID
	parseJSON(s, url, &endgames)

	return
}

func (s *Session) UserFortresses(userID string) (fortresses UserProgressFortresses) {

	var url string = "https://www.hackthebox.com/api/v4/profile/progress/fortress/" + userID
	parseJSON(s, url, &fortresses)

	return
}

func (s *Session) UserProLabs(userID string) (proLabs UserProgressProLabs) {

	var url string = "https://www.hackthebox.com/api/v4/profile/progress/prolab/" + userID
	parseJSON(s, url, &proLabs)

	return
}

func (s *Session) UserActivity(userID string) (activities UserAllActivities) {

	var url string = "https://www.hackthebox.com/api/v4/profile/activity/" + userID
	parseJSON(s, url, &activities)

	return
}

func (s *Session) UserBloods(userID string) (bloods UserAllBloods) {

	var url string = "https://www.hackthebox.com/api/v4/profile/bloods/" + userID
	parseJSON(s, url, &bloods)

	return
}

func (s *Session) UserSubmissions(userID string) (submissions UserAllSubmissions) {

	var url string = "https://www.hackthebox.com/api/v4/profile/content/" + userID
	parseJSON(s, url, &submissions)

	return
}

func (s *Session) UserAchievements(userID string, duration string) (acheivements UserAllAchievements) {

	var url string = "https://www.hackthebox.com/api/v4/profile/graph/" + duration + "/" + userID
	parseJSON(s, url, &acheivements)

	return
}

func (s *Session) UserPathOwns(userID string) (owns UserAllAttackPaths) {

	var url string = "https://www.hackthebox.com/api/v4/profile/chart/machines/attack/" + userID
	parseJSON(s, url, &owns)

	return
}

func (s *Session) UserProfile(userID string) (profile UserProfile) {

	var url string = "https://www.hackthebox.com/api/v4/profile/" + userID
	parseJSON(s, url, &profile)

	return
}

func (s *Session) UserBadges(userID string) (badges UserAllBadges) {

	var url string = "https://www.hackthebox.com/api/v4/profile/badges/" + userID
	parseJSON(s, url, &badges)

	return
}