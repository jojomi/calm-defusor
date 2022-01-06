package main

import "github.com/jojomi/calm-defusor/modules"

func main() {
	var mod modules.Module

	mod = modules.NewPasswordModule()
	mod = modules.NewSimpleWiresModule()
	err := mod.Solve()
	if err != nil {
		panic(err)
	}
}
