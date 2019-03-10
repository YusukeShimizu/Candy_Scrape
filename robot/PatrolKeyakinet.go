package robot

import (
	"fmt"
	"log"
	"net/url"
	"reflect"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func (r *Robot) PatrolKeyakinet() error {

	err := r.bow.Open("http://www.city.setagaya.lg.jp/kurashi/107/165/819/index.html")
	if err != nil {
		return err
	}
	err = r.bow.Click("#content > div > p > a")
	if err != nil {
		return err
	}
	values := url.Values{}
	r.bow.Find(`#form1 > input[type="hidden"]`).Each(func(arg1 int, arg2 *goquery.Selection) {
		key, _ := arg2.Attr("name")
		value, _ := arg2.Attr("value")
		values.Set(key, value)
	})
	values.Set("__EVENTTARGET", "ssCategory")
	values.Set("__EVENTARGUMENT", "06")
	values.Set("radioPurposeLarge", "01")
	values.Set("radioShisetsuLarge", "01")
	values.Set("shisetsuNameTxt", "")
	values.Set("dummy", "")

	err = r.bow.PostForm("https://setagaya.keyakinet.net/Web/Home/WgR_ModeSelect", values)
	if err != nil {
		return err
	}
	values.Set("__EVENTTARGET", "next")
	values.Set("__EVENTARGUMENT", "")
	values.Set("checkShisetsu", "112101")
	values.Set("HyojiMode", "filterAll")

	err = r.bow.PostForm("https://setagaya.keyakinet.net/Web/Yoyaku/WgR_ShisetsuKensaku", values)
	if err != nil {
		return err
	}

	r.bow.Find("#body > div.content_body > div.item_body > div > div.scroll-div.clearfix > table > tbody > tr:nth-child(1)").Find("td").Each(func(_ int, s *goquery.Selection) {
		input := s.Find("input")
		val, exist := input.Attr("value")
		log.Println(s.Text(), val)

		if exist {
			day := format(val)
			stringDay := day.Format("2006-01-02")
			status, before, err := r.redis.Get(stringDay)
			if err != nil {
				log.Println(err)
			}
			r.redis.Set(stringDay, s.Text())
			if !(status == s.Text()) {
				r.notifier.Notify(fmt.Sprintf("%vの世田谷公園野球場の予約状態が%vに変更されました。", day.Format("2006-01-02"), s.Text()))
				if before && (reflect.DeepEqual([]byte(s.Text()), []byte("△")) || reflect.DeepEqual([]byte(s.Text()), []byte("○"))) && day.Weekday() == time.Saturday {
					r.notifier.Notify(fmt.Sprintf("%vの世田谷公園野球場に空きがでました。", stringDay))
				}
			}
		}
	})
	return err
}

func format(val string) time.Time {
	timeLayout := val[:4] + "-" + val[4:6] + "-" + val[6:8]
	t, err := time.Parse("2006-01-02", timeLayout)
	if err != nil {
		log.Fatal(err)
	}
	return t
}
