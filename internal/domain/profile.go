package domain

import "github.com/jamieaitken/twitter-lambda/api"

type Profile struct {
	ID                   int64
	IDStr                string
	Name                 string
	ScreenName           string
	Location             *string
	Derived              string
	URL                  *string
	Description          *string
	Protected            bool
	Verified             bool
	FollowersCount       int
	FriendsCount         int
	ListedCount          int
	FavouritesCount      int
	StatusesCount        int
	CreatedAt            string
	ProfileBannerURL     string
	ProfileImageURLHTTPS string
	DefaultProfile       bool
	DefaultProfileImage  bool
	WithheldInCountries  []string
	WithheldScope        string
}

func (p Profile) ToPresentation() api.Profile {
	return api.Profile(p)
}
