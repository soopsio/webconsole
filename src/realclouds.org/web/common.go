package web

//Result *
type Result struct {
	Ok      bool        `json:"ok" xml:"ok"`
	Msg     string      `json:"msg" xml:"msg"`
	ErrCode string      `json:"err_code" xml:"err_code"`
	Data    interface{} `json:"data" xml:"data"`
}
