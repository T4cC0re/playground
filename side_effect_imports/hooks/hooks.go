package hooks

import (
		"github.com/pkg/errors"
)

var hooks = map[string]Hook{}

type Hook func(payload interface{}) (interface{}, error)

var EExists = errors.New("Hook already registered")
var ENoHook = errors.New("Hook does not exist")
var EInvalidPayload = errors.New("Payload invalid")


func AddHook(name string, hook Hook) (error) {
	if _, ok := hooks[name]; ok {
		return EExists
	}
	hooks[name] = hook
	return nil
}

func Run(name string, payload interface{}) (interface{}, error) {
	if hook, ok := hooks[name]; ok {
		return hook(payload)
	}
	return nil, ENoHook
}

func Hooks() map[string]Hook {
	return hooks
}
