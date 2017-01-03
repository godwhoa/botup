package cache

import (
	"github.com/godwhoa/random-shit/botup.me/botup"
)

type Map struct {
	Store map[string]string
}

func (m Map) Get(key string) (string, error) {
	val, ok := m.Store[key]
	if !ok {
		return "", botup.KeyDoesntExist
	}
	return val, nil
}

func (m Map) Set(key string, value string) error {
	_, ok := m.Store[key]
	if !ok {
		m.Store[key] = value
	} else {
		return botup.KeyAlreadyExist
	}
	return nil
}

func (m Map) SetForce(key string, value string) {
	m.Store[key] = value
}

func (m Map) Delete(key string) error {
	if _, ok := m.Store[key]; ok {
		delete(m.Store, key)
	} else {
		return botup.KeyDoesntExist
	}
	return nil
}
