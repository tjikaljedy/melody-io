package main

import (
	"fmt"
	midware "melody-io/core/internal"
	"melody-io/core/internal/auth"
	"melody-io/core/pkg/log"
	"os"
	"time"
)

var intro = `Running "todo" server with the following options:
	TODO_DEBOUNCE: %s
`
var logger = log.New()

func main() {
	debounce := parseDebounce()

	fmt.Printf(intro, debounceText(debounce))

	var setup midware.Setup

	ctx, cancel := setup.Context()
	defer cancel()

	ebus, estore, ereg, disconnect := setup.Events(ctx, "server")
	defer disconnect()

	cbus, _ := setup.Commands(ereg, ebus)

	repo := setup.Aggregates(estore)

	commandErrors := auth.HandleCommands(ctx, cbus, repo)
	logger.Error(ctx, commandErrors)
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
