//go:generate go-bind-plugin -output-name DemoPlugin -plugin-package=../demo_plugin -rebuild -interface -output-package demoplugin -output-path demoplugin.so

package main

import (
	"./hooks"
	_ "./sideeffects"
	"./plugins"
	"github.com/prometheus/common/log"
)

// Resource: https://medium.com/learning-the-go-programming-language/writing-modular-go-programs-with-plugins-ec46381ee1a9

func main() {
	hookList := []string{"se1", "se2", "nope"}

	log.Debug(plugins.Run("demo", "Hello!"))

	for _, hook := range hookList {
		_, e := hooks.Run(hook, nil)
		if e == hooks.ENoHook {
			log.Errorln("Hook " + hook + " not found")
			continue
		}
		if e != nil {
			panic(e)
		}
	}
}
