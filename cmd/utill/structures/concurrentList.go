package structures

import "sync"

type ConcurrentMapString struct {
	m  *sync.Mutex
	mp map[string]bool
}

func NewConcurrentMapString() *ConcurrentMapString {
	return &ConcurrentMapString{
		m:  &sync.Mutex{},
		mp: make(map[string]bool),
	}
}

func (l *ConcurrentMapString) Append(item string) {
	l.m.Lock()
	defer l.m.Unlock()

	l.mp[item] = true
}

func (l *ConcurrentMapString) Contains(item string) bool {
	l.m.Lock()
	defer l.m.Unlock()

	_, ok := l.mp[item]

	return ok
}

func (l *ConcurrentMapString) Remove(item string) {
	l.m.Lock()
	defer l.m.Unlock()

	delete(l.mp, item)
}
