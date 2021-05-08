package urlsizesetter

import (
	"fmt"
	"strings"

	"github.com/jamieaitken/twitter-lambda/internal/domain"
)

type Setter struct{}

func New() Setter {
	return Setter{}
}

func (s Setter) ProfileImageSize(profile domain.Profile, size int) domain.Profile {
	profile.ProfileImageURLHTTPS = strings.Replace(
		profile.ProfileImageURLHTTPS,
		"_normal",
		fmt.Sprintf("_%vx%v", size, size),
		1)

	return profile
}
