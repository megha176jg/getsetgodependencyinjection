package cache

import "bitbucket.org/junglee_games/getsetgo/clients/redis"

type Factory struct {
	config Config
}

func NewCacheFactory(config Config) *Factory {
	return &Factory{config: config}
}

//GetCache
func (cf *Factory) GetCache(name string) (Cache, error) {
	switch name {
	case REDIS:
		return redis.NewRedisClient(cf.config.GetRedisHost()), nil
	}
	return nil, ErrCacheNotFound
}
