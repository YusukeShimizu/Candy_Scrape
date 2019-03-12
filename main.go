package main

import (
	"log"

	"github.com/Candy_Scrape/env"
	notify "github.com/Candy_Scrape/notify"
	"github.com/Candy_Scrape/redisdb"
	robot "github.com/Candy_Scrape/robot"
	"github.com/robfig/cron"
)

func main() {

	shutdown := make(chan interface{})
	config, err := env.Process()
	if err != nil {
		log.Fatal(err)
	}
	notifyer, err := notify.NewNotifyer(&config)
	if err != nil {
		log.Fatal(err)
	}
	redis, err := redisdb.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	robot := robot.NewRobot(*notifyer, *redis)
	cron := cron.New()
	cron.AddFunc(config.Pace, func() {
		err := robot.PatrolSetagayaPark()
		if err != nil {
			shutdown <- err
		}
	})
	cron.Start()
	<-shutdown
	log.Fatal(shutdown)

}
