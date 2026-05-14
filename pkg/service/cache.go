package service

import (
	"log"
	"sync"
	"time"
)

type Cache interface {
	Get() (NowPlayingTrack, bool)
	Set(track NowPlayingTrack)
	Clear()
}

type localCache struct {
	lock     sync.Mutex
	hasValue bool
	track    NowPlayingTrack
	ttl      time.Time
}

func NewLocalCache() Cache {
	cache := &localCache{}
	cache.Clear()

	return cache
}

func (c *localCache) Get() (NowPlayingTrack, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if !c.hasValue {
		log.Println("cache: miss")
		return NowPlayingTrack{}, false
	}

	if time.Now().After(c.ttl) {
		log.Println("cache: expired")
		return NowPlayingTrack{}, false
	}

	log.Println("cache: hit")
	return c.track, true
}

func (c *localCache) Set(track NowPlayingTrack) {
	log.Println("cache: update")

	c.lock.Lock()
	defer c.lock.Unlock()

	c.hasValue = true
	c.track = track
	c.ttl = time.Now().Add(30 * time.Second)
}

func (c *localCache) Clear() {
	log.Println("cache: clear")

	c.lock.Lock()
	defer c.lock.Unlock()

	c.hasValue = false
	c.track = NowPlayingTrack{}
	c.ttl = time.Time{}
}
