// Package botmanager provide ...
package botmanager

import (
	"errors"
	"sync"
)

var (
	ErrDuplicationToken = errors.New("duplicate token")
	ErrNotFound         = errors.New("bot not found")
)

type Runner interface {
	Start(token string) error
	Stop(token string) error
}

type Bot struct {
	Name  string
	Token string
}

type Manager struct {
	mu     sync.RWMutex
	bots   map[string]Bot
	runner Runner
}

func NewManager(r Runner) *Manager {
	return &Manager{
		bots:   make(map[string]Bot),
		runner: r,
	}
}

func (m *Manager) Register(name, token string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.bots[token]; exists {
		return ErrDuplicationToken
	}

	if err := m.runner.Start(token); err != nil {
		return err
	}

	m.bots[token] = Bot{
		Name:  name,
		Token: token,
	}

	return nil
}

func (m *Manager) List() []Bot {
	m.mu.RLock()
	defer m.mu.RUnlock()

	result := make([]Bot, 0, len(m.bots))
	for _, bot := range m.bots {
		result = append(result, bot)
	}
	return result
}

func (m *Manager) Remove(token string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.bots[token]; !exists {
		return ErrNotFound
	}

	if err := m.runner.Stop(token); err != nil {
		return err
	}

	delete(m.bots, token)
	return nil
}

func (m *Manager) Bot(token string) (Bot, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	bot, ok := m.bots[token]
	return bot, ok
}
