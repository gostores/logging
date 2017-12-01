// The Test package is used for testing logging. It is here for backwards
// compatibility from when logging' organization was upper-case. Please use
// lower-case logging and the `null` package instead of this one.
package test

import (
	"io/ioutil"
	"sync"

	"github.com/gostores/logging"
)

// Hook is a hook designed for dealing with logs in test scenarios.
type Hook struct {
	// Entries is an array of all entries that have been received by this hook.
	// For safe access, use the AllEntries() method, rather than reading this
	// value directly.
	Entries []*logging.Entry
	mu      sync.RWMutex
}

// NewGlobal installs a test hook for the global logger.
func NewGlobal() *Hook {

	hook := new(Hook)
	logging.AddHook(hook)

	return hook

}

// NewLocal installs a test hook for a given local logger.
func NewLocal(logger *logging.Logger) *Hook {

	hook := new(Hook)
	logger.Hooks.Add(hook)

	return hook

}

// NewNullLogger creates a discarding logger and installs the test hook.
func NewNullLogger() (*logging.Logger, *Hook) {

	logger := logging.New()
	logger.Out = ioutil.Discard

	return logger, NewLocal(logger)

}

func (t *Hook) Fire(e *logging.Entry) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.Entries = append(t.Entries, e)
	return nil
}

func (t *Hook) Levels() []logging.Level {
	return logging.AllLevels
}

// LastEntry returns the last entry that was logged or nil.
func (t *Hook) LastEntry() *logging.Entry {
	t.mu.RLock()
	defer t.mu.RUnlock()
	i := len(t.Entries) - 1
	if i < 0 {
		return nil
	}
	// Make a copy, for safety
	e := *t.Entries[i]
	return &e
}

// AllEntries returns all entries that were logged.
func (t *Hook) AllEntries() []*logging.Entry {
	t.mu.RLock()
	defer t.mu.RUnlock()
	// Make a copy so the returned value won't race with future log requests
	entries := make([]*logging.Entry, len(t.Entries))
	for i, entry := range t.Entries {
		// Make a copy, for safety
		e := *entry
		entries[i] = &e
	}
	return entries
}

// Reset removes all Entries from this test hook.
func (t *Hook) Reset() {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.Entries = make([]*logging.Entry, 0)
}
