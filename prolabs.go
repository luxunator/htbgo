package htbgo

// ProlabMaster
type ProlabMaster struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Thumb string `json:"avatar_thumb"`
}

// Current Prolabs List
// https://www.hackthebox.com/api/v4/prolabs

// ProLabsActiveList
type ProLabsActiveList struct {
	Status bool                   `json:"status"`
	Data   *ProLabsActiveListItem `json:"data"`
}

// ProLabsActiveListItem
type ProLabsActiveListItem struct {
	Count int                         `json:"count"`
	Labs  []*ProLabsActiveListItemLab `json:"labs"`
}

// ProLabsActiveListItemLab
type ProLabsActiveListItemLab struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	ReleaseAt         string `json:"release_at"`
	MachineCount      int    `json:"pro_machines_count"`
	FlagCount         int    `json:"pro_flags_count"`
	Ownership         int    `json:"ownership"`
	IsEligibleForCert bool   `json:"user_eligible_for_certificate"`
	IsNew             bool   `json:"new"`
	SkillLevel        string `json:"skill_level"`
	Category          string `json:"designated_category"`
	Team              string `json:"team"`
	Level             int    `json:"level"`
	ServersCount      int    `json:"lab_servers_count"`
	Cover             string `json:"cover_img_url"`
}

// Prolab Profile
// https://www.hackthebox.com/api/v4/prolab/{prolabID}/info

// ProLabInfo
type ProLabInfo struct {
	Status bool            `json:"status"`
	Data   *ProLabInfoItem `json:"data"`
}

// ProLabInfoItem
type ProLabInfoItem struct {
	ID           int           `json:"id"`
	Name         string        `json:"name"`
	Version      string        `json:"version"`
	EntryPoints  []string      `json:"entry_points"`
	Description  string        `json:"description"`
	Video        string        `json:"video_url"`
	MachineCount int           `json:"pro_machines_count"`
	FlagCount    int           `json:"pro_flags_count"`
	Cover        string        `json:"cover_image_url"`
	ServersCount int           `json:"lab_servers_count"`
	ActiveUsers  int           `json:"active_users"`
	LabMaster    *ProlabMaster `json:"lab_master"`
}

// Prolab Overview
// https://www.hackthebox.com/api/v4/prolab/{prolabID}/overview

// ProLabOverviewInfo
type ProLabOverviewInfo struct {
	Status bool                    `json:"status"`
	Data   *ProLabOverviewInfoItem `json:"data"`
}

// ProLabOverviewInfoItem
type ProLabOverviewInfoItem struct {
	ID              int                                    `json:"id"`
	Name            string                                 `json:"name"`
	Version         string                                 `json:"version"`
	Excerpt         interface{}                            `json:"excerpt"`
	MachinsCount    int                                    `json:"pro_machines_count"`
	FlagCount       int                                    `json:"pro_flags_count"`
	SocialLinks     *ProLabOverviewInfoItemSocialLinks     `json:"social_links"`
	NewVersion      bool                                   `json:"new_version"`
	OverviewImg     string                                 `json:"overview_image_url"`
	SkillLevel      string                                 `json:"skill_level"`
	DesignatedLevel *ProLabOverviewInfoItemDesignatedLevel `json:"designated_level"`
	LabMaster       *ProlabMaster                          `json:"lab_master"`
}

// ProLabOverviewInfoItemSocialLinks
type ProLabOverviewInfoItemSocialLinks struct {
	Discord string `json:"discord"`
	Forum   string `json:"forum"`
}

// ProLabOverviewInfoItemDesignatedLevel
type ProLabOverviewInfoItemDesignatedLevel struct {
	Category    string `json:"category"`
	Level       int    `json:"level"`
	Description string `json:"description"`
	Team        string `json:"team"`
}

// Prolab Machines List
// https://www.hackthebox.com/api/v4/prolab/{prolabID}/machines

// ProLabMachinesList
type ProLabMachinesList struct {
	Status bool                      `json:"status"`
	Data   []*ProLabMachinesListItem `json:"data"`
}

// ProLabMachinesListItem
type ProLabMachinesListItem struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	OS    string `json:"os"`
	Thumb string `json:"avatar_thumb_url"`
}

// Prolab Flags List
// https://www.hackthebox.com/api/v4/prolab/{prolabID}/flags

// ProLabFlagsList
type ProLabFlagsList struct {
	Status bool                   `json:"status"`
	Data   []*ProLabFlagsListItem `json:"data"`
}

// ProLabFlagsListItem
type ProLabFlagsListItem struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Points   int    `json:"points"`
	HasOwned bool   `json:"owned"`
}

// Prolab FAQ
// https://www.hackthebox.com/api/v4/prolab/{prolabID}/faq

// ProLabFAQList
type ProLabFAQList struct {
	Status bool                 `json:"status"`
	Data   []*ProLabFAQListItem `json:"data"`
}

// ProLabFAQListItem
type ProLabFAQListItem struct {
	Question  string `json:"question"`
	Answer    string `json:"answer"`
	IsGeneric bool   `json:"generic"`
}

// Prolab Selected Reviews
// https://www.hackthebox.com/api/v4/prolab/{prolabID}/reviews_overview

// ProLabSelectedReviewsList
type ProLabSelectedReviewsList struct {
	Status bool                           `json:"status"`
	Data   *ProLabSelectedReviewsListItem `json:"data"`
}

