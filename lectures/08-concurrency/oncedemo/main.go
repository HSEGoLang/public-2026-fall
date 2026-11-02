package main

import (
	"fmt"
	"sync"
)

// START ICONS OMIT

// START LOAD_ICONS OMIT
func loadIcons() {
	fmt.Println("loading icons")

	icons = map[string]string{
		"spades.png":   "♠",
		"hearts.png":   "♥",
		"diamonds.png": "♦",
		"clubs.png":    "♣",
	}
}

// END LOAD_ICONS OMIT

// START ICON_ONCE OMIT
var (
	icons         map[string]string
	loadIconsOnce sync.Once
)

func Icon(name string) string {
	loadIconsOnce.Do(loadIcons)
	return icons[name]
}

// END ICON_ONCE OMIT
// END ICONS OMIT

// START ICON_MAIN OMIT
func main() {
	var wg sync.WaitGroup

	values := make([]string, 2)
	for i := range values {
		wg.Add(1)
		go func() {
			defer wg.Done()

			values[i] = Icon("spades.png")
		}()
	}

	wg.Wait()
	fmt.Println(values[0], values[1])
}

// END ICON_MAIN OMIT
