package steam

import (
	"fmt"
	"github.com/mmcdole/gofeed"
)

type FeedWalkerFunc func(f *gofeed.Item) bool

func WalkFeed(appId string, walkerFunc FeedWalkerFunc) error {
	feed, err := gofeed.NewParser().ParseURL(fmt.Sprintf("https://store.steampowered.com/feeds/news/app/%s", appId))
	if err != nil {
		return fmt.Errorf("invalid app ID %s: %w", appId, err)
	}

	for _, v := range feed.Items {
		if !walkerFunc(v) {
			return nil
		}
	}

	return nil
}
