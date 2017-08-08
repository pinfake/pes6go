package server

import (
	"sync"
)

type IdMap struct {
	mu     sync.RWMutex
	lastId uint32
	data   map[uint32]interface{}
}

func (m *IdMap) RLock() {
	m.mu.RLock()
}

func (m *IdMap) RUnlock() {
	m.mu.RUnlock()
}

func (m *IdMap) Get(id uint32) interface{} {
	return m.data[id]
}

func (m *IdMap) GetNewId() uint32 {
	defer m.mu.Unlock()
	m.mu.Lock()
	m.lastId++
	return m.lastId
}

func (m *IdMap) Add(id uint32, e interface{}) {
	defer m.mu.Unlock()
	m.mu.Lock()
	m.data[id] = e
}

func (m *IdMap) Delete(id uint32) {
	defer m.mu.Unlock()
	m.mu.Lock()
	delete(m.data, id)
}

func NewIdMap() *IdMap {
	return &IdMap{
		data: make(map[uint32]interface{}),
	}
}
