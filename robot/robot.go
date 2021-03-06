package robot

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/YusukeShimizu/Candy_Scrape/notify"
	"github.com/YusukeShimizu/Candy_Scrape/redisdb"
	"github.com/headzoo/surf/browser"
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
	val    string
}

type result struct {
	time      []string
	dimention []value
}

type value struct {
	status []string
}

var wdays = [...]string{"日", "月", "火", "水", "木", "金", "土"}

func NewRobot(notifyer notify.Notifyer, redis redisdb.Redis) *Robot {
	r := Robot{}
	r.notifier = notifyer
	r.bow = surf.NewBrowser()
	r.redis = redis
	return &r
}

func (r *Robot) filterTarget(place string, findeds []finded) ([]finded, error) {
	filtered := []finded{}
	for _, finded := range findeds {
		status, err := r.redis.HGet(place, finded.date.Format("2006-01-02"))
		if err != nil {
			return findeds, err
		}
		if !(status == finded.status) {
			r.redis.HSet(place, finded.date.Format("2006-01-02"), finded.status)
			log.Println(fmt.Sprintf("%vの%vの予約状態が%vに変更されました。", finded.date, place, finded.status))
			if reflect.DeepEqual([]byte(finded.status), []byte("△")) || reflect.DeepEqual([]byte(finded.status), []byte("○")) {
				//if !goholiday.IsBusinessDay(finded.date) {
				log.Println(fmt.Sprintf("%v %v曜日の%vに空きがでました。", finded.date.Format("2006-01-02"), wdays[finded.date.Weekday()], place))
				filtered = append(filtered, finded)
				//}
			}
		}
	}
	return filtered, nil
}

func (r *Robot) patrol(park string, results []result) error {
	for _, re := range results {
		for _, d := range re.dimention {
			for i, s := range d.status {
				if reflect.DeepEqual([]byte(s), []byte("○")) {
					morning, err := time.Parse("15:04", "12:00")
					hitTime, err := time.Parse("15:04", strings.Split(re.time[i], "～")[0])
					fmt.Println(morning, hitTime)
					if err != nil {
						return err
					}
					if hitTime.Before(morning) {
						r.notifier.Notify(fmt.Sprintf("%v %vの%vに空きができました。明日12時予約します？", re.time[0], re.time[i], park))
					}
				}
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
