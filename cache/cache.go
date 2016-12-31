package cache

import (
	"github.com/godwhoa/random-shit/botup.me/botup"
)

type Map struct {
	store map[string]string
}

func (m *Map) Get(key string) (string, error) {
	if val, ok := m.store[key]; !ok {
		return botup.KeyDoesntExist
	}
	return val, nil
}

func (m *Map) Set(key string, value string) error {
	if m.store[key] == nil {
		m.store[key] = value
	} else {
		return botup.KeyAlreadyExist
	}
	return nil
}

func (m *Map) Delete(key string) error {
	if val, ok := m.store[key]; ok {
		m.store[key] = nil
	} else {
		return botup.KeyDoesntExist
	}
	return nil
}
