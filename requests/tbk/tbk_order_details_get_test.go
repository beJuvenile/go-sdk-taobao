package tbk_test

import (
	"encoding/json"
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
	req.SetQueryType(1)
	req.SetOrderScene(2)
	req.SetJumpType(1)
	req.SetPageNo(1)
	req.SetStartTime("2020-07-13 11:06:00")
	req.SetEndTime("2020-07-13 11:07:00")
	body, err := c.Exec(req)
	if err != nil {
		log.Fatalln(err)
	}
	var result tbk.TbkOrderDetailsGetData
	result, err = tbk.TbkOrderDetailsGetResult(body)
	if err != nil {
		log.Fatalln(err)
	}
	ret, _ := json.Marshal(result.Results)
	log.Println(string(ret))
}
