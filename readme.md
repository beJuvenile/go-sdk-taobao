## TaoBao SDK <sub>Glang</sub>

#### 使用示例
```
1、初始化client
c := opentaobao.New()
c.AppKey = "2526218"
c.AppSecret = "da2e7dd98976df40fae3899afab4bfe"

2、初始化request
req := tbk.OrderDetailsGetRequest()
req.SetStartTime("2020-05-27 11:00:00")
req.SetEndTime("2020-05-27 18:00:00")
req.SetOrderScene(2)

3、执行请求
body, err := c.Exec(req)
if err != nil {
    log.Fatalln(err)
}

4、解析结果
var result tbk.OrderDetailsGetData
result, err = tbk.OrderDetailsGetResult(body)
if err != nil {
    log.Fatalln(err)
}
log.Println(result.Results)
```
 
 
#### 参考
1. [淘宝开放平台](https://open.taobao.com)