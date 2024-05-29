package widget

import (
	"context"
	"html/template"
	"time"

	"github.com/raopx/glance/internal/assets"
	"github.com/raopx/glance/internal/feed"
)

type RSS struct {
	widgetBase      `yaml:",inline"`
	FeedRequests    []feed.RSSFeedRequest `yaml:"feeds"`
	Style           string                `yaml:"style"`
	ThumbnailHeight float64               `yaml:"thumbnail-height"`
	CardHeight      float64               `yaml:"card-height"`
	Items           feed.RSSFeedItems     `yaml:"-"`
	Limit           int                   `yaml:"limit"`
	CollapseAfter   int                   `yaml:"collapse-after"`
}

func (widget *RSS) Initialize() error {
	widget.withTitle("RSS Feed").withCacheDuration(1 * time.Hour)

	if widget.Limit <= 0 {
		widget.Limit = 25
	}

	if widget.CollapseAfter == 0 || widget.CollapseAfter < -1 {
		widget.CollapseAfter = 5
	}

	if widget.ThumbnailHeight < 0 {
		widget.ThumbnailHeight = 0
	}

	if widget.CardHeight < 0 {
		widget.CardHeight = 0
	}

	return nil
}

func (widget *RSS) Update(ctx context.Context) {
	items, err := feed.GetItemsFromRSSFeeds(widget.FeedRequests)

	if !widget.canContinueUpdateAfterHandlingErr(err) {
		return
	}

	if len(items) > widget.Limit {
		items = items[:widget.Limit]
	}

	widget.Items = items
}

func (widget *RSS) Render() template.HTML {
	if widget.Style == "horizontal-cards" {
		return widget.render(widget, assets.RSSHorizontalCardsTemplate)
	}

	if widget.Style == "horizontal-cards-2" {
		return widget.render(widget, assets.RSSHorizontalCards2Template)
	}

	return widget.render(widget, assets.RSSListTemplate)
}
