package data_structures

import (
	"log"
)

type Set map[any]struct{}

func NewSet() *Set {
	return &Set{}
}

func (s *Set) Add(v string) {
	(*s)[v] = struct{}{}
}

func (s *Set) Exists(v string) bool {
	_, exists := (*s)[v]
	return exists
}

func (s *Set) Remove(v string) string {
	if s.Exists(v) {
		delete((*s), v)
		return v
	} else {
		log.Println("Key does not exists. Empty string returned.")
		return ""
	}
}
