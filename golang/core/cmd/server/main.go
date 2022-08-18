package main

import (
	"fmt"
	"melody-io/core/cmd"
	"melody-io/core/internal/auth"
	"melody-io/core/pkg"
	"os"
	"time"
)

var intro = `Running "todo" server with the following options:
	TODO_DEBOUNCE: %s
`

func main() {
	debounce := parseDebounce()

	fmt.Printf(intro, debounceText(debounce))

	var setup cmd.Setup

	ctx, cancel := setup.Context()
	defer cancel()

	ebus, estore, ereg, disconnect := setup.Events(ctx, "server")
	defer disconnect()

	cbus, _ := setup.Commands(ereg, ebus)

	repo := setup.Aggregates(estore)

	commandErrors := auth.HandleCommands(ctx, cbus, repo)
	pkg.LogErrors(ctx, commandErrors)
}

func parseDebounce() time.Duration {
	if d, err := time.ParseDuration(os.Getenv("TODO_DEBOUNCE")); err == nil {
		return d
	}
	return 0
}

func debounceText(dur time.Duration) string {
	if dur == 0 {
		return fmt.Sprintf("%s (disabled)", dur)
	}
	return dur.String()
}
