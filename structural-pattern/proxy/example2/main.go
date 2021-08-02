package main

import (
	"fmt"
	"time"
)

type ThirdPartyYoutube interface {
	ListVideo() []string
}

type ThirdPartyYoutubeImpl struct {
}

func (t ThirdPartyYoutubeImpl) ListVideo() []string {
	<-time.After(200 * time.Millisecond)

	return []string{
		"https://youtube.com/watch/1",
		"https://youtube.com/watch/2",
	}
}

type CacheYoutube struct {
	service ThirdPartyYoutube
	cache   []string
}

func (c *CacheYoutube) run() {
	for {
		select {
		case <-time.After(5 * time.Second):
			c.cache = []string{}
		}
	}
}

func (c *CacheYoutube) ListVideo() []string {

	if len(c.cache) == 0 {
		c.cache = c.service.ListVideo()
		fmt.Println("call real api to get list of videos")
	}

	fmt.Println("return list of videos")

	return c.cache
}

func main() {

	impl := ThirdPartyYoutubeImpl{}
	cache := CacheYoutube{
		service: &impl,
	}

	go cache.run()

	fmt.Println(cache.ListVideo())
	fmt.Println(cache.ListVideo())

	<-time.After(7 * time.Second)

	fmt.Println(cache.ListVideo())
}
