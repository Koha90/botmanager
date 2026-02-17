package botmanager

import "context"

type Runner interface {
	Run(ctx context.Context, token string) error
}
