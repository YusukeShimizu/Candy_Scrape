package robot

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func (r *Robot) PatrolSetagayaPark() error {

	err := r.bow.Open("https://www.city.setagaya.lg.jp/mokuji/kusei/010/003/index.html")
	if err != nil {
		return err
	}
	err = r.bow.Click("#content > div:nth-child(3) > div.article > div > p:nth-child(1) > a")
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
				val:    val,
			})
		}
	})
	filters, err := r.filterTarget("世田谷公園", findeds)
	if err != nil {
		return err
	}
	if len(filters) == 0 {
		return nil
	}
	for _, filter := range filters {
		values.Add("checkdate", filter.val)
	}
	values.Set("hyoujiOpenCloseFlg", "close")
	values.Set("textDate", time.Now().Format("2006/01/02"))
	values.Set("radioPeriod", "2週間")
	values.Set("radioDisplay", "false")
	values.Set("radioJikan", "all")
	values.Del("checkShisetsu")
	values.Del("HyojiMode")
	err = r.bow.PostForm("https://setagaya.keyakinet.net/Web/Yoyaku/WgR_ShisetsubetsuAkiJoukyou", values)
	if err != nil {
		return err
	}
	results := []result{}
	r.bow.Find("#body > div.content_body > div.item_body > div > div").Each(func(i int, s *goquery.Selection) {
		time := []string{}
		s.Find("table > thead > tr > th").Each(func(j int, s2 *goquery.Selection) {
			time = append(time, s2.Text())
		})
		dimention := []value{}
		s.Find("table > tbody > tr").Each(func(k int, s3 *goquery.Selection) {
			status := []string{}
			s3.Find("td").Each(func(l int, s4 *goquery.Selection) {
				status = append(status, s4.Text())
			})
			dimention = append(dimention, value{status: status})
		})
		results = append(results, result{
			time:      time,
			dimention: dimention,
		})
	})
	fmt.Println(results)
	err = r.patrol("世田谷公園", results)
	if err != nil {
		return err
	}
	return nil
}

func (r *Robot) PatrolHanegiPark() error {

	err := r.bow.Open("http://www.city.setagaya.lg.jp/kurashi/107/165/819/index.html")
	if err != nil {
		return err
	}
	err = r.bow.Click("#content > div:nth-child(3) > div.article > div > p:nth-child(1) > a")
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
	values.Set("checkShisetsu", "112102")
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
				val:    val,
			})
		}
	})
	filters, err := r.filterTarget("羽根木公園", findeds)
	if err != nil {
		return err
	}
	if len(filters) == 0 {
		return nil
	}

	for _, filter := range filters {
		values.Add("checkdate", filter.val)
	}
	values.Set("hyoujiOpenCloseFlg", "close")
	values.Set("textDate", time.Now().Format("2006/01/02"))
	values.Set("radioPeriod", "2週間")
	values.Set("radioDisplay", "false")
	values.Set("radioJikan", "all")
	values.Del("checkShisetsu")
	values.Del("HyojiMode")
	err = r.bow.PostForm("https://setagaya.keyakinet.net/Web/Yoyaku/WgR_ShisetsubetsuAkiJoukyou", values)
	if err != nil {
		return err
	}
	results := []result{}
	r.bow.Find("#body > div.content_body > div.item_body > div > div").Each(func(i int, s *goquery.Selection) {
		time := []string{}
		s.Find("table > thead > tr > th").Each(func(j int, s2 *goquery.Selection) {
			time = append(time, s2.Text())
		})
		dimention := []value{}
		s.Find("table > tbody > tr").Each(func(k int, s3 *goquery.Selection) {
			status := []string{}
			s3.Find("td").Each(func(l int, s4 *goquery.Selection) {
				status = append(status, s4.Text())
			})
			dimention = append(dimention, value{status: status})
		})
		results = append(results, result{
			time:      time,
			dimention: dimention,
		})
	})
	err = r.patrol("羽根木公園", results)
	if err != nil {
		return err
	}
	return nil
}

