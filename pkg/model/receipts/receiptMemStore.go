package receipts

import (
	"sync"
)

type MemStore struct {
	data sync.Map
}

func NewMemStore() MemStore {
	return MemStore{
		data: sync.Map{},
	}
}

func (m *MemStore) Add(id string, receipt Receipt) bool {
	m.data.Store(id, receipt)
	return true
}

func (m *MemStore) Get(id string) (Receipt, bool) {
	if val, ok := m.data.Load(id); ok {
		return val.(Receipt), true
	}
	return Receipt{}, false
}

func (m *MemStore) Update(id string, receipt Receipt) bool {
	if _, ok := m.data.Load(id); ok {
		m.data.Store(id, receipt)
		return true
	}
	return false
}

func (m *MemStore) Remove(id string) bool {
	_, ok := m.data.LoadAndDelete(id)
	return ok
}
