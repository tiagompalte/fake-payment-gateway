package protocols

import "context"

type Job interface {
	Execute(ctx context.Context, args ...any) error
}
