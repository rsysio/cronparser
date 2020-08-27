package main

import "fmt"

func Printer(schedule *CronSchedule) {

	fmt.Printf("%-14v%v\n", minute.name, schedule.Minute)
	fmt.Printf("%-14v%v\n", hour.name, schedule.Hour)
	fmt.Printf("%-14v%v\n", dom.name, schedule.DoM)
	fmt.Printf("%-14v%v\n", month.name, schedule.Month)
	fmt.Printf("%-14v%v\n", dow.name, schedule.DoW)
	fmt.Printf("%-14v%v\n", "command", schedule.Command)
}
