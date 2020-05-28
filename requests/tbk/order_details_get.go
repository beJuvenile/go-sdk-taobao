package tbk

import (
	"encoding/json"
	"errors"

	opentaobao "github.com/beJuvenile/go-sdk-taobao"
)

// 淘宝客-推广者-所有订单查询
// Link: https://open.taobao.com/api.htm?docId=43328&docType=2&scopeId=16175
// Author: Ken.Zhang
// Email: kenphp@yeah.net

type OrderDetailsGetResponse struct {
	Data      OrderDetailsGetData `json:"data"`
	RequestID string              `json:"request_id"`
}

type SuccessResult struct {
	Data      OrderDetailsGetData `json:"data"`
	RequestID string              `json:"request_id"`
}

type OrderDetailsGetData struct {
	HasNext       bool   `json:"has_next"`
	HasPre        bool   `json:"has_pre"`
	PageNo        int    `json:"page_no"`
	PageSize      int    `json:"page_size"`
	PositionIndex string `json:"position_index"`
	Results       []struct {
		AdzoneID                           int64  `json:"adzone_id"`
		AdzoneName                         string `json:"adzone_name"`
		AlimamaRate                        string `json:"alimama_rate"`
		AlimamaShareFee                    string `json:"alimama_share_fee"`
		AlipayTotalPrice                   string `json:"alipay_total_price"`
		ClickTime                          string `json:"click_time"`
		DepositPrice                       string `json:"deposit_price"`
		FlowSource                         string `json:"flow_source"`
		IncomeRate                         string `json:"income_rate"`
		ItemCategoryName                   string `json:"item_category_name"`
		ItemImg                            string `json:"item_img"`
		ItemNum                            int    `json:"item_num"`
		ItemTitle                          string `json:"item_title"`
		OrderType                          string `json:"order_type"`
		PayPrice                           string `json:"pay_price"`
		PubID                              int    `json:"pub_id"`
		PubShareFee                        string `json:"pub_share_fee"`
		PubSharePreFee                     string `json:"pub_share_pre_fee"`
		PubShareRate                       string `json:"pub_share_rate"`
		RefundTag                          int    `json:"refund_tag"`
		RelationID                         int    `json:"relation_id"`
		SellerNick                         string `json:"seller_nick"`
		SellerShopTitle                    string `json:"seller_shop_title"`
		SiteID                             int    `json:"site_id"`
		SiteName                           string `json:"site_name"`
		SubsidyFee                         string `json:"subsidy_fee"`
		SubsidyRate                        string `json:"subsidy_rate"`
		SubsidyType                        string `json:"subsidy_type"`
		TbDepositTime                      string `json:"tb_deposit_time"`
		TbPaidTime                         string `json:"tb_paid_time"`
		TerminalType                       string `json:"terminal_type"`
		TkCommissionFeeForMediaPlatform    string `json:"tk_commission_fee_for_media_platform"`
		TkCommissionPreFeeForMediaPlatform string `json:"tk_commission_pre_fee_for_media_platform"`
		TkCommissionRateForMediaPlatform   string `json:"tk_commission_rate_for_media_platform"`
		TkCreateTime                       string `json:"tk_create_time"`
		TkDepositTime                      string `json:"tk_deposit_time"`
		TkEarningTime                      string `json:"tk_earning_time"`
		TkOrderRole                        int    `json:"tk_order_role"`
		TkPaidTime                         string `json:"tk_paid_time"`
		TkStatus                           int    `json:"tk_status"`
		TkTotalRate                        string `json:"tk_total_rate"`
		TotalCommissionFee                 string `json:"total_commission_fee"`
		TotalCommissionRate                string `json:"total_commission_rate"`
		TradeID                            string `json:"trade_id"`
		TradeParentID                      string `json:"trade_parent_id"`
	} `json:"results"`
}

type OrderDetailsGet struct {
	params opentaobao.Parameter // 参数
}

func OrderDetailsGetRequest() *OrderDetailsGet {
	r := new(OrderDetailsGet)
	r.params = make(opentaobao.Parameter)
	return r
}

func OrderDetailsGetResult(data []byte) (OrderDetailsGetData, error) {
	var result SuccessResult
	err := json.Unmarshal(data, &result)
	if err != nil {
		return OrderDetailsGetData{}, err
	}
	if result.RequestID == "" {
		var errResult opentaobao.ErrorResult
		err = json.Unmarshal(data, &errResult)
		return OrderDetailsGetData{}, errors.New(errResult.ErrorResponse.SubMsg)
	}

	return result.Data, nil
}

func (r *OrderDetailsGet) SetQueryType(value int) {
	r.params["query_type"] = value
}

func (r *OrderDetailsGet) SetOrderScene(value int) {
	r.params["order_scene"] = value
}

func (r *OrderDetailsGet) SetPositionIndex(value string) {
	r.params["position_index"] = value
}

func (r *OrderDetailsGet) SetMemberType(value int) {
	r.params["member_type"] = value
}

func (r *OrderDetailsGet) SetTkStatus(value int) {
	r.params["tk_status"] = value
}

func (r *OrderDetailsGet) SetStartTime(value string) {
	r.params["start_time"] = value
}

func (r *OrderDetailsGet) SetEndTime(value string) {
	r.params["end_time"] = value
}

func (r *OrderDetailsGet) SetJumpType(value int) {
	r.params["jump_type"] = value
}

func (r *OrderDetailsGet) SetPageNo(value int) {
	r.params["page_no"] = value
}

func (r *OrderDetailsGet) SetPageSize(value int) {
	r.params["page_size"] = value
}

func (r *OrderDetailsGet) GetMethod() string {
	return "taobao.tbk.order.details.get"
}

func (r *OrderDetailsGet) GetApiParams() opentaobao.Parameter {
	return r.params
}
