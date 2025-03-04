package cache

import (
	"context"
	"errors"
	"time"

	"github.com/tiagompalte/fake-payment-gateway/pkg/healthcheck"
)

var ErrItemNotFound = errors.New("cache: item not found")

type Cache interface {
	healthcheck.HealthCheck
	Set(ctx context.Context, key string, value any, ttl time.Duration) error
	Get(ctx context.Context, key string, value any) error
	Clear(ctx context.Context, key string) error
	ClearAll(ctx context.Context) error
}
