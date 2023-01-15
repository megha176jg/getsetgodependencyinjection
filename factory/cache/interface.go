package cache

type Cache interface {
	PushInSet(setName string, values ...string) error
	RemoveFromSet(setName string, values ...string) error
	GetSetValues(setName string) ([]string, error)
	Enqueue(queueName string, values ...string) error
	Dequeue(queueName string) (string, error)
	Peek(queueName string) (string, error)
}

type Config interface {
	GetRedisHost() string
}

type UnimplementedConfig struct {
}

func (c *UnimplementedConfig) GetRedisHost() string {
	return "unimplemented"
}
