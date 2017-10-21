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

	devMode := utils.GetENVToBool("DEV_MODE")
	if devMode {
		E.Logger.SetLevel(log.DEBUG)
	}

	E.Use(rcmw.MwVersion, rcmw.MwContext)

	E.Use(middleware.CORS(), middleware.RequestID(), middleware.Logger(), middleware.Recover())

	E.Use(middleware.GzipWithConfig(middleware.GzipConfig{Level: 9}))

	store, err := NewSessionStore([]byte("#Www-RealClouds_Org_Or_wWw.rEalcLouds_iO*"))
	if nil != err {
		E.Logger.Fatalf("New session store error: %v", err.Error())
	}

	E.Use(session.Sessions("rcid", store))

	// E.Use(rcmw.MySQL().MwMySQL)

	pd := utils.GetProjectDir()

	E.Static("/static", pd+"/static")

	E.Renderer = rcmw.MwRender(rcmw.RenderOpt{
		Directory: pd + "/templates/default",
		DevMode:   devMode,
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

	tls := utils.GetENVToBool("TLS")
	if tls {
		cert := utils.GetENV("CERT")
		key := utils.GetENV("KEY")
		if len(cert) > 0 && len(key) > 0 {
			E.Logger.Fatal(E.StartTLS(":"+tlsPort, cert, key))
		} else {
			tlsHost := utils.GetENV("TLS_HOST")
			if len(tlsHost) > 0 {
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
