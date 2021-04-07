package main

import (
<<<<<<< HEAD
	"./service"
	"./util"
=======
	"PugServer/common"
	"PugServer/service"
	"PugServer/util"
>>>>>>> ef357e05a2d4ca9fc4ce18fb59ced7e8f0a71238
	log "github.com/cihub/seelog"
	"time"
)

<<<<<<< HEAD
func initLog() {
=======
func InitLog() {
>>>>>>> ef357e05a2d4ca9fc4ce18fb59ced7e8f0a71238
	logger, err := log.LoggerFromConfigAsFile("conf/seelog.xml")
	if err != nil {
		panic(err)
	}

	log.ReplaceLogger(logger)
	log.Debug("############################################################################")
	log.Debug("#						      PugServer									  #")
	log.Debug("############################################################################")
}

<<<<<<< HEAD
func main() {
	defer log.Flush()

	initLog()
=======
func InitConfig() {
	common.LoadConfig()
}

func main() {
	defer log.Flush()

	InitLog()
	InitConfig()
>>>>>>> ef357e05a2d4ca9fc4ce18fb59ced7e8f0a71238

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
