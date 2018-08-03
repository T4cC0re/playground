package plugin_definitions

import "errors"

type Plugin string
type PluginEntrypoint func(command string, payload interface{}) (interface{}, error)
type PluginInfo struct {
	Name    string
	Version string
}

var ENoFunction = errors.New("Plugin does not exist")
var EInvalidPayload = errors.New("Payload invalid")
