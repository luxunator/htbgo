package htbgo

type ProlabMaster struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	AvatarThumb string `json:"avatar_thumb"`
}

// Current Prolabs List
// https://www.hackthebox.com/api/v4/prolabs

type ProLabsActiveList struct {
	Status bool `json:"status"`
	Data   struct {
		Count int `json:"count"`
		Labs  []struct {
			ID                         int    `json:"id"`
			Name                       string `json:"name"`
			ReleaseAt                  string `json:"release_at"`
			ProMachinesCount           int    `json:"pro_machines_count"`
			ProFlagsCount              int    `json:"pro_flags_count"`
			Ownership                  int    `json:"ownership"`
			UserEligibleForCertificate bool   `json:"user_eligible_for_certificate"`
			New                        bool   `json:"new"`
			SkillLevel                 string `json:"skill_level"`
			DesignatedCategory         string `json:"designated_category"`
			Team                       string `json:"team"`
			Level                      int    `json:"level"`
			LabServersCount            int    `json:"lab_servers_count"`
			CoverImageURL              string `json:"cover_img_url"`
		} `json:"labs"`
	} `json:"data"`
}

// Prolab Profile
// https://www.hackthebox.com/api/v4/prolab/{prolabID}/info

type ProLabInfo struct {
	Status bool `json:"status"`
	Data   struct {
		ID               int          `json:"id"`
		Name             string       `json:"name"`
		Version          string       `json:"version"`
		EntryPoints      []string     `json:"entry_points"`
		Description      string       `json:"description"`
		VideoURL         string       `json:"video_url"`
		ProMachinesCount int          `json:"pro_machines_count"`
		ProFlagsCount    int          `json:"pro_flags_count"`
		CoverImageURL    string       `json:"cover_image_url"`
		LabServersCount  int          `json:"lab_servers_count"`
		ActiveUsers      int          `json:"active_users"`
		LabMaster        ProlabMaster `json:"lab_master"`
	} `json:"data"`
}

// Prolab Overview
// https://www.hackthebox.com/api/v4/prolab/{prolabID}/overview

type ProLabOverviewInfo struct {
	Status bool `json:"status"`
	Data   struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Version string `json:"version"`
		//Excerpt `json:"excerpt"` TODO
		ProMachinesCount int `json:"pro_machines_count"`
		ProFlagsCount    int `json:"pro_flags_count"`
		SocialLinks      struct {
			Discord string `json:"discord"`
			Forum   string `json:"forum"`
		} `json:"social_links"`
		NewVersion       bool   `json:"new_version"`
		OverviewImageURL string `json:"overview_image_url"`
		SkillLevel       string `json:"skill_level"`
		DesignatedLevel  struct {
			Category    string `json:"category"`
			Level       int    `json:"level"`
			Description string `json:"description"`
			Team        string `json:"team"`
		} `json:"designated_level"`
		LabMaster ProlabMaster `json:"lab_master"`
	} `json:"data"`
}

// Prolab Machines List
// https://www.hackthebox.com/api/v4/prolab/{prolabID}/machines

type ProLabMachinesList struct {
	Status bool `json:"status"`
	Data   []struct {
		ID             int    `json:"id"`
		Name           string `json:"name"`
		OS             string `json:"os"`
		AvatarThumbURL string `json:"avatar_thumb_url"`
	} `json:"data"`
}

// Prolab Flags List
// https://www.hackthebox.com/api/v4/prolab/{prolabID}/flags

type ProLabFlagsList struct {
	Status bool `json:"status"`
	Data   []struct {
		ID     int    `json:"id"`
		Title  string `json:"title"`
		Points int    `json:"points"`
		Owned  bool   `json:"owned"`
	} `json:"data"`
}

// Prolab FAQ
// https://www.hackthebox.com/api/v4/prolab/{prolabID}/faq

type ProLabFAQList struct {
	Status bool `json:"status"`
	Data   []struct {
		Question string `json:"question"`
		Answer   string `json:"answer"`
		Generic  bool   `json:"generic"`
	} `json:"data"`
}

