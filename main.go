package main

import (
	"context"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jamieaitken/twitter-lambda/api"
	"github.com/jamieaitken/twitter-lambda/internal/pkg/twitter"
	"github.com/jamieaitken/twitter-lambda/internal/pkg/urlsizesetter"
	"github.com/jamieaitken/twitter-lambda/internal/profile"
)

func HandleRequest(ctx context.Context) (api.Profile, error) {
	logger := log.Default()

	client := twitter.New(logger, http.DefaultClient)

	sizeSetter := urlsizesetter.New()

	controller := profile.New(logger, client, sizeSetter)

	account, err := controller.GetProfile(ctx)
	if err != nil {
		return api.Profile{}, err
	}

	return account.ToPresentation(), nil
}

func main() {
	lambda.Start(HandleRequest)
}
