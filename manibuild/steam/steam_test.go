package steam

import (
	"github.com/mmcdole/gofeed"
	"github.com/stretchr/testify/suite"
	"testing"
)

type SteamTestSuite struct {
	suite.Suite
}

func (s *SteamTestSuite) TestWalk() {
	_ = WalkFeed("1623730", func(f *gofeed.Item) bool {
		return true
	})
}

func TestSteamTestSuite(t *testing.T) {
	suite.Run(t, new(SteamTestSuite))
}
