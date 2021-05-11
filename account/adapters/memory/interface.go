package imemory

import "time"

type IMemoryStore interface {
	SetLoginAttemps(key string, value int, exp time.Duration) error
	GetLoginAttemps(key string) (int, error)
}
