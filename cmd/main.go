package main

import (
	"context"
	"log"

	"github.com/sskender/spotify/pkg/auth"
	"github.com/sskender/spotify/pkg/service"
	"github.com/sskender/spotify/pkg/spotify"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var svc *service.Provider

func init() {
	ctx := context.TODO()

	tokenStore, err := auth.NewTokenManager(ctx)
	if err != nil {
		log.Fatalf("failed to create token manager: %v", err)
	}

	httpClient := spotify.NewHTTPClient(ctx, tokenStore)
	spotifyClient := spotify.NewSpotifyClient(httpClient)

	cache := service.NewLocalCache()
	svc = service.NewProvider(spotifyClient, cache)
}

func handler(ctx context.Context, request events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	track := svc.GetNowPlayingTrack()

	if request.RawPath == "/image.svg" {
		return events.LambdaFunctionURLResponse{
			StatusCode: 200,
			Headers: map[string]string{
				"Content-Type":                "image/svg+xml",
				"Cache-Control":               "public, max-age=60",
				"Access-Control-Allow-Origin": "*",
			},
			Body: track.GenerateSvg(),
		}, nil
	}

	return events.LambdaFunctionURLResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
		},
		Body: track.GenerateJson(),
	}, nil
}

func main() {
	lambda.Start(handler)
}
