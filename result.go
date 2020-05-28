package opentaobao

type ErrorResult struct {
	ErrorResponse ErrorResponse `json:"error_response"`
}

type ErrorResponse struct {
	Code      int64  `json:"code"`
	Msg       string `json:"msg"`
	SubCode   string `json:"sub_code"`
	SubMsg    string `json:"sub_msg"`
	RequestID string `json:"request_id"`
}
