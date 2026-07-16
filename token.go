package snbgo

import "sync"
import "time"

type TokenManager struct {
	mu         sync.Mutex
	token      string
	expiryTime int64 // unix millis
}

func (tm *TokenManager) Set(token string, expiryTime int64) {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	tm.token = token
	tm.expiryTime = expiryTime
}

func (tm *TokenManager) Get() (string, bool) {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	if tm.token == "" {
		return "", false
	}
	if time.Now().UnixMilli() >= tm.expiryTime {
		return "", false
	}
	return tm.token, true
}

func (tm *TokenManager) IsExpired() bool {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	if tm.token == "" {
		return true
	}
	return time.Now().UnixMilli() >= tm.expiryTime
}
