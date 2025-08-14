package store

import (
	"sync/atomic"
)

type command[T any] struct {
	action     string
	id         uint64
	value      T
	predicate  func(id uint64, value T) bool
	result     chan<- T
	resultOk   chan<- bool
	resultId   chan<- uint64
	resultList chan<- map[uint64]T
}

type InMemoryStore[T any] struct {
	items         map[uint64]T
	commands      chan command[T]
	serialCounter uint64
}

func NewInMemoryStore[T any]() *InMemoryStore[T] {
	store := &InMemoryStore[T]{
		items:    make(map[uint64]T),
		commands: make(chan command[T]),
	}
	go store.loop()
	return store
}

func (s *InMemoryStore[T]) loop() {
	for cmd := range s.commands {
		switch cmd.action {
		case "get":
			val, ok := s.items[cmd.id]
			cmd.result <- val
			cmd.resultOk <- ok
		case "set":
			s.items[cmd.id] = cmd.value
		case "filter":
			result := make(map[uint64]T)
			for id, val := range s.items {
				if cmd.predicate(id, val) {
					result[id] = val
				}
			}
			cmd.resultList <- result
		}
	}
}

func (s *InMemoryStore[T]) Get(id uint64) (T, bool) {
	result := make(chan T)
	resultOk := make(chan bool)
	s.commands <- command[T]{action: "get", id: id, result: result, resultOk: resultOk}
	val := <-result
	ok := <-resultOk
	return val, ok
}

func (s *InMemoryStore[T]) Add(value T) uint64 {
	id := atomic.AddUint64(&s.serialCounter, 1)
	s.commands <- command[T]{action: "set", value: value, id: id}
	return id
}

func (s *InMemoryStore[T]) Filter(predicate func(id uint64, value T) bool) map[uint64]T {
	result := make(chan map[uint64]T)
	s.commands <- command[T]{
		action:     "filter",
		predicate:  predicate,
		resultList: result,
	}
	return <-result
}
