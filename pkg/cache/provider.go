package cache

import (
	"github.com/tiagompalte/fake-payment-gateway/configs"
)

func ProviderSet(
	config configs.Config,
) Cache {
	switch config.Cache.DriverName {
	case configs.CacheMemory:
		return NewMemoryCache()
	case configs.CacheRedis:
		return NewRedisCache(
			config.Cache.Redis.Host,
			config.Cache.Redis.Port,
			config.Cache.Redis.DB,
			config.Cache.Redis.Pass,
			config.Cache.Redis.Prefix,
		)
	default: // configs.CacheMock
		return NewMockCache()
	}
}
