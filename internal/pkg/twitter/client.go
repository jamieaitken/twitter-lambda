package twitter

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/jamieaitken/twitter-lambda/internal/domain"
)

type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	Logger *log.Logger
	Doer   Doer
}

func New(logger *log.Logger, doer Doer) Client {
	return Client{
		Logger: logger,
		Doer:   doer,
	}
}

func (c Client) GetProfile(ctx context.Context) (domain.Profile, error) {
	req, err := c.buildRequest(ctx)
	if err != nil {
		c.Logger.Println(err)

		return domain.Profile{}, err
	}

	resp, err := c.Doer.Do(req)
	if err != nil {
		c.Logger.Println(err)

		return domain.Profile{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.Logger.Println(err)

		return domain.Profile{}, err
	}

	var profile Profile

	err = json.Unmarshal(body, &profile)
	if err != nil {
		c.Logger.Println(err)

		return domain.Profile{}, err
	}

	return profile.ToDomain(), nil
}

func (c Client) buildRequest(ctx context.Context) (*http.Request, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		"https://api.twitter.com/1.1/users/show.json?screen_name=Jamie__Aitken&id=24401846",
		nil,
	)
	if err != nil {
		c.Logger.Println(err)

		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", os.Getenv("TWITTER_TOKEN")))

	return req, nil
}
