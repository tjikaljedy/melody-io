package main

import (
	"fmt"
	"log"
	"math/rand"
	"melody-io/core-es/cmd"
	"melody-io/core-es/internal/auth"
	"time"

	"github.com/google/uuid"
)

func main() {
	var setup cmd.Setup

	ctx, cancel := setup.Context()
	defer cancel()

	ebus, _, ereg, disconnect := setup.Events(ctx, "client")
	defer disconnect()

	cbus, _ := setup.Commands(ereg, ebus)

	// Wait a bit to ensure that the todo server is running before dispatching commands.
	<-time.After(3 * time.Second)

	// Create a new todo list and add some tasks.
	userID := uuid.New()
	for i := 0; i < 10; i++ {
		sleepRandom()

		cmd := auth.UserSigninTask(userID, fmt.Sprintf("Task %d", i+1))
		if err := cbus.Dispatch(ctx, cmd.Any()); err != nil {
			log.Panicf("Failed to dispatch command: %v [cmd=%v, task=%q]", err, cmd.Name(), cmd.Payload())
		}
	}

}

func sleepRandom() {
	dur := time.Duration(rand.Intn(1000)) * time.Millisecond
	log.Printf("Waiting %s before dispatching next command ...", dur)
	time.Sleep(dur)
}
