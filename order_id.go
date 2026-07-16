package snbgo

import (
	"fmt"
	"sync"
	"time"
)

type OrderIDGenerator struct {
	mu     sync.Mutex
	lastTS int64
	seq    int
}

func (g *OrderIDGenerator) Next() (string, error) {
	g.mu.Lock()
	defer g.mu.Unlock()

	now := time.Now().Unix()
	if now == g.lastTS {
		if g.seq >= 999 {
			return "", fmt.Errorf("snbgo: too many order IDs generated in one second")
		}
		g.seq++
	} else {
		g.lastTS = now
		g.seq = 0
	}
	return fmt.Sprintf("%d%03d", g.lastTS, g.seq), nil
}
