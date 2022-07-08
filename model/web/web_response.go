package web

type WebResponse struct {
	Error       interface{} `json:"error"`
	ErrorDetail interface{} `json:"error_detail"`
	Code        int         `json:"code"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
}
