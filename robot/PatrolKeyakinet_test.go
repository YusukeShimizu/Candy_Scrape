package robot

import (
	"log"
	"testing"

	"github.com/YusukeShimizu/Candy_Scrape/env"
	"github.com/YusukeShimizu/Candy_Scrape/notify"
	"github.com/YusukeShimizu/Candy_Scrape/redisdb"
)

func TestPatrolSetagayaPark(t *testing.T) {

	config, err := env.Process()
	if err != nil {
		log.Fatal(err)
	}
	notifyer, err := notify.NewNotifyer(&config)
	redis, err := redisdb.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	robot := NewRobot(*notifyer, *redis)
	err = robot.PatrolSetagayaPark()
	if err != nil {
		t.Fatal(err)
	}

}

func TestPatrolHanegiPark(t *testing.T) {

	config, err := env.Process()
	if err != nil {
		log.Fatal(err)
	}
	notifyer, err := notify.NewNotifyer(&config)
	redis, err := redisdb.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	robot := NewRobot(*notifyer, *redis)
	err = robot.PatrolHanegiPark()
	if err != nil {
		t.Fatal(err)
	}

}

func TestPatrolNogemachiPark(t *testing.T) {

	config, err := env.Process()
	if err != nil {
		log.Fatal(err)
	}
	notifyer, err := notify.NewNotifyer(&config)
	redis, err := redisdb.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	robot := NewRobot(*notifyer, *redis)
	err = robot.PatrolNogemachiPark()
	if err != nil {
		t.Fatal(err)
	}

}

func TestPatrolSougouPark(t *testing.T) {

	config, err := env.Process()
	if err != nil {
		log.Fatal(err)
	}
	notifyer, err := notify.NewNotifyer(&config)
	redis, err := redisdb.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	robot := NewRobot(*notifyer, *redis)
	err = robot.PatrolSougouPark()
	if err != nil {
		t.Fatal(err)
	}

}
