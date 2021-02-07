package main

import (
	"PugServer/common"
	"PugServer/service"
	"PugServer/util"
	log "github.com/cihub/seelog"
	"time"
)

func InitLog() {
	logger, err := log.LoggerFromConfigAsFile("conf/seelog.xml")
	if err != nil {
		panic(err)
	}

	log.ReplaceLogger(logger)
	log.Debug("############################################################################")
	log.Debug("#						      PugServer									  #")
	log.Debug("############################################################################")
}

func InitConfig() {
	common.LoadConfig()
}

func main() {
	defer log.Flush()

	InitLog()
	InitConfig()

	addr, err := util.Cfg.GetValue("", "ServerAddr")
	if err != nil {
		log.Error("can't find addr:", err)
		panic(err)
	}

	go service.Run(addr)

	for {
		time.Sleep(1 * time.Second)
	}
}
