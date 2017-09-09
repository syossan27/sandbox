package action

import "fmt"

type Action struct {
	Application application
}

type application struct {
	Name string
}

func (action Action) ShowName() {
	fmt.Println(action.Application.Name)
}

func NewAction(appConfig Action) Action {
	var action Action
	action = appConfig
	return action
}
