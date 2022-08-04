package htbgo

type ReportAreas []struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// List of BadgesList
// https://www.hackthebox.com/api/v4/category/badges
type BadgesList struct {
	Categories *BadgesListCategories `json:"categories"`
}

type BadgesListCategories []struct {
	ID     int               `json:"id"`
	Name   string            `json:"name"`
	Badges *BadgesListBadges `json:"badges"`
}

type BadgesListBadges []struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description_en"`
	Color         string `json:"color"`
	Icon          string `json:"icon"`
	CategoryID    int    `json:"badge_category_id"`
	BadgeableType string `json:"badgeable_type"`
	BadgeableID   int    `json:"badgeable_id"`
	Percentage    int    `json:"percentage"`
	UsersCount    int    `json:"users_count"`
}

// List Bug Feedback Areas
// https://www.hackthebox.com/api/v4/user/feedback/bug/areas
type ReportBugAreasList struct {
	Info *ReportAreas `json:"info"`
}

// Machine And Challenge Stats
// https://www.hackthebox.com/api/v4/content/stats
type LabsStatsInfo struct {
	Machines         int                          `json:"machines"`
	Challenges       int                          `json:"challenges"`
	Users            int                          `json:"users"`
	RecruitmentUsers int                          `json:"recruitment_users"`
	PlatformHours    int                          `json:"platform_hours"`
	ChallengeTypes   *LabsStatsInfoChallengeTypes `json:"challenge_counts"`
	Prolabs          *LabsStatsInfoProlabs        `json:"prolabs"`
	Universities     int                          `json:"universities"`
	Endgames         int                          `json:"endgames"`
	Fortresses       int                          `json:"fortresses"`
	DedicatedLabs    *LabsStatsInfoDedicatedLabs  `json:"dedicated_labs"`
	Companies        int                          `json:"companies"`
}

type LabsStatsInfoChallengeTypes []struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Count int    `json:"challenges_count"`
}

type LabsStatsInfoProlabs struct {
	Info    *LabsStatsInfoInfo `json:"info"`
	Servers int                `json:"servers"`
	Users   int                `json:"users"`
}

type LabsStatsInfoInfo []struct {
	ID       int `json:"id"`
	Flags    int `json:"flags"`
	Machines int `json:"machines"`
}

type LabsStatsInfoDedicatedLabs struct {
	Labs  int `json:"labs"`
	Users int `json:"users"`
}

// ChangelogsList
// https://www.hackthebox.com/api/v4/changelogs
type ChangelogsList struct {
	Changelogs *ChangelogsListChangelogs `json:"changelogs"`
}

type ChangelogsListChangelogs []struct {
	ID          int                    `json:"id"`
	Version     string                 `json:"version"`
	Description string                 `json:"description"`
	CreatedAt   string                 `json:"created_at"`
	Changes     *ChangelogsListChanges `json:"changes"`
}

type ChangelogsListChanges []struct {
	ID          int    `json:"id"`
	ChangelogID int    `json:"changelog_id"`
	Type        string `json:"type"`
	Text        string `json:"text"`
}

// List Hackthebox Servers
// https://www.hackthebox.com/api/v4/lab/list
type ServersList struct {
	LabCategories    *ServersListLabCategories    `json:"lab_categories"`
	LabCategoryCode  string                       `json:"lab_category_code"`
	Servers          *ServersListServers          `json:"servers"`
	ReleaseArenaLabs *ServersListReleaseArenaLabs `json:"release_arena_labs"`
	ServerID         int                          `json:"server_id"`
	Disabled         bool                         `json:"disabled"`
}

type ServersListLabCategories []struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

type ServersListServers []struct {
	ID                   int    `json:"id"`
	Name                 string `json:"friendly_name"`
	CurrentClientsActive int    `json:"current_clients"`
}

type ServersListReleaseArenaLabs []struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Code     string `json:"code"`
}

// List Improvement Feedback Areas
// https://www.hackthebox.com/api/v4/user/feedback/improvement/areas
type ReportImprovementAreasList struct {
	Info *ReportAreas `json:"info"`
}

// Various Stats
// https://www.hackthebox.com/api/v4/user/dashboard
type StatsInfo struct {
	DashboardPlayers *StatsInfoDashboardPlayers `json:"dashboard_players"`
}

type StatsInfoDashboardPlayers struct {
	Online string `json:"online_players"`
}

// Sidebar changelogs
// https://www.hackthebox.com/api/v4/sidebar/changelog

type ChangelogsSidebarInfo struct {
	Changelog *ChangelogsSidebarInfoChangelog `json:"changelog"`
}

type ChangelogsSidebarInfoChangelog struct {
	ID        int    `json:"id"`
	Version   string `json:"version"`
	CreatedAt string `json:"created_at"`
}

// Announcement
// https://www.hackthebox.com/api/v4/sidebar/announcement

type AnnouncementInfo struct {
	Announcement *AnnouncementInfoAnnouncement `json:"announcement"`
}

type AnnouncementInfoAnnouncement struct {
	ID        int    `json:"id"`
	UpdatedAt string `json:"updated_at"`
	Title     string `json:"title"`
	CreatedAt string `json:"created_at"`
}

func (s *Session) Badges() (badges *BadgesList, err error) {

	var url string = "https://www.hackthebox.com/api/v4/category/badges"
	err = parseJSON(s, url, &badges)

	return
}

func (s *Session) ReportBugAreas() (areas *ReportBugAreasList, err error) {

	var url string = "https://www.hackthebox.com/api/v4/user/feedback/bug/areas"
	err = parseJSON(s, url, &areas)

	return
}

func (s *Session) LabsStats() (stats *LabsStatsInfo, err error) {

	var url string = "https://www.hackthebox.com/api/v4/content/stats"
	err = parseJSON(s, url, &stats)

	return
}

func (s *Session) Changelogs() (changelogs *ChangelogsList, err error) {

	var url string = "https://www.hackthebox.com/api/v4/changelogs"
	err = parseJSON(s, url, &changelogs)

	return
}

func (s *Session) Servers() (servers *ServersList, err error) {

	var url string = "https://www.hackthebox.com/api/v4/lab/list"
	err = parseJSON(s, url, &servers)

	return
}

func (s *Session) ReportImprovementAreas() (areas *ReportImprovementAreasList, err error) {

	var url string = "https://www.hackthebox.com/api/v4/user/feedback/improvement/areas"
	err = parseJSON(s, url, &areas)

	return
}

func (s *Session) Stats() (stats *StatsInfo, err error) {

	var url string = "https://www.hackthebox.com/api/v4/user/dashboard"
	err = parseJSON(s, url, &stats)

	return
}

func (s *Session) ChangelogsSidebar() (changelogs *ChangelogsSidebarInfo, err error) {

	var url string = "https://www.hackthebox.com/api/v4/sidebar/changelog"
	err = parseJSON(s, url, &changelogs)

	return
}

func (s *Session) Announcement() (announcement *AnnouncementInfo, err error) {

	var url string = "https://www.hackthebox.com/api/v4/sidebar/announcement"
	err = parseJSON(s, url, &announcement)

	return
}
