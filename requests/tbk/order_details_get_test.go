package tbk_test

import (
	"log"
	"testing"

	"github.com/beJuvenile/go-sdk-taobao/requests/tbk"

	opentaobao "github.com/beJuvenile/go-sdk-taobao"
)

func TestOrderDetailsGet(t *testing.T) {
	c := opentaobao.New()
	c.AppKey = "2526218"
	c.AppSecret = "da2e7dd98976df40fae3899afab4bfe"
	req := tbk.OrderDetailsGetRequest()
	req.SetStartTime("2020-05-27 11:00:00")
	req.SetEndTime("2020-05-27 18:00:00")
	req.SetOrderScene(2)
	body, err := c.Exec(req)
	if err != nil {
		log.Fatalln(err)
	}
	var result tbk.OrderDetailsGetData
	result, err = tbk.OrderDetailsGetResult(body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(result.Results)
}
