package robot

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func (r *Robot) PatrolKeyakinet() error {
	values := url.Values{}
	values.Set("__RequestVerificationToken", "omjzY/noGvso1jqvd0Hdm2uYg+Eiuk/XReMedacCMswgzslekCOa2kSj/Wcv7rjbHoA5Hsj0Udp9qOnGLor+JlubRpOhwsP3pfNRNmlbHVZOBZ0cL5V8bB6dRKHrIEkdzv+FqA==")
	values.Set("__EVENTTARGET", "next")
	values.Set("__EVENTARGUMENT", "")
	values.Set("checkShisetsu", "112101")
	values.Set("HyojiMode", "filterAll")
	values.Set("as_sfid", "AAAAAAVkHhDSsIvsBzQl5WBr0n7IehmQ-HirPBzw687zy4h0Th_UuhF7tcRoDGwPPXeJI59l40UhcDwybnas39lTIqlDise5IcaFLL-4CL-g09gE3Rq16kVnvF3iAJPQHYH_W34=")
	values.Set("as_fid", "1c579f9054dbad63c7bdb6e328e8b09df4e66feb")
	req, err := http.NewRequest("POST", "https://setagaya.keyakinet.net/Web/Yoyaku/WgR_ShisetsubetsuAkiJoukyou", strings.NewReader(values.Encode()))
	if err != nil {
		panic(err)
	}
	req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	req.Header.Add("accept-encoding", "br")
	req.Header.Add("accept-language", "ja-JP,ja;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Add("cache-control", "max-age=0")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("cookie", "ASP.NET_SessionId=quxotw550a1gdm45pdwjog2j; Setagaya=jAOHJ4LuAepE1ir9QdPnCgDl8vo0000; __RequestVerificationToken_L1dlYg__=c5/nokDg7oF1msjk+oSoQSsZLs566D7g//YncmKzD61ImxlU69rEbmG/9rsxJexKMjL4s8QvA6T4+b2hJeOguLAWBwu4mCqKAm0l78NJfHqgmYoM6tFAS7N2vxJDeS7gTqeEtA==; Setagaya_172.16.81.3_%2F_wat=AAAAAAXMQswAoLXIAmvRduoAlMSUyD6qxAen0NblrpRYPF5YKpqP6hX9PGEHgGANQ9-qK0WkKiBBWJTBvP6SU6hG1oShX317noSDoeuLJHRKiAnstw==&AAAAAAVTvFLVx7QMojfIUD6F7KZCylUCgc0bk8VyxpWe6EpRSD-Qls3uWP62l5qafT-IkYxNlV2wImklOESndSAl3HIo048uZJOl456qnK7HR9C_t3fnGKVGCE4i7RNw-3ybu2c=&")
	req.Header.Add("origin", "https://setagaya.keyakinet.net")
	req.Header.Add("referer", "https://setagaya.keyakinet.net/Web/Yoyaku/WgR_ShisetsuKensaku")
	req.Header.Add("upgrade-insecure-requests", "1")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.121 Safari/537.36")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	doc, err := goquery.NewDocumentFromResponse(resp)

	if err != nil {
		return err
	}

	doc.Find("#body > div.content_body > div.item_body > div > div.scroll-div.clearfix > table > tbody > tr:nth-child(1)").Find("td").Each(func(_ int, s *goquery.Selection) {
		input := s.Find("input")
		val, exist := input.Attr("value")
		log.Println(input, val)
		if exist {
			day := format(val)
			if input.Text() == "△" || input.Text() == "○" {
				if day.Weekday() == time.Saturday {
					r.notifier.Notify(fmt.Sprintf("%vの世田谷公園野球場に空きがでました。", day.Format("2006-01-02")))
				}
			}
		}
	})

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
