### 一些工具的集合

##### 支付的排序工具
 - 首先定义一个结构体,并给结构体字段打上tag
 - 设置完值后,调用sort函数,返回的就是排序好的字符串了
 
```go
type UnifyOrderReq struct {
	Appid            string `sort:"appid"`
	Body             string `sort:"body"`
	Mch_id           string `sort:"mch_id"`
	Nonce_str        string `sort:"nonce_str"`
	Notify_url       string `sort:"notify_url"`
	Trade_type       string `sort:"trade_type"`
	Spbill_create_ip string `sort:"spbill_create_ip"`
	Total_fee        int    `sort:"total_fee"`
	Out_trade_no     string `sort:"out_trade_no"`
	Sign             string `sort:"sign"`
}

```
```go
func TestSort(t *testing.T) {
	var yourReq UnifyOrderReq
	AppSecret := "12345678910111213141516171819202"
	yourReq.Appid = "wxxxxxxxxxxxxxxxxx" //微信开放平台我们创建出来的app的app id
	yourReq.Body = "shangpinmiaoshu"
	yourReq.Mch_id = "1111111111" //商户号
	yourReq.Nonce_str = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXx"
	yourReq.Notify_url = "http://www.baidu.com/"
	yourReq.Trade_type = "NATIVE"
	yourReq.Trade_type = "JSAPI"
	yourReq.Spbill_create_ip = "0.0.0.0"
	yourReq.Total_fee = 1                           //单位是分，这里是1毛钱

	currentTime := time.Now().UnixNano()
	r := rand.New(rand.NewSource(currentTime))
	yourReq.Out_trade_no = strconv.Itoa(int(currentTime)) + strconv.Itoa(r.Intn(1000000)) //后台系统单号
	pp := Sort(&yourReq, "sort") + "&key=" + AppSecret
	log.Println(pp)
}

```

##### sql语句的生成工具
```go
func GetReportDataList(SiteId string, indexId string, sTime, eTime string) (uDatas []ReportFc, err error) {
	sql, err := sqlBuilder.Select("表名").
		Where_(indexId, "index_id").
		Where_(sTime, "day_time>=").
		Where_(eTime, "day_time<=").
		Where_(SiteId, "site_id").String()
	if err != nil {
		return
	}
	o := orm.NewOrm()
	o.Using("xxx")
	_, err = o.Raw(sql).QueryRows(&uDatas)
	return
}

```