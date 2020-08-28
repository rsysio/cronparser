package main

import (
	"log"
	"os"
)

func main() {

	// example from email
	// "*/15 0 1,15 * 1-5 /usr/bin/find"

	args := os.Args
	if len(os.Args) < 2 {
		log.Fatal("Provide cron string as arg 1")
	}

	cron := NewCronSchedule(args[1])
	cron.Process()

	Printer(cron)
}
