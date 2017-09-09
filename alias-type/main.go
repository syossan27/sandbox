package main

import (
	"sandbox/alias-type/action"
	"sandbox/alias-type/config"
)

func main() {
	applicationConfig := config.GetConfig()
	action := action.NewAction(applicationConfig)
	action.ShowName()
}
