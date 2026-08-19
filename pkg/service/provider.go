package service

import (
	"log"
	"strings"
)

type Provider struct {
	client Client
	cache  Cache
}

func NewProvider(client Client, cache Cache) *Provider {
	return &Provider{
		client: client,
		cache:  cache,
	}
}

func (p *Provider) GetNowPlayingTrack() NowPlayingTrack {
	track, ok := p.cache.Get()
	if ok {
		return track
	}

	spotifyTrack, err := p.client.GetCurrentlyPlaying()
	if err != nil {
		log.Printf("error fetching current track: %v", err)
	}

	if spotifyTrack == nil {
		track = NowPlayingTrack{IsPlaying: false}
		p.cache.Set(track)

		return track
	}

	artists := make([]string, 0, len(spotifyTrack.Item.Artists))
	for _, a := range spotifyTrack.Item.Artists {
		artists = append(artists, a.Name)
	}

	track = NowPlayingTrack{
		IsPlaying: spotifyTrack.IsPlaying,
		Title:     spotifyTrack.Item.Name,
		Author:    strings.Join(artists, ", "),
		Link:      spotifyTrack.Item.ExternalURLs["spotify"],
	}

	p.cache.Set(track)

	return track
}
