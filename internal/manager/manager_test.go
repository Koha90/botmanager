package manager

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

type fakeRunner struct {
	started chan string
	done    chan string
}

func newFakeRunner() *fakeRunner {
	return &fakeRunner{
		started: make(chan string, 100),
		done:    make(chan string, 100),
	}
}

func (f *fakeRunner) Run(ctx context.Context, token string) error {
	f.started <- token

	<-ctx.Done()

	f.done <- token

	return nil
}

type ctxRunner struct {
	done chan string
}

func newCtxRunner() *ctxRunner {
	return &ctxRunner{
		done: make(chan string, 100),
	}
}

func (c *ctxRunner) Run(ctx context.Context, token string) error {
	<-ctx.Done()
	c.done <- token
	return nil
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

	for i := 0; i < 100; i++ {
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

	for i := 0; i < workers; i++ {
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
				_ = manager.List()
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

func TestRegisterStartBot(t *testing.T) {
	runner := newFakeRunner()
	manager := NewManager(runner)

	_ = manager.Register("bot1", "token123")

	time.Sleep(10 * time.Millisecond)

	select {
	case tok := <-runner.started:
		if tok != "token123" {
			t.Fatal("wrong token")
		}
	case <-time.After(time.Second):
		t.Fatal("runner did not start")
	}
}

func TestRemoveStopsBot(t *testing.T) {
	runner := newFakeRunner()
	manager := NewManager(runner)

	_ = manager.Register("bot1", "token123")
	_ = manager.Remove("token123")

	select {
	case tok := <-runner.done:
		if tok != "token123" {
			t.Fatalf("unexpected token %s", tok)
		}
	case <-time.After(time.Second):
		t.Fatal("runner did not stopped")
	}
}

func TestStopAll(t *testing.T) {
	r := newFakeRunner()
	m := NewManager(r)

	for i := 0; i < 10; i++ {
		_ = m.Register(
			fmt.Sprintf("bot%d", i),
			fmt.Sprintf("token%d", i),
		)
	}

	m.StopAll()

	for i := 0; i < 10; i++ {
		select {
		case <-r.done:
		case <-time.After(time.Second):
			t.Fatal("runner was not stopped")
		}
	}
}

func TestRemoveCancelContextRunner(t *testing.T) {
	r := newCtxRunner()
	m := NewManager(r)

	_ = m.Register("bot", "token123")

	if err := m.Remove("token123"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	select {
	case token := <-r.done:
		if token != "token123" {
			t.Fatalf("unexpected token %s", token)
		}
	default:
		t.Fatal("context runner was not canceled")
	}
}

func TestStopAllCancelsAllContextRunners(t *testing.T) {
	r := newCtxRunner()
	m := NewManager(r)

	for i := 0; i < 5; i++ {
		_ = m.Register("bot", fmt.Sprintf("token%d", i))
	}

	m.StopAll()

	for i := 0; i < 5; i++ {
		select {
		case <-r.done:
		default:
			t.Fatal("not all runners were canceled")
		}
	}
}
