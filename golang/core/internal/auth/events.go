package auth

import "github.com/modernice/goes/codec"

// Events
const (
	TaskUserSigin = "auth.task_usersignin"
)

// ListEvents are all events of a todo list.
var ListEvents = [...]string{
	TaskUserSigin,
}

// RegisterEvents registers events into a registry.
func RegisterEvents(r codec.Registerer) {
	codec.Register[string](r, TaskUserSigin)

}
