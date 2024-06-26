package steam

import (
	"fmt"
	"github.com/mmcdole/gofeed"
	"manibuild/gen/go/manifest"
	"regexp"
	"time"
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

type Version struct {
	Name        string
	Date        time.Time
	Link        string
	Title       string
	Description string
}

func FetchLatestVersion(mani *manifest.BuildTrigger_SteamFeed) (ver *Version, err error) {
	vp := mani.GetSpec()
	if vp == nil {
		return nil, fmt.Errorf("missing version parser")
	}

	titleMatcher := func(s string) string {
		return ""
	}
	bodyMatcher := func(s string) string {
		return ""
	}

	if vp.TitleMatcher != nil {
		m := vp.GetTitleMatcher()
		re, err := regexp.Compile(m)
		if err != nil {
			return nil, fmt.Errorf("failed to parse title matcher %s: %w", m, err)
		}
		titleMatcher = func(s string) string {
			m := re.FindStringSubmatch(s)
			if len(m) > 0 {
				return m[len(m)-1]
			}
			return ""
		}
	}

	if vp.BodyMatcher != nil {
		m := vp.GetBodyMatcher()
		re, err := regexp.Compile(m)
		if err != nil {
			return nil, fmt.Errorf("failed to parse body matcher %s: %w", m, err)
		}
		bodyMatcher = func(s string) string {
			m := re.FindStringSubmatch(s)
			if len(m) > 0 {
				return m[len(m)-1]
			}
			return ""
		}
	}

	ver = &Version{}
	err = WalkFeed(mani.GetAppId(), func(f *gofeed.Item) bool {
		title := titleMatcher(f.Title)
		body := bodyMatcher(f.Description)
		if title != "" || body != "" {
			if title != "" {
				ver.Name = title
			} else {
				ver.Name = body
			}

			if f.PublishedParsed != nil {
				ver.Date = *f.PublishedParsed
			} else if f.UpdatedParsed != nil {
				ver.Date = *f.UpdatedParsed
			}

			ver.Link = f.Link
			ver.Title = f.Title
			ver.Description = f.Description
			return false
		}

		return true
	})

	return
}
