package opentaobao_test

import (
	"log"
	"testing"

	opentaobao "github.com/beJuvenile/go-sdk-taobao"
	"github.com/beJuvenile/go-sdk-taobao/requests/tbk"
)

func TestClient(t *testing.T) {
	c := opentaobao.New()
	c.AppKey = opentaobao.AppKey
	c.AppSecret = opentaobao.AppSecret
	req := tbk.TbkOrderDetailsGetRequest()
	req.SetStartTime("2020-05-27 11:00:00")
	req.SetEndTime("2020-05-27 18:00:00")
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
