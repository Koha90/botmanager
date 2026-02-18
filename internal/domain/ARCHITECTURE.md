# Domain Layer Rules

- Domain does not depend on infrastructure.
- All business invariants are enforced inside aggregates.
- State transitions are allowed only through methods.
- Domain events are generated inside aggregates.
- Version field is used for optimistic locking.
