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

func (s *Session) Badges() (badges Badges) {
	var url string = "https://www.hackthebox.com/api/v4/category/badges"
	parseJSON(s, url, &badges)
	return
}
