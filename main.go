package main

import (
	"context"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jamieaitken/twitter-lambda/api"
	"github.com/jamieaitken/twitter-lambda/internal/pkg/twitter"
)

func HandleRequest(ctx context.Context) (api.Profile, error) {
	logger := log.Default()

	client := twitter.New(logger, http.DefaultClient)
	profile, err := client.GetProfile(ctx)
	if err != nil {
		return api.Profile{}, err
	}

	return profile.ToPresentation(), nil
}

func main() {
	lambda.Start(HandleRequest)
}
