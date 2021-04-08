package main

import (
	"AuthService/handle"
	"AuthService/service"
	"fmt"
	ini "github.com/Unknwon/goconfig"
	log "github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
)

var (
	CONF_SEELOG_FILE string = "conf/seelog.xml"
	INI_CONFIG_FILE  string = "conf/conf.ini"

	LocalAddr string
	DBAddr    string
	RedisAddr string
)

func ServiceInit() {
	defer log.Flush()

	// init log
	logger, err := log.LoggerFromConfigAsFile(CONF_SEELOG_FILE)
	if err != nil {
		panic(fmt.Sprintf("load %v err %v", CONF_SEELOG_FILE, err))
		return
	} else {
		log.ReplaceLogger(logger)
		log.Info("load log success")
	}

	// init config
	cfg, err := ini.LoadConfigFile(INI_CONFIG_FILE)
	if err != nil {
		log.Errorf(err.Error())
		panic(err)
	}

	LocalAddr = cfg.MustValue("Server", "LocalAddr", "0.0.0.0:9020")
	DBAddr = cfg.MustValue("Server", "DBAddr", "127.0.0.1:3306")
	RedisAddr = cfg.MustValue("Server", "RedisAddr", "127.0.0.1:6379")
}

func main() {
	ServiceInit()

	router := gin.Default()

	v1 := router.Group("v1")
	{
		v1.POST("/login", handle.LoginEndPoint)
		v1.POST("/checkauth, handle.CheckAuthEndPoint")
		v1.GET("/version", handle.VersionEndPoint)
	}

	router.Run(LocalAddr)
}
