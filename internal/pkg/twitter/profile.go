package twitter

import "github.com/jamieaitken/twitter-lambda/internal/domain"

type Profile struct {
	ID                   int64    `json:"id"`
	IDStr                string   `json:"id_str"`
	Name                 string   `json:"name"`
	ScreenName           string   `json:"screen_name"`
	Location             *string  `json:"location"`
	Derived              string   `json:"derived"`
	URL                  *string  `json:"url"`
	Description          *string  `json:"description"`
	Protected            bool     `json:"protected"`
	Verified             bool     `json:"verified"`
	FollowersCount       int      `json:"followers_count"`
	FriendsCount         int      `json:"friends_count"`
	ListedCount          int      `json:"listed_count"`
	FavouritesCount      int      `json:"favourites_count"`
	StatusesCount        int      `json:"statuses_count"`
	CreatedAt            string   `json:"created_at"`
	ProfileBannerURL     string   `json:"profile_banner_url"`
	ProfileImageURLHTTPS string   `json:"profile_image_url_https"`
	DefaultProfile       bool     `json:"default_profile"`
	DefaultProfileImage  bool     `json:"default_profile_image"`
	WithheldInCountries  []string `json:"withheld_in_countries"`
	WithheldScope        string   `json:"withheld_scope"`
}

func (p Profile) ToDomain() domain.Profile {
	return domain.Profile(p)
}
