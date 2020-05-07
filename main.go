package main

import (
	"./service"
	"./util"
	log "github.com/cihub/seelog"
	"time"
)

func initLog() {
	logger, err := log.LoggerFromConfigAsFile("conf/seelog.xml")
	if err != nil {
		panic(err)
	}

	log.ReplaceLogger(logger)
	log.Debug("############################################################################")
	log.Debug("#						      PugServer									  #")
	log.Debug("############################################################################")
}

func main() {
	defer log.Flush()

	initLog()

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
