package sideeffects

import (
	"../hooks"
	"log"
)

func init() {
	hooks.AddHook("se1", sideEffect)
}

func sideEffect(payload interface{}) (a interface{}, b error) {
	log.Println("I am a hook in sideeffect1")
	return
}
