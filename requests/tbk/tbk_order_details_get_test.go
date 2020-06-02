package tbk_test

import (
	"log"
	"testing"

	"github.com/beJuvenile/go-sdk-taobao/requests/tbk"

	opentaobao "github.com/beJuvenile/go-sdk-taobao"
)

func TestOrderDetailsGet(t *testing.T) {
	c := opentaobao.New()
	c.AppKey = opentaobao.AppKey
	c.AppSecret = opentaobao.AppSecret
	req := tbk.TbkOrderDetailsGetRequest()
	req.SetStartTime("2020-05-27 11:00:00")
	req.SetEndTime("2020-05-27 12:00:00")
	req.SetOrderScene(2)
	body, err := c.Exec(req)
	if err != nil {
		log.Fatalln(err)
	}
	var result tbk.TbkOrderDetailsGetData
	result, err = tbk.TbkOrderDetailsGetResult(body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(result.Results)
}
