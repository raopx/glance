package widget

import (
	"context"
	"html/template"
	"time"

	"github.com/raopx/glance/internal/assets"
	"github.com/raopx/glance/internal/feed"
)

type TwitchChannels struct {
	widgetBase      `yaml:",inline"`
	ChannelsRequest []string             `yaml:"channels"`
	Channels        []feed.TwitchChannel `yaml:"-"`
	CollapseAfter   int                  `yaml:"collapse-after"`
}

func (widget *TwitchChannels) Initialize() error {
	widget.withTitle("Twitch Channels").withCacheDuration(time.Minute * 10)

	if widget.CollapseAfter == 0 || widget.CollapseAfter < -1 {
		widget.CollapseAfter = 5
	}

	return nil
}

func (widget *TwitchChannels) Update(ctx context.Context) {
	channels, err := feed.FetchChannelsFromTwitch(widget.ChannelsRequest)

	if !widget.canContinueUpdateAfterHandlingErr(err) {
		return
	}

	channels.SortByViewers()
	widget.Channels = channels
}

func (widget *TwitchChannels) Render() template.HTML {
	return widget.render(widget, assets.TwitchChannelsTemplate)
}
