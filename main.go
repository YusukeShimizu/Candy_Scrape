package main

import (
	"log"
	"time"

	"github.com/YusukeShimizu/Candy_Scrape/env"
	notify "github.com/YusukeShimizu/Candy_Scrape/notify"
	"github.com/YusukeShimizu/Candy_Scrape/redisdb"
	robot "github.com/YusukeShimizu/Candy_Scrape/robot"
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
	go func() {
		for {
			err := robot.PatrolSetagayaPark()
			if err != nil {
				shutdown <- err
			}
			err = robot.PatrolHanegiPark()
			if err != nil {
				shutdown <- err
			}
			err = robot.PatrolNogemachiPark()
			if err != nil {
				shutdown <- err
			}
			err = robot.PatrolSougouPark()
			if err != nil {
				shutdown <- err
			}
			time.Sleep(600 * time.Second)
		}
	}()
	if err != nil {
		log.Fatal(err)
	}
	go notifyer.Wait()

	<-shutdown
	log.Fatal(shutdown)

}
