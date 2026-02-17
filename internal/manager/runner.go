package manager

import "context"

type Runner interface {
	Run(ctx context.Context, token string) error
}
