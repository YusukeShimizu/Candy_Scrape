package robot

import (
	"log"
	"testing"

	"github.com/Candy_Scrape/notify"
	"github.com/Candy_Scrape/redisdb"
)

func TestPatrolKetakinet(t *testing.T) {

	notifyer := notify.NewNotifyer()
	redis, err := redisdb.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	robot := NewRobot(*notifyer, *redis)
	err = robot.PatrolKeyakinet()
	if err != nil {
		t.Fatal(err)
	}

}
