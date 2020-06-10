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

	for i := 0; i < 150; i++ {
		req := tbk.TbkOrderDetailsGetRequest()
		req.SetStartTime("2020-05-27 11:00:00")
		req.SetEndTime("2020-05-27 11:10:00")
		req.SetOrderScene(2)
		body, err := c.Exec(req)
		if err != nil {
			log.Fatalln(err)
		}
		var result tbk.TbkOrderDetailsGetData
		result, err = tbk.TbkOrderDetailsGetResult(body)
		if err != nil {
			if err.Error() == "接口返回错误" {
				errRet, err := tbk.TbkOrderDetailGetError(body)
				if err == nil {
					log.Println(errRet)
				}
			} else {
				log.Fatalln(err)
			}
		}
		log.Println(result.Results)
	}

}
