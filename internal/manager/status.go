package manager

type BotStatus int

const (
	StatusStarting BotStatus = iota
	StatusRunning
	StatusStopping
	StatusStopped
	StatusFailed
)
