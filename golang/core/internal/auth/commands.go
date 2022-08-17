package auth

import (
	"context"

	"github.com/google/uuid"
	"github.com/modernice/goes/aggregate"
	"github.com/modernice/goes/codec"
	"github.com/modernice/goes/command"
	"github.com/modernice/goes/command/handler"
)

// Commands
const (
	UserSigninCmd = "todo.auth.usersignin_task"
)

func UserSigninTask(userID uuid.UUID, task string) command.Cmd[string] {
	return command.New(UserSigninCmd, task, command.Aggregate(LoginAggregate, userID))
}

// RegisterCommands registers commands into a registry.
func RegisterCommands(r codec.Registerer) {
	codec.Register[string](r, UserSigninCmd)
}

// HandleCommands handles todo list commands that are dispatched over the
// provided command bus until ctx is canceled. Command errors are sent into
// the returned error channel.
func HandleCommands(ctx context.Context, bus command.Bus, repo aggregate.Repository) <-chan error {
	return handler.New(New, repo, bus).MustHandle(ctx)
}
