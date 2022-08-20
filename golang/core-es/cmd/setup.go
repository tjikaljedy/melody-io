package cmd

import (
	"context"
	"log"
	"melody-io/core-es/internal/auth"
	"os"
	"os/signal"
	"syscall"

	"github.com/modernice/goes/aggregate/repository"
	"github.com/modernice/goes/backend/mongo"
	"github.com/modernice/goes/backend/nats"
	"github.com/modernice/goes/codec"
	"github.com/modernice/goes/command"
	"github.com/modernice/goes/command/cmdbus"
	"github.com/modernice/goes/event"
	"github.com/modernice/goes/event/eventstore"
)

type Setup struct{}

func (s *Setup) Context() (context.Context, context.CancelFunc) {
	return signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGTERM)
}

func (s *Setup) Events(ctx context.Context, serviceName string) (_ event.Bus, _ event.Store, _ *codec.Registry, disconnect func()) {
	log.Printf("Setting up events ...")

	r := event.NewRegistry()
	auth.RegisterEvents(r)

	bus, disconnect := s.EventBus(ctx, r, serviceName)
	store := eventstore.WithBus(mongo.NewEventStore(r), bus)

	return bus, store, r, disconnect
}

func (s *Setup) EventBus(ctx context.Context, enc codec.Encoding, serviceName string) (_ event.Bus, disconnect func()) {
	bus := nats.NewEventBus(enc)

	return bus, func() {
		log.Printf("Disconnecting from NATS ...")

		if err := bus.Disconnect(ctx); err != nil {
			log.Panicf("Failed to disconnect from NATS: %v", err)
		}
	}
}

func (s *Setup) Commands(ereg *codec.Registry, ebus event.Bus) (command.Bus, *codec.Registry) {
	log.Printf("Setting up commands ...")

	r := command.NewRegistry()
	auth.RegisterCommands(r)

	cmdbus.RegisterEvents(ereg)

	return cmdbus.New(r, ebus), r
}

func (s *Setup) Aggregates(estore event.Store) *repository.Repository {
	log.Printf("Setting up aggregates ...")

	return repository.New(estore)
}
