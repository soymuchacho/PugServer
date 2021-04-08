package main

import (
	"AuthService/handle"
	"AuthService/service"
	"PugCommon"
	"fmt"
	ini "github.com/Unknwon/goconfig"
	log "github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
	"time"
)

var (
	CONF_SEELOG_FILE string = "conf/seelog.xml"
	INI_CONFIG_FILE  string = "conf/conf.ini"

	LocalAddr string
	DBAddr    string
	RedisAddr string
	EtcdAddr  string
	EtcdTTL   int64

	ServiceType string
	GinMode     string
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
	EtcdAddr = cfg.MustValue("Server", "EtcdAddr", "127.0.0.1:2379")

	ServiceType = cfg.MustValue("Server", "ServiceType", "/PugServer/AuthService")
	GinMode = cfg.MustValue("Server", "GinMode", "release")
	EtcdTTL = cfg.MustInt64("Server", "EtcdTTL", 20)
}

func RegisterService() {
	sername, err := PugCommon.GenServiceName()
	if err != nil {
		panic(fmt.Errorf("Gen server name err %v", err))
	}
	log.Debugf("register service : %v/%v", ServiceType, sername)

	serinfo := PugCommon.ServiceInfo{
		ServiceName: sername,
		ServiceType: ServiceType,
		ServiceAddr: LocalAddr,
		Version:     service.ServiceVersion,
		Load:        0,
	}

	ser, err := service.NewService(EtcdAddr)
	if err != nil {
		panic("gen service error")
	}

	go func() {
		for {
			err := ser.RegisterWithKeep(serinfo, EtcdTTL)
			if err != nil {
				log.Errorf("register with keep error")
			} else {
				log.Errorf("unregistered")
			}

			// 尝试重新注册
			time.Sleep(5 * time.Second)
			log.Debugf("register again %v", serinfo)
		}
	}()
}

func main() {
	ServiceInit()
	RegisterService()

	router := gin.Default()

	v1 := router.Group("v1")
	{
		v1.POST("/login", handle.LoginEndPoint)
		v1.POST("/checkauth, handle.CheckAuthEndPoint")
		v1.GET("/version", handle.VersionEndPoint)
	}

	router.Run(LocalAddr)
}
