package day10

import (
	"strings"
	"sync"
)

type StringStack struct {
	items []string
	lock  sync.RWMutex
}

func (s *StringStack) Push(i string) {
	if s.items == nil {
		s.items = []string{}
	}

	s.lock.Lock()
	s.items = append(s.items, i)
	s.lock.Unlock()
}

func (s *StringStack) Pop() *string {
	if s.items == nil {
		return nil
	}
	s.lock.Lock()
	item := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1]
	s.lock.Unlock()
	return &item
}

func (s *StringStack) Size() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.items)
}

func (s *StringStack) All() []string {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.items
}

func (s *StringStack) IsEmpty() bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.items) == 0
}

func (s *StringStack) ToString() string {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return strings.Join(s.items, "")
}
