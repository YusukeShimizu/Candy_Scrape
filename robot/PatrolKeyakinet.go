package robot

import (
	"log"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

func (r *Robot) PatrolSetagayaPark() error {

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
	findeds := []finded{}
	r.bow.Find("#body > div.content_body > div.item_body > div > div.scroll-div.clearfix > table > tbody > tr:nth-child(1)").Find("td").Each(func(_ int, s *goquery.Selection) {
		input := s.Find("input")
		val, exist := input.Attr("value")
		log.Println(s.Text(), val)
		if exist {
			findeds = append(findeds, finded{
				status: s.Text(),
				date:   format(val),
			})
		}
	})
	err = r.patrol("世田谷公園", findeds)
	if err != nil {
		return err
	}
	return nil
}
