// Package botmanager provide ...
package botmanager

import (
	"context"
	"sync"
)

type Manager struct {
	mu     sync.RWMutex
	bots   map[string]*botEntry
	runner Runner
}

func NewManager(r Runner) *Manager {
	return &Manager{
		bots:   make(map[string]*botEntry),
		runner: r,
	}
}

func (m *Manager) Register(name, token string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.bots[token]; exists {
		return ErrDuplicationToken
	}

	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})

	entry := &botEntry{
		bot: Bot{
			Name:  name,
			Token: token,
		},
		cancel: cancel,
		done:   done,
	}

	m.bots[token] = entry

	go func() {
		defer close(done)
		_ = m.runner.Run(ctx, token)
	}()

	return nil
}

func (m *Manager) Remove(token string) error {
	m.mu.Lock()

	entry, ok := m.bots[token]
	if !ok {
		m.mu.Unlock()
		return ErrNotFound
	}
	delete(m.bots, token)
	m.mu.Unlock()

	entry.cancel()
	<-entry.done

	return nil
}

func (m *Manager) Bot(token string) (Bot, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	entry, ok := m.bots[token]
	if !ok {
		return Bot{}, false
	}
	return entry.bot, ok
}

func (m *Manager) List() []Bot {
	m.mu.RLock()
	defer m.mu.RUnlock()

	result := make([]Bot, 0, len(m.bots))
	for _, entry := range m.bots {
		result = append(result, entry.bot)
	}
	return result
}

func (m *Manager) StopAll() {
	m.mu.Lock()
	entries := make([]*botEntry, 0, len(m.bots))
	for _, entry := range m.bots {
		entries = append(entries, entry)
	}
	m.bots = make(map[string]*botEntry)
	m.mu.Unlock()

	for _, entry := range entries {
		entry.cancel()
	}

	for _, entry := range entries {
		<-entry.done
	}
}
