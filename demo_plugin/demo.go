package main

import . "../side_effect_imports/plugin_definitions"

func Entrypoint(command string, payload interface{}) (interface{}, error) {
	if command == "info" {
		return PluginInfo{Name: "DemoPlugin", Version: "1.0"}, nil
	}
	return nil, ENoFunction
}
