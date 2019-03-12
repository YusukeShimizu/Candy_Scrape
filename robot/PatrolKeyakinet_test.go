package robot

import (
	"log"
	"testing"

	"github.com/Candy_Scrape/env"
	"github.com/Candy_Scrape/notify"
	"github.com/Candy_Scrape/redisdb"
)

func TestPatrolKetakinet(t *testing.T) {

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
