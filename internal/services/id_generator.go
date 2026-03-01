package services

import "context"

type IDGenerator interface {
	NextOrderID(ctx context.Context) (int, error)
}
