package main

import "sync"

// START MAILBOX_TYPE OMIT
type Mailbox struct {
	mu      sync.Mutex
	changed *sync.Cond
	value   int
	full    bool
}

// END MAILBOX_TYPE OMIT

// START MAILBOX_NEW OMIT
func NewMailbox() *Mailbox {
	m := &Mailbox{}
	m.changed = sync.NewCond(&m.mu)

	return m
}

// END MAILBOX_NEW OMIT

// START MAILBOX_PUT OMIT
func (m *Mailbox) Put(value int) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for m.full {
		m.changed.Wait()
	}

	m.value = value
	m.full = true
	m.changed.Broadcast()
}

// END MAILBOX_PUT OMIT

// START MAILBOX_GET OMIT
func (m *Mailbox) Get() int {
	m.mu.Lock()
	defer m.mu.Unlock()

	for !m.full {
		m.changed.Wait()
	}

	value := m.value
	m.full = false
	m.changed.Broadcast()

	return value
}

// END MAILBOX_GET OMIT
