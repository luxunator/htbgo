package htbgo

// ReportArea
type ReportArea struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// List of BadgesList
// https://www.hackthebox.com/api/v4/category/badges

// BadgesList
type BadgesList struct {
	Categories []*BadgesListItem `json:"categories"`
}


// BadgesListItem
type BadgesListItem struct {
	ID     int                    `json:"id"`
	Name   string                 `json:"name"`
	Badges []*BadgesListItemBadge `json:"badges"`
}

// BadgesListItemBadge
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

// ReportBugAreasList
type ReportBugAreasList struct {
	Info []*ReportArea `json:"info"`
}

// Machine And Challenge Stats
// https://www.hackthebox.com/api/v4/content/stats

// LabsStatsInfo
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

// LabsStatsInfoChallengeType
type LabsStatsInfoChallengeType struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Count int    `json:"challenges_count"`
}

// LabsStatsInfoProlabs
type LabsStatsInfoProlabs struct {
	Info    []*LabsStatsInfoProlabsItem `json:"info"`
	Servers int                         `json:"servers"`
	Users   int                         `json:"users"`
}

// LabsStatsInfoProlabsItem
type LabsStatsInfoProlabsItem struct {
	ID       int `json:"id"`
	Flags    int `json:"flags"`
	Machines int `json:"machines"`
}

// LabsStatsInfoDedicatedLabs
type LabsStatsInfoDedicatedLabs struct {
	Labs  int `json:"labs"`
	Users int `json:"users"`
}

// ChangelogsList
// https://www.hackthebox.com/api/v4/changelogs

// ChangelogsList
type ChangelogsList struct {
	Changelogs []*ChangelogsListItem `json:"changelogs"`
}

// ChangelogsListItem
type ChangelogsListItem struct {
	ID          int                         `json:"id"`
	Version     string                      `json:"version"`
	Description string                      `json:"description"`
	CreatedAt   string                      `json:"created_at"`
	Changes     []*ChangelogsListItemChange `json:"changes"`
}

// ChangelogsListItemChange
type ChangelogsListItemChange struct {
	ID          int    `json:"id"`
	ChangelogID int    `json:"changelog_id"`
	Type        string `json:"type"`
	Text        string `json:"text"`
}

// List Hackthebox Servers
// https://www.hackthebox.com/api/v4/lab/list

// ServersList
type ServersList struct {
	LabCategories    []*ServersListLabCategory     `json:"lab_categories"`
	LabCategoryCode  string                        `json:"lab_category_code"`
	Servers          []*ServersListServer          `json:"servers"`
	ReleaseArenaLabs []*ServersListReleaseArenaLab `json:"release_arena_labs"`
	ServerID         int                           `json:"server_id"`
	Disabled         bool                          `json:"disabled"`
}

// ServersListLabCategory
type ServersListLabCategory struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

// ServersListServer
type ServersListServer struct {
	ID                   int    `json:"id"`
	Name                 string `json:"friendly_name"`
	CurrentClientsActive int    `json:"current_clients"`
}

// ServersListReleaseArenaLab
type ServersListReleaseArenaLab struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Code     string `json:"code"`
}

// List Improvement Feedback Areas
// https://www.hackthebox.com/api/v4/user/feedback/improvement/areas

// ReportImprovementAreasList
type ReportImprovementAreasList struct {
	Info []*ReportArea `json:"info"`
}

// Various Stats
// https://www.hackthebox.com/api/v4/user/dashboard

// StatsInfo
type StatsInfo struct {
	DashboardPlayers *StatsInfoItem `json:"dashboard_players"`
}

// StatsInfoItem
type StatsInfoItem struct {
	Online string `json:"online_players"`
}

// Sidebar changelogs
// https://www.hackthebox.com/api/v4/sidebar/changelog

// ChangelogsSidebarInfo
type ChangelogsSidebarInfo struct {
	Changelog *ChangelogsSidebarInfoItem `json:"changelog"`
}

// ChangelogsSidebarInfoItem
type ChangelogsSidebarInfoItem struct {
	ID        int    `json:"id"`
	Version   string `json:"version"`
	CreatedAt string `json:"created_at"`
}

// Announcement
// https://www.hackthebox.com/api/v4/sidebar/announcement

// AnnouncementInfo
type AnnouncementInfo struct {
	Announcement *AnnouncementInfoItem `json:"announcement"`
}

// AnnouncementInfoItem
type AnnouncementInfoItem struct {
	ID        int    `json:"id"`
	UpdatedAt string `json:"updated_at"`
	Title     string `json:"title"`
	CreatedAt string `json:"created_at"`
}

// Badges returns a list of all the current badges
func (s *Session) Badges() (badges *BadgesList, err error) {

	var url string = "https://www.hackthebox.com/api/v4/category/badges"
	err = parseJSON(s, url, &badges)

	return
}

// ReportBugAreas returns a list of categories of bugs that can be reported
func (s *Session) ReportBugAreas() (areas *ReportBugAreasList, err error) {

	var url string = "https://www.hackthebox.com/api/v4/user/feedback/bug/areas"
	err = parseJSON(s, url, &areas)

	return
}

// LabsStats returns a list of various stats related to labs
func (s *Session) LabsStats() (stats *LabsStatsInfo, err error) {

	var url string = "https://www.hackthebox.com/api/v4/content/stats"
	err = parseJSON(s, url, &stats)

	return
}

// Changelogs returns a list of changes to the platform
func (s *Session) Changelogs() (changelogs *ChangelogsList, err error) {

	var url string = "https://www.hackthebox.com/api/v4/changelogs"
	err = parseJSON(s, url, &changelogs)

	return
}

// Servers returns a list of the currently available servers
func (s *Session) Servers() (servers *ServersList, err error) {

	var url string = "https://www.hackthebox.com/api/v4/lab/list"
	err = parseJSON(s, url, &servers)

	return
}

// ReportImprovementAreas returns a list of categories of improvements that can be requested
func (s *Session) ReportImprovementAreas() (areas *ReportImprovementAreasList, err error) {

	var url string = "https://www.hackthebox.com/api/v4/user/feedback/improvement/areas"
	err = parseJSON(s, url, &areas)

	return
}

// Stats returns statistics that are related to the dashboard
func (s *Session) Stats() (stats *StatsInfo, err error) {

	var url string = "https://www.hackthebox.com/api/v4/user/dashboard"
	err = parseJSON(s, url, &stats)

	return
}

// ChangelogsSidebar returns the current version of and date of update located in the sidebar
func (s *Session) ChangelogsSidebar() (changelogs *ChangelogsSidebarInfo, err error) {

	var url string = "https://www.hackthebox.com/api/v4/sidebar/changelog"
	err = parseJSON(s, url, &changelogs)

	return
}

// Announcement returns the current announcement to be shown on the sidebard
func (s *Session) Announcement() (announcement *AnnouncementInfo, err error) {

	var url string = "https://www.hackthebox.com/api/v4/sidebar/announcement"
	err = parseJSON(s, url, &announcement)

	return
}
