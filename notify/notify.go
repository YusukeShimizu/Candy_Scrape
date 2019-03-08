package notify

import "log"

type Notifyer struct{}

func NewNotifyer() *Notifyer {
	return &Notifyer{}
}

func (n *Notifyer) Notify(message string) {
	log.Println(message)
}
