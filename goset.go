package goset

import (
	"sort"
	"sync"
)

type Set struct {
	m *sync.Map
}

func New() *Set {
	return &Set{m: &sync.Map{}}
}

func NewSetWithStrings(strs []string) *Set {
	s := New()
	for i := range strs {
		s.Add(strs[i])
	}
	return s
}

func (s *Set) Exists(v interface{}) bool {
	_, ok := s.m.Load(v)
	return ok
}

func (s *Set) Add(v interface{}) {
	s.m.Store(v, struct{}{})
}

func (s *Set) AddString(v ...string) *Set {
	for i := range v {
		s.Add(v[i])
	}
	return s
}

func (s *Set) Strings() []string {
	dst := make([]string, 0)
	s.m.Range(func(key, value interface{}) bool {
		str, ok := key.(string)
		if ok {
			dst = append(dst, str)
		}
		return true
	})
	return dst
}

func (s *Set) SortStrings() []string {
	dst := s.Strings()
	sort.Strings(dst)
	return dst
}

func (s *Set) Delete(v interface{}) {
	s.m.Delete(v)
}

func (s *Set) DeleteString(v ...string) *Set {
	for i := range v {
		s.Delete(v[i])
	}
	return s
}
