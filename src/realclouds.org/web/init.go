package web

import (
	"fmt"
	"time"

	"github.com/ipfans/echo-session"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/acme/autocert"
	rcmw "realclouds.org/middleware"
	"realclouds.org/utils"
)

var (
	//E Echo
	E = echo.New()
)

//Run 启动服务
func Run() {
	utils.RegGob(time.Time{})

	E.HideBanner = true
	E.Logger.SetLevel(log.DEBUG)

	E.Use(rcmw.MwVersion, rcmw.MwContext)

	E.Use(middleware.RequestID(), middleware.Logger(), middleware.Recover(),
		middleware.GzipWithConfig(middleware.GzipConfig{Level: 9}))

	redisAddr := utils.GetENV("REDIS_ADDR")
	if len(redisAddr) == 0 {
		redisAddr = "127.0.0.1:6379"
	}

	redisPwd := utils.GetENV("REDIS_PWD")
	if len(redisPwd) == 0 {
		redisPwd = ""
	}

	var store session.Store
	var err error

	redisSession := utils.GetENV("REDIS_SESSION")
	if len(redisSession) > 0 && ("true" == redisSession || "1" == redisSession) {
		store, err = session.NewRedisStore(32, "tcp", redisAddr, redisPwd, []byte("#_www-A^yer*Dudu_CoM_$-&==2017.17+20-87"))
		if err != nil {
			E.Logger.Fatal(err)
		}
	} else {
		store = session.NewCookieStore([]byte("#_www-A^yer*Dudu_CoM_$-&==2017.17+20-87"))
	}

	sessionTimeout := utils.GetENV("SESSION_TIMEOUT")
	if len(sessionTimeout) == 0 {
		sessionTimeout = "1800"
	}

	sessionTimeoutInt, err := utils.StringUtils(sessionTimeout).Int()
	if nil != err {
		sessionTimeoutInt = 1800
	}
	store.MaxAge(sessionTimeoutInt)

	E.Use(session.Sessions("rcid", store))

	dbHost := utils.GetENV("DB_HOST")
	if len(dbHost) == 0 {
		dbHost = "127.0.0.1:3306"
	}

	dbUserName := utils.GetENV("DB_USERNAME")
	if len(dbUserName) == 0 {
		dbUserName = "ayerdudu"
	}

	dbPassword := utils.GetENV("DB_PASSWORD")
	if len(dbPassword) == 0 {
		dbPassword = "AyerDudu888"
	}

	dbDataBase := utils.GetENV("DB_DATABASE")
	if len(dbDataBase) == 0 {
		dbDataBase = "ayerdudu"
	}

	dbMaxIdleConns := utils.GetENV("DB_MAXIDLECONNS")
	if len(dbMaxIdleConns) == 0 {
		dbMaxIdleConns = "10"
	}

	dbMaxOpenConns := utils.GetENV("DB_MAXOPENCONNS")
	if len(dbDataBase) == 0 {
		dbMaxOpenConns = "100"
	}

	maxIdleConns, err := utils.StringUtils(dbMaxIdleConns).Int()
	if nil != err {
		maxIdleConns = 10
	}

	maxOpenConns, err := utils.StringUtils(dbMaxOpenConns).Int()
	if nil != err {
		maxOpenConns = 100
	}

	if mySQL, err := rcmw.MySQLConf(
		rcmw.DBConfig{
			Addr:         dbHost,
			MaxIdleConns: maxIdleConns,
			MaxOpenConns: maxOpenConns,
			UserName:     dbUserName,
			Password:     dbPassword,
			DataBase:     dbDataBase,
			LogMode:      true,
		},
	); nil != err {
		E.Logger.Fatalf("%v", err)
	} else {
		em := rcmw.MwMySQL(mySQL)
		E.Use(em.MySQL)
	}

	pd := utils.GetProjectDir()

	E.Static("/static", pd+"/static")

	// E.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
	// 	TokenLookup: "form:X-XSRF-TOKEN",
	// }))

	E.Renderer = rcmw.MwRender(rcmw.RenderOpt{
		Directory: pd + "/templates/default",
	})

	RouterInit()

	port := utils.GetENV("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	tlsPort := utils.GetENV("TLS_PORT")
	if len(tlsPort) == 0 {
		tlsPort = "8443"
	}

	tls := utils.GetENV("TLS")
	if len(tls) > 0 && ("true" == tls || "1" == tls) {
		cert := utils.GetENV("CERT")
		key := utils.GetENV("KEY")
		if len(cert) > 0 && len(key) > 0 {
			E.Logger.Fatal(E.StartTLS(":"+tlsPort, cert, key))
		} else {
			tlsHost := utils.GetENV("TLS_HOST")
			if len(tlsHost) != 0 {
				E.AutoTLSManager.Cache = autocert.DirCache(pd + "/.cache")
				E.Logger.Fatal(E.StartAutoTLS(tlsHost + ":" + tlsPort))
			} else {
				E.Logger.Fatal(fmt.Errorf("%s", "Invalid certificate configuration."))
			}
		}
	} else {
		E.Logger.Fatal(E.Start(":" + port))
	}
}
