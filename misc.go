package htbgo

// List of Badges
// https://www.hackthebox.com/api/v4/category/badges
type Badges struct {
	Categories []struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Badges []struct {
			ID              int    `json:"id"`
			Name            string `json:"name"`
			Description     string `json:"description_en"`
			Color           string `json:"color"`
			Icon            string `json:"icon"`
			BadgeCategoryID int    `json:"badge_category_id"`
			BadgeableType   string `json:"badgeable_type"`
			BadgeableID     int    `json:"badgeable_id"`
			Percentage      int    `json:"percentage"`
			UsersCount      int    `json:"users_count"`
		} `json:"badges"`
	} `json:"categories"`
}

// List Bug Feedback Areas
// https://www.hackthebox.com/api/v4/user/feedback/bug/areas
type BugFeedbackAreasList struct {
	Info []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"info"`
}

// Machine And Challenge Stats
// https://www.hackthebox.com/api/v4/content/stats
type MachineAndChallengeStats struct {
	Machines         int `json:"machines"`
	Challenges       int `json:"challenges"`
	Users            int `json:"users"`
	RecruitmentUsers int `json:"recruitment_users"`
	PlatformHours    int `json:"platform_hours"`
	ChallengeTypes   []struct {
		ID                 int    `json:"id"`
		Name               string `json:"name"`
		NumberOfChallenges int    `json:"challenges_count"`
	} `json:"challenge_counts"`
	Prolabs struct {
		Info []struct {
			ID       int `json:"id"`
			Flags    int `json:"flags"`
			Machines int `json:"machines"`
		} `json:"info"`
		Servers int `json:"servers"`
		Users   int `json:"users"`
	} `json:"prolabs"`
	Universities  int `json:"universities"`
	Endgames      int `json:"endgames"`
	Fortresses    int `json:"fortresses"`
	DedicatedLabs struct {
		Labs  int `json:"labs"`
		Users int `json:"users"`
	} `json:"dedicated_labs"`
	Companies int `json:"companies"`
}

// Changelogs
// https://www.hackthebox.com/api/v4/changelogs
type Changelogs struct {
	Changelogs []struct {
		ID          int    `json:"id"`
		Version     string `json:"version"`
		Description string `json:"description"`
		CreatedAt   string `json:"created_at"`
		Changes     []struct {
			ID          int    `json:"id"`
			ChangelogID int    `json:"changelog_id"`
			Type        string `json:"type"`
			Text        string `json:"text"`
		} `json:"changes"`
	} `json:"changelogs"`
}

// List Hackthebox Servers
// https://www.hackthebox.com/api/v4/lab/list
type HacktheboxServers struct {
	LabCategories []struct {
		Code     string `json:"code"`
		Name     string `json:"name"`
		Location string `json:"location"`
	} `json:"lab_categories"`
	LabCategoryCode string `json:"lab_category_code"`
	Servers         []struct {
		ID                   int    `json:"id"`
		Name                 string `json:"friendly_name"`
		CurrentClientsActive int    `json:"current_clients"`
	} `json:"servers"`
	ReleaseArenaLabs []struct {
		Name     string `json:"name"`
		Location string `json:"location"`
		Code     string `json:"code"`
	} `json:"release_arena_labs"`
	ServerID int  `json:"server_id"`
	Disabled bool `json:"disabled"`
}

// List Improvement Feedback Areas
// https://www.hackthebox.com/api/v4/user/feedback/improvement/areas
type ImprovementFeedbackAreasList struct {
	Info []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"info"`
}

// Various Stats
// https://www.hackthebox.com/api/v4/user/dashboard
type VariousStats struct {
	DashboardPlayers struct {
		OnlinePlayers string `json:"online_players"`
	} `json:"dashboard_players"`
}

// Sidebar changelogs
// https://www.hackthebox.com/api/v4/sidebar/changelog

type SidebarChangelogs struct {
	Changelog struct {
		ID        int    `json:"id"`
		Version   string `json:"version"`
		CreatedAt string `json:"created_at"`
	} `json:"changelog"`
}

// Announcement
// https://www.hackthebox.com/api/v4/sidebar/announcement

type SidebarAnnouncement struct {
	Announcement struct {
		ID        int    `json:"id"`
		UpdatedAt string `json:"updated_at"`
		Title     string `json:"title"`
		CreatedAt string `json:"created_at"`
	} `json:"announcement"`
}

func (s *Session) Badges() (badges Badges, err error) {

	var url string = "https://www.hackthebox.com/api/v4/category/badges"
	err = parseJSON(s, url, &badges)

	return
}

func (s *Session) BugFeedbackAreas() (areas BugFeedbackAreasList, err error) {

	var url string = "https://www.hackthebox.com/api/v4/user/feedback/bug/areas"
	err = parseJSON(s, url, &areas)

	return
}

func (s *Session) MachineAndChallengeStats() (stats MachineAndChallengeStats, err error) {

	var url string = "https://www.hackthebox.com/api/v4/content/stats"
	err = parseJSON(s, url, &stats)

	return
}

func (s *Session) Changelogs() (changelogs Changelogs, err error) {

	var url string = "https://www.hackthebox.com/api/v4/changelogs"
	err = parseJSON(s, url, &changelogs)

	return
}

func (s *Session) Servers() (servers HacktheboxServers, err error) {

	var url string = "https://www.hackthebox.com/api/v4/lab/list"
	err = parseJSON(s, url, &servers)

	return
}

func (s *Session) ImprovementFeedbackAreas() (areas ImprovementFeedbackAreasList, err error) {

	var url string = "https://www.hackthebox.com/api/v4/user/feedback/improvement/areas"
	err = parseJSON(s, url, &areas)

	return
}

func (s *Session) VariousStats() (stats VariousStats, err error) {

	var url string = "https://www.hackthebox.com/api/v4/user/dashboard"
	err = parseJSON(s, url, &stats)

	return
}

func (s *Session) SidebarChangelogs() (changelogs SidebarChangelogs, err error) {

	var url string = "https://www.hackthebox.com/api/v4/sidebar/changelog"
	err = parseJSON(s, url, &changelogs)

	return
}

func (s *Session) Announcement() (announcement SidebarAnnouncement, err error) {

	var url string = "https://www.hackthebox.com/api/v4/sidebar/announcement"
	err = parseJSON(s, url, &announcement)
	
	return
}