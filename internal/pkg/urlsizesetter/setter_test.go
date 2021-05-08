package urlsizesetter

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jamieaitken/twitter-lambda/internal/domain"
)

func TestProfileImageSize(t *testing.T) {
	tests := []struct {
		name            string
		givenProfile    domain.Profile
		givenSize       int
		expectedProfile domain.Profile
	}{
		{
			name: "Given a size of 400, expect 400x400 image",
			givenProfile: domain.Profile{
				ProfileImageURLHTTPS: "https://twitter.com/image_normal.jpg",
			},
			givenSize: 400,
			expectedProfile: domain.Profile{
				ProfileImageURLHTTPS: "https://twitter.com/image_400x400.jpg",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			setter := New()

			profile := setter.ProfileImageSize(test.givenProfile, test.givenSize)

			if !cmp.Equal(profile, test.expectedProfile) {
				t.Fatal(cmp.Diff(profile, test.givenProfile))
			}
		})
	}
}
