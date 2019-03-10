package robot

import (
	"github.com/Candy_Scrape/notify"
	"github.com/Candy_Scrape/redisdb"
	"github.com/headzoo/surf/browser"
	surf "gopkg.in/headzoo/surf.v1"
)

type Robot struct {
	notifier notify.Notifyer
	bow      *browser.Browser
	redis    redisdb.Redis
}

func NewRobot(notifyer notify.Notifyer, redis redisdb.Redis) *Robot {
	r := Robot{}
	r.notifier = notifyer
	r.bow = surf.NewBrowser()
	r.redis = redis
	return &r
}
