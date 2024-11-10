package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	// Create a new cron instance
	c := cron.New()

	// Schedule a job to run every minute
	_, err := c.AddFunc("*/1 * * * *", func() {
		fmt.Println("Task running at:", time.Now())
	})
	if err != nil {
		fmt.Println("Error scheduling task:", err)
		return
	}

	// Start the cron scheduler
	c.Start()

	// Keep the program running
	select {}
}

//cron.New(): Creates a new cron scheduler.

// c.AddFunc("*/1 * * * *", func() { ... }): Adds a job to run every minute using a cron expression (*/1 * * * *).

// This function will print the current timestamp every minute.

// c.Start(): Starts the scheduler.

// select {}: Keeps the main function from exiting immediately, letting the cron job continue running.
