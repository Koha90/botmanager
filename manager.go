// Package botmanager provide ...
package botmanager

import (
	"context"
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

type ContextRunner interface {
	Run(ctx context.Context, token string) error
}

type Bot struct {
	Name  string
	Token string
}

type Manager struct {
	mu     sync.RWMutex
	bots   map[string]Bot
	runner Runner

	cancels map[string]context.CancelFunc
	wg      sync.WaitGroup
}

func NewManager(r Runner) *Manager {
	return &Manager{
		bots:    make(map[string]Bot),
		cancels: make(map[string]context.CancelFunc),
		runner:  r,
	}
}

func (m *Manager) Register(name, token string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.bots[token]; exists {
		return ErrDuplicationToken
	}

	// если Runner поддерживает Run(ctx)
	if r, ok := m.runner.(ContextRunner); ok {
		ctx, cancel := context.WithCancel(context.Background())
		m.cancels[token] = cancel

		m.wg.Add(1)

		go func() {
			defer m.wg.Done()
			_ = r.Run(ctx, token)
		}()
	} else {
		// fallback для старых тестов
		if err := m.runner.Start(token); err != nil {
			return err
		}
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

	if cancel, ok := m.cancels[token]; ok {
		cancel()
		delete(m.cancels, token)
	} else {
		if err := m.runner.Stop(token); err != nil {
			return err
		}
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

func (m *Manager) StopAll() {
	m.mu.Lock()
	defer m.mu.Unlock()

	for token, cancel := range m.cancels {
		cancel()
		delete(m.cancels, token)
	}

	m.wg.Wait()
}