// ProLabSelectedReviewsListItem
type ProLabSelectedReviewsListItem struct {
	RatingsCount int                                      `json:"total_number_of_ratings"`
	RatingAvg    int                                      `json:"users_average_rating"`
	Feedbacks    []*ProLabSelectedReviewsListItemFeedback `json:"feedbacks"`
}

// ProLabSelectedReviewsListItemFeedback
type ProLabSelectedReviewsListItemFeedback struct {
	Rating int                                        `json:"rating"`
	Text   string                                     `json:"text"`
	User   *ProLabSelectedReviewsListItemFeedbackUser `json:"user"`
}

// ProLabSelectedReviewsListItemFeedbackUser
type ProLabSelectedReviewsListItemFeedbackUser struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Prolab Paginated Reviews
// https://www.hackthebox.com/api/v4/prolab/{prolabID}/reviews?page={page}

// ProLabPaginatedReviewsList
type ProLabPaginatedReviewsList struct {
	Data  []*ProLabPaginatedReviewsListData `json:"data"`
	Links *ProLabPaginatedReviewsListLinks  `json:"links"`
	Meta  *ProLabPaginatedReviewsListMeta   `json:"meta"`
}

// ProLabPaginatedReviewsListData
type ProLabPaginatedReviewsListData struct {
	ID                 int                                 `json:"id"`
	Rating             int                                 `json:"rating"`
	Difficulty         int                                 `json:"difficulty"`
	CreatedAt          string                              `json:"created_at"`
	Text               string                              `json:"text"`
	HelpfulReviewCount int                                 `json:"helpful_pro_feedbacks_count"`
	InHelpfulReview    bool                                `json:"user_in_helpful_pro_feedbacks"`
	User               *ProLabPaginatedReviewsListDataUser `json:"user"`
}

// ProLabPaginatedReviewsListDataUser
type ProLabPaginatedReviewsListDataUser struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Thumb string `json:"avatar_thumb"`
}

// ProLabPaginatedReviewsListLinks
type ProLabPaginatedReviewsListLinks struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Prev  string `json:"prev"`
	Next  string `json:"next"`
}

// ProLabPaginatedReviewsListMeta
type ProLabPaginatedReviewsListMeta struct {
	CurrentPage int                                   `json:"current_page"`
	From        int                                   `json:"from"`
	LastPage    int                                   `json:"last_page"`
	Links       []*ProLabPaginatedReviewsListMetaLink `json:"links"`
	Path        string                                `json:"path"`
	PerPage     int                                   `json:"per_page"`
	To          int                                   `json:"to"`
	Total       int                                   `json:"total"`
}

// ProLabPaginatedReviewsListMetaLink
type ProLabPaginatedReviewsListMetaLink struct {
	URL      string `json:"url"`
	Label    string `json:"label"`
	IsActive bool   `json:"active"`
}

// ProLabsActive returns a list of the currently active prolabs
func (s *Session) ProLabsActive() (proLabs *ProLabsActiveList, err error) {

	var url string = "https://www.hackthebox.com/api/v4/prolabs"
	err = parseJSON(s, url, &proLabs)

	return
}

// ProLab returns the information of a prolab, given a prolab ID
func (s *Session) ProLab(proLabID int) (info *ProLabInfo, err error) {

	proLabIDString, err := toPositiveIntString(proLabID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/prolab/" + proLabIDString + "/info"
	err = parseJSON(s, url, &info)

	return
}

// ProLabOverview returns the overview of a prolab, given a prolab ID
func (s *Session) ProLabOverview(proLabID int) (overview *ProLabOverviewInfo, err error) {

	proLabIDString, err := toPositiveIntString(proLabID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/prolab/" + proLabIDString + "/overview"
	err = parseJSON(s, url, &overview)

	return
}

// ProLabMachines returns a list of the machines that a prolab contains, given an prolab ID
func (s *Session) ProLabMachines(proLabID int) (machines *ProLabMachinesList, err error) {

	proLabIDString, err := toPositiveIntString(proLabID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/prolab/" + proLabIDString + "/machines"
	err = parseJSON(s, url, &machines)

	return
}

// ProLabFlags returns a list of information on the flags within a prolab, given an prolab ID
func (s *Session) ProLabFlags(proLabID int) (flags *ProLabFlagsList, err error) {

	proLabIDString, err := toPositiveIntString(proLabID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/prolab/" + proLabIDString + "/flags"
	err = parseJSON(s, url, &flags)

	return
}

// ProLabFAQ returns a list of FAQ's of a prolab, given a prolab ID
func (s *Session) ProLabFAQ(proLabID int) (faq *ProLabFAQList, err error) {

	proLabIDString, err := toPositiveIntString(proLabID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/prolab/" + proLabIDString + "/faq"
	err = parseJSON(s, url, &faq)

	return
}

// ProLabSelectedReviews returns a list of selected reviews of a prolab, given a prolab ID
func (s *Session) ProLabSelectedReviews(proLabID int) (reviews *ProLabSelectedReviewsList, err error) {

	proLabIDString, err := toPositiveIntString(proLabID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/prolab/" + proLabIDString + "/reviews_overview"
	err = parseJSON(s, url, &reviews)

	return
}

// ProLabPaginatedReviews returns a list of reviews of a machine, given a machine ID and page number
func (s *Session) ProLabPaginatedReviews(proLabID int, page int) (reviews *ProLabPaginatedReviewsList, err error) {

	proLabIDString, err := toPositiveIntString(proLabID)
	if err != nil {
		return
	}

	pageString, err := toPositiveIntString(page)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/prolab/" + proLabIDString + "/reviews?page=" + pageString
	err = parseJSON(s, url, &reviews)

	return
}
