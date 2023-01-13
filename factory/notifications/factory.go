package notifications

import "bitbucket.org/junglee_games/getsetgo/clients/slack"

type Factory struct {
	config Config
}

func NewNotifierFactory(config Config) *Factory {
	return &Factory{config: config}
}

func (f *Factory) GetNotifier(name string) (Notifier, error) {
	switch name {
	case SLACK:
		return slack.NewSlackClient(f.config.GetSlackURL(), f.config.IsSlackEnabled()), nil
	default:
		return nil, ErrInvalidNotificationClient
	}
}