// Prolab Selected Reviews
// https://www.hackthebox.com/api/v4/prolab/{prolabID}/reviews_overview

type ProLabSelectedReviewsList struct {
	Status bool `json:"status"`
	Data   struct {
		TotalNumberOfRatings int `json:"total_number_of_ratings"`
		UsersAverageRating   int `json:"users_average_rating"`
		Feedbacks            []struct {
			Rating int    `json:"rating"`
			Text   string `json:"text"`
			User   struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"user"`
		} `json:"feedbacks"`
	} `json:"data"`
}

// Prolab Paginated Reviews
// https://www.hackthebox.com/api/v4/prolab/{prolabID}/reviews?page={page}

type ProLabPaginatedReviewsList struct {
	Data []struct {
		ID                        int    `json:"id"`
		Rating                    int    `json:"rating"`
		Difficulty                int    `json:"difficulty"`
		CreatedAt                 string `json:"created_at"`
		Text                      string `json:"text"`
		HelpfulProFeedbacksCount  int    `json:"helpful_pro_feedbacks_count"`
		UserInHelpfulProFeedbacks bool   `json:"user_in_helpful_pro_feedbacks"`
		User                      struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			AvatarThumb string `json:"avatar_thumb"`
		} `json:"user"`
	} `json:"data"`
	Links struct {
		First string `json:"first"`
		Last  string `json:"last"`
		Prev  string `json:"prev"`
		Next  string `json:"next"`
	} `json:"links"`
	Meta struct {
		CurrentPage int `json:"current_page"`
		From        int `json:"from"`
		LastPage    int `json:"last_page"`
		Links       []struct {
			URL    string `json:"url"`
			Label  string `json:"label"`
			Active bool   `json:"active"`
		} `json:"links"`
		Path    string `json:"path"`
		PerPage int    `json:"per_page"`
		To      int    `json:"to"`
		Total   int    `json:"total"`
	} `json:"meta"`
}

func (s *Session) ProLabsActive() (proLabs ProLabsActiveList, err error) {

	var url string = "https://www.hackthebox.com/api/v4/prolabs"
	err = parseJSON(s, url, &proLabs)

	return
}

func (s *Session) ProLab(proLabID int) (info ProLabInfo, err error) {

	proLabIDString, err := toPositiveIntString(proLabID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/prolab/" + proLabIDString + "/info"
	err = parseJSON(s, url, &info)

	return
}

func (s *Session) ProLabOverview(proLabID int) (overview ProLabOverviewInfo, err error) {

	proLabIDString, err := toPositiveIntString(proLabID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/prolab/" + proLabIDString + "/overview"
	err = parseJSON(s, url, &overview)

	return
}

func (s *Session) ProLabMachines(proLabID int) (machines ProLabMachinesList, err error) {

	proLabIDString, err := toPositiveIntString(proLabID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/prolab/" + proLabIDString + "/machines"
	err = parseJSON(s, url, &machines)

	return
}

func (s *Session) ProLabFlags(proLabID int) (flags ProLabFlagsList, err error) {

	proLabIDString, err := toPositiveIntString(proLabID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/prolab/" + proLabIDString + "/flags"
	err = parseJSON(s, url, &flags)

	return
}

func (s *Session) ProLabFAQ(proLabID int) (faq ProLabFAQList, err error) {

	proLabIDString, err := toPositiveIntString(proLabID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/prolab/" + proLabIDString + "/faq"
	err = parseJSON(s, url, &faq)

	return
}

func (s *Session) ProLabSelectedReviews(proLabID int) (reviews ProLabSelectedReviewsList, err error) {

	proLabIDString, err := toPositiveIntString(proLabID)
	if err != nil {
		return
	}

	var url string = "https://www.hackthebox.com/api/v4/prolab/" + proLabIDString + "/reviews_overview"
	err = parseJSON(s, url, &reviews)

	return
}

func (s *Session) ProLabPaginatedReviews(proLabID int, page int) (reviews ProLabPaginatedReviewsList, err error) {

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
