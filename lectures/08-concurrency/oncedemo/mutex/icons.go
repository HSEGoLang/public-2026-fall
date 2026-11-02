// Package mutex demonstrates lazy initialization with a mutex.
package mutex

import "sync"

// START MUTEX_ICONS OMIT
type Icons struct {
	mu    sync.Mutex
	icons map[string]string
}

func (i *Icons) Icon(name string) string {
	i.mu.Lock()
	defer i.mu.Unlock()

	if i.icons == nil {
		i.icons = map[string]string{"spades.png": "♠"}
	}

	return i.icons[name]
}

// END MUTEX_ICONS OMIT
