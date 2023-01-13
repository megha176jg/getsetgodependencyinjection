package notifications

type Notifier interface {
	SendMessage(msg string) error
}

type Config interface {
	IsSlackEnabled() bool
	GetSlackURL() string
}

type UnimplementedConfig struct {
}

func (uc *UnimplementedConfig) IsSlackEnabled() bool {
	return false
}

func (uc *UnimplementedConfig) GetSlackURL() string {
	return "unimplemented"
}
