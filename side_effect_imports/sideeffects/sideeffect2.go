package sideeffects

import (
	"../hooks"
	"log"
)

func init() {
	hooks.AddHook("se2", sideEffect2)
}

func sideEffect2(payload interface{}) (a interface{}, b error) {
	log.Println("I am a hook in sideeffect2")
	return
}
