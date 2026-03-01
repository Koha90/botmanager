package domain

// AggregateRoot represent a domain aggregate.
//
// All aggregates:
//   - maintain a version for optimistic locking
//   - produce domain events
//   - expose events for application layer publishing
type AggregateRoot interface {
	Version() int
	PullEvevnts() []DomainEvent
}

// BaseAggregate provides common functionality
// for all aggregates.
//
// It handles:
//   - optimistic locking version
//   - domain events buffering
//
// Concrete aggregates should embed this struct.
type BaseAggregate struct {
	version int
	events  []DomainEvent
}

// Version returns aggregate version.
func (a *BaseAggregate) Version() int {
	return a.version
}

// incrementVersion increases aggregate version.
// Should be called on every state mutation.
func (a *BaseAggregate) incrementVersion() {
	a.version++
}

// addEvent stores domain event inside aggregate.
func (a *BaseAggregate) addEvent(event DomainEvent) {
	a.events = append(a.events, event)
}

// PullEvevnts returns accumulated events and clears buffer.
func (a *BaseAggregate) PullEvents() []DomainEvent {
	result := make([]DomainEvent, len(a.events))
	copy(result, a.events)

	a.events = nil
	return result
}

// setInitialVersion sets initial aggregate version.
// Intended for constructors.
func (a *BaseAggregate) setInitialVersion(v int) {
	a.version = v
}
