package profile

import (
	"context"
	"log"

	"github.com/jamieaitken/twitter-lambda/internal/domain"
)

type AccountProvider interface {
	GetProfile(ctx context.Context) (domain.Profile, error)
}

type URLSizeSetter interface {
	ProfileImageSize(profile domain.Profile, size int) domain.Profile
}

type Controller struct {
	Logger        *log.Logger
	Account       AccountProvider
	URLSizeSetter URLSizeSetter
}

func New(logger *log.Logger, provider AccountProvider, setter URLSizeSetter) Controller {
	return Controller{
		Logger:        logger,
		Account:       provider,
		URLSizeSetter: setter,
	}
}

func (c Controller) GetProfile(ctx context.Context) (domain.Profile, error) {
	profile, err := c.Account.GetProfile(ctx)
	if err != nil {
		c.Logger.Println(err)

		return domain.Profile{}, err
	}

	return c.URLSizeSetter.ProfileImageSize(profile, 400), nil
}
