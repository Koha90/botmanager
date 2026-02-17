package botmanager

import "context"

type Bot struct {
	Name  string
	Token string
}

type botEntry struct {
	bot    Bot
	cancel context.CancelFunc
	done   chan struct{}
}
