package botmanager

import (
	"errors"
	"fmt"
	"sync"
	"testing"
)

type fakeRunner struct {
	started  map[string]bool
	stopped  map[string]bool
	startErr error
	stopErr  error
}

func newFakeRunner() *fakeRunner {
	return &fakeRunner{
		started: make(map[string]bool),
		stopped: make(map[string]bool),
	}
}

func TestRegisterBot(t *testing.T) {
	runner := newFakeRunner()
	manager := NewManager(runner)

	err := manager.Register("bot1", "token123")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(manager.List()) != 1 {
		t.Fatalf("expected 1 bot, got %d", len(manager.List()))
	}
}

func TestRegisterDuplicateToken(t *testing.T) {
	runner := newFakeRunner()
	manager := NewManager(runner)

	err := manager.Register("bot1", "token123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	err = manager.Register("bot2", "token123")
	if err == nil {
		t.Fatal("expected error for dublicate token")
	}
}

func TestRemoveBot(t *testing.T) {
	runner := newFakeRunner()
	manager := NewManager(runner)

	err := manager.Register("bot1", "token123")
	if err != nil {
		t.Fatalf("unknown error: %v", err)
	}

	err = manager.Remove("token123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(manager.List()) != 0 {
		t.Fatal("expected 0 bots after removal")
	}
}

func TestRemoveUnknownBot(t *testing.T) {
	runner := newFakeRunner()
	manager := NewManager(runner)

	err := manager.Remove("unknown")
	if err != ErrNotFound {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}

func TestConcurrentRegister(t *testing.T) {
	runner := newFakeRunner()
	manager := NewManager(runner)

	done := make(chan struct{})

	for i := range 100 {
		go func(i int) {
			_ = manager.Register(
				fmt.Sprintf("bot%d", i),
				fmt.Sprintf("token1%d", i),
			)
			done <- struct{}{}
		}(i)
	}

	for i := 0; i < 100; i++ {
		<-done
	}
}

func TestStressRegisterRemove(t *testing.T) {
	runner := newFakeRunner()
	manager := NewManager(runner)

	var wg sync.WaitGroup

	workers := 1000

	for i := range workers {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			token := fmt.Sprintf("token%d", i)

			switch i % 3 {
			case 0:
				_ = manager.Register("bot", token)
			case 1:
				_ = manager.Remove(token)
			case 2:
				_ = manager.List
			}
		}(i)
	}

	wg.Wait()
}

func TestBootLookup(t *testing.T) {
	runner := newFakeRunner()
	manager := NewManager(runner)

	_ = manager.Register("bot1", "token123")

	bot, ok := manager.Bot("token123")
	if !ok {
		t.Fatal("expected bot to exist")
	}

	if bot.Name != "bot1" {
		t.Fatalf("expected bot name bot1, got %s", bot.Name)
	}
}

func TestBootLookupUnknown(t *testing.T) {
	runner := newFakeRunner()
	manager := NewManager(runner)

	_, ok := manager.Bot("unknown")
	if ok {
		t.Fatal("expected bot not to exists")
	}
}

func (f *fakeRunner) Start(token string) error {
	if f.startErr != nil {
		return f.startErr
	}
	f.started[token] = true
	return nil
}

func (f *fakeRunner) Stop(token string) error {
	if f.stopErr != nil {
		return f.stopErr
	}
	f.stopped[token] = true
	return nil
}

func TestRegisterStartBot(t *testing.T) {
	runner := newFakeRunner()
	manager := NewManager(runner)

	err := manager.Register("bot1", "token123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !runner.started["token123"] {
		t.Fatal("expected bot to be stopped")
	}
}

func TestRemoveStopsBot(t *testing.T) {
	runner := newFakeRunner()
	manager := NewManager(runner)

	_ = manager.Register("bot1", "token123")

	err := manager.Remove("token123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !runner.stopped["token123"] {
		t.Fatal("expected bot to be stopped")
	}
}

func TestRegisterFailsIfRunnerStartFails(t *testing.T) {
	runner := newFakeRunner()
	runner.startErr = errors.New("start failed")

	manager := NewManager(runner)

	err := manager.Register("bot", "token123")
	if err == nil {
		t.Fatal("expected error from runner start")
	}

	if len(manager.List()) != 0 {
		t.Fatal("bot should not be registered if start fails")
	}
}

func TestRemoveFailsIfRunnerStopFails(t *testing.T) {
	runner := newFakeRunner()
	manager := NewManager(runner)

	_ = manager.Register("bot1", "token123")

	runner.stopErr = errors.New("stop failed")

	err := manager.Remove("token123")
	if err == nil {
		t.Fatal("expected error from runner stop")
	}

	if _, ok := manager.Bot("token123"); !ok {
		t.Fatal("bot should remain registered if stop fails")
	}
}
