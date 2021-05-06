package twitter

import (
	"bytes"
	"context"
	"errors"
	"io"
	"log"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/jamieaitken/twitter-lambda/internal/domain"
)

func TestClient_GetProfile(t *testing.T) {
	tests := []struct {
		name            string
		givenDoer       Doer
		expectedProfile domain.Profile
	}{
		{
			name: "Given a response from Twitter, expect a domain profile",
			givenDoer: mockDoer{
				GivenResponse: &http.Response{
					Body:   io.NopCloser(bytes.NewReader([]byte(profileResponse))),
					Status: http.StatusText(http.StatusOK),
				},
			},
			expectedProfile: domain.Profile{
				ID:                  1,
				IDStr:               "id1",
				Name:                "Golang",
				ScreenName:          "@golang",
				URL:                 &URL,
				FollowersCount:      5,
				FriendsCount:        5,
				ListedCount:         1,
				FavouritesCount:     1,
				StatusesCount:       2,
				CreatedAt:           "Sat Mar 14 17:46:49 +0000 2009",
				ProfileBannerURL:    "https://twitter.com",
				DefaultProfile:      true,
				DefaultProfileImage: true,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := New(log.Default(), test.givenDoer)

			actualProfile, err := client.GetProfile(context.Background())
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(actualProfile, test.expectedProfile) {
				t.Fatal(cmp.Diff(actualProfile, test.expectedProfile))
			}
		})
	}
}

func TestClient_GetProfile_Fails(t *testing.T) {
	failedTwitterRequest := errors.New("failed to get data")

	tests := []struct {
		name          string
		givenDoer     Doer
		expectedError error
	}{
		{
			name: "Given a failed response from Twitter, expect an error to be raised",
			givenDoer: mockDoer{
				GivenResponse: &http.Response{
					Status: http.StatusText(http.StatusUnauthorized),
				},
				GivenError: failedTwitterRequest,
			},
			expectedError: failedTwitterRequest,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := New(log.Default(), test.givenDoer)

			_, err := client.GetProfile(context.Background())
			if err == nil {
				t.Fatalf("expected %v, got nil", test.expectedError)
			}

			if !cmp.Equal(err, test.expectedError, cmpopts.EquateErrors()) {
				t.Fatal(cmp.Diff(err, test.expectedError, cmpopts.EquateErrors()))
			}
		})
	}
}

type mockDoer struct {
	GivenResponse *http.Response
	GivenError    error
}

func (m mockDoer) Do(_ *http.Request) (*http.Response, error) {
	return m.GivenResponse, m.GivenError
}

var profileResponse = `{
	"id":1,
	"id_str":"id1",
	"name":"Golang",
	"screen_name":"@golang",
	"location":null,
	"derived":"",
	"url":"https://twitter.com/user",
	"description":null,
	"protected":false,
	"verified":false,
	"followers_count":5,
	"friends_count":5,
	"listed_count":1,
	"favourites_count":1,
	"statuses_count":2,
	"created_at":"Sat Mar 14 17:46:49 +0000 2009",
	"profile_banner_url":"https://twitter.com",
	"profile_image_url_https":"",
	"default_profile":true,
	"default_profile_image":true,
	"withheld_in_countries":null,
	"withheld_scope":""
}`

var URL = "https://twitter.com/user"
