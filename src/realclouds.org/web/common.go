package web

import (
	session "github.com/ipfans/echo-session"
	"realclouds.org/utils"
)

//Result *
type Result struct {
	Ok      bool        `json:"ok" xml:"ok"`
	Msg     string      `json:"msg" xml:"msg"`
	ErrCode string      `json:"err_code" xml:"err_code"`
	Data    interface{} `json:"data" xml:"data"`
}

//NewSessionStore 创建一个 Session store 实例
func NewSessionStore(keyPairs ...[]byte) (store session.Store, err error) {

	redisAddr := utils.GetENV("REDIS_ADDR")
	if len(redisAddr) < 1 {
		redisAddr = "127.0.0.1:6379"
	}

	redisPwd := utils.GetENV("REDIS_PWD")
	if len(redisPwd) < 1 {
		redisPwd = ""
	}

	redisSession := utils.GetENVToBool("REDIS_SESSION")
	if redisSession {
		store, err = session.NewRedisStore(32, "tcp", redisAddr, redisPwd, keyPairs...)
	} else {
		store = session.NewCookieStore(keyPairs...)
	}

	sessionTimeout := utils.GetENV("SESSION_TIMEOUT")
	if len(sessionTimeout) < 1 {
		sessionTimeout = "1800"
	}

	sessionTimeoutInt, err := utils.StringUtils(sessionTimeout).Int()
	if nil != err {
		sessionTimeoutInt = 1800
	}

	store.MaxAge(sessionTimeoutInt)

	return
}
