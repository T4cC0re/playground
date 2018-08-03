package main

import (
	. "../side_effect_imports/plugin_definitions"
	"fmt"
)

func Entrypoint(command string, payload interface{}) (interface{}, error) {
	if command == "info" {
		return PluginInfo{Name: "DemoPlugin", Version: "1.0"}, nil
	}
	return nil, ENoFunction
}

// SayHello says hello :)
func SayHello(who string) {
	fmt.Printf("Hello %s\n", who)
}

func main() {
  SayHello("a")
  Entrypoint("i", nil)
}
