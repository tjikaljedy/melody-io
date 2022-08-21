package auth

import (
	"log"
	"strings"

	"github.com/google/uuid"
	"github.com/modernice/goes/aggregate"
	"github.com/modernice/goes/command"
	"github.com/modernice/goes/command/handler"
	"github.com/modernice/goes/event"
)

const LoginAggregate = "auth.login"

// Aggregate
type Login struct {
	*aggregate.Base
	*handler.BaseHandler

	tasks   []string
	archive []string
}

func New(id uuid.UUID) *Login {
	var login *Login
	login = &Login{
		Base: aggregate.New(LoginAggregate, id),
		BaseHandler: handler.NewBase(
			handler.BeforeHandle(func(ctx command.Context) error {
				log.Printf("Handling %q command ... [login=%s]", ctx.Name(), id)
				return nil
			}),
			handler.AfterHandle(func(command.Context) {
				login.print()
			}),
		),
	}

	// Register the event appliers for each of the aggregate events.
	event.ApplyWith(login, login.signin, TaskUserSigin)
	// Register the commands handlers.
	command.ApplyWith(login, login.Signin, UserSigninCmd)

	return login

}

func (login *Login) Tasks() []string {
	return login.tasks
}

// Archive returns the completed tasks.
func (login *Login) Archive() []string {
	return login.archive
}

func (login *Login) signin(evt event.Of[string]) {
	login.tasks = append(login.tasks, evt.Data())
}

func (login *Login) Signin(task string) error {
	for _, t := range login.tasks {
		if strings.EqualFold(t, task) {
			return nil
		}
	}

	aggregate.Next(login, TaskUserSigin, task)

	return nil
}

func (login *Login) print() {
	log.Printf("[List:%s] Tasks: %v", login.ID, login.Tasks())
	log.Printf("[List:%s] Archive: %v", login.ID, login.Archive())
}
