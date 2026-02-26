package memory

import "context"

type MemoryIDGenerator struct {
	next int
}

func NewMemoryIDGenerator() *MemoryIDGenerator {
	return &MemoryIDGenerator{next: 1}
}

func (g *MemoryIDGenerator) NextOrderID(ctx context.Context) (int, error) {
	id := g.next
	g.next++
	return id, nil
}
