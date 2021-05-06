package api

type Profile struct {
	ID                   int64    `json:"id"`
	IDStr                string   `json:"idStr"`
	Name                 string   `json:"name"`
	ScreenName           string   `json:"screenName"`
	Location             *string  `json:"location"`
	Derived              string   `json:"derived"`
	URL                  *string  `json:"url"`
	Description          *string  `json:"description"`
	Protected            bool     `json:"protected"`
	Verified             bool     `json:"verified"`
	FollowersCount       int      `json:"followersCount"`
	FriendsCount         int      `json:"friendsCount"`
	ListedCount          int      `json:"listedCount"`
	FavouritesCount      int      `json:"favouritesCount"`
	StatusesCount        int      `json:"statusesCount"`
	CreatedAt            string   `json:"createdAt"`
	ProfileBannerURL     string   `json:"profileBannerURL"`
	ProfileImageURLHTTPS string   `json:"profileImageURLHTTPS"`
	DefaultProfile       bool     `json:"defaultProfile"`
	DefaultProfileImage  bool     `json:"defaultProfileImage"`
	WithheldInCountries  []string `json:"withheldInCountries"`
	WithheldScope        string   `json:"withheldScope"`
}
