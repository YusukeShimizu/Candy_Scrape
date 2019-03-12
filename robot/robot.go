package robot

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/Candy_Scrape/notify"
	"github.com/Candy_Scrape/redisdb"
	"github.com/headzoo/surf/browser"
	"github.com/yut-kt/goholiday"
	surf "gopkg.in/headzoo/surf.v1"
)

type Robot struct {
	notifier notify.Notifyer
	bow      *browser.Browser
	redis    redisdb.Redis
}

type finded struct {
	status string
	date   time.Time
}

var wdays = [...]string{"日", "月", "火", "水", "木", "金", "土"}

func NewRobot(notifyer notify.Notifyer, redis redisdb.Redis) *Robot {
	r := Robot{}
	r.notifier = notifyer
	r.bow = surf.NewBrowser()
	r.redis = redis
	return &r
}

func (r *Robot) patrol(place string, findeds []finded) error {
	for _, finded := range findeds {
		status, before, err := r.redis.Get(finded.date.Format("2006-01-02"))
		if err != nil {
			return err
		}
		if !(status == finded.status) {
			r.redis.Set(finded.date.Format("2006-01-02"), finded.status)
			log.Println(fmt.Sprintf("%vの%vの予約状態が%vに変更されました。", finded.date, place, finded.status))
			if before && (reflect.DeepEqual([]byte(finded.status), []byte("△")) || reflect.DeepEqual([]byte(finded.status), []byte("○"))) &&
				!goholiday.IsBusinessDay(finded.date) {
				r.notifier.Notify(fmt.Sprintf("%v %v曜日の%vに空きがでました。", finded.date.Format("2006-01-02"), wdays[finded.date.Weekday()], place))
			}
		}
	}
	return nil
}

func format(val string) time.Time {
	timeLayout := val[:4] + "-" + val[4:6] + "-" + val[6:8]
	t, err := time.Parse("2006-01-02", timeLayout)
	if err != nil {
		log.Fatal(err)
	}
	return t
}
