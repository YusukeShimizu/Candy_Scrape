package main

import (
	"log"

	robot "github.com/Candy_Scrape/Robot"
	notify "github.com/Candy_Scrape/notify"
)

func main() {
	notifyer := notify.NewNotifyer()
	robot := robot.NewRobot(*notifyer)
	err := robot.PatrolKeyakinet()
	if err != nil {
		log.Fatal(err)
	}
}
