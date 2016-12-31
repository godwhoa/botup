package botup

import "errors"

var KeyDoesntExist = errors.New("Key doesn't exist")
var KeyAlreadyExist = errors.New("Key already exists")

type CacheService interface {
	Get(key string) (string, error)
	Set(key string, value string) error
	Delete(key string) error
}
