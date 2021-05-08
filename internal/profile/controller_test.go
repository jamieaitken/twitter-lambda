package profile

import (
	"context"
	"errors"
	"log"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/jamieaitken/twitter-lambda/internal/domain"
)

func TestGetProfile(t *testing.T) {
	tests := []struct {
		name                 string
		givenAccountProvider AccountProvider
		givenURLSizeSetter   URLSizeSetter
		expectedProfile      domain.Profile
	}{
		{
			name: "given an account fetch from twitter, expect desired image size",
			givenAccountProvider: mockAccountProvider{
				GivenProfile: domain.Profile{
					ProfileImageURLHTTPS: "https://twitter.com/image_normal.jpg",
				},
			},
			givenURLSizeSetter: mockURLSizeSetter{
				GivenProfile: domain.Profile{
					ProfileImageURLHTTPS: "https://twitter.com/image_400x400.jpg",
				},
			},
			expectedProfile: domain.Profile{
				ProfileImageURLHTTPS: "https://twitter.com/image_400x400.jpg",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			controller := New(log.Default(), test.givenAccountProvider, test.givenURLSizeSetter)

			profile, err := controller.GetProfile(context.Background())
			if err != nil {
				t.Fatalf("expected nil, got %v", err)
			}

			if !cmp.Equal(profile, test.expectedProfile) {
				t.Fatal(cmp.Diff(profile, test.expectedProfile))
			}
		})
	}
}

func TestGetProfile_Fails(t *testing.T) {
	twitterError := errors.New("failed to fetch from twitter")

	tests := []struct {
		name                 string
		givenAccountProvider AccountProvider
		givenURLSizeSetter   URLSizeSetter
		expectedError        error
	}{
		{
			name: "given a failed account fetch from twitter, expect error",
			givenAccountProvider: mockAccountProvider{
				GivenError: twitterError,
			},
			givenURLSizeSetter: mockURLSizeSetter{},
			expectedError:      twitterError,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			controller := New(log.Default(), test.givenAccountProvider, test.givenURLSizeSetter)

			_, err := controller.GetProfile(context.Background())
			if err == nil {
				t.Fatalf("expected %v, got nil", test.expectedError)
			}

			if !cmp.Equal(err, test.expectedError, cmpopts.EquateErrors()) {
				t.Fatal(cmp.Diff(err, test.expectedError, cmpopts.EquateErrors()))
			}
		})
	}
}

type mockAccountProvider struct {
	GivenProfile domain.Profile
	GivenError   error
}

func (m mockAccountProvider) GetProfile(_ context.Context) (domain.Profile, error) {
	return m.GivenProfile, m.GivenError
}

type mockURLSizeSetter struct {
	GivenProfile domain.Profile
}

func (m mockURLSizeSetter) ProfileImageSize(_ domain.Profile, _ int) domain.Profile {
	return m.GivenProfile
}
