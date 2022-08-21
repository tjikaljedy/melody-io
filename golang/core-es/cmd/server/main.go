package main

import (
	"context"
	"flag"
	"fmt"
	"melody-io/core-es/cmd"
	"melody-io/core-es/internal/auth"
	"melody-io/lib-es/pkg/config"
	"melody-io/lib-es/pkg/log"
	"os"
	"time"
)

var flagConfig = flag.String("config", "../../config/local.yml", "config file")
var logger = log.New().With(context.Background())
var intro = `Running server with the following options: DEBOUNCE: %s`

func init() {
	flag.Parse()
}

func main() {
	cfg, err := config.Load(*flagConfig, logger)
	if err != nil {
		logger.Errorf("failed to load application configuration: %s", err)
		os.Exit(-1)
	}

	var setup cmd.Setup
	ctx, cancel := setup.InitalContext(cfg)
	defer cancel()

	debounce := parseDebounce()
	fmt.Printf(intro, debounceText(debounce))
	ebus, estore, ereg, disconnect := setup.Events(ctx, "server")
	defer disconnect()

	cbus, _ := setup.Commands(ereg, ebus)
	repo := setup.Aggregates(estore)

	commandErrors := auth.HandleCommands(ctx, cbus, repo)
	log.LogErrors(ctx, commandErrors)
}

func parseDebounce() time.Duration {
	if d, err := time.ParseDuration(os.Getenv("DEBOUNCE")); err == nil {
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