func (r *Robot) PatrolNogemachiPark() error {

	err := r.bow.Open("http://www.city.setagaya.lg.jp/kurashi/107/165/819/index.html")
	if err != nil {
		return err
	}
	err = r.bow.Click("#content > div:nth-child(3) > div.article > div > p:nth-child(1) > a")
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
	values.Set("checkShisetsu", "112103")
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
				val:    val,
			})
		}
	})
	filters, err := r.filterTarget("野毛町公園", findeds)
	if err != nil {
		return err
	}
	fmt.Println(filters)
	if len(filters) == 0 {
		return nil
	}

	for _, filter := range filters {
		values.Add("checkdate", filter.val)
	}
	values.Set("hyoujiOpenCloseFlg", "close")
	values.Set("textDate", time.Now().Format("2006/01/02"))
	values.Set("radioPeriod", "2週間")
	values.Set("radioDisplay", "false")
	values.Set("radioJikan", "all")
	values.Del("checkShisetsu")
	values.Del("HyojiMode")
	err = r.bow.PostForm("https://setagaya.keyakinet.net/Web/Yoyaku/WgR_ShisetsubetsuAkiJoukyou", values)
	if err != nil {
		return err
	}
	results := []result{}
	r.bow.Find("#body > div.content_body > div.item_body > div > div").Each(func(i int, s *goquery.Selection) {
		time := []string{}
		s.Find("table > thead > tr > th").Each(func(j int, s2 *goquery.Selection) {
			time = append(time, s2.Text())
		})
		dimention := []value{}
		s.Find("table > tbody > tr").Each(func(k int, s3 *goquery.Selection) {
			status := []string{}
			s3.Find("td").Each(func(l int, s4 *goquery.Selection) {
				status = append(status, s4.Text())
			})
			dimention = append(dimention, value{status: status})
		})
		results = append(results, result{
			time:      time,
			dimention: dimention,
		})
	})
	err = r.patrol("野毛町公園", results)
	if err != nil {
		return err
	}
	return nil
}

func (r *Robot) PatrolSougouPark() error {

	err := r.bow.Open("http://www.city.setagaya.lg.jp/kurashi/107/165/819/index.html")
	if err != nil {
		return err
	}
	err = r.bow.Click("#content > div:nth-child(3) > div.article > div > p:nth-child(1) > a")
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
	values.Set("__EVENTARGUMENT", "07")
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
	values.Set("checkShisetsu", "112104")
	values.Set("HyojiMode", "filterAll")
	err = r.bow.PostForm("https://setagaya.keyakinet.net/Web/Yoyaku/WgR_ShisetsuKensaku", values)
	if err != nil {
		return err
	}
	findeds := []finded{}
	r.bow.Find("#body > div.content_body > div.item_body > div > div.scroll-div.clearfix > table > tbody > tr:nth-child(15)").Find("td").Each(func(_ int, s *goquery.Selection) {
		input := s.Find("input")
		val, exist := input.Attr("value")
		log.Println(s.Text(), val)
		if exist {
			findeds = append(findeds, finded{
				status: s.Text(),
				date:   format(val),
				val:    val,
			})
		}
	})
	filters, err := r.filterTarget("総合運動場", findeds)
	if err != nil {
		return err
	}
	fmt.Println(filters)
	if len(filters) == 0 {
		return nil
	}

	for _, filter := range filters {
		values.Add("checkdate", filter.val)
	}
	values.Set("hyoujiOpenCloseFlg", "close")
	values.Set("textDate", time.Now().Format("2006/01/02"))
	values.Set("radioPeriod", "2週間")
	values.Set("radioDisplay", "false")
	values.Set("radioJikan", "all")
	values.Del("checkShisetsu")
	values.Del("HyojiMode")
	err = r.bow.PostForm("https://setagaya.keyakinet.net/Web/Yoyaku/WgR_ShisetsubetsuAkiJoukyou", values)
	if err != nil {
		return err
	}
	results := []result{}
	r.bow.Find("#body > div.content_body > div.item_body > div > div").Each(func(i int, s *goquery.Selection) {
		time := []string{}
		s.Find("table > thead > tr > th").Each(func(j int, s2 *goquery.Selection) {
			time = append(time, s2.Text())
			fmt.Println(s2.Text())
		})
		dimention := []value{}
		s.Find("table > tbody > tr").Each(func(k int, s3 *goquery.Selection) {
			status := []string{}
			s3.Find("td").Each(func(l int, s4 *goquery.Selection) {
				status = append(status, s4.Text())
			})
			dimention = append(dimention, value{status: status})
		})
		results = append(results, result{
			time:      time,
			dimention: dimention,
		})
	})
	err = r.patrol("総合運動場", results)
	if err != nil {
		return err
	}
	return nil
}
