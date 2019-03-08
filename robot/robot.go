package robot

import "github.com/Candy_Scrape/notify"

type Robot struct {
	notifier notify.Notifyer
}

func NewRobot(notifyer notify.Notifyer) *Robot {
	r := Robot{}
	r.notifier = notifyer
	return &r
}
