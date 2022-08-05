package htbgo

type ReportArea struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// List of BadgesList
// https://www.hackthebox.com/api/v4/category/badges
type BadgesList struct {
	Categories []*BadgesListItem `json:"categories"`
}

type BadgesListItem struct {
	ID     int                    `json:"id"`
	Name   string                 `json:"name"`
	Badges []*BadgesListItemBadge `json:"badges"`
}

type BadgesListItemBadge struct {
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
	Info []*ReportArea `json:"info"`
}

// Machine And Challenge Stats
// https://www.hackthebox.com/api/v4/content/stats
type LabsStatsInfo struct {
	Machines         int                           `json:"machines"`
	Challenges       int                           `json:"challenges"`
	Users            int                           `json:"users"`
	RecruitmentUsers int                           `json:"recruitment_users"`
	PlatformHours    int                           `json:"platform_hours"`
	ChallengeTypes   []*LabsStatsInfoChallengeType `json:"challenge_counts"`
	Prolabs          *LabsStatsInfoProlabs         `json:"prolabs"`
	Universities     int                           `json:"universities"`
	Endgames         int                           `json:"endgames"`
	Fortresses       int                           `json:"fortresses"`
	DedicatedLabs    *LabsStatsInfoDedicatedLabs   `json:"dedicated_labs"`
	Companies        int                           `json:"companies"`
}

type LabsStatsInfoChallengeType struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Count int    `json:"challenges_count"`
}

type LabsStatsInfoProlabs struct {
	Info    []*LabsStatsInfoProlabsItem `json:"info"`
	Servers int                         `json:"servers"`
	Users   int                         `json:"users"`
}

type LabsStatsInfoProlabsItem struct {
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
	Changelogs []*ChangelogsListItem `json:"changelogs"`
}

type ChangelogsListItem struct {
	ID          int                         `json:"id"`
	Version     string                      `json:"version"`
	Description string                      `json:"description"`
	CreatedAt   string                      `json:"created_at"`
	Changes     []*ChangelogsListItemChange `json:"changes"`
}

type ChangelogsListItemChange struct {
	ID          int    `json:"id"`
	ChangelogID int    `json:"changelog_id"`
	Type        string `json:"type"`
	Text        string `json:"text"`
}

// List Hackthebox Servers
// https://www.hackthebox.com/api/v4/lab/list
type ServersList struct {
	LabCategories    []*ServersListLabCategory     `json:"lab_categories"`
	LabCategoryCode  string                        `json:"lab_category_code"`
	Servers          []*ServersListServer          `json:"servers"`
	ReleaseArenaLabs []*ServersListReleaseArenaLab `json:"release_arena_labs"`
	ServerID         int                           `json:"server_id"`
	Disabled         bool                          `json:"disabled"`
}

type ServersListLabCategory struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

type ServersListServer struct {
	ID                   int    `json:"id"`
	Name                 string `json:"friendly_name"`
	CurrentClientsActive int    `json:"current_clients"`
}

type ServersListReleaseArenaLab struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Code     string `json:"code"`
}

// List Improvement Feedback Areas
// https://www.hackthebox.com/api/v4/user/feedback/improvement/areas
type ReportImprovementAreasList struct {
	Info []*ReportArea `json:"info"`
}

// Various Stats
// https://www.hackthebox.com/api/v4/user/dashboard
type StatsInfo struct {
	DashboardPlayers *StatsInfoItem `json:"dashboard_players"`
}

type StatsInfoItem struct {
	Online string `json:"online_players"`
}

// Sidebar changelogs
// https://www.hackthebox.com/api/v4/sidebar/changelog

type ChangelogsSidebarInfo struct {
	Changelog *ChangelogsSidebarInfoItem `json:"changelog"`
}

type ChangelogsSidebarInfoItem struct {
	ID        int    `json:"id"`
	Version   string `json:"version"`
	CreatedAt string `json:"created_at"`
}

// Announcement
// https://www.hackthebox.com/api/v4/sidebar/announcement

type AnnouncementInfo struct {
	Announcement *AnnouncementInfoItem `json:"announcement"`
}

type AnnouncementInfoItem struct {
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
